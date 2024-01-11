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

        if (blockChain) {
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
    // console.log(tweetArray)
    const newIDs = [];
    const localTweets = tweetArray.map(tweet => {
        const tw_data = TwitterBasicInfo.loadTwBasicInfo(tweet.twitter_id)
        // const blockchain = BlockChainData.load(tweet.web3_id);
        const obj = new TweetToShowOnWeb(tweet, tw_data, null);
        TweetToShowOnWeb.syncToDb(obj);
        newIDs.push(obj.create_time);
        return obj;
    });

    populateLatestTweets(localTweets, refreshNewest).then(async r => {
        await queryLastStatusInfo(newIDs)
    });
    updateLocalCachedTweetList(newIDs)
}

async function queryLastStatusInfo(ids, callback) {
    try {
        const res = await PostToSrvByJson("/tweetStatusRealTime", {create_time: ids})
        // console.log(res);
        const statusMap = JSON.parse(res);
        for (let key in statusMap) {
            let status = statusMap[key];
            // console.log("Create Time:", status.create_time, "Vote Count:", status.vote_count);
            updateTweetCardVoteNo(status.create_time, status.vote_count);
            if (callback) {
                callback(status.create_time, status.vote_count);
            }
        }
    } catch (err) {
        console.log(err)
    }
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
        return TweetToShowOnWeb.load(tweetID)
    }).filter(tweet => tweet !== null);

    if (tweets.length === 0) {
        return;
    }

    populateLatestTweets(tweets, false).then(async r => {
        await queryLastStatusInfo(localTweetsIds)
    });
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

async function populateLatestTweets(newCachedTweet, insertAtHead) {
    const tweetsPark = document.querySelector('.tweets-park');
    let maxCreateTime = BigInt(0);
    let minCreateTime = minTweetIdCurShowed;
    for (const tweet of newCachedTweet) {
        const tweetCard = document.getElementById('tweetTemplate').cloneNode(true);
        tweetCard.style.display = '';

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
        const createTime = BigInt(timeSuffix);
        if (createTime > maxCreateTime) {
            maxCreateTime = createTime;
        }
        if (createTime < minCreateTime || minCreateTime === BigInt(0)) {
            minCreateTime = createTime;
        }
        tweetCard.id = "tweet-card-info-" + timeSuffix;
        tweetCard.dataset.rawObj = JSON.stringify(tweet);
        tweetCard.dataset.createTime = tweet.create_time;
        tweetCard.dataset.signature = tweet.signature;
        tweetCard.dataset.prefixedHash = tweet.prefixed_hash;
        tweetCard.dataset.paidStatus = tweet.payment_status;

        tweetCard.querySelector('.tweet-header').id = "tweet-header-" + tweet.create_time;
        tweetCard.querySelector('.twitterAvatar').src = tweet.profile_image_url;
        tweetCard.querySelector('.twitterName').textContent = tweet.name;
        tweetCard.querySelector('.twitterUserName').textContent = '@' + tweet.username;
        tweetCard.querySelector('.tweetCreateTime').textContent = formatTime(tweet.create_time);
        tweetCard.querySelector('.tweet-content').textContent = tweet.text;

        const voteBtn = tweetCard.querySelector('.tweet-action-vote');
        voteBtn.textContent = `打赏(${voteContractMeta.votePriceInEth} eth)`;
        voteBtn.onclick = () => voteToThisTweet(tweet.create_time);
        const statusElem = tweetCard.querySelector('.tweetPaymentStatus');
        statusElem.textContent = TXStatus.Str(tweet.payment_status);
        const retryButton = tweetCard.querySelector('.tweetPaymentRetry')

        setupTweetPaymentStatus(tweet, retryButton, statusElem);

        tweetCard.querySelector('.vote-number').textContent = 0;

        if (insertAtHead) {
            tweetsPark.insertBefore(tweetCard, tweetsPark.firstChild);
        } else {
            tweetsPark.appendChild(tweetCard);
        }
    }

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
    tweetCard.dataset.paidStatus = newStatus;
    paymentStatusElement.textContent = TXStatus.Str(newStatus);
    const retryButton = tweetCard.querySelector('.tweetPaymentRetry')
    retryButton.classList.remove('show');
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
        checkMetamaskErr(err);
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

        changeLoadingTips("waiting for blockchain packaging:" + txResponse.hash);
        updateTweetPaymentStatus(create_time, TXStatus.Pending, txResponse.hash);

        const txReceipt = await txResponse.wait();
        console.log("Transaction Receipt: ", txReceipt);

        const txStatus = txReceipt.status ? TXStatus.Success : TXStatus.Failed;
        updateTweetPaymentStatus(create_time, txStatus, txResponse.hash);

        hideLoading();
        showDialog("transaction " + (txReceipt.status ? "confirmed" : "failed"));
    } catch (err) {
        const newErr = checkMetamaskErr(err);
        if (newErr && newErr.includes("duplicate post")) {
            updateTweetPaymentStatus(create_time, TXStatus.Success, prefixed_hash);
        }
    }
}

function updateTweetPaymentStatus(create_time, status, hash) {
    PostToSrvByJson("/updateTweetPaymentStatus", {
        create_time: create_time, status: status, hash: hash
    }).then(resp => {
        console.log(resp);
        updateTweetCardWhenStatusChanged(create_time, status);
        const obj = TweetToShowOnWeb.load(create_time);
        obj.payment_status = status;
        TweetToShowOnWeb.syncToDb(obj);
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

function updateTweetCardVoteNo(createTime, totalNo) {
    const tweetCardId = "tweet-card-info-" + createTime;
    const tweetCard = document.getElementById(tweetCardId);
    if (!tweetCard) {
        console.error('Tweet card not found for ID:', tweetCardId);
        return;
    }
    const paymentStatusElement = tweetCard.querySelector('.vote-number');
    if (!paymentStatusElement) {
        console.error('Tweet payment status element not found');
        return;
    }

    paymentStatusElement.textContent = totalNo;
}

function updateTweetVoteStatic(create_time, voteCount) {
    PostToSrvByJson("/updateTweetVoteStatic", {
        create_time: create_time, vote_count: Number(voteCount),
    }).then(resp => {
        // console.log(resp);
        const countFormSrv = JSON.parse(resp);
        updateTweetCardVoteNo(create_time, countFormSrv.vote_count);
        loadUserGameInfo().then(r => {
        });
    }).catch(err => {
        console.log(err);
        showWaiting("error", "update payment status failed:" + err.toString());
    })
}

async function startToVote(voteCount, prefixedHash, createTime) {
    try {
        showWaiting("prepare to pay");

        const amount = voteContractMeta.postPrice.mul(voteCount);

        const txResponse = await tweetVoteContract.voteToTweets(
            prefixedHash,
            voteCount,
            {value: amount}
        );
        console.log("Transaction Response: ", txResponse);
        changeLoadingTips("waiting for blockchain packaging:" + txResponse.hash);

        const txReceipt = await txResponse.wait();
        console.log("Transaction Receipt: ", txReceipt);

        showDialog("Transaction: " + txReceipt.status ? "success" : "failed");

        if (!txReceipt.status) {
            return;
        }

        hideLoading();
        updateTweetVoteStatic(createTime, voteCount);
        loadCurGameMeta();
    } catch (err) {
        checkMetamaskErr(err);
    }
}

async function voteToThisTweet(create_time) {
    console.log(create_time);
    const obj = TweetToShowOnWeb.load(create_time)
    if (!obj){
        showDialog("tips","please reload page")
        return;
    }
    const paidStatus = obj.payment_status;
    const createTime = Number(create_time);
    const prefixedHash = obj.prefixed_hash;

    if (Number(paidStatus) !== 2) {
        showDialog("tips", "can't vote to unpaid tweet")
        return;
    }

    openVoteModal(function (voteCount) {
        // console.log("用户选择的票数:", voteCount);
        startToVote(voteCount, prefixedHash, createTime).then(r => {
            const detail = document.querySelector('#tweet-detail');
            const origNo = detail.querySelector('.vote-number').textContent;
            detail.querySelector('.vote-number').textContent = (Number(origNo) + Number(voteCount)).toString();
        });
    });
}

async function payThisTweetAgain(tweet) {
    try {

        const createTime = Number(tweet.create_time);
        const prefixedHash = tweet.prefixed_hash;
        const signature = tweet.signature;

        console.log("Create Time: ", createTime);
        console.log("signature: ", signature);
        console.log("Prefixed Hash: ", prefixedHash);

        showWaiting("prepare to pay");

        await processTweetPayment(createTime, prefixedHash, signature)
    } catch (err) {
        checkMetamaskErr(err);
    }
}

function showUserVotedTweets() {

}


function showUserPostedTweets() {

}


async function withdrawFromUserTweetIncome() {
    const bInEth = await loadKolTweetIncome();

    if (bInEth <= 0) {
        showDialog("tips", "balance too low");
        return;
    }

    await withdrawAction(tweetVoteContract);

    await loadKolTweetIncome();
}