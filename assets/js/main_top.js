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

        const tweetHeader = team_card.querySelector(".team-leader");

        team_card.id = "team-header=" + teamDetails.tweetHash;
        // team_card.dataset.tweetHash = teamDetails.tweetHash;

        team_card.querySelector('.team-id-txt').innerText = teamDetails.tweetHash;
        const tweet = __globalTweetMemCacheByHash.get(teamDetails.tweetHash);
        if (!tweet) {
            __queryAndFillTeamHeader(tweetHeader, teamDetails.tweetHash);
        } else {
            await __setOnlyHeader(tweetHeader, tweet.twitter_id);
        }

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
        const allMates = await lotteryGameContract.teamMembers(gameContractMeta.curRound, tweetHash);
        console.log(allMates);
        if (allMates.memNo === 0) {
            hideLoading();
            showDialog("tips", "empty members")
            return;
        }
       const memberPark = team_card.querySelector('.team-member-park');
        memberPark.innerHTML='';
        for (let i = 0; i < allMates.members.length; i++) {
            const memberCard = document.getElementById('team-members-template').cloneNode(true);
            memberCard.style.display='';

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

function switchToTopKol() {
    curScrollContentID = 13;
    initTopDivStatus("top-hot-Kol");
}

function switchToTopVoter() {
    curScrollContentID = 14;
    initTopDivStatus("top-hot-voter");
}

const cachedTopVotedTweets = new MemCachedTweets();

function initTopPage() {
    curScrollContentID = 1;
    initTopDivStatus("top-most-voted-tweet");
}