const cachedUserTweets = new MemCachedTweets();

async function loadTweetsUserPosted() {
    curScrollContentID = 2;

    const tweetsDiv = document.getElementById('tweets-post-by-user');
    tweetsDiv.style.display = 'block';
    const votedDiv = document.getElementById('tweets-voted-by-user');
    votedDiv.style.display = 'none';

    await __loadTweetAtUserPost(true, ninjaUserObj.eth_addr);
}

async function olderPostedTweets() {
    await __loadTweetAtUserPost(false, ninjaUserObj.eth_addr);
}

async function __loadTweetAtUserPost(newest, web3ID) {
    const param = new TweetQueryParam("", newest, web3ID, []);
    if (newest) {
        param.start_id = cachedGlobalTweets.MaxID;
    } else {
        param.start_id = cachedGlobalTweets.MinID;
    }

    const needUpdateUI = await TweetsQuery(param, cachedUserTweets);
    if (needUpdateUI) {
        fillUserPostedTweetsList(newest);
        cachedUserTweets.CachedItem = [];
    }
}

function fillUserPostedTweetsList(newest) {
    const tweetsDiv = document.getElementById('tweets-post-by-user');

    for (const tweet of cachedUserTweets.CachedItem) {

        const tweetCard = document.getElementById('tweetTemplateForUserSelf').cloneNode(true);
        tweetCard.style.display = '';
        tweetCard.querySelector('.tweet-header').id = "tweet-header-for-user-" + tweet.create_time;
        tweetCard.id = "tweet-card-for-user-" + tweet.create_time;
        setupCommonTweetHeader(tweetCard, tweet);

        const contentArea = tweetCard.querySelector('.tweet-content');
        contentArea.textContent = tweet.text;

        tweetCard.querySelector('.vote-number').textContent = 0;//TODO:: refactor this logic.

        const statusElem = tweetCard.querySelector('.tweetPaymentStatus');
        statusElem.textContent = TXStatus.Str(tweet.payment_status);

        const retryButton = tweetCard.querySelector('.tweetPaymentRetry')
        if (tweet.payment_status === TXStatus.NoPay) {
            retryButton.classList.add('show');
            retryButton.onclick = () => payThisTweetAgain(tweet.create_time);
        }

        if (newest) {
            tweetsDiv.insertBefore(tweetCard, tweetsDiv.firstChild);
        } else {
            tweetsDiv.appendChild(tweetCard);
        }
    }
}


const cachedUserVotedTweets = new MemCachedTweets();

async function loadTweetsUserVoted() {
    curScrollContentID = 22;
    const tweetsDiv = document.getElementById('tweets-post-by-user');
    tweetsDiv.style.display = 'none';
    const votedDiv = document.getElementById('tweets-voted-by-user');
    votedDiv.style.display = 'block';
}

async function olderVotedTweets() {
}