const cachedGlobalTweets = new MemCachedTweets();

function bindingTwitter() {
    window.location.href = "/signUpByTwitter";
}

async function __loadTweetsAtHomePage(newest) {
    try {
        const param = new TweetQueryParam(0, "", []);
        const needUpdateUI = await TweetsQuery(param, newest, cachedGlobalTweets);
        if (needUpdateUI) {
            await fillTweetParkAtHomePage(newest);
            cachedGlobalTweets.CachedItem = [];
        }
    } catch (err) {
        console.log(err);
        showDialog("error", err.toString());
    }
}

async function loadTweetsForHomePage() {
    const tweetsDiv = document.getElementById('tweets-park');
    tweetsDiv.style.display = 'block';
    showWaiting("loading.....");
    __loadTweetsAtHomePage(true).then(r => {
        console.log("load newest global tweets success");
    }).finally(r => {
        hideLoading();
    });
}

async function loadOlderTweetsForHomePage() {
    if (cachedGlobalTweets.latestID === 0) {
        console.log("no need to load older data");
        return;
    }
    return __loadTweetsAtHomePage(false)
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

async function __fillNormalTweet(clear, parkID, cached, templateId, cardID, overlap, callback) {
    const tweetsPark = document.getElementById(parkID);
    if (clear) {
        tweetsPark.innerHTML = '';
    }

    for (const tweet of cached.CachedItem) {

        const tweetCard = document.getElementById(templateId).cloneNode(true);
        tweetCard.style.display = '';
        tweetCard.id = cardID + tweet.create_time;
        tweetCard.dataset.createTime = tweet.create_time;
        const tweetHeader = document.getElementById('tweet-header-template').cloneNode(true);
        tweetHeader.style.display = '';

        const sibling = tweetCard.querySelector('.tweet-footer')
        const contentArea = await setupCommonTweetHeader(tweetHeader, tweet, overlap);
        tweetCard.insertBefore(tweetHeader, sibling);

        if (callback) {
            callback(tweetCard, tweetHeader, tweet)
        }

        tweetsPark.appendChild(tweetCard);

        const showMoreBtn = tweetCard.querySelector('.show-more');
        if (contentArea.scrollHeight <= contentArea.clientHeight) {
            showMoreBtn.style.display = 'none';
        } else {
            showMoreBtn.style.display = 'block';
        }
    }
}

async function fillTweetParkAtHomePage(clear) {

    return __fillNormalTweet(clear, 'tweets-park',
        cachedGlobalTweets, 'tweetTemplate',
        "tweet-card-for-home-", true, function (tweetCard, tweetHeader, tweet) {
            tweetCard.dataset.detailType = '1';
            tweetCard.querySelector('.vote-number').textContent = tweet.vote_count;
            __showVoteButton(tweetCard, tweet);
        });
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
    return PostToSrvByJson("/updateTweetPaymentStatus", {
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
            hideLoading();
            showDialog("error", "post tweet failed");
            return;
        }
        const basicTweet = JSON.parse(resp);
        hideLoading();
        await procPaymentForPostedTweet(basicTweet);

        await updatePaymentStatusToSrv(basicTweet)

        if (curScrollContentID === 0) {
            __loadTweetsAtHomePage(true).then(r => {
                clearDraftTweetContent();
            });
        } else if (curScrollContentID === 2) {
            __loadTweetAtUserPost(true, ninjaUserObj.eth_addr).then(r => {
                clearDraftTweetContent();
            });
        }

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


    const postBtn = document.getElementById("tweet-post-with-eth-btn-txt");
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

    if (isMore) {
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
