class TweetContentToPost {
    constructor(tweet_content, createAt, web3Id, twitterID, tweet_id, signature) {
        this.text = tweet_content;
        this.create_time = createAt;
        this.web3_id = web3Id;
        this.twitter_id = twitterID;
        this.tweet_id = tweet_id;
        this.signature = signature;
    }
}

class MemCachedTweets {
    constructor() {
        this.latestID = 0;
        this.moreOldTweets = true;
        this.isLoading = false;
        this.CachedItem = [];
    }

    canLoadMoreOldData() {
        return this.moreOldTweets && !this.isLoading
    }
}

const dbKeyCachedVoteContractMeta = "__db_key_cached_vote_contract_meta__"

class TweetVoteContractSetting {
    constructor(postPrice, votePrice, votePriceInEth, maxVote, pluginAddr, pluginStop, kolRate, feeRate) {
        this.postPrice = postPrice;
        this.votePrice = votePrice;
        this.votePriceInEth = votePriceInEth;
        this.maxVote = maxVote;
        this.pluginAddr = pluginAddr;
        this.pluginStop = pluginStop;
        this.kolRate = kolRate;
        this.feeRate = feeRate;
    }

    static sycToDb(obj) {
        localStorage.setItem(TweetVoteContractSetting.DBKey(), JSON.stringify(obj));
    }

    static DBKey() {
        return dbKeyCachedVoteContractMeta;
    }

    static load() {
        const storedVal = localStorage.getItem(TweetVoteContractSetting.DBKey());
        return storedVal ? JSON.parse(storedVal) : null;
    }
}

class GameBasicInfo {
    constructor(curRound, totalBonus, ticketNo, curBonus, dTime) {
        this.curRound = curRound;
        this.totalBonus = totalBonus;
        this.ticketNo = ticketNo;
        this.curBonus = curBonus;
        this.dTime = dTime;
    }
}

const TXStatus = Object.freeze({
    NoPay: 0, Pending: 1, Success: 2, Failed: 3, Str(val) {
        switch (val) {
            case this.NoPay:
                return "not paid";
            case this.Pending:
                return "pending";
            case this.Success:
                return "success";
            case this.Failed:
                return "failed";
            default:
                return "unknown";
        }
    }
});

const TweetDetailSource = Object.freeze({
    NoNeed:'0', HomePage: '1',  MyVoted: '2', MostVoted: '3',
    MostTeam: '4',
});

class TweetQueryParam {
    constructor(startID, web3ID, voted, hashList) {
        this.start_id = startID;
        this.web3_id = web3ID;
        this.voted_ids = voted;
        this.hash_arr = hashList;
    }
}

class TeamDetailOnBlockChain {
    constructor(tweetHash, memCount, voteCount) {
        this.tweetHash = tweetHash;
        this.memCount = memCount;
        this.voteCount = voteCount;
    }
}