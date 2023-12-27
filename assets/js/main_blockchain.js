let metamaskObj = null;

function setupMetamask() {
    metamaskObj = window.ethereum;
    metamaskObj.on('accountsChanged', metamaskAccountChanged);
    metamaskObj.on('chainChanged', metamaskChainChanged);
    metamaskObj.request({method: 'eth_chainId'}).then(chainID => {
        metamaskChainChanged(chainID).then(r => {
        });
    })
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
    // console.log('eth balance:', balance);
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