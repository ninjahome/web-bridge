let lotteryGameContract = null;
let gameSettings = null;
let currentRoundData = null;
let personalData = null

class GameSettings {
    constructor(roundNo, bonus, tickPrice, tickPriceInEth, isOpen) {
        this.roundNo = roundNo;
        this.bonus = bonus;
        this.tickPrice = tickPrice;
        this.tickPriceInEth = tickPriceInEth;
        this.isOpen = isOpen;
    }
}

class GameRoundInfo {
    constructor(hash, dTime, winner, winTeam, winTicketID, curBonus, random) {
        this.hash = hash;
        this.dTime = dTime;
        this.winner = winner;
        this.winTeam = winTeam;
        this.winTicketID = winTicketID;
        this.curBonus = curBonus;
        this.random = random;
        this.TeamCount = 0;
        this.TickCount = 0;
    }

    static fromBlockChainObj(obj) {
        const curBonusInEth = ethers.utils.formatUnits(obj.bonus, 'ether');
        const dTime = obj.discoverTime.toNumber() * 1000;
        return new GameRoundInfo(obj.randomHash, dTime, obj.winner,
            obj.winTeam, obj.winTicketID, curBonusInEth, obj.randomVal);
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
    constructor(balance, tickets, teams, map, hasIndividual) {
        this.balance = balance;
        this.tickets = tickets;
        this.teams = teams;
        this.tickMap = map;
        this.hasIndividual = hasIndividual;
    }
}

async function initGamePage() {
    await checkMetaMaskEnvironment(initGameContract);
    const address = __globalContractConf.get(__globalTargetChainNetworkID).gameLottery;
    document.querySelector('.contract-address-value').textContent = address;
    syncWinnerHistoryData().then(r => {
    });
    syncWinTeamHistoryData().then(r => {
    })
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

        loadCurrentRoundMeta().then(r => {
            setupCurrentRoundData();
        });

        loadPersonalMeta().then(r => {
            setupPersonalData();
        });
    });
}

async function loadGameSettings() {

    showWaiting("syncing system data from block chain")
    try {
        const [currentRoundNo, totalBonus, tickPriceForOuter, isOpenToOuter] =
            await lotteryGameContract.systemSettings();

        const totalBonusInEth = ethers.utils.formatUnits(totalBonus, 'ether');
        const tickPriceInEth = ethers.utils.formatUnits(tickPriceForOuter, 'ether');

        gameSettings = new GameSettings(currentRoundNo, totalBonusInEth,
            tickPriceForOuter, tickPriceInEth, isOpenToOuter);

    } catch (err) {
        console.log(err);
        showDialog(DLevel.Warning, "load game settings from block chain failed");
    } finally {
        hideLoading()
    }
}

async function loadCurrentRoundMeta() {
    try {

        const gameInfo = await lotteryGameContract.gameInfoRecord(gameSettings.roundNo);

        currentRoundData = GameRoundInfo.fromBlockChainObj(gameInfo);

        const [teamNo, voteNo] = await lotteryGameContract.allTeamInfoNo(gameSettings.roundNo);

        currentRoundData.TeamCount = teamNo;
        currentRoundData.TickCount = voteNo;

    } catch (err) {
        console.log(err);
        showDialog(DLevel.Warning, "load game data from block chain failed");
    }
}

async function loadPersonalMeta() {
    try {
        const balance = await lotteryGameContract.balance(ninjaUserObj.eth_addr);
        const balanceInEth = ethers.utils.formatUnits(balance, 'ether');

        const obj = await lotteryGameContract.tickList(gameSettings.roundNo, ninjaUserObj.eth_addr);
        if (obj[0].length === 0) {
            personalData = new PersonalData(balanceInEth, [], [], null, false);
            return;
        }

        const mapTickets = new Map();
        const mapTeams = new Map();
        for (let i = 0; i < obj[0].length; i++) {
            const tickId = obj[0][i];
            const teamHash = obj[1][i];
            mapTickets.set(tickId, teamHash);
            mapTeams.set(teamHash, true);
        }

        personalData = new PersonalData(balanceInEth, Array.from(mapTickets.keys()),
            Array.from(mapTeams.keys()), mapTickets, mapTeams.has(__noTeamID));
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

function setupCurrentRoundData() {

    document.getElementById("prize-pool-bonus-val").textContent = currentRoundData.curBonus;
    document.getElementById("prize-pool-random-hash").textContent = currentRoundData.hash;
    document.getElementById("prize-pool-team-no").textContent = currentRoundData.TeamCount;
    document.getElementById("prize-pool-tick-no").textContent = currentRoundData.TickCount;

    const elem = document.getElementById("prize-pool-discover-time");
    startCountdown(currentRoundData.dTime, function (days, hours, minutes, seconds, finished) {
        if (finished) {
            elem.innerText = i18next.t('game-status-with-draw');
            return;
        }

        elem.innerText = days + i18next.t('game-status-day') + hours + i18next.t('game-status-hour') + minutes
            + i18next.t('game-status-minute') + seconds + i18next.t('game-status-second');
    });
}

function setupPersonalData() {
    document.getElementById("personal-balance-val").textContent = personalData.balance;
    document.getElementById("personal-ticket-no-val").textContent = personalData.tickets.length;
    let teamNo = personalData.teams.length;
    if (personalData.hasIndividual) {
        teamNo -= 1;
    }
    document.getElementById("personal-team-no-val").textContent = teamNo;
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
        const tid = personalData.tickets[i];
        cell.innerHTML = tid;

        const teamID = personalData.tickMap.get(tid)
        if (teamID === __noTeamID) {
            cell.title = __noTeamTxt;
            cell.style.background = 'rgba(222, 64, 51, 0.3)';
        } else {
            cell.title = "团队: " + teamID;
        }

        counter++;
    }
}

function showTeamDetail() {
    if (!personalData || personalData.tickets.length === 0) {
        return;
    }

    const teamDiv = document.querySelector('.user-team');
    const isShowing = teamDiv.style.display === 'block';
    teamDiv.style.display = isShowing ? 'none' : 'block';

    if (isShowing) {
        const teamDetailDiv = document.querySelector('.team-detail-for-one');
        teamDetailDiv.style.display = 'none';
        return;
    }

    const tableBody = document.getElementById("team-detail-body");
    tableBody.innerHTML = '';
    for (let i = 0; i < personalData.teams.length; i++) {
        const teamHash = personalData.teams[i];
        if (teamHash === __noTeamID) {
            continue;
        }
        let row = tableBody.insertRow();
        let cell = row.insertCell();
        cell.innerHTML = teamHash
        cell = row.insertCell();
        cell.innerHTML = `<button class="team-detail-in-one-team" onclick="showOneTeamDetails('${teamHash}')">`+i18next.t('detail-btn')+`</button>`;
    }
}

function hideOneTeamDetails() {
    const teamDetailDiv = document.querySelector('.team-detail-for-one');
    teamDetailDiv.style.display = 'none';
}

async function showOneTeamDetails(team) {
    // console.log(team);
    const teamDetailDiv = document.querySelector('.team-detail-for-one');
    teamDetailDiv.style.display = 'block';

    try {
        showWaiting("syncing from block chain")
        const obj = await lotteryGameContract.teamMembers(gameSettings.roundNo, team);

        document.getElementById("team-detail-for-one-memNo").textContent = obj.memNo;
        document.getElementById("team-detail-for-one-tickNo").textContent = obj.voteNo;

        const tableBody = document.getElementById("team-detail-for-one-body");
        tableBody.innerHTML = '';
        for (let i = 0; i < obj.voteNos.length; i++) {

            let row = tableBody.insertRow();
            let cell = row.insertCell();
            cell.innerHTML = obj.members[i];

            cell = row.insertCell();
            cell.innerHTML = obj.voteNos[i];
        }
    } catch (err) {
        showDialog(DLevel.Warning, "load team detail failed")
    } finally {
        hideLoading();
    }
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
        const queryNo = Number(roundNo);
        if (queryNo > gameSettings.roundNo) {
            showDialog(DLevel.Tips, "bigger than current round no:" + gameSettings.roundNo);
            return;
        }
        showWaiting("syncing from block chain");
        const obj = await lotteryGameContract.gameInfoRecord(queryNo);
        const cardDiv = document.querySelector('.round-history');

        fullFillGameCard(obj, cardDiv);

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

function fullFillGameCard(obj, cardDiv) {
    cardDiv.style.display = 'block';
    cardDiv.querySelector('.one-round-bonus-val').textContent = ethers.utils.formatUnits(obj.bonus, 'ether');

    const dTime = new Date(obj.discoverTime * 1000);
    cardDiv.querySelector('.one-round-discover-val').textContent = dTime.toString();

    cardDiv.querySelector('.history-game-random').textContent = obj.randomVal;
    cardDiv.querySelector('.history-game-random-hash').textContent = obj.randomHash;
    cardDiv.querySelector('.history-game-winner-address').textContent = obj.winner;
    cardDiv.querySelector('.history-game-winner-team').textContent = obj.winTeam;
    cardDiv.querySelector('.history-game-winner-ticket').textContent = obj.winTicketID;
}

let __toRoundNo = 0;

async function loadHistoryData() {
    const moreBtn = document.querySelector('.history-data-list-more-btn');
    moreBtn.style.display = 'block';
    __toRoundNo = gameSettings.roundNo - 1;

    const parentDiv = document.querySelector('.history-data-list');
    parentDiv.style.display = 'block';
    parentDiv.innerHTML = '';

    document.getElementById('hide-history-data-btn').style.display = 'block'

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

function hideHistoryData() {
    const parentDiv = document.querySelector('.history-data-list');
    parentDiv.style.display = 'none';
    parentDiv.innerHTML = '';
    document.getElementById('hide-history-data-btn').style.display = 'none';
    document.querySelector('.history-data-list-more-btn').style.display = 'none';
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
            __shareVoteToTweet(0, no).then(r => {
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

async function syncWinTeamHistoryData() {
    const data = await GetToSrvByJson('/queryWinTeamHistory');
    if (!data) {
        return;
    }
    cachedWinTeamHistoryData = data;
    document.querySelector('.team-winning-count').textContent = "" + cachedWinTeamHistoryData.length;
}

function showTeamWinHistory() {
    if (cachedWinTeamHistoryData.length === 0) {
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

        for (const obj of cachedWinTeamHistoryData) {

            const winTeamCard = document.getElementById("winTeam-history-template").cloneNode(true);
            winTeamCard.style.display = 'block';
            winTeamCard.id = null;

            winTeamCard.querySelector('.one-round-bonus-val').textContent = obj.bonus;
            winTeamCard.querySelector('.one-round-round-val').textContent = obj.round_no;
            winTeamCard.querySelector('.team-id-txt.id').textContent = obj.win_team;
            winTeamCard.querySelector('.one-round-my-vote-no').textContent = obj.member_vote_no;
            winTeamCard.querySelector('.one-round-vote-no').textContent = obj.total_vote_no;
            winTeamCard.querySelector('.one-round-mem-no').textContent = obj.total_mem_no;

            const myBonus = obj.bonus / obj.total_vote_no * obj.member_vote_no;
            winTeamCard.querySelector('.one-round-bonus-for-me').textContent = myBonus;

            historyDiv.appendChild(winTeamCard);
        }

    } catch (err) {
        showDialog(DLevel.Warning, "load err:" + err.toString())
    }

    syncWinTeamHistoryData().then(r => {
    })
}

function showUserWinHistory() {
    if (cachedWinnerHistoryData.length === 0) {
        return;
    }
    const historyDiv = document.querySelector('.winTeam-history-list');
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

            if (obj.win_team === __noTeamID) {
                winnerCard.querySelector('.team-id-txt.type').textContent = __noTeamTxt;
                winnerCard.querySelector('.team-id-txt.id').textContent = '';
                winnerCard.querySelector('.one-round-bonus-for-me').textContent = obj.bonus;

            } else {
                winnerCard.querySelector('.team-id-txt.id').textContent = obj.win_team;
                winnerCard.querySelector('.team-id-txt.type').textContent =  i18next.t('winning-history-type-team');
                winnerCard.querySelector('.one-round-bonus-for-me').textContent = obj.bonus / 2;
            }

            historyDiv.appendChild(winnerCard);
        }
    } catch (err) {
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
