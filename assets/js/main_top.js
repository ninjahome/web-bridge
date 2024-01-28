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

async function switchToTopTeam() {
    try {
        curScrollContentID = 12;
        initTopDivStatus("top-hot-tweet-team", 1);

        showWaiting("syncing from block chain");
        if (!gameContractMeta) {
            await initGameContractMeta();
        }

        changeLoadingTips("querying team detail from block chain");
        const obj = await lotteryGameContract.allTeamInfo(gameContractMeta.curRound);
        const cachedTopTeam = [];

        for (let i = 0; i < obj.tweets.length; i++) {
            const teamObj = new TeamDetailOnBlockChain(obj.tweets[i], obj.memCounts[i], obj.voteCounts[i]);
            cachedTopTeam.push(teamObj);
        }

        changeLoadingTips("sorting result");
        cachedTopTeam.sort((a, b) => b.voteCount - a.voteCount);

        await fulfillTopTeam(cachedTopTeam);

    } catch (err) {
        console.log(err);
        showDialog("error", err.toString());
    } finally {
        hideLoading();
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


        team_card.querySelector('.team-voted-count').innerText = teamDetails.voteCount;
        team_card.querySelector('.team-members-count').innerText = teamDetails.memCount;

        team_card.querySelector('.join-team').onclick = () => joinTeam(tweet, teamDetails.tweetHash, team_card);
        team_card.querySelector('.show-team-mates').onclick = () => showTeammates(teamDetails.tweetHash, team_card);

        let tweet = __globalTweetMemCacheByHash.get(teamDetails.tweetHash);
        if (!tweet) {
            tweet = await __queryTweetFoTeam(tweetHeader, teamDetails.tweetHash);
        }

        if (!tweet) {
            tweetHeader.querySelector('.team-id').style.cursor = 'default';
            tweetHeader.querySelector('.twitterAvatar').src = __defaultLogo;
            tweetHeader.querySelector('.twitterName').textContent = "非本系统推文";
        } else {
            await __setOnlyHeader(tweetHeader, tweet.twitter_id);
            team_card.dataset.createTime = tweet.create_time;
            team_card.querySelector('.team-id-txt').onclick = () =>
                showTweetDetail('top-hot-tweet-team', tweet);
        }

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

    const memberPark = team_card.querySelector('.team-members')
    memberPark.innerHTML = '';

    const isShowing = memberPark.style.display === 'block';
    if (isShowing) {
        memberPark.style.display = 'none';
        return;
    }
    memberPark.style.display = 'block';
    try {
        showWaiting("syncing members from block chain");
        if (!lotteryGameContract) {
            return;
        }
        const allMates = await lotteryGameContract.teamMembers(gameContractMeta.curRound, tweetHash);
        if (allMates.memNo === 0) {
            memberPark.style.display = 'none';
            showDialog("tips", "empty members")
            return;
        }

        for (let i = 0; i < allMates.members.length; i++) {
            const memberCard = document.getElementById('team-member-card-template').cloneNode(true);
            memberCard.style.display = '';
            memberCard.id='';
            memberCard.querySelector('.user-voted-count').innerText = allMates.voteNos[i];
            memberCard.querySelector(".team-members-number").innerText = "" + i;

            const ethAddr = allMates.members[i];
            const njUsr = await loadNJUserInfoFromSrv(ethAddr, true);
            if (!njUsr.tw_id) {
                memberCard.querySelector(".team-membersAvatar").src = __defaultLogo;
                memberCard.querySelector(".team-membersName").innerText = ethAddr;
            } else {
                await __setOnlyHeader(memberCard, njUsr.tw_id);
            }

            memberPark.appendChild(memberCard);
        }
    } catch (err) {
        checkMetamaskErr(err);
    } finally {
        hideLoading();
    }
}

async function __queryTweetFoTeam(tweetHeader, tweetHash) {
    try {
        const obj = await GetToSrvByJson("/queryTweetByHash?tweet_hash=" + tweetHash);
        if (!obj) {
            return null;
        }
        __globalTweetMemCacheByHash.set(tweetHash, obj);
        __globalTweetMemCache.set(obj.create_time, obj);
        return obj;
    } catch (err) {
        console.log(err);
        return null;
    }
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
    curScrollContentID = 13;
    initTopDivStatus("top-hot-Kol", 2);
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
            const twitterObj = await __setOnlyHeader(njUsrCard, usr.tw_id);
            const hoverDiv = njUsrCard.querySelector(".team-member-card-header");
            hoverDiv.addEventListener('mouseenter', (event) => showHoverCard(event, twitterObj, usr.eth_addr));
            hoverDiv.addEventListener('mouseleave', () => hideHoverCard(hoverDiv));
        }

        const rankNo = njUsrCard.querySelector(".team-members-number");
        rankNo.innerText = userRankStartNo;
        if (userRankStartNo === 1) {
            rankNo.classList.add('team-members-topOne');
        } else if (userRankStartNo === 2) {
            rankNo.classList.add('team-members-topTwo');
        } else if (userRankStartNo === 3) {
            rankNo.classList.add('team-members-topThree');
        } else {
            rankNo.classList.add('team-members-topOther');
        }

        if (voter) {
            njUsrCard.querySelector(".user-voted-count").innerText = usr.vote_count;
        } else {
            njUsrCard.querySelector(".user-voted-count").innerText = usr.be_voted_count;
        }

        ninjaUserPark.appendChild(njUsrCard);

        userRankStartNo++;
    }
}

const cachedTopVoterUser = new MemCachedTweets();

async function switchToTopVoter() {
    curScrollContentID = 14;
    initTopDivStatus("top-hot-voter", 3);
    await __loadMostVotedKolUserInfo("top-hot-voter", cachedTopVoterUser, true, true);
}

async function loadOlderMostVoter() {
    if (cachedTopVoterUser.latestID === 0) {
        console.log("no need to load older data");
        return;
    }
    return __loadMostVotedKolUserInfo("top-hot-voter", cachedTopVoterUser, false, true);
}