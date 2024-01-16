
let tweetVoteContract;
let lotteryGameContract = null;
let voteContractMeta = TweetVoteContractSetting.load();
let gameContractMeta = null;


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

    const [currentRoundNo, totalBonus] = await lotteryGameContract.systemSettings();
    const gameInfo = await lotteryGameContract.gameInfoRecord(currentRoundNo);

    const curBonusInEth = ethers.utils.formatUnits(gameInfo.bonus, 'ether');
    const dTime = gameInfo.discoverTime.toNumber() * 1000;
    const totalBonusInEth = ethers.utils.formatUnits(totalBonus, 'ether');

    const [teamNo, memNo, voteNo]  = await lotteryGameContract.allTeamInfoNo(currentRoundNo);
    gameContractMeta = new GameBasicInfo(currentRoundNo,
        totalBonusInEth, voteNo, curBonusInEth,
        teamNo,memNo, dTime);
}

async function initBlockChainContract(provider) {
    try {
        if (!provider){
            tweetVoteContract = null;
            lotteryGameContract = null
            return
        }
        const signer = provider.getSigner(ninjaUserObj.eth_addr);
        const conf = __globalContractConf.get(__globalTargetChainNetworkID);
        tweetVoteContract = new ethers.Contract(conf.tweetVote, tweetVoteContractABI, signer);
        lotteryGameContract = new ethers.Contract(conf.gameLottery, gameContractABI, signer);

        initVoteContractMeta().then(r => {
        });

        initGameContractMeta().then(r => {
            setupGameInfo(true);
        });

    } catch (error) {
        console.error("block chain err: ", error);
        checkMetamaskErr(error);
    }
}

async function procPaymentForPostedTweet(tweet, callback) {
    if (!tweetVoteContract) {
        showDialog(DLevel.Tips, "please change metamask to arbitrum network")
        return;
    }

    try {
        showWaiting("paying for tweet post");

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

        if(txReceipt.status){
            showDialog(DLevel.Success,"transaction " + "confirmed");
        }else {
            showDialog(DLevel.Error,"transaction " + "failed");
        }

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
    showDialog(DLevel.Warning,code);
    return code;
}


async function procTweetVotePayment(voteCount, tweet, callback) {
    if (!tweetVoteContract|| !voteContractMeta) {
        showDialog(DLevel.Tips, "please wait for metamask syncing data")
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

        if(txReceipt.status){
            showDialog(DLevel.Success,"transaction " + "confirmed");
        }else {
            showDialog(DLevel.Error,"transaction " + "failed");
        }

        hideLoading();

        if (callback) {
            callback(tweet.create_time, voteCount);
        }
    } catch (err) {
        checkMetamaskErr(err);
    }
}

async function reloadGameBalance() {
    const b = await lotteryGameContract.balance(ninjaUserObj.eth_addr)
    document.getElementById('lottery-game-income').innerText = ethers.utils.formatUnits(b, 'ether');
}

async function reloadTweetBalance() {
    const b = await tweetVoteContract.balance(ninjaUserObj.eth_addr);
    document.getElementById("tweet-income-amount").innerText = ethers.utils.formatUnits(b, 'ether');
}


async function withdrawLotteryGameIncome() {
    showWaiting("prepare withdraw transaction");
    const valStr = document.getElementById('lottery-game-income').innerText;
    const balance = Number(valStr);

    if (!balance || balance <= 0) {
        showDialog(DLevel.Tips, "balance invalid");
        hideLoading();
        return;
    }

    await withdrawAction(lotteryGameContract);
    await reloadGameBalance();
    hideLoading();
}

async function withdrawFromUserTweetIncome() {
    showWaiting("prepare withdraw transaction");
    const valStr = document.getElementById('lottery-game-income').innerText;
    const balance = Number(valStr);
    if (balance <= 0) {
        showDialog(DLevel.Tips, "balance too low");
        hideLoading();
        return;
    }

    await withdrawAction(tweetVoteContract);
    await reloadTweetBalance();
    hideLoading();
}

