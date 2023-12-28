// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./common.sol";

contract TweetExchangeAmin is Owner {
    uint256 public constant oneFinney = 1e6 gwei;
    uint256 public tweetPostPrice = 0.005 ether;
    uint256 public tweetVotePrice = 0.005 ether;

    uint256 public feeReceived;
    uint256 public maxVotePerTweet = 1e8;
    uint256 public kolIncomePerTweetVoteRate = 30;
    uint256 public serviceFeePerTweetVoteRate = 10;
    uint256 public serviceFeeForWithdrawRate = 2;
    address public pluginAddress;
    bool public pluginStop = false;

    event UpgradeToNewRule(address indexed recipient, uint256 amount);
    event Received(address indexed sender, uint256 amount);
    event PluginChanged(address pAddr, bool stop);
    event SystemSettingChanged(
        uint256 pricePost,
        uint256 priceVot,
        uint256 kolRate,
        uint256 feeRate,
        uint256 withDrawRate
    );

    constructor() payable {}

    function upgradeToNewRule(address payable recipient)
    public
    isOwner
    isValidAddress(recipient)
    {
        payable(this.getOwner()).transfer(feeReceived);
        uint256 balance = address(this).balance;
        if (balance > 0) {
            recipient.transfer(balance);
        }
        emit UpgradeToNewRule(recipient, balance);
    }

    receive() external payable {
        emit Received(msg.sender, msg.value);
    }

    function exchangeBalance() public view returns (uint256) {
        return address(this).balance;
    }

    function changeTweetPostPrice(uint256 newPriceInFinney) public isOwner {
        tweetPostPrice = newPriceInFinney * oneFinney;
        emit SystemSettingChanged(
            tweetPostPrice,
            tweetVotePrice,
            kolIncomePerTweetVoteRate,
            serviceFeePerTweetVoteRate,
            serviceFeeForWithdrawRate
        );
    }

    function changeTweetVotePrice(uint256 newPriceInFinney) public isOwner {
        tweetVotePrice = newPriceInFinney * oneFinney;
        emit SystemSettingChanged(
            tweetPostPrice,
            tweetVotePrice,
            kolIncomePerTweetVoteRate,
            serviceFeePerTweetVoteRate,
            serviceFeeForWithdrawRate
        );
    }

    function setKolIncomePerTweetRate(uint256 newRate) public isOwner {
        require(
            newRate + serviceFeePerTweetVoteRate < 100,
            "rate is more than 100"
        );
        kolIncomePerTweetVoteRate = newRate;
        emit SystemSettingChanged(
            tweetPostPrice,
            tweetVotePrice,
            kolIncomePerTweetVoteRate,
            serviceFeePerTweetVoteRate,
            serviceFeeForWithdrawRate
        );
    }

    function setServiceFeeRateForPerTweetVote(uint256 newRate) public isOwner {
        require(
            newRate + kolIncomePerTweetVoteRate < 100,
            "rate is more than 100"
        );
        serviceFeePerTweetVoteRate = newRate;
        emit SystemSettingChanged(
            tweetPostPrice,
            tweetVotePrice,
            kolIncomePerTweetVoteRate,
            serviceFeePerTweetVoteRate,
            serviceFeeForWithdrawRate
        );
    }

    function setServiceFeeRateForWithdraw(uint256 newRate) public isOwner {
        require(newRate > 0 && newRate < 100, "rate invalid");
        serviceFeeForWithdrawRate = newRate;
        emit SystemSettingChanged(
            tweetPostPrice,
            tweetVotePrice,
            kolIncomePerTweetVoteRate,
            serviceFeePerTweetVoteRate,
            serviceFeeForWithdrawRate
        );
    }

    function setMaxVotePerTweet(uint256 newMaxVote) public isOwner {
        require(newMaxVote > 1, "invalid max vote no");
        maxVotePerTweet = newMaxVote;
    }

    function setPluginAddr(address addr) public isOwner {
        require(PlugInI(addr).checkPluginInterface(), "invalid plugin address");
        pluginAddress = addr;
        emit PluginChanged(pluginAddress, pluginStop);
    }

    function stopPlugin(bool stop) public isOwner {
        pluginStop = stop;
        emit PluginChanged(pluginAddress, pluginStop);
    }
}

contract TweetExchange is TweetExchangeAmin {
    struct TweetInfo {
        uint256 value;
        address owner;
        uint256 votes;
    }

    mapping(bytes32 => TweetInfo) public tweetsInfo;
    mapping(address => uint256) public kolTweetVoteIncome;

    event TweetPublished(address indexed from, bytes32 tweetHash);
    event TweetRightsBought(address indexed from, uint256 price, uint256 no);
    event KolWithDraw(address indexed kol, uint256 amount);
    event ThirdPartyClaims(address indexed addr, uint256 amount);

    constructor() payable {}

    function publishTweet(bytes32 hash, bytes memory signature) public payable {
        require(msg.value == tweetPostPrice, "tweet post fee cahnged");
        require(tweetsInfo[hash].owner == address(0), "duplicate post");
        require(
            recoverSigner(hash, signature) == msg.sender,
            "Invalid signature"
        );

        tweetsInfo[hash] = TweetInfo(0, msg.sender, 0);

        feeReceived += tweetPostPrice;

        emit TweetPublished(msg.sender, hash);
    }

    function buyTweetRights(bytes32 tweetHash, uint256 voteNo) public payable {
        require(voteNo > 0 && voteNo < maxVotePerTweet, "vote no. invalid");
        uint256 amount = voteNo * tweetVotePrice;
        require(amount > 1 gwei, "amount invalid");
        require(msg.value == amount, "insufficient funds");

        TweetInfo storage tweet = tweetsInfo[tweetHash];
        require(tweet.owner != address(0), "no such tweet");

        tweet.votes += voteNo;
        tweet.value += amount;

        uint256 forKolSum = (amount / 100) * kolIncomePerTweetVoteRate;
        kolTweetVoteIncome[tweet.owner] += forKolSum;

        uint256 serviceFee = (amount / 100) * serviceFeePerTweetVoteRate;
        feeReceived += serviceFee;
        if (pluginAddress != address(0)) {
            uint256 leftVal = amount - forKolSum - serviceFee;
            PlugInI(pluginAddress).tweetBought(
                tweetHash,
                tweet.owner,
                msg.sender,
                leftVal,
                voteNo
            );
        }

        emit TweetRightsBought(msg.sender, tweetVotePrice, voteNo);
    }

    function kolWithdrawTweetIncomes(uint256 amount, bool all)
    public
    noReentrant
    {
        require(kolTweetVoteIncome[msg.sender] >= amount, "insufficient funds");
        require(amount >= 1 gwei || all, "invalid param");
        if (all) {
            amount = kolTweetVoteIncome[msg.sender];
            kolTweetVoteIncome[msg.sender] = 0;
        } else {
            kolTweetVoteIncome[msg.sender] -= amount;
        }

        uint256 serviceFee = (amount / 100) * serviceFeeForWithdrawRate;
        feeReceived += serviceFee;
        amount -= serviceFee;

        payable(msg.sender).transfer(amount);
        emit KolWithDraw(msg.sender, amount);
    }

    function recoverSigner(bytes32 prefixedHash, bytes memory signature)
    public
    pure
    returns (address)
    {
        require(signature.length == 65, "Invalid signature length");

        bytes32 r;
        bytes32 s;
        uint8 v;

        assembly {
            r := mload(add(signature, 32))
            s := mload(add(signature, 64))
            v := byte(0, mload(add(signature, 96)))
        }

        if (v < 27) {
            v += 27;
        }

        return ecrecover(prefixedHash, v, r, s);
    }
}
