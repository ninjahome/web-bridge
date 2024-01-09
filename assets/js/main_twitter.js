const cachedGlobalTweets = new MemCachedTweets();

function bindingTwitter() {
    window.location.href = "/signUpByTwitter";
}

async function __loadTweetsAtHomePage(newest) {
    try {
        let startID;
        if (newest) {
            startID = cachedGlobalTweets.MaxID;
        } else {
            startID = cachedGlobalTweets.MinID;
        }
        const param = new TweetQueryParam(startID, newest, "", []);

        const needUpdateUI = await TweetsQuery(param, cachedGlobalTweets);
        if (needUpdateUI) {
            fillTweetParkAtHomePage(newest);
            cachedGlobalTweets.CachedItem = [];
        }

    } catch (err) {
        console.log(err);
        showDialog("error", err.toString());
    }
}

async function loadTweetsForHomePage() {
    await __loadTweetsAtHomePage(true);
}

async function loadOlderTweetsForHomePage() {
    await __loadTweetsAtHomePage(false);
}

async function loadTwitterUserInfoFromSrv(twitterID, useCache, syncFromTwitter) {
    try {
        if (syncFromTwitter) {
            useCache = false;
        }
        if (useCache) {
            let tw_data = TwitterBasicInfo.loadTwBasicInfo(twitterID)
            if (tw_data) {
                return tw_data;
            }
        }
        const response = await GetToSrvByJson("/queryTwBasicById?twitterID=" + twitterID + "&&forceSync=" + syncFromTwitter);
        if (!response.ok) {
            console.log("query twitter basic info failed")
            return null;
        }

        const text = await response.text();
        return TwitterBasicInfo.cacheTwBasicInfo(text);
    } catch (err) {
        console.log("queryTwBasicById err:", err)
        return null;
    }
}

function fillTweetParkAtHomePage(newest) {
    const tweetsPark = document.getElementById('tweets-park');

    for (const tweet of cachedGlobalTweets.CachedItem) {

        const tweetCard = document.getElementById('tweetTemplate').cloneNode(true);
        tweetCard.style.display = '';

        tweetCard.querySelector('.tweet-header').id = "tweet-card-header-for-home-" + tweet.create_time;
        tweetCard.id = "tweet-card-for-home-" + tweet.create_time;

        setupCommonTweetHeader(tweetCard, tweet);

        const voteBtn = tweetCard.querySelector('.tweet-action-vote');
        if (voteContractMeta) {
            voteBtn.textContent = `投票(${voteContractMeta.votePriceInEth} eth)`;
            voteBtn.onclick = () => voteToThisTweet(tweet.create_time);
        }

        tweetCard.querySelector('.vote-number').textContent = 0;//TODO:: refactor this logic.

        const contentArea = tweetCard.querySelector('.tweet-content');
        contentArea.textContent = tweet.text;

        if (newest) {
            tweetsPark.insertBefore(tweetCard, tweetsPark.firstChild);
        } else {
            tweetsPark.appendChild(tweetCard);
        }

        const showMoreBtn = tweetCard.querySelector('.show-more');
        if (contentArea.scrollHeight <= contentArea.clientHeight) {
            showMoreBtn.style.display = 'none';
        } else {
            showMoreBtn.style.display = 'block';
        }
    }
}

async function preparePostMsg() {
    const content = document.getElementById("tweets-content-txt-area").value.trim();
    if (!content) {
        showDialog("tips", "content can't be empty")
        return null;
    }

    const tweet = new TweetContentToPost(content,
        (new Date()).getTime(), ninjaUserObj.eth_addr, ninjaUserObj.tw_id);
    const message = JSON.stringify(tweet);

    const signature = await metamaskObj.request({
        method: 'personal_sign', params: [message, ninjaUserObj.eth_addr],
    });
    if (!signature) {
        showDialog("tips", "empty signature")
        return null;
    }
    return new SignDataForPost(message, signature);
}

function updatePaymentStatusToSrv(tweet) {
    PostToSrvByJson("/updateTweetPaymentStatus", {
        create_time: tweet.create_time,
        status: tweet.payment_status,
        hash: tweet.prefixed_hash
    }).then(r => {
        console.log(r);
    })
}

async function postTweetWithPayment() {
    try {
        const tweetObj = await preparePostMsg();
        showWaiting("posting to twitter");
        const resp = await PostToSrvByJson("/postTweet", tweetObj);
        if (!resp) {
            showDialog("error", "post tweet failed");
            return;
        }
        const basicTweet = JSON.parse(resp);

        await procPaymentForPostedTweet(basicTweet, updatePaymentStatusToSrv);

        __loadTweetsAtHomePage(true).then(r => {
            clearDraftTweetContent();
        });
    } catch (err) {
        checkMetamaskErr(err);
    } finally {
        closePostTweetDiv();
    }
}

async function showPostTweetDiv() {
    if (!metamaskProvider) {
        showDialog("tips", "please change metamask to arbitrum network")
        return;
    }

    if (!ninjaUserObj.tw_id) {
        showDialog("tips", "bind twitter first", bindingTwitter);
        return;
    }

    const modal = document.querySelector('.modal-for-tweet-post');
    modal.style.display = 'block';
    document.getElementById('modal-overlay').style.display = 'block';


    const postBtn = document.getElementById("tweet-post-with-eth-btn");
    postBtn.innerText = "发布推文(" + voteContractMeta.votePriceInEth + " eth)"
}

function closePostTweetDiv() {
    const modal = document.querySelector('.modal-for-tweet-post');
    modal.style.display = 'none';
    document.getElementById('modal-overlay').style.display = 'none';
}

function clearDraftTweetContent() {
    document.getElementById("tweets-content-txt-area").value = '';
}

function showFullTweetContent() {
    const tweetCard = this.closest('.tweet-card');
    const tweetContent = tweetCard.querySelector('.tweet-content');
    const isMore = this.getAttribute('data-more') === 'true';

    if (isMore)  {
        tweetContent.style.display = 'block';
        tweetContent.classList.remove('tweet-content-collapsed');
        tweetCard.style.maxHeight = 'none';
        this.innerText = "更少";
        this.setAttribute('data-more', 'false');
    } else {
        tweetContent.style.display = '-webkit-box';
        tweetContent.classList.add('tweet-content-collapsed');
        tweetCard.style.maxHeight = '400px';
        this.setAttribute('data-more', 'true');
        this.innerText = "更多";
    }
}
