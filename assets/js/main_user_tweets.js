const cachedUserTweets = new MemCachedTweets();

async function loadTweetsUserPosted() {
    curScrollContentID = 2;

    const tweetsDiv = document.getElementById('tweets-post-by-user');
    tweetsDiv.style.display = 'block';
    const votedDiv = document.getElementById('tweets-voted-by-user');
    votedDiv.style.display = 'none';
    const detail = document.querySelector('#tweet-detail');
    detail.style.display = 'none';
    __loadTweetAtUserPost(true, ninjaUserObj.eth_addr).then(r => {
        console.log("load newest tweets of user posted success");
    });
}

async function olderPostedTweets() {
    __loadTweetAtUserPost(false, ninjaUserObj.eth_addr).then(r => {
        console.log("load older tweets of user posted success");
    });
}

async function __loadTweetAtUserPost(newest, web3ID) {
    const param = new TweetQueryParam("", newest, web3ID, []);
    if (newest) {
        param.start_id = cachedUserTweets.MaxID;
    } else {
        param.start_id = cachedUserTweets.MinID;
    }

    const needUpdateUI = await TweetsQuery(param, newest, cachedUserTweets);
    if (needUpdateUI) {
        await fillUserPostedTweetsList(newest);
        cachedUserTweets.CachedItem = [];
    }
}

function __checkPayment(tweet, retryButton, statusElem) {
    if (tweet.payment_status !== TXStatus.NoPay) {
        return;
    }
    retryButton.classList.add('show');
    retryButton.onclick = () => procPaymentForPostedTweet(tweet, function (newObj) {
        updatePaymentStatusToSrv(newObj).then();
        __globalTweetMemCache.set(newObj.create_time, newObj);
        if (newObj.payment_status !== TXStatus.NoPay) {
            retryButton.classList.remove('show');
            statusElem.textContent = TXStatus.Str(newObj.payment_status);
        }
    });
}

async function fillUserPostedTweetsList(newest) {
    const tweetsDiv = document.getElementById('tweets-post-by-user');

    for (const tweet of cachedUserTweets.CachedItem) {

        const tweetCard = document.getElementById('tweetTemplateForUserSelf').cloneNode(true);
        tweetCard.style.display = '';

        tweetCard.querySelector('.tweet-header').id = "tweet-header-for-user-" + tweet.create_time;
        tweetCard.id = "tweet-card-for-user-" + tweet.create_time;

        tweetCard.dataset.createTime = tweet.create_time;

        await setupCommonTweetHeader(tweetCard, tweet);

        const contentArea = tweetCard.querySelector('.tweet-content');
        contentArea.textContent = tweet.text;

        tweetCard.querySelector('.vote-number').textContent = tweet.vote_count;

        const statusElem = tweetCard.querySelector('.tweetPaymentStatus');
        statusElem.textContent = TXStatus.Str(tweet.payment_status);

        const retryButton = tweetCard.querySelector('.tweetPaymentRetry')
        __checkPayment(tweet, retryButton, statusElem);

        if (newest) {
            tweetsDiv.insertBefore(tweetCard, tweetsDiv.firstChild);
        } else {
            tweetsDiv.appendChild(tweetCard);
        }
    }
}

async function loadTweetsUserVoted() {

    curScrollContentID = 22;
    const tweetsDiv = document.getElementById('tweets-post-by-user');
    tweetsDiv.style.display = 'none';
    const votedDiv = document.getElementById('tweets-voted-by-user');
    votedDiv.style.display = 'block';
    const detail = document.querySelector('#tweet-detail');
    detail.style.display = 'none';

    await __loadTweetIDsUserVoted(true);
}

async function olderVotedTweets() {
    console.log('lod old data trigger')
}

const cachedUserVotedTweets = new MemCachedTweets();
const cachedVoteStatusForUser = new Map()

async function __loadTweetIDsUserVoted(newest) {

    const param = new TweetQueryParam("", newest, ninjaUserObj.eth_addr, []);
    if (newest) {
        param.start_id = cachedUserVotedTweets.MaxID;
    } else {
        param.start_id = cachedUserVotedTweets.MinID;
    }
    const resp = await PostToSrvByJson("/votedTweetIds", param);
    if (!resp) {
        return;
    }
    console.log(resp);
    let status = JSON.parse(resp);
    if (status.length === 0) {
        return;
    }

    const currentIds = [];
    status.forEach(obj => {
            cachedVoteStatusForUser.set(obj.create_time, obj.vote_count);
            currentIds.push(obj.create_time);
        }
    );
    console.log(currentIds)

    const paramForDetail = new TweetQueryParam(0, newest, "", currentIds);
    const needUpdateUI = await TweetsQuery(paramForDetail, newest, cachedUserVotedTweets);
    if (needUpdateUI) {
        await fillUserVotedTweetsList(newest);
        cachedUserVotedTweets.CachedItem = [];
    }
}

async function fillUserVotedTweetsList(newest) {
    const tweetsDiv = document.getElementById('tweets-voted-by-user');

    for (const tweet of cachedUserVotedTweets.CachedItem) {

        const tweetCard = document.getElementById('tweetTemplateForVoted').cloneNode(true);
        tweetCard.style.display = '';

        tweetCard.querySelector('.tweet-header').id = "tweet-header-for-vote-" + tweet.create_time;
        tweetCard.id = "tweet-card-for-vote-" + tweet.create_time;

        tweetCard.dataset.createTime = tweet.create_time;

        await setupCommonTweetHeader(tweetCard, tweet);

        const contentArea = tweetCard.querySelector('.tweet-content');
        contentArea.textContent = tweet.text;

        tweetCard.querySelector('.total-vote-number').textContent = tweet.vote_count;
        const userVoteCounter = tweetCard.querySelector('.user-vote-number');

        userVoteCounter.textContent = cachedVoteStatusForUser.get(tweet.create_time) ?? 0;

        __showVoteButton(tweetCard, tweet);

        if (newest) {
            tweetsDiv.insertBefore(tweetCard, tweetsDiv.firstChild);
        } else {
            tweetsDiv.appendChild(tweetCard);
        }
    }
}

