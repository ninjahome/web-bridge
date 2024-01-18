const __globalTweetMemCache = new Map()
const __globalTweetMemCacheByHash = new Map()

window.onscroll = function () {
    throttle(contentScroll, 200);
}

let throttleTimer;

function throttle(callback, time) {
    if (throttleTimer) return;

    throttleTimer = setTimeout(() => {
        callback();
        clearTimeout(throttleTimer);
        throttleTimer = null;
    }, time);
}

function contentScroll() {
    let cacheObj;
    let uiCallback;

    switch (curScrollContentID) {
        case 0:
            cacheObj = cachedGlobalTweets;
            uiCallback = loadOlderTweetsForHomePage;
            break;
        case 1:
            cacheObj = cachedUserTweets;
            uiCallback = loadOlderMostVotedTweet;
            break;
        case 13:
            cacheObj = cachedTopVotedKolUser;
            uiCallback = loadOlderMostVotedKol;
            break;
        case 14:
            cacheObj = cachedTopVoterUser;
            uiCallback = loadOlderMostVoter;
            break;
        case 2:
            cacheObj = cachedUserTweets;
            uiCallback = olderPostedTweets;
            break;
        case 22:
            cacheObj = cachedUserVotedTweets;
            uiCallback = olderVotedTweets;
            break;
        default:
            return;
    }

    if (!cacheObj.canLoadMoreOldData()) {
        return;
    }

    const documentHeight = Math.max(document.body.scrollHeight, document.documentElement.scrollHeight);
    if (window.innerHeight + window.scrollY < documentHeight - 100) {
        return;
    }

    cacheObj.isLoading = true;
    uiCallback().then(r => {
        console.log("common load latest older data");
    }).finally(r => {
        cacheObj.isLoading = false;
    });
}

function clearCachedData() {
    localStorage.clear();
    sessionStorage.clear();
    window.location.href = "/signIn";
}

async function showHoverCard(event, web3ID) {

    const hoverCard = document.getElementById('hover-card');
    const rect = event.currentTarget.getBoundingClientRect();
    const avatar = event.currentTarget.querySelector('img').src;
    const name = event.currentTarget.querySelector('.name').textContent;
    const userName = event.currentTarget.querySelector('.username').textContent;

    const njUsrInfo = await loadNJUserInfoFromSrv(web3ID, true);

    document.getElementById('hover-avatar').src = avatar;
    document.getElementById('hover-name').textContent = name;
    document.getElementById('hover-user-name').textContent = userName;

    hoverCard.style.display = 'block';
    hoverCard.style.left = `${rect.left}px`;
    hoverCard.style.top = `${rect.bottom + window.scrollY}px`;

    if (!njUsrInfo) {
        console.log("failed to load web3 user:", web3ID);
        return;
    }
    document.getElementById('buy-key-button').onclick = () => {
        hoverCard.style.display = 'none';
        showUserProfile(njUsrInfo)
    };
    document.getElementById('hover-tweet-count').textContent = njUsrInfo.tweet_count;
    document.getElementById('hover-vote-count').textContent = njUsrInfo.vote_count;
    document.getElementById('hover-voted-count').textContent = njUsrInfo.be_voted_count;
}

function hideHoverCard(obj) {

    const hoverCard = document.getElementById('hover-card');
    setTimeout(() => {
        if (!hoverCard.matches(':hover') && !obj.matches(':hover')) {
            hoverCard.style.display = 'none';
        }
    }, 300);
}

function cachedToMem(tweetArray, cacheObj) {
    tweetArray.map(tweet => {
        __globalTweetMemCache.set(tweet.create_time, tweet);
        __globalTweetMemCacheByHash.set(tweet.prefixed_hash, tweet);
        if (tweet.create_time < cacheObj.latestID || cacheObj.latestID === 0) {
            cacheObj.latestID = tweet.create_time;
        }
        cacheObj.CachedItem.push(tweet);
    });
}

async function TweetsQuery(param, newest, cacheObj) {
    try {
        param.start_id = newest ? 0 : cacheObj.latestID;
        if (newest) {
            cacheObj.latestID = 0;
        }
        const resp = await PostToSrvByJson("/tweetQuery", param);
        if (!resp) {
            return false;
        }
        const tweetArray = JSON.parse(resp);

        cacheObj.moreOldTweets = tweetArray.length !== 0 || newest;
        cachedToMem(tweetArray, cacheObj);

        return cacheObj.CachedItem.length > 0;
    } catch (err) {
        throw new Error(err);
    }
}

async function __setOnlyHeader(tweetHeader, twitter_id) {
    const twitterObj = TwitterBasicInfo.loadTwBasicInfo(twitter_id);
    if (!twitterObj) {
        const newObj = await loadTwitterUserInfoFromSrv(twitter_id, true)
        if (!newObj) {
            console.log("failed load twitter user info");
            return;
        }
        tweetHeader.querySelector('.twitterAvatar').src = newObj.profile_image_url;
        tweetHeader.querySelector('.twitterName').textContent = newObj.name;
        tweetHeader.querySelector('.twitterUserName').textContent = '@' + newObj.username;

    } else {
        tweetHeader.querySelector('.twitterAvatar').src = twitterObj.profile_image_url;
        tweetHeader.querySelector('.twitterName').textContent = twitterObj.name;
        tweetHeader.querySelector('.twitterUserName').textContent = '@' + twitterObj.username;
    }
}

async function setupCommonTweetHeader(tweetHeader, tweet, overlap) {
    tweetHeader.querySelector('.tweetCreateTime').textContent = formatTime(tweet.create_time);
    await __setOnlyHeader(tweetHeader, tweet.twitter_id);

    const contentArea = tweetHeader.querySelector('.tweet-content');
    contentArea.textContent = tweet.text;
    const wrappedHeader = tweetHeader.querySelector('.tweet-header');

    if (overlap) {
        const tweetCard = wrappedHeader.parentNode;
        wrappedHeader.addEventListener('mouseenter', (event) => showHoverCard(event, tweet.web3_id));
        wrappedHeader.addEventListener('mouseleave', (event) => hideHoverCard(wrappedHeader));
    }
    return contentArea;
}

function refreshTwitterInfo() {
    showWaiting("tips", "loading from twitter server");
    loadTwitterUserInfoFromSrv(ninjaUserObj.tw_id, false, true).then(async twInfo => {
        hideLoading();
        await setupUserBasicInfoInSetting();
    })
}

function quitFromService() {
    fetch("/signOut", {method: 'GET'}).then(r => {
        window.location.href = "/signIn";
    }).catch(err => {
        console.log(err)
        window.location.href = "/signIn";
    })
}

async function showTweetDetail(parentName) {
    const detail = document.querySelector('#tweet-detail');
    detail.style.display = 'block';

    const tweetCard = this.closest(parentName);
    tweetCard.parentNode.style.display = 'none';

    const create_time = Number(tweetCard.dataset.createTime);
    const detailType = Number(tweetCard.dataset.detailType);

    const obj = __globalTweetMemCache.get(create_time)
    if (!obj) {
        showDialog(DLevel.Error, "can't find tweet obj");
        return;
    }

    detail.querySelector('.tweetCreateTime').textContent = formatTime(obj.create_time);
    await __setOnlyHeader(detail, obj.twitter_id);

    detail.querySelector('.tweet-text').textContent = obj.text;
    detail.querySelector('#tweet-prefixed-hash').textContent = obj.prefixed_hash;
    detail.querySelector('.back-button').onclick = () => {
        tweetCard.parentNode.style.display = 'block';
        detail.style.display = 'none';
    }

    const counter = detail.querySelector('.vote-number');
    counter.textContent = obj.vote_count;
    __showVoteButton(detail, obj, function (newVote) {
        counter.textContent = newVote.vote_count;
    });

    const statusElem = detail.querySelector('.tweetPaymentStatus');
    statusElem.textContent = TXStatus.Str(obj.payment_status);
    if (detailType !== 3 && obj.payment_status !== TXStatus.NoPay) {
        detail.querySelector('.tweetRemoveUnPaid').style.display = 'none';
    }
}

function __showVoteButton(tweetCard, tweet, callback) {
    const voteBtn = tweetCard.querySelector('.tweet-action-vote');
    if (!voteContractMeta) {
        return;
    }
    voteBtn.textContent = `投票(${voteContractMeta.votePriceInEth} eth)`;
    voteBtn.onclick = () => voteToTheTweet(tweet, callback);
}

async function __updateVoteNumberForTweet(tweetObj, newVote) {

    let tweetCard = document.getElementById("tweet-card-for-vote-" + tweetObj.create_time)
    if (tweetCard) {
        tweetCard.querySelector('.total-vote-number').textContent = tweetObj.vote_count;
        if (newVote) {
            const userVoteCounter = tweetCard.querySelector('.user-vote-number');
            userVoteCounter.textContent = newVote.user_vote_count;
            cachedVoteStatusForUser.set(newVote.create_time, newVote.user_vote_count);
        }
    }

    tweetCard = document.getElementById("tweet-card-for-user-" + tweetObj.create_time)
    if (tweetCard) {
        tweetCard.querySelector('.vote-number').textContent = tweetObj.vote_count;
    }

    tweetCard = document.getElementById("tweet-card-for-home-" + tweetObj.create_time)
    if (tweetCard) {
        tweetCard.querySelector('.vote-number').textContent = tweetObj.vote_count;
    }

    tweetCard = document.getElementById("tweet-card-for-njusr-vote-" + tweetObj.create_time)
    if (tweetCard) {
        tweetCard.querySelector('.total-vote-count').textContent = tweetObj.vote_count;
    }
}

async function voteToTheTweet(obj, callback) {

    if (Number(obj.payment_status) !== TXStatus.Success) {
        showDialog(DLevel.Warning, "can't vote to unpaid tweet")
        return;
    }

    openVoteModal(function (voteCount, shareToTweet) {
        procTweetVotePayment(voteCount, obj, async function (create_time, vote_count) {
            const newVote = await updateVoteStatusToSrv(create_time, vote_count);
            obj.vote_count = newVote.vote_count;
            __updateVoteNumberForTweet(obj, newVote).then(r => {
            });
            if (shareToTweet) {
                __shareVoteToTweet(create_time, vote_count).then(r => {
                });
            }
            if (callback) {
                callback(newVote);
            }
        });
    });
}

async function updateVoteStatusToSrv(create_time, vote_count) {
    const resp = await PostToSrvByJson("/updateTweetVoteStatus", {
        create_time: create_time,
        vote_count: Number(vote_count),
    });
    console.log(resp);
    return JSON.parse(resp);
}

async function loadNJUserInfoFromSrv(ethAddr, useCache) {
    try {

        if (useCache) {
            let nj_data = NJUserBasicInfo.loadNjBasic(ethAddr);
            if (nj_data) {
                return nj_data;
            }
        }

        const response = await GetToSrvByJson("/queryNjBasicByID?web3_id=" + ethAddr.toLowerCase());
        if (!response.ok) {
            console.log("query twitter basic info failed")
            return null;
        }

        const obj = await response.json();
        NJUserBasicInfo.cacheNJUsrObj(obj).then(r => {
        })

        return obj;
    } catch (err) {
        console.log("queryTwBasicById err:", err)
        return null;
    }
}


async function withdrawAction(contract) {
    try {
        const txResponse = await contract.withdraw("0x00", true);
        console.log("Transaction Response: ", txResponse);
        showWaiting("prepare to withdraw:" + txResponse.hash);

        const txReceipt = await txResponse.wait();
        console.log("Transaction Receipt: ", txReceipt);

        showDialog(DLevel.Tips, "Transaction: " + txReceipt.status ? "success" : "failed");
        hideLoading();
    } catch (err) {
        checkMetamaskErr(err);
    }
}