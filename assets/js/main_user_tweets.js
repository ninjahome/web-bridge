const cachedUserTweets = new MemCachedTweets();

function __myPostOrVotedTweetsStatus(isPosted) {
    const tweetsDiv = document.getElementById('tweets-post-by-user');
    tweetsDiv.style.display = isPosted ? 'block' : 'none';
    const votedDiv = document.getElementById('tweets-voted-by-user');
    votedDiv.style.display = isPosted ? 'none' : 'block';
    const detail = document.querySelector('#tweet-detail');
    detail.style.display = 'none';

    const switchArea = document.getElementById("my-vote-or-post-tweet-switch-area")
    switchArea.querySelectorAll(".top-topic-btn").forEach(c => c.classList.remove('active'));

    if (isPosted) {
        switchArea.querySelector('.my-posted-tweets').classList.add('active')
    } else {
        switchArea.querySelector('.my-voted-tweets').classList.add('active')
    }
}

async function loadTweetsUserPosted() {
    try {
        curScrollContentID = 2;
        showWaiting("loading...");
        __myPostOrVotedTweetsStatus(true);
        await __loadTweetAtUserPost(true, ninjaUserObj.eth_addr, cachedUserTweets, fillUserPostedTweetsList);

    } catch (err) {
        console.log(err);
        showDialog(DLevel.Warning, err.toString());
    } finally {
        hideLoading();
    }
}

async function olderPostedTweets() {
    if (cachedUserTweets.latestID === 0) {
        console.log("no need to load older posted data");
        return;
    }
    return __loadTweetAtUserPost(false, ninjaUserObj.eth_addr, cachedUserTweets, fillUserPostedTweetsList);
}

async function __loadTweetAtUserPost(newest, web3ID, cache, callback) {
    const param = new TweetQueryParam(0, web3ID, []);

    const needUpdateUI = await TweetsQuery(param, newest, cache);
    if (needUpdateUI) {
        await callback(param.start_id === 0);
        cache.CachedItem = [];
    }
}

function __checkPayment(tweetCard, tweet) {
    if (tweet.payment_status !== TXStatus.NoPay) {
        return;
    }

    const retryButton = tweetCard.querySelector('.tweetPaymentRetry');
    const deleteButton = tweetCard.querySelector('.tweetPaymentDelete');

    retryButton.classList.add('show');
    retryButton.onclick = () => procPaymentForPostedTweet(tweet, function (newObj, txHash) {
        __globalTweetMemCache.set(newObj.create_time, newObj);
        if (newObj.payment_status !== TXStatus.Success) {
            return;
        }

        retryButton.classList.remove('show');
        deleteButton.classList.remove('show');
        updatePaymentStatusToSrv(newObj, txHash).then();
    });

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
        showWaiting("deleting tweet......")
        await PostToSrvByJson("/removeUnpaidTweet",
            {create_time: createTime, status: TXStatus.NoPay});
        __globalTweetMemCache.delete(createTime);
        return true;
    } catch (e) {
        showDialog(DLevel.Error, "remove unpaid tweet failed:" + e.toString());
        return false;
    } finally {
        hideLoading();
    }
}

async function fillUserPostedTweetsList(clear) {

    return __fillNormalTweet(clear, 'tweets-post-by-user',
        cachedUserTweets.CachedItem,
        'tweetTemplateForUserSelf', "tweet-card-for-user-", false,
        TweetDetailSource.NoNeed, function (tweetCard, tweetHeader, tweet) {
            tweetCard.querySelector('.vote-number').textContent = tweet.vote_count;
            tweetCard.querySelector('.tweet-content').style.cursor = "default";
            __checkPayment(tweetCard, tweet);
        });
}

async function loadTweetsUserVoted() {
    try {
        curScrollContentID = 22;
        showWaiting("loading...");

        __myPostOrVotedTweetsStatus(false);
        await __loadTweetIDsUserVoted(true, ninjaUserObj.eth_addr, cachedUserVotedTweets, cachedVoteStatusForUser, fillUserVotedTweetsList);
    } catch (e) {
        console.log(e);
        showDialog(e.toString());
    } finally {
        hideLoading()
    }
}

async function olderVotedTweets() {
    console.log('lod old data trigger')
    if (cachedUserVotedTweets.latestID === 0) {
        console.log("no need to load older data");
        return;
    }
    return __loadTweetIDsUserVoted(false, ninjaUserObj.eth_addr, cachedUserVotedTweets, cachedVoteStatusForUser, fillUserVotedTweetsList);
}

const cachedUserVotedTweets = new MemCachedTweets();
const cachedVoteStatusForUser = new Map()

async function __loadTweetIDsUserVoted(newest, web3ID, cache, voteStatusCache, callback) {

    const param = new TweetQueryParam(0, web3ID, []);
    if (!newest) {
        param.start_id = cache.latestID;
    }
    const status = await PostToSrvByJson("/votedTweetIds", param);
    if (!status || status.length === 0) {
        if (!newest) {
            cache.moreOldTweets = false;
        }
        return;
    }

    const currentIds = [];
    status.forEach(obj => {
            voteStatusCache.set(obj.create_time, obj.vote_count);
            currentIds.push(obj.create_time);
        }
    );
    // console.log(currentIds)

    const paramForDetail = new TweetQueryParam(0, "", currentIds);
    const needUpdateUI = await TweetsQuery(paramForDetail, newest, cache);
    if (needUpdateUI && callback) {
        await callback(newest);
    }
    cache.CachedItem = [];
}


async function fillUserVotedTweetsList(clear) {
    return __fillNormalTweet(clear, 'tweets-voted-by-user',
        cachedUserVotedTweets.CachedItem,
        'tweetTemplateForVoted',
        "tweet-card-for-vote-",
        true, TweetDetailSource.MyVoted,
        function (tweetCard, tweetHeader, tweet) {

            tweetCard.querySelector('.total-vote-number').textContent = tweet.vote_count;
            const userVoteCounter = tweetCard.querySelector('.user-vote-number');

            userVoteCounter.textContent = cachedVoteStatusForUser.get(tweet.create_time) ?? 0;
            __showVoteButton(tweetCard, tweet);
        });
}

async function showUserProfile(njUser) {
    if (njUser.eth_addr === ninjaUserObj.eth_addr) {
        showDialog(DLevel.Tips, "This is yourself");
        return;
    }
    currentNinjaUsrLoading = njUser;
    const detail = document.getElementById('nj-user-profile');
    detail.style.display = 'block';
    let parentNode;
    document.querySelectorAll('.content-in-middle-area').forEach(c => {
        if (c.classList.contains('active')) {
            parentNode = c;
        }
        c.classList.remove('active')
    });

    detail.querySelector(".back-button").onclick = function () {
        if (parentNode) {
            parentNode.classList.add('active');
        }
        detail.style.display = 'none';
    }

    detail.querySelector(".web3id").textContent = njUser.eth_addr;
    const header = detail.querySelector(".tweet-header-in-profile")
    await __setOnlyHeader(header, njUser.tw_id, njUser.eth_addr);
    await loadPostedTweetsOfNjUsr();
}

const cachedNinjaUserPostedTweets = new MemCachedTweets();

async function olderNinjaUsrPostedTweets() {
    await __loadTweetAtUserPost(false, currentNinjaUsrLoading.eth_addr,
        cachedNinjaUserPostedTweets, fillNinjaUserPostedTweetsList)
}

function __njUserVoteOrPostedTweetsStatus(isPosted) {
    const postedDiv = document.getElementById('nj-user-posted-tweets');
    postedDiv.style.display = isPosted ? 'block' : 'none';

    const votedDiv = document.getElementById('nj-user-vote-tweets');
    votedDiv.style.display = isPosted ? 'none' : 'block';

    const switchArea = document.getElementById('twee-switch-area-for-nj-user')
    switchArea.querySelectorAll('.top-topic-btn').forEach(btn => btn.classList.remove('active'));

    if (isPosted) {
        switchArea.querySelector('.tweetsPostedByNjUser').classList.add('active')
    } else {
        switchArea.querySelector('.tweetsVotedNjUser').classList.add('active')
    }
}

async function loadPostedTweetsOfNjUsr() {
    showWaiting("loading");
    curScrollContentID = 51;
    __njUserVoteOrPostedTweetsStatus(true);
    await __loadTweetAtUserPost(true, currentNinjaUsrLoading.eth_addr, cachedNinjaUserPostedTweets, fillNinjaUserPostedTweetsList)
    hideLoading();
}

async function fillNinjaUserPostedTweetsList(clear) {
    return __fillNormalTweet(clear, 'nj-user-posted-tweets',
        cachedNinjaUserPostedTweets.CachedItem,
        'tweetTemplateForNjUsrProfile', "tweet-card-for-njusr-post-", false,
        TweetDetailSource.NoNeed,
        function (tweetCard, tweetHeader, tweet) {
            tweetCard.querySelector('.total-vote-count').textContent = tweet.vote_count;
            tweetCard.querySelector('.tweet-content').style.cursor = "default";
            tweetCard.querySelector('.vote-count').style.display = 'none';
            __showVoteButton(tweetCard, tweet);
        });
}

const cachedNinjaUserVotedTweets = new MemCachedTweets();
const cachedNinjaVoteStatusForUser = new Map()
let currentNinjaUsrLoading = null

async function olderNinjaUsrVotedTweets() {
    await __loadTweetIDsUserVoted(false, currentNinjaUsrLoading.eth_addr,
        cachedNinjaUserVotedTweets, cachedNinjaVoteStatusForUser, fillNinjaUserVotedTweetsList);
}

async function loadVotedTweetsOfNjUsr() {
    showWaiting("loading");
    curScrollContentID = 52;
    __njUserVoteOrPostedTweetsStatus(false);
    await __loadTweetIDsUserVoted(true, currentNinjaUsrLoading.eth_addr,
        cachedNinjaUserVotedTweets, cachedNinjaVoteStatusForUser, fillNinjaUserVotedTweetsList);
    hideLoading();
}


async function fillNinjaUserVotedTweetsList(clear) {
    return __fillNormalTweet(clear, 'nj-user-vote-tweets',
        cachedNinjaUserVotedTweets.CachedItem,
        'tweetTemplateForNjUsrProfile',
        "tweet-card-for-njusr-vote-",
        false,
        TweetDetailSource.NoNeed,
        function (tweetCard, tweetHeader, tweet) {

            tweetCard.querySelector('.total-vote-count').textContent = tweet.vote_count;
            tweetCard.querySelector('.tweet-content').style.cursor = "default";

            const userVoteCounter = tweetCard.querySelector('.nj-user-vote-count');
            userVoteCounter.textContent = cachedNinjaVoteStatusForUser.get(tweet.create_time) ?? 0;

            __showVoteButton(tweetCard, tweet);
        });
}