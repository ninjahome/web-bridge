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
        console.log(obj);
        return new GameRoundInfo(obj.randomHash, dTime, obj.winner,
            obj.winTeam, obj.winTicketID, curBonusInEth, obj.randomVal);
    }
}

class PersonalData {
    constructor(balance, tickets, teams, map) {
        this.balance = balance;
        this.tickets = tickets;
        this.teams = teams;
        this.tickMap = map;
    }
}

async function initGamePage() {
    await checkMetaMaskEnvironment(initGameContract)
}

async function initGameContract(provider) {
    if (!provider) {
        lotteryGameContract = null;
        return;
    }

    const signer = provider.getSigner(ninjaUserObj.eth_addr);
    const conf = __globalContractConf.get(__globalTargetChainNetworkID);
    lotteryGameContract = new ethers.Contract(conf.gameLottery, gameContractABI, signer);

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

        console.log(gameSettings);

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
        console.log(teamNo, voteNo);

        currentRoundData.TeamCount = teamNo;
        currentRoundData.TickCount = voteNo;

        console.log(currentRoundData);

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
            personalData = new PersonalData(balanceInEth, [], [], null);
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
            Array.from(mapTeams.keys()), mapTickets);

        console.log(personalData);
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
            elem.innerText = "开奖中";
            return;
        }

        elem.innerText = days + " 天" + hours + " 时" + minutes + " 分" + seconds + " 秒";
    });
}

function setupPersonalData() {
    document.getElementById("personal-balance-val").textContent = personalData.balance;
    document.getElementById("personal-ticket-no-val").textContent = personalData.tickets.length;
    document.getElementById("personal-team-no-val").textContent = personalData.teams.length;
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
        cell.title = "团队: " + personalData.tickMap.get(personalData.tickets[i]);
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
        teamDetailDiv.style.display =  'none';
        return;
    }

    const tableBody = document.getElementById("team-detail-body");
    tableBody.innerHTML = '';
    for (let i = 0; i < personalData.teams.length; i++) {
        let row = tableBody.insertRow();
        let cell = row.insertCell();
        const teamHash = personalData.teams[i];
        cell.innerHTML = teamHash
        cell = row.insertCell();
        cell.innerHTML = `<button class="team-detail-in-one-team" onclick="showOneTeamDetails('${teamHash}')">详情</button>`;
    }
}
function hideOneTeamDetails(){
    const teamDetailDiv = document.querySelector('.team-detail-for-one');
    teamDetailDiv.style.display =  'none';
}

async function showOneTeamDetails(team) {
    console.log(team);
    const teamDetailDiv = document.querySelector('.team-detail-for-one');
    teamDetailDiv.style.display =  'block';

    try {
        showWaiting("syncing from block chain")
        const obj = await lotteryGameContract.teamMembers(gameSettings.roundNo, team);
        console.log(obj);
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

async function buyTicket() {
    if (!gameSettings) {
        await loadGameSettings();
    }

    if (!gameSettings.isOpen) {
        showDialog(DLevel.Tips, "not open for personal user");
        return;
    }
}

function showUserWinHistory(){
    showDialog(DLevel.Tips,"not ok now");
    // const historyDiv = document.querySelector('.winning-history');
    // const isShowing = historyDiv.style.display === 'block';
    // historyDiv.style.display = isShowing ? 'none' : 'block';
    // if (isShowing){
    //     return;
    // }
}

function showGameRule(className){
    const gameRuleDiv = document.querySelector(className);
    gameRuleDiv.style.display = gameRuleDiv.style.display === 'none' ? 'block' : 'none';
}