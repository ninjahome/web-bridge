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
        changeLoadingTips("querying team detail from block chain")
        console.log(allTeams);
        for (const teamHash of allTeams) {
            // queryNjBasicByID
            const obj = lotteryGameContract.tweetTeamMap(gameContractMeta.curRound, teamHash)
        }


    } catch (err) {
        console.log(err);
        showDialog("error", err.toString());
    }
}

function fulfillTopTeam(){

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