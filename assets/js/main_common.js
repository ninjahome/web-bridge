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

function contentScroll(){
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
    const tweetCard = this.closest('.tweet-card');
    const obj = JSON.parse(tweetCard.dataset.rawObj);

    const hoverCard = document.getElementById('hover-card');
    const rect = this.getBoundingClientRect();
    const avatar = this.querySelector('img').src;
    const name = this.querySelector('.name').textContent;
    const tweetCount = '0'; // obj.tweet_no;
    const voteCount = '0'; // obj.vote_count;

    // 设置悬浮卡片内容
    document.getElementById('hover-avatar').src = avatar;
    document.getElementById('hover-name').textContent = name;
    document.getElementById('hover-tweet-count').textContent = tweetCount;
    document.getElementById('hover-vote-count').textContent = voteCount;

    // 设置悬浮卡片的位置
    hoverCard.style.display = 'block';
    hoverCard.style.left = `${rect.left}px`;
    hoverCard.style.top = `${rect.bottom + window.scrollY}px`;
}

function hideHoverCard(obj) {
    // console.log(obj);
    if(obj){
        obj.style.display = 'none';
        return;
    }
    // 检查鼠标是否在 hover-card 或 tweet-header 上
    const hoverCard = document.getElementById('hover-card');
    setTimeout(() => {
        if (!hoverCard.matches(':hover') && !this.matches(':hover')) {
            hoverCard.style.display = 'none';
        }
    }, 300);
}

function cachedToMem(tweetArray, cacheObj) {
    tweetArray.map(tweet => {
        const exist = cacheObj.TweetMaps.get(tweet.create_time);
        cacheObj.TweetMaps.set(tweet.create_time, tweet);
        if (exist){
            return;
        }
        cacheObj.CachedItem.push(tweet);

        if (tweet.create_time > cacheObj.MaxID) {
            cacheObj.MaxID = tweet.create_time;
        }

        if (tweet.create_time < cacheObj.MinID || cacheObj.MinID === BigInt(0)) {
            cacheObj.MinID = tweet.create_time;
        }
    });
}

async function TweetsQuery(param, cacheObj) {
    try {
        const resp = await PostToSrvByJson("/tweetQuery", param);
        if (!resp) {
            return false;
        }
        const tweetArray = JSON.parse(resp);
        if (tweetArray.length === 0) {
            return false;
        }

        cachedToMem(tweetArray, cacheObj);
        return cacheObj.CachedItem.length > 0;
    } catch (err) {
        throw new Error(err);
    }
}

function setupCommonTweetHeader(tweetCard, tweet){
    tweetCard.querySelector('.tweet-header').id = "tweet-header-" + tweet.create_time;
    tweetCard.querySelector('.tweetCreateTime').textContent = formatTime(tweet.create_time);
    tweetCard.querySelector('.tweet-content').textContent = tweet.text;

    const twitterObj = TwitterBasicInfo.loadTwBasicInfo(tweet.twitter_id);
    if (!twitterObj) {
        loadTwitterUserInfoFromSrv(tweet.twitter_id, true).then(newObj => {
            if (!newObj) {
                console.log("failed load twitter user info");
                return;
            }
            tweetCard.querySelector('.twitterAvatar').src = newObj.profile_image_url;
            tweetCard.querySelector('.twitterName').textContent = newObj.name;
            tweetCard.querySelector('.twitterUserName').textContent = '@' + newObj.username;
        });
    } else {
        tweetCard.querySelector('.twitterAvatar').src = twitterObj.profile_image_url;
        tweetCard.querySelector('.twitterName').textContent = twitterObj.name;
        tweetCard.querySelector('.twitterUserName').textContent = '@' + twitterObj.username;
    }
}