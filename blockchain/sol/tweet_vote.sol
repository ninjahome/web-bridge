// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

import "./common.sol";

/********************************************************************************
 *                       admin logic
 *********************************************************************************/
contract TweetVoteAmin is ServiceFeeForWithdraw {
    uint256 public constant oneFinney = 1e6 gwei;

    uint256 public tweetPostPrice = 0.005 ether;
    uint256 public tweetVotePrice = 0.005 ether;

    uint256 public maxVotePerTweet = 1e8;

    uint8 public kolIncomePerTweetVoteRate = 30;
    uint8 public serviceFeePerTweetVoteRate = 10;

    address public pluginAddress;
    bool public pluginStop = true;

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

    function adminSetKolIncomePerTweetRate(uint8 newRate) public isOwner {
        require(
            newRate + serviceFeePerTweetVoteRate <= 100,
            "rate is more than 100"
        );
        kolIncomePerTweetVoteRate = newRate;
        emit SystemRateChanged(newRate, "kol_income_per_tweet_vote_rate");
    }

    function adminSetServiceFeeRateForPerTweetVote(uint8 newRate)
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

    function adminSetMaxVotePerTweet(uint256 newMaxVote) public isOwner {
        require(newMaxVote >= 1, "invalid max vote no");
        maxVotePerTweet = newMaxVote;
        emit SystemRateChanged(newMaxVote, "max_vote_number_once");
    }

    function adminSetPluginAddr(address addr) public isOwner {
        require(
            TweetVotePlugInI(addr).checkPluginInterface(),
            "invalid plugin address"
        );
        pluginAddress = addr;
        pluginStop = false;
        emit PluginChanged(pluginAddress, pluginStop);
    }

    function adminStopPlugin(bool stop) public isOwner {
        pluginStop = stop;
        emit PluginChanged(pluginAddress, pluginStop);
    }
}

/********************************************************************************
 *                       business logic
 *********************************************************************************/

contract TweetVote is TweetVoteAmin {
    mapping(bytes32 => address) public ownersOfAllTweets;
    event KolRightsBought(address kolAddr, address buyer, uint256 rightsNo);

    event TweetPublished(address indexed from, bytes32 tweetHash);
    event TweetVoted(
        bytes32 tweetHash,
        address voter,
        uint256 pricePerVote,
        uint256 voteNo
    );
    event KolWithdraw(address indexed kol, uint256 amount);

    constructor() payable {}

    function publishTweet(bytes32 hash, bytes memory signature)
    public
    payable
    inRun
    {
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

    function voteToTweets(bytes32 tweetHash, uint256 voteNo)
    public
    payable
    noReentrant
    inRun
    {
        require(voteNo > 0 && voteNo < maxVotePerTweet, "vote no. invalid");
        uint256 amount = voteNo * tweetVotePrice;
        require(amount > __minValCheck, "amount invalid");
        require(msg.value == amount, "vote price has changed");

        address tweetOwner = ownersOfAllTweets[tweetHash];
        require(tweetOwner != address(0), "no such tweet");

        uint256 forKolSum = (amount / 100) * kolIncomePerTweetVoteRate;
        balance[tweetOwner] += forKolSum;

        uint256 serviceFee = (amount / 100) * serviceFeePerTweetVoteRate;
        recordServiceFee(serviceFee);

        uint256 leftVal = amount - forKolSum - serviceFee;

        if (
            pluginAddress != address(0) &&
            pluginStop == false &&
            leftVal > __minValCheck
        ) {
            TweetVotePlugInI(pluginAddress).tweetBought{value: leftVal}(
                tweetHash,
                tweetOwner,
                msg.sender,
                voteNo
            );
        }

        emit TweetVoted(tweetHash, msg.sender, tweetVotePrice, voteNo);
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
