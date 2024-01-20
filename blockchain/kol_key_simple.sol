// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./common.sol";

contract KolKeySimple is ServiceFeeForWithdraw, KolIncomeToPoolI {
    struct KeySettings {
        uint256 price;
        uint256 nonce;
        uint256 totalVal;
        uint256 totalNo;
    }

    struct KolKey {
        mapping(uint256 => uint256) amount;
        uint256[] nonce;
    }

    struct MapArray {
        mapping(address => uint256) filter;
        address[] list;
    }

    uint8 public kolIncomeRatePerKeyBuy = 90;
    uint8 public serviceFeeRatePerKeyBuy = 10;
    uint256 public maxKeyNoForKol = 100;

    mapping(address => KeySettings) public KeySettingsRecord;
    mapping(address => mapping(address => KolKey)) keyBalance;
    mapping(address => mapping(uint256 => uint256)) public incomePerNoncePerKey;

    mapping(address => MapArray) private keyHoldersOfKol;
    mapping(address => MapArray) private kolsOfKeyHolder;
    address[] public allKolInSystem;

    event InvestorWithdrawByOneNonce(
        address investor,
        address kol,
        uint256 nonce,
        uint256 val
    );

    event InvestorWithdrawByOneKol(address investor, address kol, uint256 val);
    event InvestorWithdrawAllIncome(address investor, uint256 val);
    event KolKeyOpened(address kol, uint256 pricePerKey);
    event KolIncomeToPoolAction(
        int8 sourceID,
        address sourceConract,
        address kol,
        uint256 keyNo,
        uint256 keyNonce,
        uint256 amount,
        uint256 valPerKey
    );
    event KolKeyBought(
        address kolAddr,
        address buyer,
        uint256 keyNo,
        uint256 curNonce,
        uint256 KoltotalNo
    );
    event KeyTransfer(
        address from,
        address to,
        address kol,
        uint256 nonce,
        uint256 amount
    );
    event KeyTransferAll(address from, address to, address kol);
    event KeyRebound(address from, address to);
    event SystemSet(uint256 num, string op);

    receive() external payable {}

    /********************************************************************************
     *                       admin operation
     *********************************************************************************/

    function adminSetKolIncomeRate(uint8 newRate) public isOwner {
        require(newRate >= 0 && newRate <= 100, "invalid rate");
        kolIncomeRatePerKeyBuy = newRate;
        emit SystemSet(newRate, "kol_key_income_rate_changed");
    }

    function adminSetKeyFeeRate(uint8 newRate) public isOwner {
        require(newRate >= 0 && newRate <= 100, "invalid rate");
        serviceFeeRatePerKeyBuy = newRate;
        emit SystemSet(newRate, "kol_key_fee_rate_changed");
    }

    function adminSetMaxKolKeyNo(uint256 keyNo) public isOwner {
        require(maxKeyNoForKol != keyNo, "no need to change");
        maxKeyNoForKol = keyNo;
        emit SystemSet(keyNo, "kol_key_max_no_changed");
    }

    /********************************************************************************
     *                       kol operation
     *********************************************************************************/
    function kolOpenKeySale(uint256 priceInFin) public inRun {
        require(priceInFin * 1e6 gwei >= __minValCheck, "too low price");

        KeySettings storage ks = KeySettingsRecord[msg.sender];
        require(ks.nonce == 0, "duplicate operation");

        ks.price = priceInFin * 1e6 gwei;
        ks.totalNo = 0;
        ks.nonce = 1;

        allKolInSystem.push(msg.sender);

        emit KolKeyOpened(msg.sender, ks.price);
    }

    /********************************************************************************
     *                       value to kol key pool
     *********************************************************************************/

    function kolOpenKeyPool(address sourceAddr) external view returns (bool) {
        return KeySettingsRecord[sourceAddr].nonce > 0;
    }

    function kolGotIncome(int8 sourceID, address kolAddr)
    external
    payable
    noReentrant
    isValidAddress(kolAddr)
    inRun
    {
        uint256 val = msg.value;
        require(val > __minValCheck, "invalid msg value");

        KeySettings storage ks = KeySettingsRecord[kolAddr];
        ks.totalVal += val;

        if (ks.totalNo == 0) {
            balance[kolAddr] += val;
            return;
        }

        uint256 valPerKey = val / ks.totalNo;
        incomePerNoncePerKey[kolAddr][ks.nonce] = valPerKey;
        ks.nonce += 1;

        emit KolIncomeToPoolAction(
            sourceID,
            kolAddr,
            kolAddr,
            ks.totalNo,
            ks.nonce,
            val,
            valPerKey
        );
    }

    function checkPluginInterface() external pure returns (bool) {
        return true;
    }

    function buyKolKey(address kolAddr, uint256 keyNo)
    public
    payable
    isValidAddress(kolAddr)
    noReentrant
    inRun
    {
        require(keyNo >= 1, "invalid key count");

        KeySettings storage ks = KeySettingsRecord[kolAddr];
        require(ks.nonce >= 1 && ks.price > __minValCheck, "key not open");
        uint256 amount = ks.price * keyNo;
        require(msg.value == amount, "price of kol's key has changed");
        require(
            maxKeyNoForKol >= ks.totalNo + keyNo,
            "more than key limitation"
        );

        ks.totalNo += keyNo;

        KolKey storage kk = keyBalance[msg.sender][kolAddr];
        if (kk.amount[ks.nonce] == 0) {
            kk.nonce.push(ks.nonce);
        }
        kk.amount[ks.nonce] += keyNo;

        MapArray storage investors = keyHoldersOfKol[kolAddr];
        if (investors.filter[msg.sender] == 0) {
            investors.list.push(msg.sender);
            investors.filter[msg.sender] = investors.list.length;
        }

        MapArray storage kols = kolsOfKeyHolder[msg.sender];
        if (kols.filter[kolAddr] == 0) {
            kols.list.push(kolAddr);
            kols.filter[kolAddr] = kols.list.length;
        }

        uint256 fee = (amount / 100) * serviceFeeRatePerKeyBuy;
        recordServiceFee(fee);

        uint256 incomeForKol = amount - fee;
        balance[kolAddr] += incomeForKol;

        emit KolKeyBought(kolAddr, msg.sender, keyNo, ks.nonce, ks.totalNo);
    }

    /********************************************************************************
     *                       key auction
     *********************************************************************************/

    function transferKey(
        address to,
        address kol,
        uint256 nonce,
        uint256 amount
    ) public isValidAddress(kol) inRun {
        require(amount > 0, "invalid amount");
        KolKey storage key = keyBalance[msg.sender][kol];
        require(key.amount[nonce] >= amount, "no enough key to transfer");

        KolKey storage toKey = keyBalance[to][kol];

        if (toKey.amount[nonce] == 0) {
            toKey.nonce.push(nonce);
        }
        key.amount[nonce] -= amount;
        toKey.amount[nonce] += amount;

        emit KeyTransfer(msg.sender, to, kol, nonce, amount);
    }

    function transferAllKey(address to, address kol)
    public
    isValidAddress(kol)
    inRun
    {
        KolKey storage key = keyBalance[msg.sender][kol];

        for (uint256 idx = 0; idx < key.nonce.length; idx++) {
            uint256 non = key.nonce[idx];

            if (key.amount[non] == 0) {
                continue;
            }

            KolKey storage toKey = keyBalance[to][kol];

            if (toKey.amount[non] == 0) {
                toKey.nonce.push(non);
            }

            toKey.amount[non] += key.amount[non];
            delete key.amount[non];
        }

        delete key.nonce;
        delete keyBalance[msg.sender][kol];

        emit KeyTransferAll(msg.sender, to, kol);
    }

    function rebindKolKey(address to) public {
        require(to != address(0), "invalid to address");

        KeySettings storage fromSetting = KeySettingsRecord[msg.sender];
        KeySettings storage toSetting = KeySettingsRecord[to];

        require(fromSetting.nonce >= 1, "not key for kol");
        require(toSetting.nonce == 0, "to address has open");

        for (uint256 y = 0; y < fromSetting.nonce; y++) {
            incomePerNoncePerKey[to][y] = incomePerNoncePerKey[msg.sender][y];
            delete incomePerNoncePerKey[msg.sender][y];
        }

        KeySettingsRecord[to] = fromSetting;
        delete KeySettingsRecord[msg.sender];

        MapArray storage from_holders = keyHoldersOfKol[msg.sender];
        MapArray storage to_holders = keyHoldersOfKol[msg.sender];

        for (uint256 i = 0; i < from_holders.list.length; i++) {
            address holder = from_holders.list[i];

            to_holders.list.push(holder);
            to_holders.filter[holder] = to_holders.list.length;

            KolKey storage key_from = keyBalance[holder][msg.sender];
            KolKey storage key_to = keyBalance[holder][to];

            for (uint256 x = 0; x < key_from.nonce.length; x++) {
                uint256 nonce = key_from.nonce[x];
                uint256 amount = key_from.amount[nonce];
                if (amount == 0) {
                    continue;
                }
                key_to.nonce.push(nonce);
                key_to.amount[nonce] = amount;
            }
            delete keyBalance[holder][msg.sender];

            uint256 idx = kolsOfKeyHolder[holder].filter[msg.sender];
            if (idx > 0) {
                kolsOfKeyHolder[holder].list[idx - 1] = to;
                kolsOfKeyHolder[holder].filter[to] = idx;
            }
        }

        delete keyHoldersOfKol[msg.sender];

        allKolInSystem.push(to);

        emit KeyRebound(msg.sender, to);
    }

    /********************************************************************************
     *                       income withdraw
     *********************************************************************************/

    function withdrawFromOneNonce(address kol, uint256 nonce)
    public
    noReentrant
    isValidAddress(kol)
    inRun
    {
        KeySettings storage ks = KeySettingsRecord[kol];
        KolKey storage key = keyBalance[msg.sender][kol];

        require(ks.nonce >= 1, "no income for kol");
        require(key.amount[nonce] > 0, "no key in this nonce");

        uint256 val = 0;
        for (uint256 idx = nonce; idx <= (ks.nonce - 1); idx++) {
            uint256 valPerKey = incomePerNoncePerKey[kol][idx];
            if (valPerKey == 0) {
                continue;
            }
            val += valPerKey * key.amount[nonce];
        }
        require(val <= address(this).balance, "insufficient funds");
        require(val > __minValCheck, "too small funds");

        if (key.amount[ks.nonce] == 0) {
            key.nonce.push(ks.nonce);
        }
        key.amount[ks.nonce] += key.amount[nonce];

        delete key.amount[nonce];
        uint256 reminders = minusWithdrawFee(val);

        payable(msg.sender).transfer(reminders);

        emit InvestorWithdrawByOneNonce(msg.sender, kol, nonce, reminders);
    }

    function withdrawFromOneKol(address kol)
    public
    noReentrant
    isValidAddress(kol)
    inRun
    {
        privateWithdrawFromOneKol(kol, msg.sender, true);
    }

    function privateWithdrawFromOneKol(
        address kol,
        address investor,
        bool once
    ) internal returns (uint256) {
        KeySettings storage ks = KeySettingsRecord[kol];
        KolKey storage key = keyBalance[investor][kol];

        uint256 val = 0;
        uint256 totalKeyNo = 0;

        for (uint256 kIdx = 0; kIdx < key.nonce.length; kIdx++) {
            uint256 non = key.nonce[kIdx];
            if (key.amount[non] == 0) {
                continue;
            }

            for (
                uint256 nonceIdx = non;
                nonceIdx <= (ks.nonce - 1);
                nonceIdx++
            ) {
                uint256 valPerKey = incomePerNoncePerKey[kol][nonceIdx];
                if (valPerKey == 0) {
                    continue;
                }
                val += valPerKey * key.amount[non];
            }

            totalKeyNo += key.amount[non];
            delete key.amount[non];
        }
        require(
            val > __minValCheck && val <= address(this).balance,
            "invalid val withdraw"
        );

        delete key.nonce;
        if (key.amount[ks.nonce] == 0) {
            key.nonce.push(ks.nonce);
        }
        key.amount[ks.nonce] += totalKeyNo;

        if (once) {
            uint256 reminders = minusWithdrawFee(val);
            payable(msg.sender).transfer(reminders);
            emit InvestorWithdrawByOneKol(msg.sender, kol, reminders);
        }
        return val;
    }

    function withdrawAllIncome() public noReentrant inRun {
        MapArray storage kol = kolsOfKeyHolder[msg.sender];
        require(kol.list.length > 0, "no investment");

        uint256 val = 0;
        for (uint256 idx = 0; idx < kol.list.length; idx++) {
            val += privateWithdrawFromOneKol(kol.list[idx], msg.sender, false);
        }

        uint256 reminders = minusWithdrawFee(val);

        payable(msg.sender).transfer(reminders);
        emit InvestorWithdrawAllIncome(msg.sender, reminders);
    }

    /********************************************************************************
     *                       income query
     *********************************************************************************/

    function IncomeOfOneNonceByAmount(
        address kol,
        uint256 nonce,
        uint256 amount
    ) public view returns (uint256) {
        KeySettings memory ks = KeySettingsRecord[kol];
        if (amount == 0 || ks.nonce <= 1) {
            return 0;
        }

        uint256 sumForThisNonce = 0;
        for (
            uint256 startNonce = nonce;
            startNonce <= (ks.nonce - 1);
            startNonce++
        ) {
            uint256 valPerKey = incomePerNoncePerKey[kol][startNonce];
            sumForThisNonce += valPerKey * amount;
        }

        return sumForThisNonce;
    }

    function IncomeOfOneNonce(
        address kol,
        uint256 nonce,
        address investor
    ) public view returns (uint256) {
        KeySettings memory ks = KeySettingsRecord[kol];

        KolKey storage record = keyBalance[investor][kol];
        if (record.amount[nonce] == 0 || ks.nonce <= 1) {
            return 0;
        }

        uint256 sumForThisNonce = 0;
        for (
            uint256 startNonce = nonce;
            startNonce <= (ks.nonce - 1);
            startNonce++
        ) {
            uint256 valPerKey = incomePerNoncePerKey[kol][startNonce];
            sumForThisNonce += valPerKey * record.amount[nonce];
        }

        return sumForThisNonce;
    }

    function AllIncomeOfOneKol(address kol, address investor)
    public
    view
    returns (uint256)
    {
        KolKey storage record = keyBalance[investor][kol];
        uint256 sumIncome = 0;
        for (uint256 idx = 0; idx < record.nonce.length; idx++) {
            uint256 nonce = record.nonce[idx];
            if (record.amount[nonce] == 0) {
                continue;
            }
            uint256 oneIncome = IncomeOfOneNonceByAmount(
                kol,
                nonce,
                record.amount[nonce]
            );
            sumIncome += oneIncome;
        }
        return sumIncome;
    }

    function AllIncomeOfAllKol(address investor) public view returns (uint256) {
        uint256 sumIncome = 0;
        MapArray storage record = kolsOfKeyHolder[investor];
        for (uint256 idx = 0; idx < record.list.length; idx++) {
            address kol = record.list[idx];
            uint256 oneIncome = AllIncomeOfOneKol(kol, investor);
            sumIncome += oneIncome;
        }
        return sumIncome;
    }

    /********************************************************************************
     *                       basic query
     *********************************************************************************/
    function InvestorAllKeysOfKol(address investor, address kol)
    public
    view
    returns (uint256[] memory nonce, uint256[] memory amounts)
    {
        KolKey storage record = keyBalance[investor][kol];
        if (record.nonce.length == 0) {
            return (new uint256[](0), new uint256[](0));
        }

        nonce = new uint256[](record.nonce.length);
        amounts = new uint256[](record.nonce.length);
        for (uint256 idx = 0; idx < record.nonce.length; idx++) {
            uint256 n = record.nonce[idx];
            nonce[idx] = n;
            amounts[idx] = record.amount[n];
        }
        return (nonce, amounts);
    }

    function InvestorOfKol(address kol) public view returns (address[] memory) {
        MapArray storage investors = keyHoldersOfKol[kol];
        return investors.list;
    }

    function KolOfOneInvestor(address investor)
    public
    view
    returns (address[] memory)
    {
        MapArray storage kol = kolsOfKeyHolder[investor];
        return kol.list;
    }

    function AllKolAddr() public view returns (address[] memory) {
        return allKolInSystem;
    }
}
