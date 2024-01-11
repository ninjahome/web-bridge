let metamaskObj = null;
let metamaskProvider;
let tweetVoteContract;
let lotteryGameContract;
let voteContractMeta = TweetVoteContractSetting.load();
let gameContractMeta;
let userBasicGameInfo;

async function checkMetaMaskEnvironment() {

    if (typeof window.ethereum === 'undefined') {
        window.location.href = "/signIn";
        return
    }

    metamaskObj = window.ethereum;
    metamaskObj.on('accountsChanged', metamaskAccountChanged);
    metamaskObj.on('chainChanged', checkCurrentChainID);
    const chainID = await metamaskObj.request({method: 'eth_chainId'});

    await checkCurrentChainID(chainID);
}

async function initVoteContractMeta() {
    const [
        postPrice, votePrice, maxVote, pluginAddr, pluginStop, kolRate, feeRate
    ] = await tweetVoteContract.systemSettings();

    const votePriceInEth = ethers.utils.formatUnits(votePrice, 'ether');
    voteContractMeta = new TweetVoteContractSetting(postPrice, votePrice, votePriceInEth,
        maxVote.toNumber(), pluginAddr, pluginStop, kolRate, feeRate);
    TweetVoteContractSetting.sycToDb(voteContractMeta);
}

async function initGameContractMeta() {

    const [currentRoundNo, totalBonus, ticketNo, _] = await lotteryGameContract.systemSettings();
    const gameInfo = await lotteryGameContract.gameInfoRecord(currentRoundNo);

    const curBonusInEth = ethers.utils.formatUnits(gameInfo.bonus, 'ether');
    const dTime = gameInfo.discoverTime.toNumber() * 1000;
    const totalBonusInEth = ethers.utils.formatUnits(totalBonus, 'ether');

    const allTickets = await lotteryGameContract.ticketsOfBuyer(currentRoundNo, ninjaUserObj.eth_addr);
    gameContractMeta = new GameBasicInfo(currentRoundNo,
        totalBonusInEth, ticketNo,curBonusInEth,
        allTickets.length,dTime,gameInfo.randomHash);
}

async function initBlockChainContract() {
    try {
        metamaskProvider = new ethers.providers.Web3Provider(metamaskObj);
        const signer = metamaskProvider.getSigner(ninjaUserObj.eth_addr);
        const conf = __globalContractConf.get(__globalTargetChainNetworkID);

        tweetVoteContract = new ethers.Contract(conf.tweetVote, tweetVoteContractABI, signer);
        lotteryGameContract = new ethers.Contract(conf.gameLottery, gameContractABI, signer);

        initVoteContractMeta().then(r => { });

        initGameContractMeta().then(r => {
            setupGameInfo();
        });

    } catch (error) {
        console.error("Error getting system settings: ", error);
        checkMetamaskErr(error);
    }
}

async function checkCurrentChainID(chainId) {
    if (__globalTargetChainNetworkID === chainId) {
        await initBlockChainContract();
        return;
    }

    showDialog("tips", "switch to arbitrum", switchToWorkChain, function () {
        metamaskProvider = null;
    });
}

async function switchChain(chainId) {
    try {
        await metamaskObj.request({
            method: 'wallet_switchEthereumChain',
            params: [{chainId}],
        });
        location.reload();
        return {switched: true, needAdd: false};
    } catch (error) {

        if (error.code === 4902) {
            return {switched: false, needAdd: true};
        } else {
            showDialog("error", "Failed switching to Arbitrum network");
            return {switched: false, needAdd: false};
        }
    }
}

async function addChain(chainId) {
    try {
        const chainParams = __globalMetaMaskNetworkParam.get(chainId);
        await metamaskObj.request({
            method: 'wallet_addEthereumChain',
            params: [chainParams],
        });
        location.reload();
    } catch (addError) {
        showDialog("error", "Add to network failed: " + addError.toString());
    }
}

async function switchToWorkChain() {
    const result = await switchChain(__globalTargetChainNetworkID);
    if (result.needAdd) {
        await addChain(__globalTargetChainNetworkID);
    }
}

function metamaskAccountChanged(accounts) {
    if (accounts.length === 0) {
        window.location.href = "/signOut";
        return;
    }
    window.location.href = "/signOut";
}

async function procPaymentForPostedTweet(tweet, callback) {
    if (!metamaskProvider) {
        showDialog("tips", "please change metamask to arbitrum network")
        return;
    }

    try {
        changeLoadingTips("paying for tweet post");

        const txResponse = await tweetVoteContract.publishTweet(
            tweet.prefixed_hash,
            tweet.signature,
            {value: voteContractMeta.postPrice}
        );
        console.log("Transaction Response: ", txResponse);

        changeLoadingTips("waiting for blockchain packaging:" + txResponse.hash);

        const txReceipt = await txResponse.wait();
        console.log("Transaction Receipt: ", txReceipt);

        const txStatus = txReceipt.status ? TXStatus.Success : TXStatus.Failed;

        hideLoading();

        showDialog("transaction " + (txReceipt.status ? "confirmed" : "failed"));

        tweet.payment_status = txStatus;
    } catch (err) {
        const newErr = checkMetamaskErr(err);
        if (newErr && newErr.includes("duplicate post")) {
            tweet.payment_status = TXStatus.Success;
        }
    } finally {
        if (callback) {
            callback(tweet);
        }
    }
}

function checkMetamaskErr(err) {
    console.error("Transaction error: ", err);
    hideLoading();

    if (err.code === 4001) {
        return null;
    }

    let code = err.code;
    if (!err.data || !err.data.message) {
        code = code + err.message;
    } else {
        code = "code:" + err.data.code + " " + err.data.message
    }
    if (code.includes("duplicate post")) {
        return code;
    }
    showDialog(code);
    return code;
}


async function procTweetVotePayment(voteCount, tweet, callback) {
    if (!metamaskProvider) {
        showDialog("tips", "please change metamask to arbitrum network")
        return;
    }

    try {
        showWaiting("prepare to pay");

        const amount = voteContractMeta.postPrice.mul(voteCount);

        const txResponse = await tweetVoteContract.voteToTweets(
            tweet.prefixed_hash,
            voteCount,
            {value: amount}
        );
        console.log("Transaction Response: ", txResponse);
        changeLoadingTips("waiting for blockchain packaging:" + txResponse.hash);

        const txReceipt = await txResponse.wait();
        console.log("Transaction Receipt: ", txReceipt);
        showDialog("Transaction: " + txReceipt.status ? "success" : "failed");

        hideLoading();

        if (callback) {
            callback(tweet.create_time, voteCount);
        }
    } catch (err) {
        checkMetamaskErr(err);
    }
}
