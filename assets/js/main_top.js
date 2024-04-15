function initTopDivStatus(showID, idx) {
    const parent = document.getElementById("middle-div-leaderboard");
    parent.querySelectorAll(".top-div").forEach(r => r.style.display = 'none')
    document.getElementById(showID).style.display = 'block';
    if (!idx) {
        idx = 0;
    }
    const buttons = parent.querySelectorAll(".top-topic-btn");
    buttons.forEach(r => r.classList.remove("active"));
    buttons[idx].classList.add("active");
}

const cachedTopVotedTweets = new MemCachedTweets();

async function initTopPage() {
    try {
        showWaiting("loading...");
        curScrollContentID = 1;
        initTopDivStatus("top-most-voted-tweet", 0);

        await __loadMostVotedTweets(true);
    } catch (err) {
        console.log(err);
        showDialog(DLevel.Error, err.toString());
    } finally {
        hideLoading();
    }
}

async function fillMostVotedTweet(clear, tweetArray) {

    return __fillNormalTweet(clear, "top-most-voted-tweet", tweetArray,
        "tweetTemplateForTop", "tweet-card-for-most-voted-",
        true, TweetDetailSource.MostVoted,
        function (tweetCard, tweetHeader, tweet) {
            tweetCard.querySelector('.vote-number').textContent = tweet.vote_count;
            __showVoteButton(tweetCard, tweet, function (newVote) {
                tweetCard.querySelector('.vote-number').textContent = newVote.vote_count;
            });
        });
}

async function loadOlderMostVotedTweet() {
    if (cachedTopVotedTweets.latestID === 0) {
        console.log("no need to load older data");
        return;
    }
    return __loadMostVotedTweets(false);
}

async function __loadMostVotedTweets(newest) {
    if (newest) {
        cachedTopVotedTweets.latestID = 0;
    }

    const param = new TweetQueryParam(cachedTopVotedTweets.latestID, "", []);
    const tweetArray = await PostToSrvByJson("/mostVotedTweet", param);
    if (!tweetArray || tweetArray.length === 0) {
        if (!newest) {
            cachedTopVotedTweets.moreOldTweets = false;
        }
        return;
    }

    cachedTopVotedTweets.latestID = tweetArray[tweetArray.length - 1].vote_count;
    await fillMostVotedTweet(newest, tweetArray);
}

const cachedTopVotedKolUser = new MemCachedTweets();

async function switchToTopKol() {
    curScrollContentID = 12;
    initTopDivStatus("top-hot-Kol", 1);
    await loadMostVoterOrVotedUser("top-hot-Kol", cachedTopVotedKolUser, true, false);
}

async function loadOlderMostVotedKol() {
    return loadMostVoterOrVotedUser("top-hot-Kol", cachedTopVotedKolUser, false, false);
}

async function loadMostVoterOrVotedUser(parkID, cache, newest, voter) {
    if (newest) {
        cache.latestID = 0;
    }
    const param = new TweetQueryParam(cache.latestID, "", []);
    if (voter) (
        param.voted_ids.push(1)
    )
    const userArray = await PostToSrvByJson("/mostVotedKol", param);
    if (!userArray || userArray.length === 0) {
        if (!newest) {
            cache.moreOldTweets = false;
        }
        return;
    }
    if (voter) {
        cache.latestID = userArray[userArray.length - 1].vote_count;
    } else {
        cache.latestID = userArray[userArray.length - 1].be_voted_count;
    }
    await fillMostKolOrVoterPark(parkID, newest, userArray, voter);
}

let userRankStartNo = 1;

async function fillMostKolOrVoterPark(parkID, clear, data, voter) {
    const ninjaUserPark = document.getElementById(parkID);
    if (clear) {
        ninjaUserPark.innerHTML = '';
        userRankStartNo = 1;
    }

    for (const usr of data) {
        NJUserBasicInfo.cacheNJUsrObj(usr);
        const njUsrCard = document.getElementById("team-member-card-template").cloneNode(true);
        njUsrCard.style.display = '';
        const avatarImg = njUsrCard.querySelector(".twitterAvatar");
        if (!usr.tw_id) {
            avatarImg.src = __defaultLogo;
            njUsrCard.querySelector(".twitterName").innerText = usr.eth_addr;
        } else {
            const twitterObj = await __setOnlyHeader(njUsrCard, usr.tw_id, usr.eth_addr);
            const hoverDiv = njUsrCard.querySelector(".team-member-card-header");
            hoverDiv.addEventListener('mouseenter', (event) => showHoverCard(event, twitterObj, usr.eth_addr));
            hoverDiv.addEventListener('mouseleave', () => hideHoverCard(hoverDiv));
        }

        const rankNo = njUsrCard.querySelector(".user-voted-count");
        // rankNo.innerText = userRankStartNo;
        if (userRankStartNo === 1) {
            rankNo.classList.add('vote-numberOne');
        } else if (userRankStartNo === 2) {
            rankNo.classList.add('vote-numberTwo');
        } else if (userRankStartNo === 3) {
            rankNo.classList.add('vote-numberThree');
        } else {
            rankNo.classList.add('vote-numberOther');
        }

        if (voter) {
            rankNo.innerText = usr.vote_count;
        } else {
            rankNo.innerText = usr.be_voted_count;
        }

        ninjaUserPark.appendChild(njUsrCard);

        userRankStartNo++;
    }
}

const cachedTopVoterUser = new MemCachedTweets();

async function switchToTopVoter() {
    curScrollContentID = 13;
    initTopDivStatus("top-hot-voter", 2);
    await loadMostVoterOrVotedUser("top-hot-voter", cachedTopVoterUser, true, true);
}

async function loadOlderMostVoter() {
    return loadMostVoterOrVotedUser("top-hot-voter", cachedTopVoterUser, false, true);
}