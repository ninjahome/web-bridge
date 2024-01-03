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
    constructor(postPrice, votePrice, maxVote, pluginAddr, pluginStop, kolRate, feeRate) {
        this.postPrice = postPrice;
        this.votePrice = votePrice;
        this.maxVote = maxVote;
        this.pluginAddr = pluginAddr;
        this.pluginStop = pluginStop;
        this.kolRate = kolRate;
        this.feeRate = feeRate;
    }
}

let voteContractMeta = null;

function setupMetamask() {
    metamaskObj = window.ethereum;
    metamaskObj.on('accountsChanged', metamaskAccountChanged);
    metamaskObj.on('chainChanged', metamaskChainChanged);
    metamaskObj.request({method: 'eth_chainId'}).then(chainID => {
        metamaskChainChanged(chainID).then(r => {
        });
    })
}

function initializeContract() {
    metamaskProvider = new ethers.providers.Web3Provider(metamaskObj);
    const signer = metamaskProvider.getSigner(ninjaUserObj.eth_addr);
    const conf = __globalContractConf.get(__globalTargetChainNetworkID);

    if (!conf || !conf.tweetVote) {
        return false;
    }

    const postPrice = ethers.utils.parseEther(conf.postPrice);
    const votePrice = ethers.utils.parseEther(conf.votePrice);
    voteContractMeta = new SmartContractSettings(postPrice, votePrice);
    tweetVoteContract = new ethers.Contract(conf.tweetVote, conf.tweetVoteAbi, signer);
    lotteryGameContract = new ethers.Contract(conf.gameLottery, conf.gameLotteryAbi, signer);

    loadVoteContractMeta().then(r=>{});
    return true;
}

async function loadVoteContractMeta() {
    try {
        const [
            postPrice, votePrice, maxVote, pluginAddr, pluginStop, kolRate, feeRate
        ] = await tweetVoteContract.systemSettings();
        voteContractMeta = new SmartContractSettings(postPrice, votePrice,
            maxVote.toNumber(), pluginAddr, pluginStop, kolRate, feeRate);

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
