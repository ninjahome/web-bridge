async function setupGameInfo(startCounter) {

    const gameArea = document.getElementById("lottery-game-info");

    gameArea.querySelector('.lottery-game-round-no').textContent = gameContractMeta.curRound;
    gameArea.querySelector('.lottery-amount').textContent = gameContractMeta.curBonus;
    document.getElementById('lottery-total-amount').textContent = gameContractMeta.totalBonus;
    document.getElementById('lottery-user-ticket-amount').textContent = gameContractMeta.userTickNo;
    document.getElementById('lottery-current-ticket-amount').textContent = gameContractMeta.ticketNo;

    if (!startCounter) {
        return;
    }

    let apiCounter = 0;
    startCountdown(gameContractMeta.dTime, function (days,hours,minutes,seconds, finished) {
        const countDown = document.getElementById("lottery-count-down");
        const countDownResult = document.getElementById("countdown-result");
        if (finished) {
            countDownResult.style.display ='';
            countDown.style.display = 'none';
            initGameContractMeta().then(r => {
                setupGameInfo(true);
            });
            return;
        }

        apiCounter += 1;
        countDown.style.display = '';
        countDownResult.style.display ='none';
        countDown.querySelector(".days").textContent = days;
        countDown.querySelector(".hours").textContent = hours;
        countDown.querySelector(".minutes").textContent = minutes;
        countDown.querySelector(".seconds").textContent = seconds;

        if (apiCounter >= 20) {
            apiCounter = 0;
            initGameContractMeta().then(r => {
                setupGameInfo(false);
            });
        }
    });
}
