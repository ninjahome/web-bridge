function formatTime(createTime) {
    const date = new Date(createTime);

    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');
    const seconds = date.getSeconds().toString().padStart(2, '0');
    const day = date.getDate().toString().padStart(2, '0');
    const month = (date.getMonth() + 1).toString().padStart(2, '0');

    return `${hours}:${minutes}:${seconds} ${day}/${month}`;
}


function toHex(number) {
    return '0x' + number.toString(16);
}

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
                        console.log(text)
                        throw new Error('Server responded with an error:' + response.statusText);
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
        tweetVote: "0x6500Cda46979F1956a46486B1a88768cb425E23a",
        tweetVoteAbi: tweetVoteContractABI,
        gameLottery: "0x842e751E5D5aCc9fB172812B7499e18Ac5Fb3F59",
        gameLotteryAbi: gameContractABI,
        kolKey: "",
        kolKeyAbi: "",
        postPrice: "0.005",
        votePrice: "0.005"
    }],
    [toHex(42161), {
        tweetVote: "",
        tweetVoteAbi: "",
        gameLottery: "",
        gameLotteryAbi: "",
        kolKey: "",
        kolKeyAbi: "",
        postPrice: "0.005",
        votePrice: "0.005"
    }]]);


function createModalElement() {
    const modal = document.createElement('div');
    modal.id = 'loading-modal';
    modal.style.position = 'fixed';
    modal.style.top = '0';
    modal.style.left = '0';
    modal.style.width = '100%';
    modal.style.height = '100%';
    modal.style.display = 'flex';
    modal.style.justifyContent = 'center';
    modal.style.alignItems = 'center';
    modal.style.backgroundColor = 'rgba(0, 0, 0, 0.5)';
    modal.style.zIndex = '10000';


    modal.innerHTML = '<div id="loading-message" style="background: white; padding: 20px; border-radius: 5px;">Loading...</div>';

    // 添加加载动画
    const spinner = document.createElement('div');
    spinner.id = 'loading-spinner';
    spinner.style.border = '4px solid #f3f3f3';
    spinner.style.borderTop = '4px solid #3498db';
    spinner.style.borderRadius = '50%';
    spinner.style.width = '40px';
    spinner.style.height = '40px';
    spinner.style.animation = 'spin 2s linear infinite';
    modal.appendChild(spinner);

    modal.innerHTML += '<div id="loading-message" style="margin-top: 10px;">Loading...</div>';

    return modal;
}

// 定义旋转动画
const style = document.createElement('style');
style.type = 'text/css';
style.innerHTML = `
    @keyframes spin {
        0% { transform: rotate(0deg); }
        100% { transform: rotate(360deg); }
    }
`;
document.head.appendChild(style);

function showWaiting(message, timeout) {
    const modal = createModalElement();
    document.body.appendChild(modal);

    const loadingMessage = document.getElementById('loading-message');
    loadingMessage.textContent = message;

    if (timeout) {
        setTimeout(() => {
            document.body.removeChild(modal);
        }, timeout * 1000);
    }
}

function changeLoadingTips(content) {
    const loadingMessage = document.getElementById('loading-message');
    if (loadingMessage) {
        loadingMessage.textContent = content;
    }
}

function hideLoading() {
    const modal = document.getElementById('loading-modal');
    if (modal) {
        document.body.removeChild(modal);
    }
}


function createDialogElement() {
    const dialog = document.createElement('div');
    dialog.id = 'custom-dialog';
    dialog.style.position = 'fixed';
    dialog.style.top = '0';
    dialog.style.left = '0';
    dialog.style.width = '100%';
    dialog.style.height = '100%';
    dialog.style.display = 'flex';
    dialog.style.justifyContent = 'center';
    dialog.style.alignItems = 'center';
    dialog.style.backgroundColor = 'rgba(0, 0, 0, 0.5)';
    dialog.style.zIndex = '10000';
    dialog.innerHTML = `
        <div style="background: white; padding: 20px; border-radius: 5px; max-width: 500px; width: 100%;">
            <h2 id="dialog-title" style="margin-top: 0;">Title</h2>
            <p id="dialog-message">Message</p>
            <button id="dialog-close" style="margin-right: 10px;">Close</button>
            <button id="dialog-confirm">Confirm</button>
        </div>
    `;
    return dialog;
}

function showDialog(title, msg, callback) {
    const dialog = createDialogElement();
    document.body.appendChild(dialog);

    const dialogTitle = document.getElementById('dialog-title');
    const dialogMessage = document.getElementById('dialog-message');
    const dialogCloseButton = document.getElementById('dialog-close');
    const dialogConfirmButton = document.getElementById('dialog-confirm');

    dialogTitle.textContent = title;
    dialogMessage.textContent = msg;

    // 关闭按钮的点击事件
    dialogCloseButton.addEventListener('click', function () {
        document.body.removeChild(dialog);
    });

    // 确认按钮的点击事件
    if (callback) {
        dialogConfirmButton.style.display = 'block';
        dialogConfirmButton.addEventListener('click', function () {
            callback();
            document.body.removeChild(dialog);
        });
    } else {
        dialogConfirmButton.style.display = 'none';
    }
}


class NinjaUserBasicInfo {
    constructor(addr, ethAddr, twId, createAt) {
        this.address = addr;
        this.eth_addr = ethAddr;
        this.tw_id = twId;
        this.create_at = createAt;
    }

    static syncToSessionDbForApiResponse(response) {
        const ninjaObj = JSON.parse(response)
        if (!ninjaObj.eth_addr) {
            throw new Error("invalid ninja user info")
        }
        setDataToSessionDB(sesDbKeyForCurrentUserEthAddr(), ninjaObj.eth_addr);
        sessionStorage.setItem(sesDbKeyForNjUserData(ninjaObj.eth_addr), response);
        return ninjaObj;
    }

    static loadCurrentNJUserObj() {
        const curUsrEthAddr = getDataFromSessionDB(sesDbKeyForCurrentUserEthAddr())
        const savedUserInfo = getDataFromSessionDB(sesDbKeyForNjUserData(curUsrEthAddr))
        if (!savedUserInfo) {
            return null;
        }
        return new NinjaUserBasicInfo(savedUserInfo.address, savedUserInfo.eth_addr,
            savedUserInfo.tw_id, savedUserInfo.create_at);
    }
}

class TwitterBasicInfo {
    constructor(id, name, username, avatarUrl, bio) {
        this.id = id;
        this.name = name;
        this.username = username;
        this.profile_image_url = avatarUrl;
        this.description = bio;
    }

    static loadTwBasicInfo(TwitterID) {
        const storedData = getDataFromSessionDB(sesDbKeyForTwitterUserData(TwitterID))
        if (!storedData) {
            return null
        }
        return new TwitterBasicInfo(storedData.id, storedData.name, storedData.username,
            storedData.profile_image_url, storedData.description);
    }

    static cacheTwBasicInfo(objStr) {
        const obj = JSON.parse(objStr)
        if (!obj.id) {
            throw new Error("invalid twitter basic info")
        }
        sessionStorage.setItem(sesDbKeyForTwitterUserData(obj.id), objStr);
        return obj;
    }
}

function sesDbKeyForTwitterUserData(TwitterID) {
    return "__session_database_key_for_twitter_user_data__:" + TwitterID
}

function sesDbKeyForNjUserData(ethAddr) {
    return "__session_database_key_for_ninja_user_data__:" + ethAddr
}

function sesDbKeyForCurrentUserEthAddr() {
    return "__session_database_key_for_ninja_user_current_address__"
}

function setDataToSessionDB(key, sign_data) {
    sessionStorage.setItem(key, JSON.stringify(sign_data));
}

function getDataFromSessionDB(key) {
    const storedValue = sessionStorage.getItem(key);
    return storedValue ? JSON.parse(storedValue) : null;
}

function clearSessionStorage() {
    sessionStorage.clear();
}

function lclDbKeyForBlockChainData(account) {
    return "__local_database_key_for_block_chain_data__:" + account
}

class BlockChainData {
    constructor(account) {
        this.account = account;
    }

    static load(account) {
        const storedData = localStorage.getItem(lclDbKeyForBlockChainData(account))
        const obj = storedData ? JSON.parse(storedData) : null;
        if (!obj) {
            return null
        }

        return new BlockChainData(storedData.account);
    }
}


function checkMetamaskErr(err){
    if (err.code === 4001) {
        return null;
    }
    const  code = err.code;
    if (!err.data ||!err.data.message){
        return code+err.message;
    }
    return "code:"+err.data.code+" " + err.data.message;
}