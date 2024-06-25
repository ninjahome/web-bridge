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

function calculateAnnualYield(currentPoints, totalPoints) {
    if (currentPoints <= 0 || totalPoints <= 0) {
        return 0;
    }
    const hourlyRate = (currentPoints / totalPoints) * 0.20;
    const eightHourRate = hourlyRate * 100.0;
    const dailyRate = eightHourRate * 3;
    return dailyRate * 365;
}

async function loadUserPointsInfos() {

    const userPoints = await GetToSrvByJson("/pointsForNJUsr?web3_id=" + ninjaUserObj.eth_addr.toLowerCase());
    if (!userPoints) {
        return;
    }

    document.getElementById("dessage-web3-token").innerText = userPoints.points.toFixed(2);
    document.getElementById("dessage-token-bonus").innerText = userPoints.bonus_to_win.toFixed(2);

    // console.log(userPoints);
    if (userPoints.points <= 0 || userPoints.snapshot_points <= 0 || userPoints.cur_total_points <= 0) {
        return;
    }
    const profitThisRound = userPoints.snapshot_points * pointBonusOneRound / userPoints.cur_total_points;
    const profitRate = pointBonusOneRound * 3 * 365 * 100 / userPoints.cur_total_points;

    document.getElementById("point-bonus-this-round").innerText = profitThisRound.toFixed(2);
    document.getElementById("point-bonus-annual-interest").innerText = profitRate.toFixed(2) +"%";
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