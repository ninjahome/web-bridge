function initTopDivStatus(showID) {
    const parent = document.getElementById("middle-div-leaderboard");
    parent.querySelectorAll(".tweets-park").forEach(r => r.style.display = 'none')
    document.getElementById(showID).style.display = 'block';
}

async function switchToTopTeam() {
    curScrollContentID = 12;
    initTopDivStatus("top-hot-tweet-team");

    if (!gameContractMeta) {
        showWaiting("syncing from block chain", 3);
    } else {
        showWaiting("syncing from block chain");
    }

    try {
        const allTeams = await lotteryGameContract.teamListOfRound(gameContractMeta.curRound);
        if (allTeams.length === 0) {
            console.log("no data right now");
            hideLoading();
            return;
        }
        changeLoadingTips("querying team detail from block chain");
        console.log(allTeams);
        const obj = await lotteryGameContract.allTeamInfo(gameContractMeta.curRound, allTeams);
        console.log(obj);
        const cachedTopTeam = [];

        for (let i = 0; i < allTeams.length; i++) {
            const teamObj = new TeamDetailOnBlockChain(allTeams[i], obj.memCounts[i], obj.voteCounts[i]);
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
        team_card.dataset.tweetHash = teamDetails.tweetHash;

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
        team_card.querySelector('.show-team-mates').onclick = () => showTeammates(teamDetails.tweetHash);
        parent_node.appendChild(team_card);
    }
}

async function joinTeam(obj, hash, team_card) {
    try {
        await voteToTheTweet(obj,async function (newVote) {
            switchToTopTeam().then(r=>{});
        });
    } catch (err) {
        console.log(err);
    }
}

function showTeammates(hash) {
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