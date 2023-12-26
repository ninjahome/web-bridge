const dbKeyCachedGlobalTweets = "__db_key_cached_global_tweets__"
const dbKeyCachedTweetContentById = "__db_key_cached_tweet_content__by_id__"
const maxCachedLocalTweetNo = 120
let isMoreTweetsLoading = false;
let hasMoreTweetsToLoad = true;
let maxTweetIdCurShowed = BigInt(0);
let minTweetIdCurShowed = BigInt(0);

class TweetContentToPost {
    constructor(tweet_content, createAt, web3Id, twitterID, tweet_id, signature) {
        this.text = tweet_content;
        this.create_time = createAt;
        this.web3_id = web3Id;
        this.twitter_id = twitterID;
        this.tweet_id = tweet_id;
        this.signature = signature;
    }
}

class TweetToShowOnWeb {
    constructor(njTweet, njTwitter, blockChain) {
        this.text = njTweet.text;
        this.create_time = njTweet.create_time;
        this.web3_id = njTweet.web3_id;
        this.twitter_id = njTweet.twitter_id;
        this.tweet_id = njTweet.tweet_id;
        this.signature = njTweet.signature;
        if (njTwitter) {
            this.name = njTwitter.name;
            this.username = njTwitter.username;
            this.description = njTwitter.description;
            this.profile_image_url = njTwitter.profile_image_url;
            this.verified = njTwitter.verified;
        } else {
            this.name = null;
            this.profile_image_url = DefaultAvatarSrc;
        }
    }

    static DBKey(create_time) {
        return dbKeyCachedTweetContentById + create_time;
    }
}

function showFullTweetContent() {
    const tweetCard = this.closest('.tweet-card');
    const tweetContent = tweetCard.querySelector('.tweet-content');
    if (this.innerText === "Show more") {
        tweetContent.style.display = 'block';
        tweetContent.classList.remove('tweet-content-collapsed');
        tweetCard.style.maxHeight = 'none';
        this.innerText = "Show less";
    } else {
        tweetContent.style.display = '-webkit-box';
        tweetContent.classList.add('tweet-content-collapsed');
        tweetCard.style.maxHeight = '400px';
        this.innerText = "Show more";
    }
}

function updateLocalCachedTweetList(newIDs) {
    const storedArray = JSON.parse(localStorage.getItem(dbKeyCachedGlobalTweets)) || [];
    const int64Array = storedArray.map(BigInt);
    const mergedArray = int64Array.concat(newIDs);
    mergedArray.sort((a, b) => {
        if (a > b) return -1;
        if (a < b) return 1;
        return 0;
    });

    const topNArray = mergedArray.slice(0, maxCachedLocalTweetNo);
    localStorage.setItem(dbKeyCachedGlobalTweets, JSON.stringify(topNArray.map(Number)));
}

function parseNjTweetsFromSrv(tweetArray, refreshNewest) {
    if (tweetArray.length === 0) {
        if (!refreshNewest) {
            hasMoreTweetsToLoad = false;
        }
        return;
    }
    const newIDs = [];
    const localTweets = tweetArray.map(tweet => {
        let tw_data = TwitterBasicInfo.loadTwBasicInfo(tweet.twitter_id)
        let obj = new TweetToShowOnWeb(tweet, tw_data, null);
        localStorage.setItem(TweetToShowOnWeb.DBKey(obj.create_time), JSON.stringify(obj));
        newIDs.push(obj.create_time);
        return obj;
    });
    populateLatestTweets(localTweets, refreshNewest);
    updateLocalCachedTweetList(newIDs)
}

function loadCachedGlobalTweets() {
    const storedData = localStorage.getItem(dbKeyCachedGlobalTweets)
    if (!storedData) {
        return;
    }
    let localTweetsIds = JSON.parse(storedData);
    if (localTweetsIds.length === 0) {
        return;
    }

    const tweets = localTweetsIds
        .map(tweetID => {
            const storedTweetData = localStorage.getItem(TweetToShowOnWeb.DBKey(tweetID));
            if (!storedTweetData) {
                return null;
            }
            try {
                return JSON.parse(storedTweetData);
            } catch (error) {
                console.error('Error parsing tweet data:', error);
                return null;
            }
        })
        .filter(tweet => tweet !== null);
    if (tweets.length === 0) {
        return;
    }
    populateLatestTweets(tweets, false);
}

function loadMoreTweets() {
    if (!hasMoreTweetsToLoad) {
        return;
    }
    if (isMoreTweetsLoading) {
        return;
    }

    const documentHeight = Math.max(document.body.scrollHeight, document.documentElement.scrollHeight);
    if (window.innerHeight + window.scrollY < documentHeight - 100) {
        return;
    }
    isMoreTweetsLoading = true;
    loadGlobalLatestTweetsFromSrv(false);
}

function loadGlobalLatestTweetsFromSrv(refreshNewest) {
    let startID;
    if (refreshNewest) {
        startID = maxTweetIdCurShowed
    } else {
        startID = minTweetIdCurShowed
    }

    fetch("/globalLatestTweets?startID=" + startID + "&&isRefresh=" + refreshNewest)
        .then(response => response.json())
        .then(tweetArray => {
            parseNjTweetsFromSrv(tweetArray, refreshNewest);
        })
        .catch(err => {
            showDialog("error", err.toString());
        });
}

function populateLatestTweets(newCachedTweet, insertAtHead) {
    const tweetsPark = document.querySelector('.tweets-park');

    let maxCreateTime = BigInt(0);
    let minCreateTime = BigInt(0);
    newCachedTweet.forEach(async tweet => {

        if (!tweet.name) {
            const twitterInfo = await loadTwitterInfo(tweet.twitter_id, true)
            if (!twitterInfo) {
                tweet.name = "unknown";
                tweet.username = "unknown";
            } else {
                tweet.name = twitterInfo.name;
                tweet.username = twitterInfo.username;
            }
            localStorage.setItem(TweetToShowOnWeb.DBKey(tweet.create_time), JSON.stringify(tweet));
        }

        const timeSuffix = tweet.create_time;
        const createTime = BigInt(tweet.create_time);

        if (createTime > maxCreateTime) {
            maxCreateTime = createTime;
        }
        if (createTime < minCreateTime || minCreateTime === BigInt(0)) {
            minCreateTime = createTime;
        }
        // 创建 tweet-card 元素
        const tweetCard = document.createElement('div');
        tweetCard.classList.add('tweet-card');

        // 设置 tweet-header 内容
        const tweetHeader = document.createElement('div');
        tweetHeader.classList.add('tweet-header');

        const avatarImg = document.createElement('img');
        avatarImg.src = tweet.profile_image_url;
        avatarImg.alt = "Avatar";
        avatarImg.id = `twitterAvatar-${timeSuffix}`;
        tweetHeader.appendChild(avatarImg);

        const nameSpan = document.createElement('span');
        nameSpan.classList.add('name');
        nameSpan.id = `twitterName-${timeSuffix}`;
        nameSpan.textContent = tweet.name;
        tweetHeader.appendChild(nameSpan);

        const usernameSpan = document.createElement('span');
        usernameSpan.classList.add('username');
        usernameSpan.id = `twitterUserName-${timeSuffix}`;
        usernameSpan.textContent = "@" + tweet.username;
        tweetHeader.appendChild(usernameSpan);

        const timeSpan = document.createElement('span');
        timeSpan.classList.add('time');
        timeSpan.id = `tweet-create-time-${timeSuffix}`;
        timeSpan.textContent = formatTime(tweet.create_time);
        tweetHeader.appendChild(timeSpan);

        tweetCard.appendChild(tweetHeader);

        // 设置 tweet-content
        const tweetContent = document.createElement('div');
        tweetContent.classList.add('tweet-content', 'tweet-content-collapsed');
        tweetContent.id = `tweet-content-${timeSuffix}`;
        tweetContent.textContent = tweet.text;
        tweetCard.appendChild(tweetContent);

        // 添加 Show more 按钮
        const showMoreBtn = document.createElement('button');
        showMoreBtn.classList.add('show-more');
        showMoreBtn.textContent = "Show more";
        showMoreBtn.onclick = function () {
            showFullTweetContent.call(this);
        };

        // 设置 tweet-footer
        const tweetFooter = document.createElement('div');
        tweetFooter.classList.add('tweet-footer');

        const tweetActionDiv = document.createElement('div');
        tweetActionDiv.classList.add('tweet-action');
        tweetActionDiv.innerHTML = `
            <button>0.01 eth打赏</button>
            <span>总赏额：0.23 eth 产生彩票数：68张</span>
        `;
        tweetFooter.appendChild(tweetActionDiv);

        const tweetInfoDiv = document.createElement('div');
        tweetInfoDiv.classList.add('tweet-info');
        tweetInfoDiv.innerHTML = `Web3 ID: <span id="tweet-owner-web3-id-${timeSuffix}">${tweet.web3_id}</span>`;
        tweetFooter.appendChild(tweetInfoDiv);

        tweetCard.appendChild(tweetFooter);

        if (insertAtHead) {
            tweetsPark.insertBefore(tweetCard, tweetsPark.firstChild);
        } else {
            tweetsPark.appendChild(tweetCard);
        }
        setTimeout(() => {
            if (tweetContent.scrollHeight <= tweetContent.clientHeight) {
                showMoreBtn.style.display = 'none';
            } else {
                tweetCard.appendChild(showMoreBtn); // 只有在内容溢出时才添加按钮
            }
        }, 0);
    });

    maxTweetIdCurShowed = maxCreateTime;
    minTweetIdCurShowed = minCreateTime
}

function formatTime(createTime) {
    return new Date(createTime).toLocaleString();
}

async function postTweet() {
    try {
        const content = document.getElementById("tweets-content").value.trim();
        if (!content) {
            showDialog("tips", "content can't be empty")
            return;
        }

        const twitterID = ninjaUserObj.tw_id;
        if (!twitterID) {
            showDialog("tips", "bind your twitter first")
            return;
        }
        if (!metamaskObj) {
            window.location.href = "/signIn";
            return;
        }
        const web3Id = ninjaUserObj.eth_addr;
        const tweet = new TweetContentToPost(content, (new Date()).getTime(), web3Id, twitterID);
        const message = JSON.stringify(tweet);
        const signature = await metamaskObj.request({
            method: 'personal_sign',
            params: [message, web3Id],
        })

        const sigData = new SignDataForPost(message, signature, null)
        PostToSrvByJson("/postTweet", sigData).then(resp => {
            const refreshedTweet = JSON.parse(resp)
            document.getElementById("tweets-content").value = '';
            showDialog("success", "post success");
            parseNjTweetsFromSrv([refreshedTweet], true);
        }).catch(err => {
            console.log(err);
            showDialog("error", err.toString())
        })
    } catch (err) {
        showDialog("error", err.toString())
    }
}

