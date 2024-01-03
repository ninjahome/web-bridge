const dbKeyCachedGlobalTweets = "__db_key_cached_global_tweets__"
const dbKeyCachedTweetContentById = "__db_key_cached_tweet_content__by_id__"
const maxCachedLocalTweetNo = 120
let isMoreTweetsLoading = false;
let hasMoreTweetsToLoad = true;
let maxTweetIdCurShowed = BigInt(0);
let minTweetIdCurShowed = BigInt(0);

const TXStatus = Object.freeze({
    NoPay: 0,
    Pending: 1,
    Success: 2,
    Failed: 3,
    Str(val){
        switch(val) {
            case this.NoPay: return "not paid";
            case this.Pending: return "pending";
            case this.Success: return "success";
            case this.Failed: return "failed";
            default: return "unknown";
        }
    }
});


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

class TweetToShowOnWeb {
    constructor(njTweet, njTwitter, blockChain) {

        this.text = njTweet.text;
        this.create_time = njTweet.create_time;
        this.web3_id = njTweet.web3_id;
        this.twitter_id = njTweet.twitter_id;
        this.tweet_id = njTweet.tweet_id;
        this.signature = njTweet.signature;
        this.tx_hash = njTweet.tx_hash;
        this.payment_status = njTweet.payment_status;

        if (njTwitter) {
            this.name = njTwitter.name;
            this.username = njTwitter.username;
            this.description = njTwitter.description;
            this.profile_image_url = njTwitter.profile_image_url;
            this.verified = njTwitter.verified;
        } else {
            this.name = null;
            this.profile_image_url = DefaultAvatarSrc;
        }
    }

    static DBKey(create_time) {
        return dbKeyCachedTweetContentById + create_time;
    }

    static syncToDb(obj){
        localStorage.setItem(TweetToShowOnWeb.DBKey(obj.create_time), JSON.stringify(obj));
    }
    static load(tweetID){
        const storedVal = localStorage.getItem(TweetToShowOnWeb.DBKey(tweetID));
        return storedVal?JSON.parse(storedVal):null;
    }
}

function showFullTweetContent() {
    const tweetCard = this.closest('.tweet-card');
    const tweetContent = tweetCard.querySelector('.tweet-content');
    if (this.innerText === "Show more") {
        tweetContent.style.display = 'block';
        tweetContent.classList.remove('tweet-content-collapsed');
        tweetCard.style.maxHeight = 'none';
        this.innerText = "Show less";
    } else {
        tweetContent.style.display = '-webkit-box';
        tweetContent.classList.add('tweet-content-collapsed');
        tweetCard.style.maxHeight = '400px';
        this.innerText = "Show more";
    }
}

function updateLocalCachedTweetList(newIDs) {
    const storedArray = JSON.parse(localStorage.getItem(dbKeyCachedGlobalTweets)) || [];
    const int64Array = storedArray.map(BigInt);

    const uniqueIDs = new Set(int64Array);

    newIDs.forEach(id => uniqueIDs.add(BigInt(id)));

    const mergedArray = Array.from(uniqueIDs);

    mergedArray.sort((a, b) => {
        if (a > b) return -1;
        if (a < b) return 1;
        return 0;
    });

    const topNArray = mergedArray.slice(0, maxCachedLocalTweetNo);
    localStorage.setItem(dbKeyCachedGlobalTweets, JSON.stringify(topNArray.map(Number)));
}

function parseNjTweetsFromSrv(tweetArray, refreshNewest) {
    if (tweetArray.length === 0) {
        if (!refreshNewest) {
            hasMoreTweetsToLoad = false;
        }
        return;
    }
    console.log(tweetArray)
    const newIDs = [];
    const localTweets = tweetArray.map(tweet => {
        let tw_data = TwitterBasicInfo.loadTwBasicInfo(tweet.twitter_id)
        let blockchain = BlockChainData.load(tw_data.web3_id);
        let obj = new TweetToShowOnWeb(tweet, tw_data, blockchain);
        TweetToShowOnWeb.syncToDb(obj);
        newIDs.push(obj.create_time);
        return obj;
    });
    populateLatestTweets(localTweets, refreshNewest);
    updateLocalCachedTweetList(newIDs)
}

function loadCachedGlobalTweets() {
    const storedData = localStorage.getItem(dbKeyCachedGlobalTweets)
    if (!storedData) {
        return;
    }
    let localTweetsIds = JSON.parse(storedData);
    if (localTweetsIds.length === 0) {
        return;
    }

    const tweets = localTweetsIds .map(tweetID => {
            return TweetToShowOnWeb.load(tweetID);
        }) .filter(tweet => tweet !== null);

    if (tweets.length === 0) {
        return;
    }

    populateLatestTweets(tweets, false);
}

function loadMoreTweets() {
    if (!hasMoreTweetsToLoad) {
        return;
    }
    if (isMoreTweetsLoading) {
        return;
    }

    const documentHeight = Math.max(document.body.scrollHeight, document.documentElement.scrollHeight);
    if (window.innerHeight + window.scrollY < documentHeight - 100) {
        return;
    }
    isMoreTweetsLoading = true;
    loadGlobalLatestTweetsFromSrv(false);
}

function loadGlobalLatestTweetsFromSrv(refreshNewest) {
    let startID;
    if (refreshNewest) {
        startID = maxTweetIdCurShowed
    } else {
        startID = minTweetIdCurShowed
    }

    fetch("/globalLatestTweets?startID=" + startID + "&&isRefresh=" + refreshNewest)
        .then(response => response.json())
        .then(tweetArray => {
            parseNjTweetsFromSrv(tweetArray, refreshNewest);
        })
        .catch(err => {
            showDialog("error", "api globalLatestTweets:" + err.toString());
        });
}

function populateLatestTweets(newCachedTweet, insertAtHead) {
    const tweetsPark = document.querySelector('.tweets-park');

    let maxCreateTime = BigInt(0);
    let minCreateTime = minTweetIdCurShowed;
    newCachedTweet.forEach(async tweet => {

        if (!tweet.name || tweet.name === "unknown" || tweet.profile_image_url === DefaultAvatarSrc) {
            const twitterInfo = await loadTwitterInfo(tweet.twitter_id, true)
            if (!twitterInfo) {
                tweet.name = "unknown";
                tweet.username = "unknown";
            } else {
                tweet.name = twitterInfo.name;
                tweet.username = twitterInfo.username;
                tweet.profile_image_url = twitterInfo.profile_image_url;
            }
            TweetToShowOnWeb.syncToDb(tweet);
        }

        const timeSuffix = tweet.create_time;
        const createTime = BigInt(tweet.create_time);

        if (createTime > maxCreateTime) {
            maxCreateTime = createTime;
        }
        if (createTime < minCreateTime || minCreateTime === BigInt(0)) {
            minCreateTime = createTime;
        }
        // 创建 tweet-card 元素
        const tweetCard = document.createElement('div');
        tweetCard.classList.add('tweet-card');

        // 设置 tweet-header 内容
        const tweetHeader = document.createElement('div');
        tweetHeader.classList.add('tweet-header');

        const avatarImg = document.createElement('img');
        avatarImg.src = tweet.profile_image_url;
        avatarImg.alt = "Avatar";
        avatarImg.id = `twitterAvatar-${timeSuffix}`;
        tweetHeader.appendChild(avatarImg);

        const nameSpan = document.createElement('span');
        nameSpan.classList.add('name');
        nameSpan.id = `twitterName-${timeSuffix}`;
        nameSpan.textContent = tweet.name;
        tweetHeader.appendChild(nameSpan);

        const usernameSpan = document.createElement('span');
        usernameSpan.classList.add('username');
        usernameSpan.id = `twitterUserName-${timeSuffix}`;
        usernameSpan.textContent = "@" + tweet.username;
        tweetHeader.appendChild(usernameSpan);

        const timeSpan = document.createElement('span');
        timeSpan.classList.add('time');
        timeSpan.id = `tweet-create-time-${timeSuffix}`;
        timeSpan.textContent = formatTime(tweet.create_time);
        tweetHeader.appendChild(timeSpan);

        tweetCard.appendChild(tweetHeader);

        // 设置 tweet-content
        const tweetContent = document.createElement('div');
        tweetContent.classList.add('tweet-content', 'tweet-content-collapsed');
        tweetContent.id = `tweet-content-${timeSuffix}`;
        tweetContent.textContent = tweet.text;
        tweetCard.appendChild(tweetContent);

        // 添加 Show more 按钮
        const showMoreBtn = document.createElement('button');
        showMoreBtn.classList.add('show-more');
        showMoreBtn.textContent = "Show more";
        showMoreBtn.onclick = function () {
            showFullTweetContent.call(this);
        };

        // 设置 tweet-footer
        const tweetFooter = document.createElement('div');
        tweetFooter.classList.add('tweet-footer');

        const tweetActionDiv = document.createElement('div');
        tweetActionDiv.classList.add('tweet-action');
        tweetActionDiv.innerHTML = `
            <button>0.01 eth打赏</button>
            <span>总赏额：0.23 eth 产生彩票数：68张</span>
        `;
        tweetFooter.appendChild(tweetActionDiv);

        const tweetInfoDiv = document.createElement('div');
        tweetInfoDiv.classList.add('tweet-info');
        tweetInfoDiv.innerHTML = `Payment Hash: <span id="tweet-payment-hash-${timeSuffix}">${tweet.tx_hash}</span> 
Status:<span id="tweet-payment-status-${timeSuffix}">${TXStatus.Str(tweet.payment_status)}</span>`;

        tweetFooter.appendChild(tweetInfoDiv);

        tweetCard.appendChild(tweetFooter);

        if (insertAtHead) {
            tweetsPark.insertBefore(tweetCard, tweetsPark.firstChild);
        } else {
            tweetsPark.appendChild(tweetCard);
        }
        setTimeout(() => {
            if (tweetContent.scrollHeight <= tweetContent.clientHeight) {
                showMoreBtn.style.display = 'none';
            } else {
                tweetCard.appendChild(showMoreBtn); // 只有在内容溢出时才添加按钮
            }
        }, 0);
    });

    maxTweetIdCurShowed = maxCreateTime;
    minTweetIdCurShowed = minCreateTime
}

function formatTime(createTime) {
    return new Date(createTime).toLocaleString();
}

async function postTweet() {
    try {
        const {content, twitterID, web3Id, message} = getUserInput();
        if (!content || !twitterID || !web3Id) return;

        const signature = await signMessage(message, web3Id);
        if (!signature) return;

        const tweetHash = ethers.utils.hashMessage(message);
        console.log("tweetHash=>", tweetHash, "sig=>", signature, "web3Id=>", web3Id, "message\n", message);

        showWaiting("posting to twitter");
        const resp = await PostToSrvByJson("/postTweet", new SignDataForPost(message, signature))
        const refreshedTweet = JSON.parse(resp);

        document.getElementById("tweets-content").value = '';
        parseNjTweetsFromSrv([refreshedTweet], true);


        changeLoadingTips("paying for tweet post")

        tweetVoteContract.publishTweet(tweetHash, signature, {value: tweetPostPrice})
            .then((txResponse) => {
                console.log("Transaction Response: ", txResponse);
                changeLoadingTips("waiting for blockchain packaging");
                refreshedTweet.payment_status = TXStatus.Pending;
                refreshedTweet.tx_hash = txResponse.hash;
                updateTweetPaymentStatus(refreshedTweet)
                return txResponse.wait();
            })
            .then((txReceipt) => {
                console.log("Transaction Receipt: ", txReceipt);
                hideLoading();
                refreshedTweet.payment_status = txReceipt.status ? TXStatus.Success : TXStatus.Failed;
                showDialog("transaction " + txReceipt.status ? "confirmed" : "failed");
                updateTweetPaymentStatus(refreshedTweet)
            }).catch((err) => {
            hideLoading();
            if (err.code === 4001) {
                return;
            }
            console.error("transaction for tweet: ", err);
            showDialog("transaction err:" + err.toString());
        });

    } catch (err) {
        hideLoading();
        if (err.code === 4001) {
            return;
        }
        console.error("Error publishing tweet: ", err);
        showDialog("error", "postTweet:" + err.toString());
    }
}

function updateTweetPaymentStatus(tweetObj) {

    PostToSrvByJson("/updateTweetPaymentStatus", {
        create_time: tweetObj.create_time,
        status: tweetObj.payment_status,
        hash: tweetObj.tx_hash
    }).then(resp => {
        console.log(resp);
        TweetToShowOnWeb.syncToDb(tweetObj);
        document.getElementById("tweet-payment-status-" + tweetObj.create_time).innerText = TXStatus.Str(tweetObj.payment_status);
    }).catch(err => {
        console.log(err);
        showWaiting("error", "update payment status failed:" + err.toString());
    })
}

function getUserInput() {
    const content = document.getElementById("tweets-content").value.trim();
    if (!content) showDialog("tips", "content can't be empty");

    const twitterID = ninjaUserObj.tw_id;
    if (!twitterID) showDialog("tips", "bind your twitter first");

    const web3Id = ninjaUserObj.eth_addr;
    const tweet = new TweetContentToPost(content, (new Date()).getTime(), web3Id, twitterID);
    return {content, twitterID, web3Id, message: JSON.stringify(tweet)};
}

async function signMessage(message, web3Id) {
    if (!metamaskObj) {
        window.location.href = "/signIn";
        return null;
    }
    return await metamaskObj.request({
        method: 'personal_sign',
        params: [message, web3Id],
    });
}



