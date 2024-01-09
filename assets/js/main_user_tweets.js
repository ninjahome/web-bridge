const cachedUserTweets = new MemCachedTweets();

async function loadTweetsUserPosted() {
    curScrollContentID = 2;
    await __loadTweetAtUserPost(true, ninjaUserObj.eth_addr);
}

async  function olderPostedTweets(){
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

function fillUserPostedTweetsList(newest){
    const tweetsDiv = document.querySelector('.tweets-post-by-user');

    for (const tweet of cachedUserTweets.CachedItem) {

        const tweetCard = document.getElementById('tweetTemplateForUserSelf').cloneNode(true);
        tweetCard.style.display = '';
        tweetCard.querySelector('.tweet-header').id = "tweet-header-for-user-" + tweet.create_time;
        tweetCard.id = "tweet-card-for-user-" + tweet.create_time;
        setupCommonTweetHeader(tweetCard, tweet);

        tweetCard.querySelector('.vote-number').textContent = 0;//TODO:: refactor this logic.

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
}

async function olderVotedTweets(){
}