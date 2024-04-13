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
    if (await checkIfMetaMaskSignOut() === false) {
        return;
    }
    const [currentRoundNo, totalBonus, voteNo, price, bonusPoint] = await lotteryGameContract.systemSettings();
    // console.log(price, bonusPoint);
    const gameInfo = await lotteryGameContract.gameInfoRecord(currentRoundNo);

    const curBonusInEth = ethers.utils.formatUnits(gameInfo.bonus, 'ether');
    const dTime = gameInfo.discoverTime.toNumber() * 1000;
    const totalBonusInEth = ethers.utils.formatUnits(totalBonus, 'ether');
    const bonusForPoint = ethers.utils.formatUnits(bonusPoint, 'ether');

    gameContractMeta = new GameBasicInfo(currentRoundNo,
        totalBonusInEth, voteNo, curBonusInEth, dTime, bonusForPoint);
}

async function initBlockChainContract(provider) {
    try {
        if (!provider) {
            tweetVoteContract = null;
            lotteryGameContract = null
            return
        }
        const signer = provider.getSigner(ninjaUserObj.eth_addr);
        const conf = __globalContractConf.get(__globalTargetChainNetworkID);
        tweetVoteContract = new ethers.Contract(conf.tweetVote, tweetVoteContractABI, signer);
        lotteryGameContract = new ethers.Contract(conf.gameLottery, gameContractABI, signer);

        initVoteContractMeta().then(r => {
            const postBtn = document.getElementById("tweet-post-with-eth-btn-txt-1");
            const votePriceInModal = document.getElementById("vote-price-in-modal");
            if (postBtn) {
                postBtn.innerText = i18next.t('btn-tittle-post-tweet') + "(" + voteContractMeta.votePriceInEth + " ETH)"
                votePriceInModal.innerText = voteContractMeta.votePriceInEth + " ETH"
            }
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
        showWaiting("paying for tweet");

        const txResponse = await tweetVoteContract.publishTweet(
            tweet.prefixed_hash,
            tweet.signature,
            {value: voteContractMeta.postPrice}
        );

        changeLoadingTips("packaging:" + txResponse.hash);
        const txReceipt = await txResponse.wait();

        tweet.payment_status = txReceipt.status ? TXStatus.Success : TXStatus.Failed;

    } catch (err) {
        const newErr = checkMetamaskErr(err);
        if (newErr && newErr.includes("duplicate post")) {
            tweet.payment_status = TXStatus.Success;
        }
    } finally {
        hideLoading();
        if (callback) {
            callback(tweet);
        }
    }
}

async function procTweetVotePayment(voteCount, tweet, callback) {
    if (!tweetVoteContract || !voteContractMeta || !voteContractMeta.votePrice || !voteContractMeta.votePrice.mul) {
        showDialog(DLevel.Tips, "please wait for metamask syncing data")
        return;
    }

    try {
        showWaiting("prepare to pay");
        const amount = voteContractMeta.votePrice.mul(voteCount);

        const txResponse = await tweetVoteContract.voteToTweets(
            tweet.prefixed_hash,
            voteCount,
            {value: amount}
        );
        changeLoadingTips("packaging: " + txResponse.hash);

        const txReceipt = await txResponse.wait();

        if (!txReceipt.status) {
            showDialog(DLevel.Error, "transaction failed");
            return;
        }
        showDialog(DLevel.Success, "transaction success");

        if (callback) {
            callback(tweet.create_time, voteCount);
        }
    } catch (err) {
        checkMetamaskErr(err);
    } finally {
        hideLoading();
    }
}

async function reloadGameBalance() {
    const b = await lotteryGameContract.balance(ninjaUserObj.eth_addr);
    // console.log(b);
    document.getElementById('lottery-game-income').innerText = ethers.utils.formatUnits(b, 'ether');
}

async function reloadTweetBalance() {
    const b = await tweetVoteContract.balance(ninjaUserObj.eth_addr);
    document.getElementById("tweet-income-amount").innerText = ethers.utils.formatUnits(b, 'ether');
}

async function withdrawLotteryGameIncome() {
    const valStr = document.getElementById('lottery-game-income').innerText;
    const balance = Number(valStr);

    if (!balance || balance <= 0) {
        showDialog(DLevel.Tips, "balance invalid");
        return;
    }

    await withdrawAction(lotteryGameContract);
    await reloadGameBalance();
}

async function withdrawFromUserTweetIncome() {
    const valStr = document.getElementById('tweet-income-amount').innerText;
    const balance = Number(valStr);
    if (balance <= 0) {
        showDialog(DLevel.Tips, "balance too low");
        return;
    }

    await withdrawAction(tweetVoteContract);
    await reloadTweetBalance();
}

function incomeWithdrawHistory() {
    __incomeWithdrawHistory(ninjaUserObj.eth_addr);
}

