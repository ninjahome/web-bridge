// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./common.sol";

contract TweetExchangeAmin is Owner {
    uint256 public tweetPostPrice = 0.005 ether;
    uint256 public tweetBuyOneRightsPrice = 0.005 ether;

    uint256 public feeReceived;
    uint256 public maxVotePerTweet = 1e8;
    uint256 public kolIncomePerTweetVoteRate = 30;
    uint256 public serviceFeePerTweetVoteRate = 5;
    address public thirdEntityAddress;
    uint256 public thirdPartyBalance;
    bool public thirdPartyStopped;

    event SystemWithDrawIncomes(address indexed recipient, uint256 amount);
    event Received(address indexed sender, uint256 amount);

    modifier isValidAddress(address addr) {
        require(addr != address(0), "invalid addrrss");
        _;
    }

    function systemDrawIncomes(address payable recipient)
    public
    isOwner
    isValidAddress(recipient)
    {
        uint256 balance = address(this).balance;
        recipient.transfer(balance);
        emit SystemWithDrawIncomes(recipient, balance);
    }

    receive() external payable {
        feeReceived += msg.value;
        emit Received(msg.sender, msg.value);
    }

    function changeTweetPostPrice(uint256 newPrice) public isOwner {
        tweetPostPrice = newPrice;
    }

    function changeTweetOneRightPrice(uint256 newPrice) public isOwner {
        tweetBuyOneRightsPrice = newPrice;
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

    function changeThirdPartyEntityAddress(address entity) public isOwner {
        require(address(0) != entity);
        thirdEntityAddress = entity;
    }

    function setThirdPartyStatus(bool stopped) public isOwner {
        thirdPartyStopped = stopped;
    }
}

contract TweetExchange is TweetExchangeAmin, ExchangeI {
    struct TweetInfo {
        uint256 value;
        address owner;
        uint256 votes;
    }

    struct LotteryTicket {
        uint256 tid;
    }

    mapping(bytes32 => TweetInfo) public tweetsInfo;
    mapping(address => bytes32[]) public tweetsOfCreator;
    mapping(address => uint256) public kolTweetVoteIncome;

    event TweetPublished(address indexed from, bytes32 tweetHash);
    event TweetRightsBought(address indexed from, uint256 price, uint256 no);
    event KolWithDraw(address indexed kol, uint256 amount);
    event GameClaim(address indexed game, uint256 amount);

    constructor(address entity) {
        thirdEntityAddress = entity;
        thirdPartyStopped = false;
    }

    function publishTweet(bytes32 hash, bytes memory signature) public payable {
        require(msg.value == tweetPostPrice, "Must send 0.01 ETH");

        require(
            recoverSigner(hash, signature) == msg.sender,
            "Invalid signature"
        );

        tweetsInfo[hash] = TweetInfo(0, msg.sender, 0);
        tweetsOfCreator[msg.sender].push(hash);
        feeReceived += tweetPostPrice;

        emit TweetPublished(msg.sender, hash);
    }

    function buyTweetRights(
        bytes32 tweetHash,
        uint256 price,
        uint256 voteNo
    ) public payable {
        require(price == tweetBuyOneRightsPrice, "price has chaged");
        require(voteNo > 0 && voteNo < maxVotePerTweet, "rights no invalid");

        uint256 amount = voteNo * price;
        require(msg.value == amount, "insuficient funds");

        TweetInfo storage tweet = tweetsInfo[tweetHash];
        require(tweet.owner != address(0), "no such tweet");

        tweet.votes += voteNo;
        tweet.value += amount;

        uint256 forKolSum = (amount / 100) * kolIncomePerTweetVoteRate;
        kolTweetVoteIncome[tweet.owner] += forKolSum;

        uint256 serficeFee = (amount / 100) * serviceFeePerTweetVoteRate;
        feeReceived += serficeFee;

        if (thirdEntityAddress != address(0) && thirdPartyStopped == false) {
            uint256 thirdPatyFee = amount - forKolSum - serficeFee;
            thirdPartyBalance += thirdPatyFee;
            ThirdPartEntityI(thirdEntityAddress).rightsBought(
                msg.sender,
                price,
                voteNo
            );
        }

        emit TweetRightsBought(msg.sender, price, voteNo);
    }

    function kolWithDrawTweetIncomes(uint256 amount) public noReentrant {
        require(kolTweetVoteIncome[msg.sender] >= amount, "insufficient funds");
        payable(msg.sender).transfer(amount);
        emit KolWithDraw(msg.sender, amount);
    }

    function recoverSigner(bytes32 hash, bytes memory signature)
    internal
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

        bytes32 prefixedHash = keccak256(
            abi.encodePacked("\x19Ethereum Signed Message:\n32", hash)
        );

        return ecrecover(prefixedHash, v, r, s);
    }

    function claimIncomes() public noReentrant {
        require(thirdEntityAddress == msg.sender, "no rights");
        require(thirdPartyBalance >= 1 gwei, "no balance");
        payable(thirdEntityAddress).transfer(thirdPartyBalance);
        emit GameClaim(thirdEntityAddress, thirdPartyBalance);
    }
}
