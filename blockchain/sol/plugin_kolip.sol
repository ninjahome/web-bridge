// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./common.sol";
import "./tweet_exchange.sol";

contract KolIPGame is ServiceFeeForWithdraw, PlugInI {
    struct KolKeyStatus {
        uint256 nonce;
        uint256 totalNo;
    }

    struct BuyerKeyRecord {
        mapping(uint256 => uint256) nonceToAmount;
        uint256[] nonceList;
    }

    struct MapArray {
        mapping(address => bool) filter;
        address[] list;
    }

    mapping(address => MapArray) private __investorsOfSomeKol;
    mapping(address => MapArray) private __kolOfOneInvestor;
    mapping(address => KolKeyStatus) private __kolKeyStatusRecord;
    mapping(address => mapping(address => BuyerKeyRecord))
    private __buyersKeyRecordOfKol;
    mapping(address => mapping(uint256 => uint256)) public incomePerNoncePerKey;

    event InvestorWithdrawByOneNonce(
        address investor,
        address kol,
        uint256 nonce,
        uint256 val
    );
    event InvestorWithdrawByOneKol(address investor, address kol, uint256 val);
    event InvestorWithdrawAllIncome(address investor, uint256 val);

    /********************************************************************************
     *                       record operation
     *********************************************************************************/
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
    {
        uint256 val = msg.value;
        require(val > __minValCheck, "invalid msg value");
        require(voteNo >= 1, "invalid vote no");
        require(tweetHash != bytes32(0));
        KolKeyStatus memory recordOfKolKey = __kolKeyStatusRecord[tweetOwner];
        if (recordOfKolKey.totalNo == 0) {
            return;
        }
        uint256 valPerKey = val / recordOfKolKey.totalNo;
        incomePerNoncePerKey[tweetOwner][recordOfKolKey.nonce] += valPerKey;
    }

    function checkPluginInterface() external pure returns (bool) {
        return true;
    }

    function KolIPRightsBought(
        address kolAddr,
        address buyer,
        uint256 keyNo
    )
    external
    payable
    isValidAddress(buyer)
    isValidAddress(kolAddr)
    onlyAdmin
    noReentrant
    {
        KolKeyStatus storage recordOfKolKey = __kolKeyStatusRecord[kolAddr];
        recordOfKolKey.nonce++;
        recordOfKolKey.totalNo += keyNo;

        BuyerKeyRecord storage record = __buyersKeyRecordOfKol[buyer][kolAddr];
        record.nonceToAmount[recordOfKolKey.nonce] = keyNo;
        record.nonceList.push(recordOfKolKey.nonce);

        MapArray storage investors = __investorsOfSomeKol[kolAddr];
        if (investors.filter[buyer] == false) {
            investors.list.push(buyer);
        }

        MapArray storage kols = __kolOfOneInvestor[buyer];
        if (kols.filter[kolAddr] == false) {
            kols.list.push(kolAddr);
        }
    }

    /********************************************************************************
     *                       income withdraw
     *********************************************************************************/

    function withdrawFromOneNonce(address kol, uint256 nonce)
    public
    noReentrant
    isValidAddress(kol)
    {
        KolKeyStatus memory curStatus = __kolKeyStatusRecord[kol];
        BuyerKeyRecord storage record = __buyersKeyRecordOfKol[msg.sender][kol];

        uint256 amount = record.nonceToAmount[nonce];
        require(amount > 0, "no key no in this nonce");
        uint256 val = IncomeOfOneNonce(kol, nonce, amount);
        require(val <= address(this).balance, "insufficient funds");

        delete record.nonceToAmount[nonce];
        record.nonceToAmount[curStatus.nonce] += amount;

        payable(msg.sender).transfer(val);
        emit InvestorWithdrawByOneNonce(msg.sender, kol, nonce, val);
    }

    function withdrawFromOneKol(address kol)
    public
    noReentrant
    isValidAddress(kol)
    {
        KolKeyStatus memory curStatus = __kolKeyStatusRecord[kol];
        BuyerKeyRecord storage record = __buyersKeyRecordOfKol[msg.sender][kol];
        uint256 val = AllIncomeOfOneKol(kol, msg.sender);
        require(val <= address(this).balance, "insufficient funds");

        for (uint256 idx = 0; idx < record.nonceList.length; idx++) {
            uint256 nonce = record.nonceList[idx];
            uint256 amount = record.nonceToAmount[nonce];
            delete record.nonceToAmount[nonce];
            record.nonceToAmount[curStatus.nonce] += amount;
        }
        delete record.nonceList;
        record.nonceList.push(curStatus.nonce);
        payable(msg.sender).transfer(val);
        emit InvestorWithdrawByOneKol(msg.sender, kol, val);
    }

    function withDrawAllIncome() public noReentrant {
        MapArray storage kol = __kolOfOneInvestor[msg.sender];
        require(kol.list.length > 0, "no investment");
        uint256 val = AllIncomeOfAllKol(msg.sender);
        require(val <= address(this).balance, "insufficient funds");

        for (uint256 idx = 0; idx < kol.list.length; idx++) {
            withdrawFromOneKol(kol.list[idx]);
        }

        payable(msg.sender).transfer(val);
        emit InvestorWithdrawAllIncome(msg.sender, val);
    }

    /********************************************************************************
     *                       income query
     *********************************************************************************/

    function IncomeOfOneNonce(
        address kol,
        uint256 nonce,
        uint256 amount
    ) public view returns (uint256) {
        KolKeyStatus memory status = __kolKeyStatusRecord[kol];
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
        BuyerKeyRecord storage record = __buyersKeyRecordOfKol[investor][kol];
        uint256 sumIncome = 0;
        for (uint256 idx = 0; idx < record.nonceList.length; idx++) {
            uint256 nonce = record.nonceList[idx];
            uint256 amount = record.nonceToAmount[nonce];
            if (amount == 0) {
                continue;
            }
            uint256 oneIncome = IncomeOfOneNonce(kol, nonce, amount);
            sumIncome += oneIncome;
        }
        return sumIncome;
    }

    function AllIncomeOfAllKol(address investor) public view returns (uint256) {
        uint256 sumIncome = 0;
        MapArray storage record = __kolOfOneInvestor[investor];
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
        BuyerKeyRecord storage record = __buyersKeyRecordOfKol[investor][kol];
        if (record.nonceList.length == 0) {
            return (new uint256[](0), new uint256[](0));
        }

        nonce = new uint256[](record.nonceList.length);
        amounts = new uint256[](record.nonceList.length);
        for (uint256 idx = 0; idx < record.nonceList.length; idx++) {
            uint256 n = record.nonceList[idx];
            uint256 a = record.nonceToAmount[n];
            nonce[idx] = n;
            amounts[idx] = a;
        }
        return (nonce, amounts);
    }

    function InvestorOfKol(address kol) public view returns (address[] memory) {
        MapArray storage investors = __investorsOfSomeKol[kol];
        return investors.list;
    }

    function KolOfOneInvestor(address investor)
    public
    view
    returns (address[] memory)
    {
        MapArray storage kol = __kolOfOneInvestor[investor];
        return kol.list;
    }

    function KeyStatusOfKol(address kol)
    public
    view
    returns (KolKeyStatus memory)
    {
        return __kolKeyStatusRecord[kol];
    }
}
