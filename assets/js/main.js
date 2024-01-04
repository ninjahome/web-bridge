const dbKeyCachedVoteContractMeta = "__db_key_cached_vote_contract_meta__"

function checkSystemEnvironment() {

    if (typeof window.ethereum === 'undefined') {
        window.location.href = "/signIn";
        return
    }
    setupMetamask();
}

function setupBasicInfo() {
    const twBtn = document.getElementById('sign-up-by-twitter-button')
    const twNameLabel = document.getElementById('basic-twitter-name')
    document.getElementById('basic-web3-id').innerText = ninjaUserObj.eth_addr;
    if (!ninjaUserObj.tw_id) {
        twNameLabel.style.display = 'none';
        twBtn.style.display = 'inline-block';
    } else {
        twBtn.style.display = 'none';
        twNameLabel.style.display = 'inline-block';
        loadTwitterInfo(ninjaUserObj.tw_id, true).then(twInfo => {
            setupTwitterElem(twInfo);
        })
    }
}

function signUpByTwitter() {
    window.location.href = "/signUpByTwitter";
}

function quitFromService() {
    fetch("/signOut", {method: 'GET'}).then(r => {
        window.location.href = "/signIn";
    }).catch(err => {
        console.log(err)
        window.location.href = "/signIn";
    })
}


let twitterUserObj = null;

function setupTwitterElem(twInfo) {
    if (!twInfo) {
        twitterUserObj = null;
        return;
    }
    const isVerifiedLabel = document.getElementById("basic-twitter-verified");
    const twNameLabel = document.getElementById('basic-twitter-name')
    twitterUserObj = twInfo;
    twNameLabel.innerText = twInfo.name;
    if (!twInfo.verified) {
        isVerifiedLabel.innerText = "Premium False";
    } else {
        isVerifiedLabel.innerText = "Premium True";
    }
    if (twInfo.profile_image_url) {
        document.getElementById('user-twitter-logo').src = twInfo.profile_image_url;
    }
}

async function loadTwitterInfo(twitterID, needCache, forceSync) {
    if (!forceSync) {
        forceSync = false;
    }

    try {
        if (needCache) {
            let tw_data = TwitterBasicInfo.loadTwBasicInfo(twitterID)
            if (tw_data) {
                return tw_data;
            }
        }
        const response = await GetToSrvByJson("/queryTwBasicById?twitterID=" + twitterID + "&&forceSync=" + forceSync);
        if (!response.ok) {
            console.log("query twitter basic info failed")
            return null;
        }

        const text = await response.text();
        return TwitterBasicInfo.cacheTwBasicInfo(text);
    } catch (err) {
        console.log("queryTwBasicById err:", err)
        return null;
    }
}

function refreshTwitterInfo() {
    loadTwitterInfo(ninjaUserObj.tw_id, false, true).then(twInfo => {
        setupTwitterElem(twInfo);
    })
}

let metamaskObj = null;
let metamaskProvider;
let tweetVoteContract;
let lotteryGameContract;

class SmartContractSettings {
    constructor(postPrice, votePrice, votePriceInEth, maxVote, pluginAddr, pluginStop, kolRate, feeRate) {
        this.postPrice = postPrice;
        this.votePrice = votePrice;
        this.votePriceInEth = votePriceInEth;
        this.maxVote = maxVote;
        this.pluginAddr = pluginAddr;
        this.pluginStop = pluginStop;
        this.kolRate = kolRate;
        this.feeRate = feeRate;
    }

    static sycToDb(obj) {
        localStorage.setItem(SmartContractSettings.DBKey(), JSON.stringify(obj));
    }

    static DBKey() {
        return dbKeyCachedVoteContractMeta;
    }

    static load() {
        const storedVal = localStorage.getItem(SmartContractSettings.DBKey());
        return storedVal ? JSON.parse(storedVal) : null;
    }
}

class GameContractMeta {
    constructor(curRound, totalBonus, ticketPrice, ticketPriceInEth) {
        this.curRound = curRound;
        this.totalBonus = totalBonus;
        this.ticketPrice = ticketPrice;
        this.ticketPriceInEth = ticketPriceInEth;
    }
}

class GameRoundInfo {
    constructor(round, randomHash, nextRoundTime, bonus, winner, winTeam, winTicket) {
        this.round = round;
        this.randomHash = randomHash;
        this.nextRoundTime = nextRoundTime;
        this.winner = winner;
        this.winTeam = winTeam;
        this.winTicket = winTicket;
        this.bonus = bonus;
    }
}

let voteContractMeta = SmartContractSettings.load();
let gameContractMeta;
let curGameMeta;

function setupMetamask() {
    metamaskObj = window.ethereum;
    metamaskObj.on('accountsChanged', metamaskAccountChanged);
    metamaskObj.on('chainChanged', metamaskChainChanged);
    metamaskObj.request({method: 'eth_chainId'}).then(chainID => {
        metamaskChainChanged(chainID).then(r => {
        });
    })
}

async function initializeContract() {
    metamaskProvider = new ethers.providers.Web3Provider(metamaskObj);
    const signer = metamaskProvider.getSigner(ninjaUserObj.eth_addr);
    const conf = __globalContractConf.get(__globalTargetChainNetworkID);

    if (!conf || !conf.tweetVote) {
        showDialog("error","blockchain setting err!")
        return false;
    }

    tweetVoteContract = new ethers.Contract(conf.tweetVote, conf.tweetVoteAbi, signer);
    lotteryGameContract = new ethers.Contract(conf.gameLottery, conf.gameLotteryAbi, signer);

    if (!voteContractMeta) {
        await loadVoteContractMeta();
    } else {
        loadVoteContractMeta().then(r => {
        });
    }

    loadGameContractMeta();

    return true;
}

function loadGameContractMeta() {

    lotteryGameContract.systemSettings().then(([currentRoundNo, totalBonus, ticketPriceForOuter]) => {
        const totalBonusInEth = ethers.utils.formatUnits(totalBonus, 'ether');
        const ticketPriceInEth = ethers.utils.formatUnits(totalBonus, 'ether');
        gameContractMeta = new GameContractMeta(currentRoundNo, totalBonusInEth, ticketPriceForOuter, ticketPriceInEth);
        console.log(JSON.stringify(gameContractMeta));
        loadCurGameMeta();
    }).catch(err => {
        console.log(err);
    })
}

function loadCurGameMeta() {
    lotteryGameContract.gameInfoRecord(gameContractMeta.curRound).then((gameInfo) => {
        const curBonusInEth = ethers.utils.formatUnits(gameInfo.bonus, 'ether');
        const dTime = gameInfo.discoverTime.toNumber() * 1000;
        curGameMeta = new GameRoundInfo(gameContractMeta.curRound, gameInfo.randomHash, dTime, curBonusInEth);
        fullFillGameBasicInfo();
        console.log(JSON.stringify(curGameMeta));
    }).catch(err => {
        console.log(err);
    })
}

function fullFillGameBasicInfo() {
    document.getElementById("current-round").innerText = curGameMeta.round;
    document.getElementById("total-prize").innerText = curGameMeta.bonus;
    document.getElementById("lottery-hash").innerText = curGameMeta.randomHash;
    document.getElementById("lottery-discovery-time").innerText = formatTime(curGameMeta.nextRoundTime);
    document.getElementById("total-awards").innerText = gameContractMeta.totalBonus;
}

async function loadVoteContractMeta() {
    try {
        const [
            postPrice, votePrice, maxVote, pluginAddr, pluginStop, kolRate, feeRate
        ] = await tweetVoteContract.systemSettings();

        const votePriceInEth = ethers.utils.formatUnits(votePrice, 'ether');
        voteContractMeta = new SmartContractSettings(postPrice, votePrice, votePriceInEth,
            maxVote.toNumber(), pluginAddr, pluginStop, kolRate, feeRate);
        SmartContractSettings.sycToDb(voteContractMeta);

        const tweetPostPriceInEth = ethers.utils.formatUnits(postPrice, 'ether');
        document.getElementById("tweet-post-with-eth-btn").innerText = "发布推文(" + tweetPostPriceInEth + " eth)"

        console.log(JSON.stringify(voteContractMeta));
    } catch (error) {
        console.error("Error getting system settings: ", error);
    }
}

async function metamaskChainChanged(chainId) {
    const chainBtn = document.getElementById('change-chain-id-button')
    const chainBalance = document.getElementById('basic-web3-balance')
    if (__globalTargetChainNetworkID === chainId) {
        chainBtn.style.display = 'none';
        chainBalance.style.display = 'inline-block';
        chainBalance.innerText = await metamaskAccBalance()
        return;
    }
    chainBalance.style.display = 'none';
    chainBtn.style.display = 'inline-block';
}

async function metamaskAccBalance() {
    const balance = await metamaskObj.request({
        method: 'eth_getBalance',
        params: [ninjaUserObj.eth_addr, 'latest'],
    });
    if (balance === "0x0") {
        return "0.00 eth";
    }
    if (balance < 10 ** 12) {
        return "< 0.000001 eth"
    }
    const balanceInEth = (parseInt(balance, 16) / 10 ** 18).toFixed(6);
    const formattedBalance = parseFloat(balanceInEth).toFixed(2);
    return formattedBalance + ' eth';
}


function metamaskAccountChanged(accounts) {
    if (accounts.length === 0) {
        window.location.href = "/signOut";
        return;
    }
    window.location.href = "/signOut";
}


function switchToWorkChain() {
    metamaskObj.request({
        method: 'wallet_switchEthereumChain',
        params: [{chainId: __globalTargetChainNetworkID}],
    }).catch((switchError) => {
        if (switchError.code !== 4902) {
            showDialog("error", "failed switching to arbitrum network");
            return;
        }

        const arbParam = __globalMetaMaskNetworkParam.get(__globalTargetChainNetworkID);
        metamaskObj.request({
            method: 'wallet_addEthereumChain',
            params: [arbParam],
        }).then(r => {
            console.log(r);
        }).catch(err => {
            showDialog("error", "add to network failed:" + err.toString());
        });
    });
}
