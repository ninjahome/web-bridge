let metamaskObj = null;
let metamaskProvider;

function checkMetaMaskEnvironment() {

    if (typeof window.ethereum === 'undefined') {
        window.location.href = "/signIn";
        return
    }

    metamaskObj = window.ethereum;
    metamaskObj.on('accountsChanged', metamaskAccountChanged);
    metamaskObj.on('chainChanged', checkCurrentChainID);
    metamaskObj.request({method: 'eth_chainId'}).then(chainID => {
        checkCurrentChainID(chainID);
    })
}

function checkCurrentChainID(chainId){
    let needShowSwitch = false;
    if (__globalTargetChainNetworkID !== chainId) {
        document.getElementById("switch-to-arbitrum").style.display = 'block';
        needShowSwitch = true;
    }else{
        document.getElementById("switch-to-arbitrum").style.display = 'none';
    }

    if (!ninjaUserObj.tw_id) {
        document.getElementById("bridge-to-twitter").style.display = 'block';
        needShowSwitch = true;
    }else{
        document.getElementById("bridge-to-twitter").style.display = 'none';
    }
    if (needShowSwitch){
        document.getElementById("web3-environment-operate-area").style.display = 'block';
    }
}

async function switchChain(chainId) {
    try {
        await metamaskObj.request({
            method: 'wallet_switchEthereumChain',
            params: [{chainId}],
        });
        return { switched: true, needAdd: false };
    } catch (error) {

        if (error.code === 4902) {
            return { switched: false, needAdd: true };
        } else {
            showDialog("error", "Failed switching to Arbitrum network");
            return { switched: false, needAdd: false };
        }
    }
}

async function addChain(chainId) {
    try {
        const chainParams = __globalMetaMaskNetworkParam.get(chainId);
        return metamaskObj.request({
            method: 'wallet_addEthereumChain',
            params: [chainParams],
        });
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