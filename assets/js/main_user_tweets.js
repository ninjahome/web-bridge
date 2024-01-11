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
    if (cachedUserTweets.latestID === 0) {
        console.log("no need to load older posted data");
        return;
    }
    return __loadTweetAtUserPost(false, ninjaUserObj.eth_addr);
}

async function __loadTweetAtUserPost(newest, web3ID) {
    const param = new TweetQueryParam(0, web3ID, []);

    const needUpdateUI = await TweetsQuery(param, newest, cachedUserTweets);
    if (needUpdateUI) {
        await fillUserPostedTweetsList(param.start_id === 0);
        cachedUserTweets.CachedItem = [];
    }
}

function __checkPayment(tweetCard, tweet) {
    const statusElem = tweetCard.querySelector('.tweetPaymentStatus');
    statusElem.textContent = TXStatus.Str(tweet.payment_status);

    if (tweet.payment_status !== TXStatus.NoPay) {
        return;
    }

    const retryButton = tweetCard.querySelector('.tweetPaymentRetry');
    retryButton.classList.add('show');
    retryButton.onclick = () => procPaymentForPostedTweet(tweet, function (newObj) {
        updatePaymentStatusToSrv(newObj).then();
        __globalTweetMemCache.set(newObj.create_time, newObj);
        if (newObj.payment_status !== TXStatus.NoPay) {
            retryButton.classList.remove('show');
            statusElem.textContent = TXStatus.Str(newObj.payment_status);
        }
    });

    const deleteButton = tweetCard.querySelector('.tweetPaymentDelete');
    deleteButton.classList.add('show');
    deleteButton.onclick = () => removeUnPaidTweets(tweet.create_time).then(r => {
        if (!r) {
            return;
        }
        const id = "tweet-card-for-user-" + tweet.create_time;
        const element = document.getElementById(id);
        if (element) {
            element.parentNode.removeChild(element);
        }
    });
}

async function removeUnPaidTweets(createTime) {
    try {
        const resp = await PostToSrvByJson("/removeUnpaidTweet",
            {create_time: createTime, status: TXStatus.NoPay});
        if (!resp) {
            return false;
        }

        console.log(resp);
        __globalTweetMemCache.delete(createTime);
        return true;
    } catch (e) {
        showDialog("err", "remove unpaid tweet failed:" + e.toString());
        return false;
    }
}

async function fillUserPostedTweetsList(clear) {

    return __fillNormalTweet(clear, 'tweets-post-by-user', cachedUserTweets,
        'tweetTemplateForUserSelf', "tweet-card-for-user-", false,
        function (tweetCard, tweetHeader, tweet) {
            tweetCard.dataset.detailType = '2';
            tweetCard.querySelector('.vote-number').textContent = tweet.vote_count;
            __checkPayment(tweetCard, tweet);
        });
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
    if (cachedUserVotedTweets.latestID === 0) {
        console.log("no need to load older data");
        return;
    }
    return __loadTweetIDsUserVoted(false);
}

const cachedUserVotedTweets = new MemCachedTweets();
const cachedVoteStatusForUser = new Map()

async function __loadTweetIDsUserVoted(newest) {

    const param = new TweetQueryParam(0, ninjaUserObj.eth_addr, []);
    if (!newest) {
        param.start_id = cachedUserVotedTweets.latestID;
    }

    const resp = await PostToSrvByJson("/votedTweetIds", param);
    if (!resp) {
        return;
    }
    console.log(resp);
    let status = JSON.parse(resp);
    if (status.length === 0) {
        if (!newest) {
            cachedUserVotedTweets.moreOldTweets = false;
        }
        return;
    }

    const currentIds = [];
    status.forEach(obj => {
            cachedVoteStatusForUser.set(obj.create_time, obj.vote_count);
            currentIds.push(obj.create_time);
        }
    );
    console.log(currentIds)

    const paramForDetail = new TweetQueryParam(0, "", currentIds);
    const needUpdateUI = await TweetsQuery(paramForDetail, newest, cachedUserVotedTweets);
    if (needUpdateUI) {
        await fillUserVotedTweetsList(newest);
        cachedUserVotedTweets.CachedItem = [];
    }
}

async function fillUserVotedTweetsList(clear) {
    return __fillNormalTweet(clear, 'tweets-voted-by-user',
        cachedUserVotedTweets,
        'tweetTemplateForVoted', "tweet-card-for-vote-", true,
        function (tweetCard, tweetHeader, tweet) {
            tweetCard.querySelector('.total-vote-number').textContent = tweet.vote_count;
            const userVoteCounter = tweetCard.querySelector('.user-vote-number');

            userVoteCounter.textContent = cachedVoteStatusForUser.get(tweet.create_time) ?? 0;
            tweetCard.dataset.detailType = '3';
            __showVoteButton(tweetCard, tweet);
        });
}

