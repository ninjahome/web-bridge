let kolKeyContract = null

async function initBlockChainContract(provider) {
    try {
        if (!provider) {
            kolKeyContract = null;
            return
        }
        const signer = provider.getSigner(ninjaUserObjWeb3ID);
        const conf = __globalContractConf.get(__globalTargetChainNetworkID);
        kolKeyContract = new ethers.Contract(conf.kolKey, kolKeyContractABI, signer);
        console.log("init key contract success")

    } catch (error) {
        console.error("block chain err: ", error);
        checkMetamaskErr(error);
    }
}

async function loadUserIncomeFromKey() {
    showWaiting("loading from blockchain");
    try {
        const balance = await kolKeyContract.AllIncomeOfAllKol(ninjaUserObjWeb3ID);
        const balanceInEth = ethers.utils.formatUnits(balance, 'ether');
        console.log(balanceInEth);
    } catch (err) {

    } finally {
        hideLoading();
    }
}

async function loadKeyDetails() {
    const kolList = document.querySelector('.kol-list-of-investor');

    const isShowing = kolList.style.display === 'block';
    if (isShowing) {
        kolList.style.display = 'none';
        return;
    }
    kolList.style.display = 'block';

    showWaiting("loading from blockchain");
    try {
        const balance = await kolKeyContract.AllIncomeOfAllKol(ninjaUserObjWeb3ID);
        const balanceInEth = ethers.utils.formatUnits(balance, 'ether');
        console.log(balanceInEth);
    } catch (err) {

    } finally {
        hideLoading();
    }
}