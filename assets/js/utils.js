function formatTime(createTime) {
    const date = new Date(createTime);

    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');
    const seconds = date.getSeconds().toString().padStart(2, '0');
    const day = date.getDate().toString().padStart(2, '0');
    const month = (date.getMonth() + 1).toString().padStart(2, '0');

    return `${hours}:${minutes}:${seconds} ${day}/${month}`;
}

function startCountdown(targetTime,callback) {
    const countdownInterval = setInterval(() => {
        const now = new Date().getTime();
        const timeLeft = targetTime - now;

        if (timeLeft <= 0) {
            clearInterval(countdownInterval);
            callback('开奖中',true);
            return;
        }

        const days = Math.floor(timeLeft / (1000 * 60 * 60 * 24));
        const hours = Math.floor((timeLeft % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
        const minutes = Math.floor((timeLeft % (1000 * 60 * 60)) / (1000 * 60));
        const seconds = Math.floor((timeLeft % (1000 * 60)) / 1000);

        let countdownText = '';
        if (days > 0) {
            countdownText += days + '天 ';
        }
        if (hours > 0) {
            countdownText += hours + '时 ';
        }
        if (minutes > 0) {
            countdownText += minutes + '分 ';
        }
        countdownText += seconds + '秒';

        callback(countdownText,false);
    }, 1000);
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
                        throw new Error('\tserver responded with an error:' + response.status);
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
        tweetVote: "0xa3a39F3415d2024834Ef22258FC14e5cdcc3E857",
        gameLottery: "0x9077B82966B428F2A0B4fC088BE59fDE7FDEcb46",
        kolKey: "",
        kolKeyAbi: "",
        postPrice: "0.005",
        votePrice: "0.005"
    }],
    [toHex(42161), {
        tweetVote: "",
        gameLottery: "",
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
    modal.style.backgroundColor = 'rgba(0, 0, 0, 0.3)';
    modal.style.zIndex = '10000';

    // Create a container for the spinner and text
    const container = document.createElement('div');
    container.style.display = 'flex';
    container.style.flexDirection = 'column';
    container.style.alignItems = 'center'; // Align items vertically
    container.style.background = 'rgb(255,255,255,1)';
    container.style.padding = '20px';
    container.style.borderRadius = '5px';

    // Create the spinner
    const spinner = document.createElement('div');
    spinner.id = 'loading-spinner';
    spinner.style.border = '6px solid #DDDDDE';
    spinner.style.borderTop = '6px solid #4848D8';
    spinner.style.borderRadius = '50%';
    spinner.style.width = '40px';
    spinner.style.height = '40px';
    spinner.style.animation = 'spin 2s ease-in-out infinite';

    // Create the loading message
    const message = document.createElement('div');
    message.id = 'loading-message';
    message.style.marginTop = '10px';
    message.textContent = 'Loading...';

    // Add text in the center of the container
    const loadingText = document.createElement('div');
    loadingText.textContent = 'Loading';
    loadingText.style.position = 'absolute';
    loadingText.style.top = '48.75%';
    loadingText.style.left = '50%';
    loadingText.style.transform = 'translate(-50%, -50%)';
    loadingText.style.fontSize = '10px'; // Adjust font size as needed
    loadingText.style.color = '#4848D8'; // Ensure text color is visible

    container.appendChild(loadingText);

    // Append spinner and message to the container
    container.appendChild(spinner);
    container.appendChild(message);

    // Append the container to the modal
    modal.appendChild(container);

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

function showDialog(title, msg, confirmCB, cancelCB) {
    const dialog = createDialogElement();
    document.body.appendChild(dialog);

    const dialogTitle = document.getElementById('dialog-title');
    const dialogMessage = document.getElementById('dialog-message');
    const dialogCloseButton = document.getElementById('dialog-close');
    const dialogConfirmButton = document.getElementById('dialog-confirm');

    dialogTitle.textContent = title;
    dialogMessage.textContent = msg;


    dialogCloseButton.addEventListener('click', function () {
        document.body.removeChild(dialog);
        if (cancelCB){
            cancelCB();
        }
    });

    if (confirmCB) {
        dialogConfirmButton.style.display = 'block';
        dialogConfirmButton.addEventListener('click', function () {
            document.body.removeChild(dialog);
            confirmCB();
        });
    } else {
        dialogConfirmButton.style.display = 'none';
    }
}

function lclDbKeyForTwitterUserData(TwitterID) {
    return "__database_key_for_twitter_user_data__:" + TwitterID
}

function DbKeyForNjUserData(ethAddr) {
    return "__database_key_for_ninja_user_data__:" + ethAddr
}

function clearSessionStorage() {
    sessionStorage.clear();
}

function lclDbKeyForBlockChainData(account) {
    return "__local_database_key_for_block_chain_data__:" + account
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
        const storedData = localStorage.getItem(lclDbKeyForTwitterUserData(TwitterID))
        if (!storedData) {
            return null
        }
        const twObj = JSON.parse(storedData);
        return new TwitterBasicInfo(twObj.id, twObj.name, twObj.username,
            twObj.profile_image_url, twObj.description);
    }

    static cacheTwBasicInfo(objStr) {
        const obj = JSON.parse(objStr)
        if (!obj.id) {
            throw new Error("invalid twitter basic info")
        }
        localStorage.setItem(lclDbKeyForTwitterUserData(obj.id), objStr);
        return obj;
    }
}

class NJUserBasicInfo{

    constructor(address, eth_addr,create_at,tw_id,update_at,
                tweet_count,vote_count,be_voted_count) {
        this.address = address;
        this.eth_addr = eth_addr;
        this.create_at = create_at;
        this.tw_id = tw_id;
        this.update_at = update_at;
        this.tweet_count = tweet_count;
        this.vote_count = vote_count;
        this.be_voted_count = be_voted_count;
    }


    static loadNjBasic(ethAddr) {
        const storedData = localStorage.getItem(DbKeyForNjUserData(ethAddr.toLowerCase()))
        if (!storedData) {
            return null
        }
        const nuObj = JSON.parse(storedData);
        return new NJUserBasicInfo(nuObj.address,nuObj.eth_addr,nuObj.create_at,nuObj.tw_id,
            nuObj.update_at,nuObj.tweet_count,nuObj.vote_count,nuObj.be_voted_count);
    }

    static cacheNJUsrObj(obj) {
        if (!obj.eth_addr) {
            throw new Error("invalid twitter basic info")
        }

        localStorage.setItem(DbKeyForNjUserData(obj.eth_addr.toLowerCase()), JSON.stringify(obj));
    }
}

