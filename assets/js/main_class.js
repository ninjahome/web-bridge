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

class MemCachedTweets{
    constructor() {
        this.MaxID = BigInt(0);
        this.MinID = BigInt(0);
        this.TweetMaps = new Map();
        this.moreOldTweets = true;
        this.isLoading = false;
        this.CachedItem = [];
    }

    canLoadMoreOldData(){
        return this.moreOldTweets && !this.isLoading
    }
}

class TwitterBasicInfo {
    constructor(id, name, username, avatarUrl, bio) {
        this.id = id;
        this.name = name;
        this.username = username;
        this.profile_image_url = avatarUrl;
        this.description = bio;
    }

    static loadTwBasicInfo(TwitterID) {
        const storedData = getDataFromSessionDB(sesDbKeyForTwitterUserData(TwitterID))
        if (!storedData) {
            return null
        }
        return new TwitterBasicInfo(storedData.id, storedData.name, storedData.username,
            storedData.profile_image_url, storedData.description);
    }

    static cacheTwBasicInfo(objStr) {
        const obj = JSON.parse(objStr)
        if (!obj.id) {
            throw new Error("invalid twitter basic info")
        }
        sessionStorage.setItem(sesDbKeyForTwitterUserData(obj.id), objStr);
        return obj;
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

class GameContractMeta {
    constructor(curRound, totalBonus,ticketNo, ticketPrice, ticketPriceInEth) {
        this.curRound = curRound;
        this.totalBonus = totalBonus;
        this.ticketNo = ticketNo;
        this.ticketPrice = ticketPrice;
        this.ticketPriceInEth = ticketPriceInEth;
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

class TweetQueryParam{
    constructor(startID, needNewest,web3ID,voted) {
        this.start_id = startID;
        this.web3_id = web3ID;
        this.newest = needNewest;
        this.voted = voted;
    }
}