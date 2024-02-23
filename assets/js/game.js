let lotteryGameContract = null;
let gameSettings = null;
let personalData = null

class GameSettings {
    constructor(roundNo, bonus, voteNo, tickPrice, tickPriceInEth, isOpen) {
        this.roundNo = roundNo;
        this.bonus = bonus;
        this.tickPrice = tickPrice;
        this.tickPriceInEth = tickPriceInEth;
        this.isOpen = isOpen;
        this.voteNo = voteNo;
    }
}

class GameRoundInfo {
    constructor(hash, dTime, winner, winTicketID, curBonus, random, winnerBonus) {
        this.hash = hash;
        this.dTime = dTime;
        this.winner = winner;
        this.winTicketID = winTicketID;
        this.curBonus = curBonus;
        this.random = random;
        this.bonusForWinner = winnerBonus;
    }

    static fromBlockChainObj(obj) {
        const curBonusInEth = ethers.utils.formatUnits(obj.bonus, 'ether');
        const dTime = obj.discoverTime.toNumber() * 1000;
        const bonusForWinner = ethers.utils.formatUnits(obj.bonusForWinner, 'ether');

        return new GameRoundInfo(obj.randomHash, dTime, obj.winner, obj.winTicketID,
            curBonusInEth, obj.randomVal, bonusForWinner);
    }
}

class WinTeamInfo {
    constructor(round_no, win_team, bonus, member_addr, member_vote_no, total_vote_no, total_mem_no) {
        this.round_no = round_no;
        this.win_team = win_team;
        this.bonus = bonus;
        this.member_addr = member_addr;
        this.member_vote_no = member_vote_no;
        this.total_vote_no = total_vote_no;
        this.total_mem_no = total_mem_no;
    }
}

class PersonalData {
    constructor(balance, tickets, map) {
        this.balance = balance;
        this.tickets = tickets;
        this.tickMap = map;
    }
}

async function initGamePage() {
    await checkMetaMaskEnvironment(initGameContract);
    const address = __globalContractConf.get(__globalTargetChainNetworkID).gameLottery;
    document.querySelector('.contract-address-value').textContent = address;
    syncWinnerHistoryData().then(r => {
    });
}

function showContractUrl() {
    const address = __globalContractConf.get(__globalTargetChainNetworkID).gameLottery;
    const url = __globalMetaMaskNetworkParam.get(__globalTargetChainNetworkID).blockExplorerUrls;
    window.open(url + "/address/" + address);
}

async function initGameContract(provider) {
    if (!provider) {
        lotteryGameContract = null;
        return;
    }

    const signer = provider.getSigner(ninjaUserObj.eth_addr);
    const conf = __globalContractConf.get(__globalTargetChainNetworkID);
    lotteryGameContract = new ethers.Contract(conf.gameLottery, gameContractABI, signer);

    __loadPageData();
}

function __loadPageData() {

    loadGameSettings().then(async r => {
        setupSystemData();

        setupCurrentRoundData().then(r => {
        });

        loadPersonalMeta().then(r => {
            setupPersonalData();
        });
    });
}

async function loadGameSettings() {

    showWaiting("syncing system data from block chain")
    try {
        const [currentRoundNo, totalBonus, voteNo, tickPriceForOuter, isOpenToOuter] =
            await lotteryGameContract.systemSettings();

        const totalBonusInEth = ethers.utils.formatUnits(totalBonus, 'ether');
        const tickPriceInEth = ethers.utils.formatUnits(tickPriceForOuter, 'ether');

        gameSettings = new GameSettings(currentRoundNo, totalBonusInEth, voteNo,
            tickPriceForOuter, tickPriceInEth, isOpenToOuter);

    } catch (err) {
        console.log(err);
        showDialog(DLevel.Warning, "load game settings from block chain failed");
    } finally {
        hideLoading()
    }
}

async function loadPersonalMeta() {
    try {
        const balance = await lotteryGameContract.balance(ninjaUserObj.eth_addr);
        const balanceInEth = ethers.utils.formatUnits(balance, 'ether');

        const tickets = await lotteryGameContract.tickList(gameSettings.roundNo, ninjaUserObj.eth_addr);
        if (tickets.length === 0) {
            personalData = new PersonalData(balanceInEth, [], []);
            return;
        }

        const mapTickets = new Map();
        for (let i = 0; i < tickets.length; i++) {
            const tickId = tickets[i];
            mapTickets.set(tickId, true);
        }

        personalData = new PersonalData(balanceInEth, tickets, mapTickets);
    } catch (err) {
        console.log(err);
        showDialog(DLevel.Warning, "load personal data from block chain failed")
    }
}

function setupSystemData() {
    document.querySelector(".history-total-bonus").textContent = gameSettings.bonus;
    document.querySelector(".round-number").textContent = gameSettings.roundNo;
    document.querySelector(".ticket-price-for-outer-user").textContent = gameSettings.tickPriceInEth;
}

async function setupCurrentRoundData() {
    try {

        const gameInfo = await lotteryGameContract.gameInfoRecord(gameSettings.roundNo);
        const data = GameRoundInfo.fromBlockChainObj(gameInfo);

        document.getElementById("prize-pool-bonus-val").textContent = data.curBonus;
        document.getElementById("prize-pool-random-hash").textContent = data.hash;
        document.getElementById("prize-pool-tick-no").textContent = gameSettings.voteNo;

        const elem = document.getElementById("prize-pool-discover-time");
        startCountdown(data.dTime, function (days, hours, minutes, seconds, finished) {
            if (finished) {
                elem.innerText = i18next.t('game-status-with-draw');
                return;
            }

            elem.innerText = days + i18next.t('game-status-day') + hours + i18next.t('game-status-hour') + minutes
                + i18next.t('game-status-minute') + seconds + i18next.t('game-status-second');
        });
    } catch (err) {
        console.log(err);
        showDialog(DLevel.Warning, "load game data from block chain failed");
    }
}

function setupPersonalData() {
    document.getElementById("personal-balance-val").textContent = personalData.balance;
    document.getElementById("personal-ticket-no-val").textContent = personalData.tickets.length;
}

function showPersonalTicket() {
    if (!personalData || personalData.tickets.length === 0) {
        return;
    }
    const ticketsDiv = document.querySelector('.user-tickets');
    const isShowing = ticketsDiv.style.display === 'block';
    ticketsDiv.style.display = isShowing ? 'none' : 'block';

    if (isShowing) {
        return;
    }

    const tableBody = ticketsDiv.querySelector(".tickets-num");
    tableBody.innerHTML = '';
    let counter = 0;
    let row = tableBody.insertRow();

    for (let i = 0; i < personalData.tickets.length; i++) {

        if (counter % 7 === 0 && counter !== 0) {
            row = tableBody.insertRow();
        }
        let cell = row.insertCell();
        cell.innerHTML = personalData.tickets[i];
        counter++;
    }
}

function hideOneTeamDetails() {
    const teamDetailDiv = document.querySelector('.team-detail-for-one');
    teamDetailDiv.style.display = 'none';
}

function showGameRule(className) {
    const gameRuleDiv = document.querySelector(className);
    gameRuleDiv.style.display = gameRuleDiv.style.display === 'none' ? 'block' : 'none';
}

async function showOneRoundGameInfo() {

    try {
        const roundNo = document.getElementById('round-input').value;
        if (!roundNo) {
            showDialog(DLevel.Tips, "invalid round no");
            return;
        }
        document.getElementById('round-input').value = '';
        const queryNo = Number(roundNo);
        if (queryNo > gameSettings.roundNo) {
            showDialog(DLevel.Tips, "bigger than current round no:" + gameSettings.roundNo);
            return;
        }
        showWaiting("syncing from block chain");
        const obj = await lotteryGameContract.gameInfoRecord(queryNo);
        const cardDiv = document.querySelector('.round-history');

        fullFillGameCard(obj, cardDiv, true);

    } catch (err) {
        showDialog(DLevel.Error, "failed to query form block chain:" + err.toString());
    } finally {
        hideLoading();
    }
}

function hideInfoOfThisRound() {
    const cardDiv = document.querySelector('.round-history');
    cardDiv.style.display = 'none';
}

function fullFillGameCard(obj, cardDiv, showHideBtn) {
    cardDiv.style.display = 'block';
    cardDiv.querySelector('.one-round-bonus-val').textContent = ethers.utils.formatUnits(obj.bonus, 'ether');

    const dTime = new Date(obj.discoverTime * 1000);
    cardDiv.querySelector('.one-round-discover-val').textContent = dTime.toString();

    cardDiv.querySelector('.history-game-random').textContent = obj.randomVal;
    cardDiv.querySelector('.history-game-random-hash').textContent = obj.randomHash;
    cardDiv.querySelector('.history-game-winner-address').textContent = obj.winner;
    cardDiv.querySelector('.history-game-winner-team').textContent = obj.winTeam;
    cardDiv.querySelector('.history-game-winner-ticket').textContent = obj.winTicketID;
    if (showHideBtn) {
        cardDiv.querySelector('.load-history-btn').style.display = 'block';
    }else{
        cardDiv.querySelector('.load-history-btn').style.display = 'none';
    }
}

let __toRoundNo = 0;

async function loadHistoryData() {
    const parentDiv = document.querySelector('.history-data-list');
    const moreBtn = document.querySelector('.history-data-list-more-btn');
    const isShowing = parentDiv.style.display === 'block';

    if (isShowing){
        this.textContent = i18next.t('all-history-query-btn');
        parentDiv.style.display = 'none';
        parentDiv.innerHTML = '';
        moreBtn.style.display = 'none';
        return;
    }
    this.textContent = i18next.t('hide-history-data-btn');

    moreBtn.style.display = 'block';
    __toRoundNo = gameSettings.roundNo - 1;

    parentDiv.style.display = 'block';
    parentDiv.innerHTML = '';

    await __loadHistoryData(parentDiv);
}

async function moreHistoryData() {
    const parentDiv = document.querySelector('.history-data-list');
    parentDiv.style.display = 'block';
    await __loadHistoryData(parentDiv);
}

async function __loadHistoryData(parentDiv) {

    try {
        if (__toRoundNo < 0) {
            showDialog(DLevel.Tips, "no more data");
            return;
        }

        const from = __toRoundNo > 20 ? (__toRoundNo - 20) : 0;
        showWaiting("syncing history game data from block chain")

        const obj = await lotteryGameContract.historyRoundInfo(from, __toRoundNo);
        let reversedArray = obj.slice().reverse();

        for (const gameInfo of reversedArray) {
            const div = document.getElementById('history-data-one-round-template').cloneNode(true);
            fullFillGameCard(gameInfo, div);
            parentDiv.appendChild(div);
        }

        __toRoundNo = from - 1;
        if (__toRoundNo < 0) {
            const moreBtn = document.querySelector('.history-data-list-more-btn');
            moreBtn.style.display = 'none';
        }

    } catch (err) {
        showDialog(DLevel.Warning, "load history data err:" + err.toString());
    } finally {
        hideLoading();
    }
}

async function buyTicket() {
    if (!gameSettings) {
        await loadGameSettings();
    }

    if (!gameSettings.isOpen) {
        showDialog(DLevel.Tips, "not open for personal user");
        return;
    }

    openVoteModal(procTicketPayment);
}

async function procTicketPayment(no, ifShare) {
    if (no === 0) {
        showDialog(DLevel.Tips, "on ticket at lest")
        return;
    }

    const val = gameSettings.tickPrice.mul(no);
    try {
        showWaiting("prepare to pay")
        const txResponse = await lotteryGameContract.buyTicketFromOuter(no, {value: val});

        changeLoadingTips("packaging:" + txResponse.hash);
        const txReceipt = await txResponse.wait();

        if (!txReceipt.status) {
            showDialog(DLevel.Error, "transaction " + "failed");
            return;
        }
        showDialog(DLevel.Success, "buy success");

        if (ifShare) {
            __shareVoteToTweet(0, no, i18next.t('voter-slogan')).then(r => {
                console.log("share to twitter success")
            });
        }
        __loadPageData();

    } catch (err) {
        checkMetamaskErr(err);
    } finally {
        hideLoading();
    }
}

let cachedWinnerHistoryData = []
let cachedWinTeamHistoryData = []

async function syncWinnerHistoryData() {
    const data = await GetToSrvByJson('/queryWinHistory');
    if (!data) {
        return;
    }
    cachedWinnerHistoryData = data;
    document.querySelector('.personal-winning-count').textContent = "" + cachedWinnerHistoryData.length;
}

function showUserWinHistory() {
    if (cachedWinnerHistoryData.length === 0) {
        return;
    }
    const historyDiv = document.querySelector('.winner-history-list');
    const isShowing = historyDiv.style.display === 'block';
    historyDiv.style.display = isShowing ? 'none' : 'block';
    historyDiv.innerHTML = '';
    if (isShowing) {
        return;
    }

    try {
        for (const obj of cachedWinnerHistoryData) {
            const winnerCard = document.getElementById("winning-history-template").cloneNode(true);
            winnerCard.style.display = 'block';
            winnerCard.id = null;

            winnerCard.querySelector('.one-round-bonus-val').textContent = obj.bonus;
            winnerCard.querySelector('.one-round-ticket-id').textContent = obj.win_ticket_id;
            winnerCard.querySelector('.one-round-round-val').textContent = obj.round_no;
            winnerCard.querySelector('.one-round-discover-val').textContent = formatTime(obj.discover_time);
            winnerCard.querySelector('.one-round-bonus-for-me').textContent = obj.bonus;

            historyDiv.appendChild(winnerCard);
        }
    } catch (err) {
        console.log(err)
        showDialog(DLevel.Warning, "load err:" + err.toString())
    }
    syncWinnerHistoryData().then(r => {

    })
}

async function withdrawBonus() {
    try {
        showWaiting("calling to block chain");

        const txResponse = await lotteryGameContract.withdraw(0, true);

        changeLoadingTips("packaging:" + txResponse.hash);
        const txReceipt = await txResponse.wait();

        if (!txReceipt.status) {
            showDialog(DLevel.Error, "transaction " + "failed");
            return;
        }

        showDialog(DLevel.Success, "withdraw success");

        loadPersonalMeta().then(r => {
            setupPersonalData();
        });

    } catch (err) {
        checkMetamaskErr(err);
    } finally {
        hideLoading();
    }
}

function incomeWithdrawHistory() {
    __incomeWithdrawHistory(ninjaUserObj.eth_addr);
}
