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
        this.MemCount = 0;
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

        loadCurrentRoundMeta().then(r=>{
            setupCurrentRoundData();
        });

        loadPersonalMeta().then(r=>{
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

        const [teamNo, memNo, voteNo] = await lotteryGameContract.allTeamInfoNo(gameSettings.roundNo);
        console.log(teamNo, memNo, voteNo);

        currentRoundData.TeamCount = teamNo;
        currentRoundData.MemCount = memNo;
        currentRoundData.TickCount = voteNo;

        console.log(currentRoundData);

    } catch (err) {
        console.log(err);
        showDialog(DLevel.Warning, "load game data from block chain failed");
    }
}

async function loadPersonalMeta() {
    const obj = await lotteryGameContract.tickList(gameSettings.roundNo, ninjaUserObj.eth_addr);
    if (obj[0].length === 0){
        return;
    }
    const map = new Map();

    for (let i = 0; i < obj[0].length; i++) {
        map.set(obj[0][i],obj[0][i]);
    }
    for (const tickID of obj[0]) {

    }
    console.log(obj[0]);
    console.log(obj[1]);
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
    document.getElementById("prize-pool-member-no").textContent = currentRoundData.MemCount;

    const elem = document.getElementById("prize-pool-discover-time");
    startCountdown(currentRoundData.dTime,function (days,hours,minutes,seconds, finished){
        if(finished){
            elem.innerText = "开奖中";
            return;
        }

        elem.innerText = days +" 天" +hours+ " 时"+minutes+" 分"+seconds+" 秒";
    });
}

function setupPersonalData(){

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