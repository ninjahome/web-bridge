let metamaskObj = null;
let metamaskProvider;
let tweetVoteContract;
let lotteryGameContract;
let voteContractMeta = TweetVoteContractSetting.load();
let gameContractMeta;
let curGameMeta;
let userGameInfo;

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

async function initBlockChainContract() {
    try {
        metamaskProvider = new ethers.providers.Web3Provider(metamaskObj);
        const signer = metamaskProvider.getSigner(ninjaUserObj.eth_addr);
        const conf = __globalContractConf.get(__globalTargetChainNetworkID);

        tweetVoteContract = new ethers.Contract(conf.tweetVote, tweetVoteContractABI, signer);
        lotteryGameContract = new ethers.Contract(conf.gameLottery, gameContractABI, signer);

        const [
            postPrice, votePrice, maxVote, pluginAddr, pluginStop, kolRate, feeRate
        ] = await tweetVoteContract.systemSettings();

        const votePriceInEth = ethers.utils.formatUnits(votePrice, 'ether');
        voteContractMeta = new TweetVoteContractSetting(postPrice, votePrice, votePriceInEth,
            maxVote.toNumber(), pluginAddr, pluginStop, kolRate, feeRate);
        TweetVoteContractSetting.sycToDb(voteContractMeta);

        const [currentRoundNo, totalBonus,ticketNo, ticketPriceForOuter] = await lotteryGameContract.systemSettings();
        const totalBonusInEth = ethers.utils.formatUnits(totalBonus, 'ether');
        const ticketPriceInEth = ethers.utils.formatUnits(totalBonus, 'ether');
        gameContractMeta = new GameContractMeta(currentRoundNo, totalBonusInEth,ticketNo, ticketPriceForOuter, ticketPriceInEth);

    } catch (error) {
        console.error("Error getting system settings: ", error);
        showDialog("error",error.toString());
    }
}

async function checkCurrentChainID(chainId) {
    if (__globalTargetChainNetworkID === chainId) {
        await initBlockChainContract();
        return;
    }
    showDialog("tips", "switch to arbitrum", switchToWorkChain);
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

async function processTweetPayment(create_time, prefixed_hash, signature) {
    try {
        changeLoadingTips("paying for tweet post");

        const txResponse = await tweetVoteContract.publishTweet(
            prefixed_hash,
            signature,
            {value: voteContractMeta.postPrice}
        );
        // console.log("Transaction Response: ", txResponse);

        changeLoadingTips("waiting for blockchain packaging:" + txResponse.hash);
        // updateTweetPaymentStatus(create_time, TXStatus.Pending, txResponse.hash);

        const txReceipt = await txResponse.wait();
        console.log("Transaction Receipt: ", txReceipt);

        const txStatus = txReceipt.status ? TXStatus.Success : TXStatus.Failed;
        // updateTweetPaymentStatus(create_time, txStatus, txResponse.hash);

        hideLoading();
        showDialog("transaction " + (txReceipt.status ? "confirmed" : "failed"));
    } catch (err) {
        const newErr = checkMetamaskErr(err);
        if (newErr && newErr.includes("duplicate post")) {
            // updateTweetPaymentStatus(create_time, TXStatus.Success, prefixed_hash);
        }
    }
}