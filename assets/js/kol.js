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


    } catch (error) {
        console.error("block chain err: ", error);
        checkMetamaskErr(error);
    }
}