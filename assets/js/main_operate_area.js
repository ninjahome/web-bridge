function setupGameInfo(startCounter) {
    const gameArea = document.getElementById("lottery-game-info");
    gameArea.querySelector('.lottery-game-round-no').textContent = gameContractMeta.curRound;
    gameArea.querySelector('.lottery-amount').textContent = gameContractMeta.curBonus;
    document.getElementById('lottery-total-amount').textContent = gameContractMeta.totalBonus;
    document.getElementById('lottery-current-ticket-amount').textContent = gameContractMeta.ticketNo;
    document.getElementById('bonus-for-points-value').textContent = gameContractMeta.bonusForPoint;
}

function setCounterInfo(countDown, days, hours, minutes, seconds) {
    countDown.querySelector(".days").textContent = days;
    countDown.querySelector(".hours").textContent = hours;
    countDown.querySelector(".minutes").textContent = minutes;
    countDown.querySelector(".seconds").textContent = seconds;
}

function initTimerOfCounterDown() {
    let apiCounter = 0;
    resetCounter(gameContractMeta.dTime);
    startCountdown(async function (days, hours, minutes, seconds, finished) {

        const countDown = document.getElementById("lottery-count-down");
        const countDownResult = document.getElementById("countdown-result-parent");

        if (finished) {
            countDownResult.style.display = 'block';
            countDown.style.display = 'none';
        } else {
            countDown.style.display = 'flex';
            countDownResult.style.display = 'none';
            setCounterInfo(countDown, days, hours, minutes, seconds);
        }

        apiCounter += 1;
        if (apiCounter < TimeIntervalForBlockChain) {
            return;
        }

        apiCounter = 0;
        await initGameContractMeta();
        resetCounter(gameContractMeta.dTime);
        setupGameInfo();
    });
}

async function loadUserPointsInfos() {

    const userPoints = await GetToSrvByJson("/pointsForNJUsr?web3_id=" + ninjaUserObj.eth_addr.toLowerCase());
    if (!userPoints) {
        return;
    }

    document.getElementById("dessage-web3-token").innerText = userPoints.points;
    document.getElementById("point-bonus-this-round").innerText = userPoints.points;
    document.getElementById("point-bonus-annual-interest").innerText = userPoints.points;
}

function showSelfReferralCode() {
    const div = document.getElementById("referral-Code-info");
    div.style.display = "block";
}

function copyReferralCode() {
    navigator.clipboard.writeText(this.innerText).then(function () {
        const div = document.getElementById("referral-Code-info");
        div.style.display = "none";
        showTmpTips("Copy Success");
    }).catch(function (err) {
        console.error("Failed to copy the text: ", err);
        showTmpTips("Copy Failed")
    });
}