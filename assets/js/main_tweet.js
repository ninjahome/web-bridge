const dbKeyCachedGlobalTweets = "__db_key_cached_global_tweets__"
const dbKeyCachedTweetContentById = "__db_key_cached_tweet_content__by_id__"
const maxCachedLocalTweetNo = 120
let isMoreTweetsLoading = false;
let hasMoreTweetsToLoad = true;
let maxTweetIdCurShowed = BigInt(0);
let minTweetIdCurShowed = BigInt(0);

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
        this.prefixed_hash = njTweet.prefixed_hash;
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

    static syncToDb(obj) {
        localStorage.setItem(TweetToShowOnWeb.DBKey(obj.create_time), JSON.stringify(obj));
    }

    static load(tweetID) {
        const storedVal = localStorage.getItem(TweetToShowOnWeb.DBKey(tweetID));
        return storedVal ? JSON.parse(storedVal) : null;
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

    const tweets = localTweetsIds.map(tweetID => {
        return TweetToShowOnWeb.load(tweetID);
    }).filter(tweet => tweet !== null);

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
    newCachedTweet.forEach(tweet => {
        const tweetCard = document.getElementById('tweetTemplate').cloneNode(true);
        tweetCard.style.display = '';

        const timeSuffix = tweet.create_time;
        const createTime = BigInt(timeSuffix);
        if (createTime > maxCreateTime) {
            maxCreateTime = createTime;
        }
        if (createTime < minCreateTime || minCreateTime === BigInt(0)) {
            minCreateTime = createTime;
        }
        tweetCard.id = "tweet-card-info-" + timeSuffix;

        tweetCard.dataset.createTime = tweet.create_time;
        tweetCard.dataset.signature = tweet.signature;
        tweetCard.dataset.prefixedHash = tweet.prefixed_hash;

        tweetCard.querySelector('.twitterAvatar').src = tweet.profile_image_url;
        tweetCard.querySelector('.twitterName').textContent = tweet.name;
        tweetCard.querySelector('.twitterUserName').textContent = '@' + tweet.username;
        tweetCard.querySelector('.tweetCreateTime').textContent = formatTime(tweet.create_time);
        tweetCard.querySelector('.tweet-content').textContent = tweet.text;

        const tweetVotePriceInEth = ethers.utils.formatUnits(voteContractMeta.votePrice, 'ether');
        tweetCard.querySelector('.tweet-action button').textContent = `打赏(${tweetVotePriceInEth} eth)`;

        tweetCard.querySelector('.tweetPaymentStatus').textContent = TXStatus.Str(tweet.payment_status);

        if (ninjaUserObj.eth_addr === tweet.web3_id && tweet.payment_status === TXStatus.NoPay) {
            const retryButton = tweetCard.querySelector('.tweetPaymentRetry')
            retryButton.classList.add('show');
        }

        if (insertAtHead) {
            tweetsPark.insertBefore(tweetCard, tweetsPark.firstChild);
        } else {
            tweetsPark.appendChild(tweetCard);
        }
    });
    maxTweetIdCurShowed = maxCreateTime;
    minTweetIdCurShowed = minCreateTime;
    handleShowMoreButtons();
}

function handleShowMoreButtons() {
    requestAnimationFrame(() => {
        document.querySelectorAll('.tweet-card').forEach(tweetCard => {
            const tweetContent = tweetCard.querySelector('.tweet-content');
            const showMoreBtn = tweetCard.querySelector('.show-more');

            if (tweetContent.scrollHeight <= tweetContent.clientHeight) {
                showMoreBtn.style.display = 'none';
            } else {
                tweetCard.appendChild(showMoreBtn);
            }
        });
    });
}

function updateTweetCardWhenStatusChanged(createTime, newStatus) {
    const tweetCardId = "tweet-card-info-" + createTime;
    const tweetCard = document.getElementById(tweetCardId);
    if (!tweetCard) {
        console.error('Tweet card not found for ID:', tweetCardId);
        return;
    }

    const paymentStatusElement = tweetCard.querySelector('.tweetPaymentStatus');
    if (!paymentStatusElement) {
        console.error('Tweet payment status element not found');
        return;
    }

    paymentStatusElement.textContent = TXStatus.Str(newStatus);
    const retryButton = tweetCard.querySelector('.tweetPaymentRetry')
    retryButton.classList.remove('show');
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
        const basicTweet = JSON.parse(resp);

        document.getElementById("tweets-content").value = '';
        parseNjTweetsFromSrv([basicTweet], true);

        await processTweetPayment(basicTweet.create_time, basicTweet.prefixed_hash, basicTweet.signature);
    } catch (err) {
        hideLoading();
        if (err.code === 4001) {
            return;
        }
        console.error("Error publishing tweet: ", err);
        showDialog("Transaction error: " + err.code+":"+err.message);
    }
}

async function processTweetPayment(create_time, prefixed_hash, signature) {
    try {
        changeLoadingTips("paying for tweet post");

        const txResponse = await tweetVoteContract.publishTweet(
            prefixed_hash,
            signature,
            {value: voteContractMeta.postPrice}
        );
        console.log("Transaction Response: ", txResponse);

        changeLoadingTips("waiting for blockchain packaging");
        updateTweetPaymentStatus(create_time, TXStatus.Pending, txResponse.hash);

        const txReceipt = await txResponse.wait();
        console.log("Transaction Receipt: ", txReceipt);

        const txStatus = txReceipt.status ? TXStatus.Success : TXStatus.Failed;
        updateTweetPaymentStatus(create_time, txStatus, txResponse.hash);

        hideLoading();
        showDialog("transaction " + (txReceipt.status ? "confirmed" : "failed"));

    } catch (err) {
        console.error("Transaction error: ", err);
        hideLoading();
        if (err.code === 4001) {
            return;
        }
        if (err.data && err.data.message.includes("duplicate post")){
            updateTweetPaymentStatus(create_time,TXStatus.Success,prefixed_hash);
            return;
        }
        showDialog("Transaction error: " + err.code+":"+err.message);
    }
}

function updateTweetPaymentStatus(create_time, status, hash) {
    PostToSrvByJson("/updateTweetPaymentStatus", {
        create_time: create_time, status: status, hash: hash
    }).then(resp => {
        console.log(resp);
        updateTweetCardWhenStatusChanged(create_time, status);
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
        method: 'personal_sign', params: [message, web3Id],
    });
}

function voteToThisTweet() {
}

async function payThisTweetAgain() {
    try {
        const tweetCard = this.closest('.tweet-card');

        const createTime = Number(tweetCard.dataset.createTime);
        const signature = tweetCard.dataset.signature;
        const prefixedHash = tweetCard.dataset.prefixedHash;
        console.log("Create Time: ", createTime);
        console.log("signature: ", signature);
        console.log("Prefixed Hash: ", prefixedHash);

        showWaiting("prepare to pay");

        await processTweetPayment(createTime, prefixedHash, signature)
    } catch (err) {
        console.error("Transaction error: ", err);
        hideLoading();
        if (err.code === 4001) {
            return;
        }
        showDialog("Transaction error: " + err.code+":"+err.message);
    }
}