function initTopDivStatus(showID) {
    const parent = document.getElementById("middle-div-leaderboard");
    parent.querySelectorAll(".tweets-park").forEach(r => r.style.display = 'none')
    document.getElementById(showID).style.display = 'block';
}

async function switchToTopTeam() {
    curScrollContentID = 12;
    initTopDivStatus("top-hot-tweet-team");

    showWaiting("syncing from block chain");

    if (!gameContractMeta) {
        await initGameContractMeta();
    }

    try {

        changeLoadingTips("querying team detail from block chain");
        const obj = await lotteryGameContract.allTeamInfo(gameContractMeta.curRound);
        console.log(obj);
        const cachedTopTeam = [];

        for (let i = 0; i < obj.tweets.length; i++) {
            const teamObj = new TeamDetailOnBlockChain(obj.tweets[i], obj.memCounts[i], obj.voteCounts[i]);
            cachedTopTeam.push(teamObj);
        }

        changeLoadingTips("sorting result");
        cachedTopTeam.sort((a, b) => b.voteCount - a.voteCount);

        fulfillTopTeam(cachedTopTeam).then(r => {
        });

        hideLoading();
    } catch (err) {
        console.log(err);
        hideLoading();
        showDialog("error", err.toString());
    }
}

async function fulfillTopTeam(cachedTopTeam) {

    const parent_node = document.getElementById("top-hot-tweet-team");
    parent_node.innerHTML = '';

    for (const teamDetails of cachedTopTeam) {

        const team_card = document.getElementById("team-card-in-top-template").cloneNode(true);
        team_card.style.display = '';

        team_card.dataset.tweetHash = teamDetails.tweetHash;

        const tweetHeader = team_card.querySelector(".team-leader");
        team_card.id = "team-header=" + teamDetails.tweetHash;

        team_card.querySelector('.team-id-txt').innerText = teamDetails.tweetHash;

        const tweet = __globalTweetMemCacheByHash.get(teamDetails.tweetHash);
        if (!tweet) {
            const newTweet = await __queryTweetFoTeam(tweetHeader, teamDetails.tweetHash);
            if (!newTweet) {
                return;
            }
            await __setOnlyHeader(tweetHeader, newTweet.twitter_id);
            team_card.dataset.createTime = newTweet.create_time;

        } else {
            await __setOnlyHeader(tweetHeader, tweet.twitter_id);
            team_card.dataset.createTime = tweet.create_time;

        }
        team_card.dataset.detailType = TweetDetailSource.MostTeam;

        team_card.querySelector('.team-voted-count').innerText = teamDetails.voteCount;
        team_card.querySelector('.team-members-count').innerText = teamDetails.memCount;

        team_card.querySelector('.join-team').onclick = () => joinTeam(tweet, teamDetails.tweetHash, team_card);
        team_card.querySelector('.show-team-mates').onclick = () => showTeammates(teamDetails.tweetHash, team_card);
        parent_node.appendChild(team_card);
    }
}

async function joinTeam(obj, hash, team_card) {
    try {
        await voteToTheTweet(obj, async function (newVote) {
            switchToTopTeam().then(r => {
            });
        });
    } catch (err) {
        console.log(err);
    }
}

async function showTeammates(tweetHash, team_card) {
    try {
        showWaiting("syncing members from block chain");
        if (!lotteryGameContract) {
            hideLoading();
            return;
        }
        const allMates = await lotteryGameContract.teamMembers(gameContractMeta.curRound, tweetHash);
        console.log(allMates);
        if (allMates.memNo === 0) {
            hideLoading();
            showDialog("tips", "empty members")
            return;
        }
        const memberPark = team_card.querySelector('.team-member-park');
        memberPark.innerHTML = '';
        for (let i = 0; i < allMates.members.length; i++) {
            const memberCard = document.getElementById('team-members-template').cloneNode(true);
            memberCard.style.display = '';

            const ethAddr = allMates.members[i];
            loadNJUserInfoFromSrv(ethAddr, true).then(njUsr => {
                if (!njUsr) {
                    console.log("query nj user failed:", ethAddr);
                    return;
                }
                __setOnlyHeader(memberCard, njUsr.tw_id);
            });

            memberCard.querySelector('.user-voted-count').innerText = allMates.voteNos[i];
            memberPark.appendChild(memberCard);
        }

        hideLoading();
    } catch (err) {
        hideLoading();
        checkMetamaskErr(err);
    }
}

function __queryAndFillTeamHeader(tweetHeader, tweetHash) {

    const response = GetToSrvByJson("/queryTwBasicByTweetHash?tweet_hash=" + tweetHash);

    response.then(r => {
        r.text().then(async obj => {
            const twObj = TwitterBasicInfo.cacheTwBasicInfo(obj);
            await __setOnlyHeader(tweetHeader, twObj.twitter_id);
        }).catch(err => {
            console.log(err);
        });
    });
}

async function __queryTweetFoTeam(tweetHeader, tweetHash) {
    try {
        const response = await GetToSrvByJson("/queryTweetByHash?tweet_hash=" + tweetHash);

        if (!response.ok) {
            console.log("query twitter basic info failed")
            return null;
        }

        const text = await response.text();
        const obj = JSON.parse(text);
        __globalTweetMemCacheByHash.set(tweetHash, obj);
        __globalTweetMemCache.set(obj.create_time, obj);
        return obj;
    } catch (err) {
        showDialog("error", "query tweet failed:" + err.toString());
        return null;
    }
}


const cachedTopVotedTweets = new MemCachedTweets();

async function initTopPage() {
    curScrollContentID = 1;
    initTopDivStatus("top-most-voted-tweet");
    await __loadMostVotedTweets(true);
}

async function fillMostVotedTweet(clear, tweetArray) {

    return __fillNormalTweet(clear, "top-most-voted-tweet", tweetArray,
        "tweetTemplateForTop", "tweet-card-for-most-voted-",
        true, TweetDetailSource.MostVoted,
        function (tweetCard, tweetHeader, tweet) {
            tweetCard.querySelector('.vote-number').textContent = tweet.vote_count;
            __showVoteButton(tweetCard, tweet);
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
    const resp = await PostToSrvByJson("/mostVotedTweet", param);
    if (!resp) {
        if (!newest) {
            cachedTopVotedTweets.moreOldTweets = false;
        }
        return;
    }
    const tweetArray = JSON.parse(resp);
    if (tweetArray.length === 0) {
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
    curScrollContentID = 13;
    initTopDivStatus("top-hot-Kol");
    await __loadMostVotedKolUserInfo("top-hot-Kol", cachedTopVotedKolUser, true, false);
}

async function loadOlderMostVotedKol() {
    if (cachedTopVotedKolUser.latestID === 0) {
        console.log("no need to load older data");
        return;
    }
    return __loadMostVotedKolUserInfo("top-hot-Kol", cachedTopVotedKolUser, false, false);
}

async function __loadMostVotedKolUserInfo(parkID, cache, newest, voter) {
    if (newest) {
        cache.latestID = 0;
    }
    const param = new TweetQueryParam(cache.latestID, "", []);
    if (voter) (
        param.voted_ids.push(1)
    )
    const resp = await PostToSrvByJson("/mostVotedKol", param);
    if (!resp) {
        return;
    }

    const userArray = JSON.parse(resp);
    if (userArray.length === 0) {
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
        NJUserBasicInfo.cacheNJUsrObj(usr).then(r => {
        });
        const njUsrCard = document.getElementById("ninjaUserCardTemplate").cloneNode(true);
        njUsrCard.style.display = '';

        if (!usr.tw_id) {
            njUsrCard.querySelector(".twitterAvatar").src = __defaultLogo;
            njUsrCard.querySelector(".twitterName").innerText = usr.eth_addr;
        } else {
            await __setOnlyHeader(njUsrCard, usr.tw_id);
        }

        njUsrCard.querySelector(".voteOrVotedRangeNo").innerText = userRankStartNo;
        if (voter) {
            njUsrCard.querySelector(".voteOrVotedNos").innerText = usr.vote_count;
        } else {
            njUsrCard.querySelector(".voteOrVotedNos").innerText = usr.be_voted_count;
        }
        ninjaUserPark.appendChild(njUsrCard);

        userRankStartNo++;
    }
}

const cachedTopVoterUser = new MemCachedTweets();

async function switchToTopVoter() {
    curScrollContentID = 14;
    initTopDivStatus("top-hot-voter");
    await __loadMostVotedKolUserInfo("top-hot-voter", cachedTopVoterUser, true, true);
}

async function loadOlderMostVoter() {
    if (cachedTopVoterUser.latestID === 0) {
        console.log("no need to load older data");
        return;
    }
    return __loadMostVotedKolUserInfo("top-hot-voter", cachedTopVoterUser, false, true);
}