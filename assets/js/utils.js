function PostToSrvByJson(url, data) {
    const requestOptions = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    };
    return new Promise((resolve, reject) => {
        fetch(url, requestOptions)
            .then(response => {
                if (!response.ok) {
                    return response.text().then(text => {
                        throw new Error(text || 'Server responded with an error');
                    });
                }
                return response.text();
            })
            .then(data => {
                resolve(data);
            })
            .catch(error => {
                reject(error);
            });
    });
}

async function GetToSrvByJson(url) {
    const requestOptions = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
    };
    return await fetch(url, requestOptions)
}


function toHex(number) {
    return '0x' + number.toString(16);
}


const __globalTargetChainNetworkID = toHex(421614);
const __globalMetaMaskNetworkParam = new Map([
    [toHex(42161), {
        chainId: toHex(42161),
        chainName: 'Arbitrum LlamaNodes',
        nativeCurrency: {
            name: 'ETH',
            symbol: 'ETH',
            decimals: 18
        },
        rpcUrls: ['https://arbitrum.llamarpc.com'],
        blockExplorerUrls: ['https://arbiscan.io'],
    }],
    [toHex(421613), {
        chainId: toHex(421613),
        chainName: 'Arbitrum Goerli',
        nativeCurrency: {
            name: 'AETH',
            symbol: 'AETH',
            decimals: 18
        },
        rpcUrls: ['https://endpoints.omniatech.io/v1/arbitrum/goerli/public'],
        blockExplorerUrls: ['https://goerli.arbiscan.io'],
    }],
    [toHex(421614), {
        chainId: toHex(421614),
        chainName: 'Arbitrum Sepolia',
        nativeCurrency: {
            name: 'ETH',
            symbol: 'ETH',
            decimals: 18
        },
        rpcUrls: ['https://sepolia-rollup.arbitrum.io/rpc'],
        blockExplorerUrls: ['https://sepolia.arbiscan.io'],
    }],
    [toHex(5777), {
        chainId: toHex(5777),
        chainName: 'loaclTest',
        nativeCurrency: {
            name: 'ETH',
            symbol: 'ETH',
            decimals: 18
        },
        rpcUrls: ['HTTP://192.168.1.122:7545'],
        blockExplorerUrls: [''],
    }]
]);

class SignDataForPost {
    constructor(msg, sig, payload) {
        this.message = msg;
        this.signature = sig;
        this.pay_load = payload;
    }
}

const DefaultAvatarSrc = "/assets/file/logo.png"

const __globalContractConf = new Map([
    [toHex(421614), {
        tweetVote: "0x63713037a9E337D7Db5D383070199B948598e0Da",
        tweetVoteAbi: tweetVoteContractABI,
        gameLottery: "0x57F0bbE85f5822911003A8fa425D5595D139FDFe",
        gameLotteryAbi: gameContractABI,
        kolKey: "",
        kolKeyAbi: ""
    }],
    [toHex(42161), {
        tweetVote: "",
        tweetVoteAbi: "",
        gameLottery: "",
        gameLotteryAbi: "",
        kolKey: "",
        kolKeyAbi: ""
    }]]);
