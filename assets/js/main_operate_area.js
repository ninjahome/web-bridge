async function setupGameInfo() {

    const gameArea = document.getElementById("lottery-game-info");

    gameArea.querySelector('.lottery-game-round-no').textContent = gameContractMeta.curRound;
    gameArea.querySelector('.lottery-amount').textContent = gameContractMeta.curBonus;
    document.getElementById('lottery-total-amount').textContent = gameContractMeta.totalBonus;
    document.getElementById('lottery-user-ticket-amount').textContent = gameContractMeta.userTickNo;
    document.getElementById('lottery-current-ticket-amount').textContent = gameContractMeta.ticketNo;
    document.getElementById('lottery-current-game-hash').textContent = gameContractMeta.dHash;

    startCountdown(gameContractMeta.dTime, function (txt) {
        document.getElementById("lottery-timer").innerText = txt;
    });
}
