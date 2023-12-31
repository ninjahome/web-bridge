// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./common.sol";

contract KolKeys is ServiceFeeForWithdraw, TweetVotePlugInI {
    struct KeySettings {
        uint256 price;
        uint256 maxNo;
        uint256 nonce;
        uint256 totalNo;
    }
    struct keymeta {
        uint256 amount;
        uint256 indx;
    }
    struct KolKey {
        mapping(uint256 => keymeta) nonceToAmount;
        uint256[] nonceList;
    }

    struct MapArray {
        mapping(address => bool) filter;
        address[] list;
    }

    struct KeyAuction {
        uint256 id;
        address seller;
        address bidder;
        address kol;
        uint256 nonce;
        uint256 amount;
        uint256 price;
        bool status;
    }

    uint256 public __feeForKolKeyOp = 0.001 ether;
    uint256 public __feeForKeyBiddig = 0.001 ether;
    uint8 public kolIncomeRatePerKeyBuy = 90;
    uint8 public serviceFeeRatePerKeyBuy = 10;

    mapping(address => KeySettings) public KeySettingsRecord;
    mapping(address => mapping(address => KolKey)) keyBalance;
    mapping(address => mapping(uint256 => uint256)) public incomePerNoncePerKey;

    mapping(address => MapArray) private keyHoldersOfKol;
    mapping(address => MapArray) private kolsOfInKeyHolder;

    uint256 public __currentAuctionID = 1e5;
    KeyAuction[] public keyInAuction;
    mapping(uint256 => uint256) public auctionIdxMap;

    event KolKeyOperation(
        address kol,
        uint256 price,
        uint256 maxKeyNo,
        string op
    );

    event InvestorWithdrawByOneNonce(
        address investor,
        address kol,
        uint256 nonce,
        uint256 val
    );

    event InvestorWithdrawByOneKol(address investor, address kol, uint256 val);
    event InvestorWithdrawAllIncome(address investor, uint256 val);
    event TweetBought(
        bytes32 tHash,
        address owner,
        address buyer,
        uint256 voteNoe
    );

    event KolIpRightBought(
        address kolAddr,
        address buyer,
        uint256 keyNo,
        uint256 curNonce,
        uint256 KoltotalNo
    );

    event SystemSet(uint256 num, string op);
    event KeyAuctionAction(
        address kol,
        address seller,
        address buyer,
        uint256 nonce,
        uint256 amount,
        uint256 pricePerKey,
        string typ
    );

    event KeyTransfered(
        address from,
        address to,
        address kol,
        uint256 nonce,
        uint256 amount
    );

    event KeyTransferedAll(address from, address to, address kol);

    receive() external payable {}

    function removeFromArray(uint256 indexPlusOne, uint256[] storage array)
    internal
    {
        if (array.length == 0) {
            return;
        }
        if (array.length == 1) {
            array.pop();
            return;
        }

        require(
            indexPlusOne >= 1 && indexPlusOne <= array.length,
            "Index out of bounds"
        );

        array[indexPlusOne - 1] = array[array.length - 1];
        array.pop();
    }

    /********************************************************************************
     *                       admin operation
     *********************************************************************************/
    function adminSetKolOperationFee(uint256 newFeeInGwei) public isOwner {
        newFeeInGwei = newFeeInGwei * 1 gwei;
        require(newFeeInGwei >= __minValCheck, "to small price");
        __feeForKolKeyOp = newFeeInGwei;

        emit SystemSet(newFeeInGwei, "kol_operation_fee_changed");
    }

    function adminSetBddingFee(uint256 newFeeInGwei) public isOwner {
        newFeeInGwei = newFeeInGwei * 1 gwei;
        require(newFeeInGwei >= __minValCheck, "to small price");
        __feeForKeyBiddig = newFeeInGwei;

        emit SystemSet(newFeeInGwei, "bidding_fee_changed");
    }

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

    /********************************************************************************
     *                       kol operation
     *********************************************************************************/
    function kolOpenKeySale(uint256 pricePerKey, int256 maxKeyNo)
    public
    payable
    inRun
    {
        require(maxKeyNo != 0, "invalid max key no");
        require(msg.value == __feeForKolKeyOp, "fee of key setting changed");
        require(pricePerKey >= __minValCheck, "too low price");

        KeySettings storage ks = KeySettingsRecord[msg.sender];
        require(ks.totalNo == 0, "duplicate operation");

        uint256 max_keys = maxKeyNo < 0 ? type(uint256).max : uint256(maxKeyNo);

        ks.price = pricePerKey;
        ks.maxNo = max_keys;
        ks.totalNo = 1;

        recordServiceFee(__feeForKolKeyOp);

        emit KolKeyOperation(
            msg.sender,
            pricePerKey,
            max_keys,
            "kol_key_openned"
        );
    }

    function kolSetKeyPrice(uint256 newPrice) public payable inRun {
        require(msg.value == __feeForKolKeyOp, "fee of key setting changed");
        require(newPrice >= __minValCheck, "too low price");
        KeySettings storage ks = KeySettingsRecord[msg.sender];
        require(ks.totalNo >= 1, "open your key sale first");

        ks.price = newPrice;

        recordServiceFee(__feeForKolKeyOp);

        emit KolKeyOperation(
            msg.sender,
            newPrice,
            ks.maxNo,
            "kol_key_price_changed"
        );
    }

    function kolAddKeySupply(uint256 amount) public payable inRun {
        require(msg.value == __feeForKolKeyOp, "fee of key setting changed");
        require(amount >= 1, "too small size");
        KeySettings storage ks = KeySettingsRecord[msg.sender];
        require(ks.totalNo >= 1, "open your key sale first");
        require(ks.maxNo < type(uint256).max, "too many keys");

        ks.maxNo += amount;
        recordServiceFee(__feeForKolKeyOp);

        emit KolKeyOperation(
            msg.sender,
            ks.price,
            amount,
            "kol_key_supply_add"
        );
    }

    /********************************************************************************
     *                       value to kol key pool
     *********************************************************************************/

    function incomeToKolPool(address kol)
    public
    payable
    isValidAddress(kol)
    noReentrant
    inRun
    {
        require(msg.value > __minValCheck, "too small funds");
        KeySettings storage ks = KeySettingsRecord[msg.sender];
        if (ks.totalNo <= 1) {
            balance[kol] += msg.value;
            return;
        }

        uint256 valPerKey = msg.value / ks.totalNo;
        incomePerNoncePerKey[kol][ks.nonce] += valPerKey;

        emit KolKeyOperation(kol, msg.value, ks.nonce, "kol_key_val_increase");
    }

    function tweetBought(
        bytes32 tweetHash,
        address tweetOwner,
        address buyer,
        uint256 voteNo
    )
    public
    payable
    isValidAddress(buyer)
    isValidAddress(tweetOwner)
    noReentrant
    onlyAdmin
    inRun
    {
        uint256 val = msg.value;
        require(val > __minValCheck, "invalid msg value");
        require(voteNo >= 1, "invalid vote no");
        require(tweetHash != bytes32(0));

        KeySettings storage ks = KeySettingsRecord[msg.sender];

        if (ks.totalNo <= 1) {
            balance[tweetOwner] += val;
            return;
        }

        uint256 valPerKey = val / ks.totalNo;

        incomePerNoncePerKey[tweetOwner][ks.nonce] += valPerKey;

        emit TweetBought(tweetHash, tweetOwner, buyer, voteNo);
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
        require(keyNo >= 1, "invalid key no");

        KeySettings storage ks = KeySettingsRecord[kolAddr];
        require(ks.totalNo >= 1, "ip right for kol not open");
        require(
            ks.maxNo >= (ks.totalNo + keyNo),
            "key no is more than kol's setting"
        );

        uint256 amount = ks.price * keyNo;
        require(msg.value == amount, "price of kol's key has changed");

        ks.nonce++;
        ks.totalNo += keyNo;

        KolKey storage b = keyBalance[msg.sender][kolAddr];
        b.nonceList.push(ks.nonce);
        b.nonceToAmount[ks.nonce] = keymeta(keyNo, b.nonceList.length);

        MapArray storage investors = keyHoldersOfKol[kolAddr];
        if (investors.filter[msg.sender] == false) {
            investors.filter[msg.sender] = true;
            investors.list.push(msg.sender);
        }

        MapArray storage kols = kolsOfInKeyHolder[msg.sender];
        if (kols.filter[kolAddr] == false) {
            kols.filter[kolAddr] = true;
            kols.list.push(kolAddr);
        }

        uint256 incomeForKol = (amount / 100) * kolIncomeRatePerKeyBuy;
        balance[kolAddr] += incomeForKol;

        uint256 fee = (amount / 100) * serviceFeeRatePerKeyBuy;
        recordServiceFee(fee);

        emit KolIpRightBought(kolAddr, msg.sender, keyNo, ks.nonce, ks.totalNo);
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
        keymeta storage km = key.nonceToAmount[nonce];
        require(km.amount >= amount, "no enough key to bid");

        KolKey storage toKey = keyBalance[to][kol];
        keymeta storage toKm = toKey.nonceToAmount[nonce];

        if (toKm.amount > 0) {
            km.amount -= amount;
            toKm.amount += amount;
        } else {
            km.amount -= amount;
            toKey.nonceList.push(nonce);
            toKey.nonceToAmount[nonce] = keymeta(
                amount,
                toKey.nonceList.length
            );
        }

        emit KeyTransfered(msg.sender, to, kol, nonce, amount);
    }

    function transferAllKey(address to, address kol)
    public
    isValidAddress(kol)
    inRun
    {
        KolKey storage key = keyBalance[msg.sender][kol];
        for (uint256 idx = 0; idx < key.nonceList.length; idx++) {
            uint256 nonce = key.nonceList[idx];

            keymeta memory km = key.nonceToAmount[nonce];

            KolKey storage toKey = keyBalance[to][kol];
            keymeta storage toKm = toKey.nonceToAmount[nonce];

            if (toKm.amount > 0) {
                toKm.amount += km.amount;
            } else {
                toKey.nonceList.push(nonce);
                toKey.nonceToAmount[nonce] = keymeta(
                    km.amount,
                    toKey.nonceList.length
                );
            }
            delete key.nonceToAmount[nonce];
        }

        delete key.nonceList;
        delete keyBalance[msg.sender][kol];

        emit KeyTransferedAll(msg.sender, to, kol);
    }

    function issueBid(
        address kol,
        uint256 nonce,
        uint256 amount,
        uint256 pricePerKey
    ) public payable isValidAddress(kol) inRun {
        require(msg.value == __feeForKeyBiddig, "bidding fee changed");
        require(amount > 0, "invalid amount");
        require(pricePerKey > __minValCheck, "too low price");

        KolKey storage key = keyBalance[msg.sender][kol];
        keymeta storage km = key.nonceToAmount[nonce];

        require(km.amount >= amount, "no enough key to bid");

        km.amount -= amount;
        KeyAuction memory ka = KeyAuction(
            __currentAuctionID,
            msg.sender,
            address(0),
            kol,
            nonce,
            amount,
            pricePerKey,
            true
        );

        keyInAuction.push(ka);
        auctionIdxMap[__currentAuctionID] = keyInAuction.length;
        __currentAuctionID++;

        recordServiceFee(__feeForKeyBiddig);

        emit KeyAuctionAction(
            kol,
            msg.sender,
            address(0),
            nonce,
            amount,
            pricePerKey,
            "new issue"
        );
    }

    function bid(uint256 bidID) public payable inRun {
        uint256 idx = auctionIdxMap[bidID];
        require(idx >= 1, "no such auction");

        KeyAuction storage ka = keyInAuction[idx - 1];
        require(
            ka.seller != address(0) && ka.id == bidID,
            "invalid bid target"
        );
        require(ka.bidder == address(0) && ka.status, "bidding has finished");

        uint256 val = ka.amount * ka.price;
        require(msg.value == val + __feeForKeyBiddig, "insufficient funds");

        ka.bidder = msg.sender;
        ka.status = false;
        KolKey storage toKey = keyBalance[msg.sender][ka.kol];
        keymeta storage toKm = toKey.nonceToAmount[ka.nonce];
        if (toKm.amount > 0) {
            toKm.amount += ka.amount;
        } else {
            toKey.nonceList.push(ka.nonce);
            toKey.nonceToAmount[ka.nonce] = keymeta(
                ka.amount,
                toKey.nonceList.length
            );
        }

        payable(ka.seller).transfer(val);

        recordServiceFee(__feeForKeyBiddig);

        emit KeyAuctionAction(
            ka.kol,
            ka.seller,
            msg.sender,
            ka.nonce,
            ka.amount,
            ka.price,
            "bid success"
        );
    }

    function revokeBid(uint256 bidID) public payable inRun {
        uint256 idx = auctionIdxMap[bidID];
        require(idx >= 1, "no such auction");

        require(msg.value == __feeForKeyBiddig, "fee needed");

        KeyAuction storage ka = keyInAuction[idx - 1];
        require(
            ka.bidder == address(0) && ka.id == bidID && ka.status,
            "can't be revoked"
        );
        require(ka.seller == msg.sender, "no rights");

        KolKey storage toKey = keyBalance[msg.sender][ka.kol];
        keymeta storage toKm = toKey.nonceToAmount[ka.nonce];
        if (toKm.amount > 0) {
            toKm.amount += ka.amount;
        } else {
            toKey.nonceList.push(ka.nonce);
            toKey.nonceToAmount[ka.nonce] = keymeta(
                ka.amount,
                toKey.nonceList.length
            );
        }

        delete auctionIdxMap[bidID];
        ka.status = false;

        recordServiceFee(__feeForKeyBiddig);

        emit KeyAuctionAction(
            ka.kol,
            ka.seller,
            msg.sender,
            ka.nonce,
            ka.amount,
            ka.price,
            "bid success"
        );
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

        keymeta memory km = key.nonceToAmount[nonce];
        require(km.amount > 0, "no key in this nonce");

        uint256 val = 0;
        for (uint256 idx = nonce; idx <= ks.nonce; idx++) {
            uint256 valPerKey = incomePerNoncePerKey[kol][idx];
            if (valPerKey == 0) {
                continue;
            }

            val += valPerKey * km.amount;
        }
        require(val <= address(this).balance, "insufficient funds");
        require(val > __minValCheck, "too small funds");

        removeFromArray(km.indx, key.nonceList);
        delete key.nonceToAmount[nonce];

        ks.nonce++;

        key.nonceList.push(ks.nonce);
        key.nonceToAmount[ks.nonce] = keymeta(km.amount, key.nonceList.length);

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
    ) internal isValidAddress(investor) {
        KeySettings storage ks = KeySettingsRecord[kol];
        KolKey storage key = keyBalance[msg.sender][kol];

        uint256 val = 0;
        uint256 totalKeyNo = 0;
        for (uint256 kIdx = 0; kIdx < key.nonceList.length; kIdx++) {
            uint256 curNonceInvestorHas = key.nonceList[kIdx];
            keymeta memory km = key.nonceToAmount[curNonceInvestorHas];

            for (
                uint256 nonceIdx = curNonceInvestorHas;
                nonceIdx <= ks.nonce;
                nonceIdx++
            ) {
                uint256 valPerKey = incomePerNoncePerKey[kol][nonceIdx];
                if (valPerKey == 0) {
                    continue;
                }

                val += valPerKey * km.amount;
            }

            totalKeyNo += km.amount;
            delete key.nonceToAmount[curNonceInvestorHas];
        }
        require(
            val > __minValCheck && val <= address(this).balance,
            "invalid val withdraw"
        );

        delete key.nonceList;
        ks.nonce++;

        key.nonceList.push(ks.nonce);
        key.nonceToAmount[ks.nonce] = keymeta(totalKeyNo, key.nonceList.length);

        if (once) {
            uint256 reminders = minusWithdrawFee(val);
            payable(msg.sender).transfer(reminders);
            emit InvestorWithdrawByOneKol(msg.sender, kol, reminders);
        }
    }

    function withdrawAllIncome() public noReentrant inRun {
        MapArray storage kol = kolsOfInKeyHolder[msg.sender];
        require(kol.list.length > 0, "no investment");

        uint256 val = AllIncomeOfAllKol(msg.sender);
        require(val <= address(this).balance, "insufficient funds");

        for (uint256 idx = 0; idx < kol.list.length; idx++) {
            privateWithdrawFromOneKol(kol.list[idx], msg.sender, false);
        }

        uint256 reminders = minusWithdrawFee(val);

        payable(msg.sender).transfer(reminders);
        emit InvestorWithdrawAllIncome(msg.sender, reminders);
    }

    /********************************************************************************
     *                       income query
     *********************************************************************************/

    function IncomeOfOneNonce(
        address kol,
        uint256 nonce,
        uint256 amount
    ) public view returns (uint256) {
        KeySettings memory status = KeySettingsRecord[kol];
        uint256 sumForThisNonce = 0;
        for (
            uint256 startNonce = nonce;
            startNonce <= status.nonce;
            startNonce++
        ) {
            uint256 valPerKey = incomePerNoncePerKey[kol][startNonce];
            sumForThisNonce += valPerKey * amount;
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
        for (uint256 idx = 0; idx < record.nonceList.length; idx++) {
            uint256 nonce = record.nonceList[idx];
            keymeta memory km = record.nonceToAmount[nonce];
            if (km.amount == 0) {
                continue;
            }
            uint256 oneIncome = IncomeOfOneNonce(kol, nonce, km.amount);
            sumIncome += oneIncome;
        }
        return sumIncome;
    }

    function AllIncomeOfAllKol(address investor) public view returns (uint256) {
        uint256 sumIncome = 0;
        MapArray storage record = kolsOfInKeyHolder[investor];
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
        if (record.nonceList.length == 0) {
            return (new uint256[](0), new uint256[](0));
        }

        nonce = new uint256[](record.nonceList.length);
        amounts = new uint256[](record.nonceList.length);
        for (uint256 idx = 0; idx < record.nonceList.length; idx++) {
            uint256 n = record.nonceList[idx];
            keymeta memory km = record.nonceToAmount[n];
            nonce[idx] = n;
            amounts[idx] = km.amount;
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
        MapArray storage kol = kolsOfInKeyHolder[investor];
        return kol.list;
    }

    function KeyStatusOfKol(address kol)
    public
    view
    returns (KeySettings memory)
    {
        return KeySettingsRecord[kol];
    }

    function AuctionData(uint256 start, uint256 no)
    public
    view
    returns (KeyAuction[] memory result)
    {
        result = new KeyAuction[](no);
        uint256 end = start + no;
        if (end >= keyInAuction.length) {
            end = keyInAuction.length - 1;
        }
        for (uint256 idx = start; idx < end; idx++) {
            result[0] = keyInAuction[idx];
        }
    }
}
