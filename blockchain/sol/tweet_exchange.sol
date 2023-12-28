// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./common.sol";

contract TweetExchangeAmin is Owner {
    uint256 public tweetPostPrice = 0.005 ether;
    uint256 public tweetVotePrice = 0.005 ether;

    uint256 public feeReceived;
    uint256 public maxVotePerTweet = 1e8;
    uint256 public kolIncomePerTweetVoteRate = 30;
    uint256 public serviceFeePerTweetVoteRate = 5;

    event UpgradeToNewRule(address indexed recipient, uint256 amount);
    event Received(address indexed sender, uint256 amount);

    function upgradeToNewRule(address payable recipient)
    public
    isOwner
    isValidAddress(recipient)
    {
        payable(this.getOwner()).transfer(feeReceived);
        uint256 balance = address(this).balance;
        recipient.transfer(balance);
        emit UpgradeToNewRule(recipient, balance);
    }

    receive() external payable {
        emit Received(msg.sender, msg.value);
    }

    function changeTweetPostPrice(uint256 newPrice) public isOwner {
        tweetPostPrice = newPrice;
    }

    function changeTweetVotePrice(uint256 newPrice) public isOwner {
        tweetVotePrice = newPrice;
    }

    function setKolIncomePerTweetRate(uint256 newRate) public isOwner {
        require(
            newRate + serviceFeePerTweetVoteRate < 100,
            "rate is more than 100"
        );
        kolIncomePerTweetVoteRate = newRate;
    }

    function setServiceFeeRateForPerTweetVote(uint256 newRate) public isOwner {
        require(
            newRate + kolIncomePerTweetVoteRate < 100,
            "rate is more than 100"
        );
        serviceFeePerTweetVoteRate = newRate;
    }

    function setMaxVotePerTweet(uint256 newMaxVote) public isOwner {
        maxVotePerTweet = newMaxVote;
    }
}

abstract contract TweetExchange is TweetExchangeAmin {
    struct TweetInfo {
        uint256 value;
        address owner;
        uint256 votes;
    }

    struct LotteryTicket {
        uint256 tid;
    }

    mapping(bytes32 => TweetInfo) public tweetsInfo;
    mapping(address => uint256) public kolTweetVoteIncome;

    event TweetPublished(address indexed from, bytes32 tweetHash);
    event TweetRightsBought(address indexed from, uint256 price, uint256 no);
    event KolWithDraw(address indexed kol, uint256 amount);
    event ThirdPartyClaims(address indexed addr, uint256 amount);

    function publishTweet(bytes32 hash, bytes memory signature) public payable {
        require(msg.value == tweetPostPrice, "Must send post fee");

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
        require(msg.value == amount, "insuficient funds");

        TweetInfo storage tweet = tweetsInfo[tweetHash];
        require(tweet.owner != address(0), "no such tweet");

        tweet.votes += voteNo;
        tweet.value += amount;

        uint256 forKolSum = (amount / 100) * kolIncomePerTweetVoteRate;
        kolTweetVoteIncome[tweet.owner] += forKolSum;

        uint256 serficeFee = (amount / 100) * serviceFeePerTweetVoteRate;
        feeReceived += serficeFee;
        uint256 leftVal = amount - forKolSum - serficeFee;
        tweetBought(msg.sender, leftVal,voteNo);
        emit TweetRightsBought(msg.sender, tweetVotePrice, voteNo);
    }

    function kolWithDrawTweetIncomes(uint256 amount, bool all)
    public
    noReentrant
    {
        require(kolTweetVoteIncome[msg.sender] >= amount, "insufficient funds");
        require(amount >= 1 gwei || all, "invalid param");
        kolTweetVoteIncome[msg.sender] -= amount;
        payable(msg.sender).transfer(amount);
        emit KolWithDraw(msg.sender, amount);
    }

    function tweetBought(
        address buyer,
        uint256 leftVal,
        uint256 voteNo
    ) internal virtual;
}
