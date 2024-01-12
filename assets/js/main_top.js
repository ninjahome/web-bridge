function initTopDivStatus(showID) {
    const parent = document.getElementById("middle-div-leaderboard");
    parent.querySelectorAll(".tweets-park").forEach(r => r.style.display = 'none')
    document.getElementById(showID).style.display = 'block';
}

async function switchToTopTeam() {
    curScrollContentID = 12;
    initTopDivStatus("top-hot-lottery-team");
    if (!gameContractMeta){
        showDialog("tips","block chain is syncing, please try later");
        return;
    }

    try {
        showWaiting("querying team list from blockchain")
        const allTeams = await lotteryGameContract.teamListOfRound(gameContractMeta.curRound);
        if (allTeams.length === 0){
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
            const teamObj = new TeamDetailOnBlockChain(allTeams[i],obj.memCounts[i],obj.voteCounts[i]);
            cachedTopTeam.push(teamObj);
        }
        changeLoadingTips("sorting result");
        cachedTopTeam.sort((a, b) => b.voteCount - a.voteCount);
        fulfillTopTeam(cachedTopTeam);
        hideLoading();
    } catch (err) {
        console.log(err);
        hideLoading();
        showDialog("error", err.toString());
    }
}

function fulfillTopTeam(cachedTopTeam){
    const  team_card = document.getElementById("team-card-in-top-template").cloneNode(true);
    team_card.display = '';

    const parent_node = document.getElementById("top-hot-tweet-team");
}


function switchToTopKol(){
    curScrollContentID = 13;
    initTopDivStatus("top-hot-Kol");
}

function switchToTopVoter(){
    curScrollContentID = 14;
    initTopDivStatus("top-hot-voter");
}

const cachedTopVotedTweets = new MemCachedTweets();
function initTopPage(){
    curScrollContentID = 1;
    initTopDivStatus("top-hot-tweet-team");
}