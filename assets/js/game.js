let lotteryGameContract = null;
let gameSettings = null;

class GameSettings {
    constructor(roundNo, bonus, tickNo) {
        this.roundNo = roundNo;
        this.bonus = bonus;
        this.tickNo = tickNo;
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
    loadGameSettings().then(r=>{
        setupCurrentRoundData();
    })
}

async function loadGameSettings() {
    showWaiting("syncing system data from block chain")
    try {
        const [currentRoundNo, totalBonus, ticketNo] = await lotteryGameContract.systemSettings();
        const totalBonusInEth = ethers.utils.formatUnits(totalBonus, 'ether');
        gameSettings = new GameSettings(currentRoundNo, totalBonusInEth, ticketNo);
        console.log(gameSettings);
    } catch (err) {
        console.log(err);
        showDialog(DLevel.Warning, "load game data from block chain failed");
    }finally {
        hideLoading()
    }
}

function setupCurrentRoundData(){

}