// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./common.sol";

contract TweetExchangeAmin is ServiceFeeForWithdraw {
    uint256 public constant oneFinney = 1e6 gwei;

    uint256 public tweetPostPrice = 0.005 ether;
    uint256 public tweetVotePrice = 0.005 ether;

    uint256 public maxVotePerTweet = 1e8;

    uint256 public kolIncomePerTweetVoteRate = 30;
    uint256 public serviceFeePerTweetVoteRate = 10;

    uint256 public kolIncomePerIPRightBuyRate = 90;
    uint256 public serviceFeePerKolIpRightRate = 10;

    address public pluginAddress;
    bool public pluginStop = false;

    event Received(address indexed sender, uint256 amount);
    event PluginChanged(address pAddr, bool stop);

    event SystemRateChanged(uint256 pricePost, string rateName);

    constructor() payable {}

    receive() external payable {
        emit Received(msg.sender, msg.value);
    }

    function exchangeBalance() public view returns (uint256) {
        return address(this).balance;
    }

    function adminSetTweetPostPrice(uint256 newPriceInFinney) public isOwner {
        tweetPostPrice = newPriceInFinney * oneFinney;
        emit SystemRateChanged(tweetPostPrice, "tweet_post_price");
    }

    function adminSetTweetVotePrice(uint256 newPriceInFinney) public isOwner {
        tweetVotePrice = newPriceInFinney * oneFinney;
        emit SystemRateChanged(tweetVotePrice, "tweet_vote_price");
    }

    function adminSetKolIncomePerTweetRate(uint256 newRate) public isOwner {
        require(
            newRate + serviceFeePerTweetVoteRate <= 100,
            "rate is more than 100"
        );
        kolIncomePerTweetVoteRate = newRate;
        emit SystemRateChanged(newRate, "kol_income_per_tweet_vote_rate");
    }

    function adminSetServiceFeeRateForPerTweetVote(uint256 newRate)
    public
    isOwner
    {
        require(
            newRate + kolIncomePerTweetVoteRate <= 100,
            "rate is more than 100"
        );
        serviceFeePerTweetVoteRate = newRate;
        emit SystemRateChanged(newRate, "service_fee_per_tweet_vote_rate");
    }

    function adminSetKolIncomeRatePerIpRight(uint256 newRate) public isOwner {
        require(
            newRate + serviceFeePerKolIpRightRate <= 100,
            "rate is more than 100"
        );
        kolIncomePerIPRightBuyRate = newRate;
        emit SystemRateChanged(newRate, "kol_income_per_kol_ip_right_rate");
    }

    function adminSetServiceFeeRatePerKolIPRight(uint256 newRate)
    public
    isOwner
    {
        require(
            newRate + kolIncomePerIPRightBuyRate <= 100,
            "rate is more than 100"
        );
        serviceFeePerKolIpRightRate = newRate;
        emit SystemRateChanged(newRate, "service_fee_per_kol_ip_right_rate");
    }

    function adminSetMaxVotePerTweet(uint256 newMaxVote) public isOwner {
        require(newMaxVote >= 1, "invalid max vote no");
        maxVotePerTweet = newMaxVote;
        emit SystemRateChanged(newMaxVote, "max_vote_number_once");
    }

    function adminSetPluginAddr(address addr) public isOwner {
        require(PlugInI(addr).checkPluginInterface(), "invalid plugin address");
        pluginAddress = addr;
        emit PluginChanged(pluginAddress, pluginStop);
    }

    function adminStopPlugin(bool stop) public isOwner {
        pluginStop = stop;
        emit PluginChanged(pluginAddress, pluginStop);
    }
}

contract TweetExchange is TweetExchangeAmin {
    mapping(bytes32 => address) public ownersOfAllTweets;
    mapping(address => uint256) public kolTweetBalance;
    mapping(address => uint256) public kolIpRightPrice;

    event KolIpRightOpen(address kol, uint256 price, bool isUpdateOp);
    event KolRightsBought(
        address kolAddr,
        address buyer,
        uint256 rightsNo,
        uint256 pricePerRight
    );

    event TweetPublished(address indexed from, bytes32 tweetHash);
    event TweetRightsBought(
        bytes32 tweetHash,
        address indexed from,
        uint256 value,
        uint256 voteNo
    );
    event KolWithdraw(address indexed kol, uint256 amount);
    event ThirdPartyClaims(address indexed addr, uint256 amount);

    constructor() payable {}

    function publishTweet(bytes32 hash, bytes memory signature) public payable {
        require(msg.value == tweetPostPrice, "tweet post fee cahnged");
        require(ownersOfAllTweets[hash] == address(0), "duplicate post");
        require(
            recoverSigner(hash, signature) == msg.sender,
            "Invalid signature"
        );

        ownersOfAllTweets[hash] = msg.sender;

        recordServiceFee(tweetPostPrice);

        emit TweetPublished(msg.sender, hash);
    }

    function buyTweetRights(bytes32 tweetHash, uint256 voteNo)
    public
    payable
    noReentrant
    {
        require(voteNo > 0 && voteNo < maxVotePerTweet, "vote no. invalid");
        uint256 amount = voteNo * tweetVotePrice;
        require(amount > __minValCheck, "amount invalid");
        require(msg.value == amount, "insufficient funds");
        address tweetOwner = ownersOfAllTweets[tweetHash];
        require(tweetOwner != address(0), "no such tweet");

        uint256 forKolSum = (amount / 100) * kolIncomePerTweetVoteRate;
        kolTweetBalance[tweetOwner] += forKolSum;

        uint256 serviceFee = (amount / 100) * serviceFeePerTweetVoteRate;
        recordServiceFee(serviceFee);

        uint256 leftVal = amount - forKolSum - serviceFee;

        if (
            pluginAddress != address(0) &&
            pluginStop == false &&
            leftVal > __minValCheck
        ) {
            PlugInI(pluginAddress).tweetBought{value: leftVal}(
                tweetHash,
                tweetOwner,
                msg.sender,
                voteNo
            );
        }

        emit TweetRightsBought(tweetHash, msg.sender, tweetVotePrice, voteNo);
    }

    function buyKolIpRights(address kolAddr, uint256 rightsNo)
    public
    payable
    noReentrant
    {
        require(kolIpRightPrice[kolAddr] > 0, "kol ip right not open");

        uint256 amount = kolIpRightPrice[kolAddr] * rightsNo;
        require(msg.value >= amount, "insufficient funds");

        uint256 kolIncome = (amount / 100) * kolIncomePerIPRightBuyRate;
        kolTweetBalance[kolAddr] += kolIncome;

        uint256 serviceFee = (amount / 100) * serviceFeePerKolIpRightRate;
        recordServiceFee(serviceFee);

        uint256 leftVal = amount - kolIncome - serviceFee;

        if (
            pluginAddress != address(0) &&
            pluginStop == false &&
            leftVal > __minValCheck
        ) {
            PlugInI(pluginAddress).KolIPRightsBought{value: leftVal}(
                kolAddr,
                msg.sender,
                rightsNo
            );
        }

        emit KolRightsBought(
            kolAddr,
            msg.sender,
            rightsNo,
            kolIpRightPrice[kolAddr]
        );
    }

    function setupKolIpRights(uint256 pricePerRight, bool update) public {
        require(pricePerRight > __minValCheck, "invalid ip right price");

        if (update) {
            require(kolIpRightPrice[msg.sender] > 0, "not open yet");
            kolIpRightPrice[msg.sender] = pricePerRight;
        } else {
            require(kolIpRightPrice[msg.sender] == 0, "duplicate operation");
            kolIpRightPrice[msg.sender] = pricePerRight;
        }

        emit KolIpRightOpen(msg.sender, pricePerRight, update);
    }

    function kolWithdrawTweetIncomes(uint256 amount, bool all)
    public
    noReentrant
    {
        require(kolTweetBalance[msg.sender] >= amount, "insufficient funds");
        require(amount >= __minValCheck || all, "invalid param");
        if (all) {
            amount = kolTweetBalance[msg.sender];
            kolTweetBalance[msg.sender] = 0;
        } else {
            kolTweetBalance[msg.sender] -= amount;
        }

        uint256 reminders = minusWithDrawFee(amount);

        payable(msg.sender).transfer(reminders);
        emit KolWithdraw(msg.sender, reminders);
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
