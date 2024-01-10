const __globalTweetMemCache = new Map()

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
    uiCallback().finally(r => {
        cacheObj.isLoading = false;
    });
}

let confirmCallback = null;

function openVoteModal(callback) {
    const modal = document.getElementById("vote-no-chose-modal");
    modal.style.display = "block";
    const voteCount = document.getElementById("voteCount");
    voteCount.value = 1;
    confirmCallback = callback;
}

function confirmVoteModal() {
    if (confirmCallback) {
        const voteCount = document.getElementById("voteCount").value;
        confirmCallback(voteCount);
    }
    closeVoteModal();
}

function closeVoteModal() {
    const modal = document.getElementById("vote-no-chose-modal");
    modal.style.display = "none";
}

function increaseVote() {
    const voteCount = document.getElementById("voteCount");
    voteCount.value = parseInt(voteCount.value) + 1;
}

function decreaseVote() {
    const voteCountElement = document.getElementById("voteCount");
    const newVoteCount = Math.max(1, parseInt(voteCountElement.value) - 1);
    voteCountElement.value = newVoteCount.toString();
}

function clearCachedData() {
    localStorage.clear();
    sessionStorage.clear();
    window.location.href = "/signIn";
}


function showHoverCard() {
    const hoverCard = document.getElementById('hover-card');
    const rect = this.getBoundingClientRect();
    const avatar = this.querySelector('img').src;
    const name = this.querySelector('.name').textContent;
    const tweetCount = '0'; // obj.tweet_no;
    const voteCount = '0'; // obj.vote_count;

    document.getElementById('hover-avatar').src = avatar;
    document.getElementById('hover-name').textContent = name;
    document.getElementById('hover-tweet-count').textContent = tweetCount;
    document.getElementById('hover-vote-count').textContent = voteCount;

    hoverCard.style.display = 'block';
    hoverCard.style.left = `${rect.left}px`;
    hoverCard.style.top = `${rect.bottom + window.scrollY}px`;
}

function hideHoverCard(obj) {
    // console.log(obj);
    if (obj) {
        obj.style.display = 'none';
        return;
    }
    const hoverCard = document.getElementById('hover-card');
    setTimeout(() => {
        if (!hoverCard.matches(':hover') && !this.matches(':hover')) {
            hoverCard.style.display = 'none';
        }
    }, 300);
}

function __calculateCacheIdx(cacheObj,tweet){
    const exist = cacheObj.TweetMaps.get(tweet.create_time);
    if (exist) {
        return;
    }
    cacheObj.TweetMaps.set(tweet.create_time, true);
    cacheObj.CachedItem.push(tweet);

    if (tweet.create_time > cacheObj.MaxID) {
        cacheObj.MaxID = tweet.create_time;
    }

    if (tweet.create_time < cacheObj.MinID || cacheObj.MinID === 0) {
        cacheObj.MinID = tweet.create_time;
    }
}

function cachedToMem(tweetArray, cacheObj) {
    tweetArray.map(tweet => {
        __globalTweetMemCache.set(tweet.create_time, tweet);

        if (!cacheObj){
            return;
        }
        __calculateCacheIdx(cacheObj, tweet);
    });
}

async function TweetsQuery(param, newest, cacheObj) {
    try {
        const resp = await PostToSrvByJson("/tweetQuery", param);
        if (!resp) {
            return false;
        }
        const tweetArray = JSON.parse(resp);
        if (tweetArray.length === 0) {
            if (!newest&&cacheObj) {
                cacheObj.moreOldTweets = false;
            }
            return false;
        }

        cachedToMem(tweetArray, cacheObj);
        return cacheObj.CachedItem.length > 0;
    } catch (err) {
        throw new Error(err);
    }
}

async function setupCommonTweetHeader(tweetCard, tweet) {
    tweetCard.querySelector('.tweetCreateTime').textContent = formatTime(tweet.create_time);

    const twitterObj = TwitterBasicInfo.loadTwBasicInfo(tweet.twitter_id);
    if (!twitterObj) {
        const newObj = await loadTwitterUserInfoFromSrv(tweet.twitter_id, true)
        if (!newObj) {
            console.log("failed load twitter user info");
            return;
        }
        tweetCard.querySelector('.twitterAvatar').src = newObj.profile_image_url;
        tweetCard.querySelector('.twitterName').textContent = newObj.name;
        tweetCard.querySelector('.twitterUserName').textContent = '@' + newObj.username;

    } else {
        tweetCard.querySelector('.twitterAvatar').src = twitterObj.profile_image_url;
        tweetCard.querySelector('.twitterName').textContent = twitterObj.name;
        tweetCard.querySelector('.twitterUserName').textContent = '@' + twitterObj.username;
    }
}

function refreshTwitterInfo() {
    loadTwitterUserInfoFromSrv(ninjaUserObj.tw_id, false, true).then(twInfo => {
        setupTwitterElem(twInfo);
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

async function showTweetDetail() {
    const detail = document.querySelector('#tweet-detail');
    detail.style.display = 'block';

    const tweetCard = this.closest('.tweet-card');
    tweetCard.parentNode.style.display = 'none';

    const create_time = Number(tweetCard.dataset.createTime);
    // console.log(create_time);

    const obj = __globalTweetMemCache.get(create_time)
    if (!obj) {
        showDialog("error", "can't find tweet obj");
        return;
    }
    await setupCommonTweetHeader(detail, obj);

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
}

function __showVoteButton(tweetCard, tweet, callback) {
    const voteBtn = tweetCard.querySelector('.tweet-action-vote');
    if (!voteContractMeta) {
        return;
    }
    voteBtn.textContent = `投票(${voteContractMeta.votePriceInEth} eth)`;
    voteBtn.onclick = () => voteToTheTweet(tweet.create_time, callback);
}

async function __updateVoteNumberAllElements(tweetObj, newVote) {

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
}


async function voteToTheTweet(create_time, callback) {
    const obj = __globalTweetMemCache.get(create_time)
    if (!obj) {
        showDialog("tips", "no such tweet obj, please reload page")
        return;
    }

    if (Number(obj.payment_status) !== TXStatus.Success) {
        showDialog("tips", "can't vote to unpaid tweet")
        return;
    }

    openVoteModal(function (voteCount) {
        procTweetVotePayment(voteCount, obj, async function (create_time, vote_count) {
            const newVote = await updateVoteStatusToSrv(create_time, vote_count);
            obj.vote_count = newVote.vote_count;
            __updateVoteNumberAllElements(obj, newVote).then(r => {
            });
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