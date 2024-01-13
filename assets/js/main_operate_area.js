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
    startCountdown(gameContractMeta.dTime, function (txt, finished) {
        if (finished) {
            initGameContractMeta().then(r => {
                setupGameInfo(true);
            });
            return;
        }

        apiCounter += 1;
        document.getElementById("lottery-timer").innerText = txt;

        if (apiCounter >= 20) {
            apiCounter = 0;
            initGameContractMeta().then(r => {
                setupGameInfo(false);
            });
        }
    });
}
