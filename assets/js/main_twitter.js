const cachedTweetsShowed = new MemCachedTweets();

function bindingTwitter() {
    window.location.href = "/signUpByTwitter";
}

function checkTwitterID(){
    if(ninjaUserObj.tw_id){
        return;
    }
    showDialog("tips","bind twitter first",bindingTwitter);
}

function cachedToMem(tweetArray, refreshNewest) {
    if (tweetArray.length === 0) {
        if (!refreshNewest) {
            cachedTweetsShowed.moreOldTweets = false;
        }
        return;
    }


    tweetArray.map(tweet => {
        cachedTweetsShowed.TweetMaps.set(tweet.create_time, tweet);
        if (tweet.create_time > cachedTweetsShowed.MaxID) {
            cachedTweetsShowed.MaxID = tweet.create_time;
        }

        if (tweet.create_time < cachedTweetsShowed.MinID || cachedTweetsShowed.MinID === BigInt(0)) {
            cachedTweetsShowed.MinID = tweet.create_time;
        }
    });
}

function loadLatestTweet(refreshNewest) {
    let startID;
    if (refreshNewest) {
        startID = cachedTweetsShowed.MaxID;
    } else {
        startID = cachedTweetsShowed.MinID;
    }

    fetch("/globalLatestTweets?startID=" + startID + "&&isRefresh=" + refreshNewest)
        .then(response => response.json())
        .then(tweetArray => {
            cachedToMem(tweetArray, refreshNewest);
            refreshTweetPark();
        })
        .catch(err => {
            showDialog("error", "api globalLatestTweets:" + err.toString());
        });
}

async function loadTwitterInfoFromSrv(twitterID, needCache, forceSync) {
    if (!forceSync) {
        forceSync = false;
    }
    try {
        if (needCache) {
            let tw_data = TwitterBasicInfo.loadTwBasicInfo(twitterID)
            if (tw_data) {
                return tw_data;
            }
        }
        const response = await GetToSrvByJson("/queryTwBasicById?twitterID=" + twitterID + "&&forceSync=" + forceSync);
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

function refreshTweetPark(){
    const sortedKeys = Array.from(cachedTweetsShowed.TweetMaps.keys()).sort((a, b) => b - a);
    const tweetsPark = document.querySelector('.tweets-park');
    tweetsPark.innerHTML = '';
    for (const key of sortedKeys) {
        const tweet = cachedTweetsShowed.TweetMaps.get(key);
        if (!tweet){
            console.log("no such obj:",key);
            continue;
        }

        const tweetCard = document.getElementById('tweetTemplate').cloneNode(true);
        tweetCard.style.display = '';
        tweetCard.id = "tweet-card-info-" + key;
        tweetCard.querySelector('.tweet-header').id = "tweet-header-" + tweet.create_time;
        tweetCard.querySelector('.tweetCreateTime').textContent = formatTime(tweet.create_time);
        tweetCard.querySelector('.tweet-content').textContent = tweet.text;

        const twitterObj = TwitterBasicInfo.loadTwBasicInfo(tweet.twitter_id);
        if (!twitterObj){
            loadTwitterInfoFromSrv().then(newObj=>{
                tweetCard.querySelector('.twitterAvatar').src = newObj.profile_image_url;
                tweetCard.querySelector('.twitterName').textContent = newObj.name;
                tweetCard.querySelector('.twitterUserName').textContent = '@' + newObj.username;
            });
        }else{
            tweetCard.querySelector('.twitterAvatar').src = twitterObj.profile_image_url;
            tweetCard.querySelector('.twitterName').textContent = twitterObj.name;
            tweetCard.querySelector('.twitterUserName').textContent = '@' + twitterObj.username;
        }

        const voteBtn = tweetCard.querySelector('.tweet-action-vote');

        if (voteContractMeta){
            voteBtn.textContent = `投票(${voteContractMeta.votePriceInEth} eth)`;
            voteBtn.onclick = () => voteToThisTweet(tweet.create_time);
        }

        const statusElem = tweetCard.querySelector('.tweetPaymentStatus');
        statusElem.textContent = TXStatus.Str(tweet.payment_status);

        const retryButton = tweetCard.querySelector('.tweetPaymentRetry')
        setupTweetPaymentStatus(tweet, retryButton, statusElem);
        if (tweet.payment_status === TXStatus.NoPay && ninjaUserObj.eth_addr === tweet.web3_id){
            retryButton.classList.add('show');
            retryButton.onclick = () => payThisTweetAgain(tweet.create_time);
        }

        tweetCard.querySelector('.vote-number').textContent = 0;
        tweetsPark.appendChild(tweetCard);
    }
    handleShowMoreButtons();
}

function handleShowMoreButtons() {
    requestAnimationFrame(() => {
        document.querySelectorAll('.tweet-card').forEach(tweetCard => {
            const tweetContent = tweetCard.querySelector('.tweet-content');
            const showMoreBtn = tweetCard.querySelector('.show-more');

            if (tweetContent.scrollHeight <= tweetContent.clientHeight) {
                showMoreBtn.style.display = 'none';
            } else {
                tweetCard.appendChild(showMoreBtn);
            }
        });
    });
}

function getUserInput() {
    const content = document.getElementById("tweets-content-txt-area").value.trim();
    if (!content) showDialog("tips", "content can't be empty");

    const twitterID = ninjaUserObj.tw_id;
    if (!twitterID) showDialog("tips", "bind your twitter first");

    const web3Id = ninjaUserObj.eth_addr;
    const tweet = new TweetContentToPost(content, (new Date()).getTime(), web3Id, twitterID);
    return {content, twitterID, web3Id, message: JSON.stringify(tweet)};
}

async function signMessage(message, web3Id) {
    if (!metamaskObj) {
        window.location.href = "/signIn";
        return null;
    }
    return await metamaskObj.request({
        method: 'personal_sign', params: [message, web3Id],
    });
}

async function postTweet() {
    try {
        const {content, twitterID, web3Id, message} = getUserInput();
        if (!content || !twitterID || !web3Id) return;

        const signature = await signMessage(message, web3Id);
        if (!signature) return;

        const tweetHash = ethers.utils.hashMessage(message);
        // console.log("tweetHash=>", tweetHash, "sig=>", signature, "web3Id=>", web3Id, "message\n", message);

        showWaiting("posting to twitter");
        const resp = await PostToSrvByJson("/postTweet", new SignDataForPost(message, signature))
        const basicTweet = JSON.parse(resp);

        document.getElementById("tweets-content-txt-area").value = '';

        await processTweetPayment(basicTweet.create_time, basicTweet.prefixed_hash, basicTweet.signature);
    } catch (err) {
        checkMetamaskErr(err);
    }
}