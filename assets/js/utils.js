
function formatTime(createTime) {
    const date = new Date(createTime);

    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');
    const seconds = date.getSeconds().toString().padStart(2, '0');
    const day = date.getDate().toString().padStart(2, '0');
    const month = (date.getMonth() + 1).toString().padStart(2, '0');

    return `${hours}:${minutes}:${seconds} ${day}/${month}`;
}

function startCountdown(targetTime, callback) {
    const countdownInterval = setInterval(() => {
        const now = new Date().getTime();
        const timeLeft = targetTime - now;

        if (timeLeft <= 0) {
            clearInterval(countdownInterval);
            callback('', '', '', '', true);
            return;
        }

        const days = Math.floor(timeLeft / (1000 * 60 * 60 * 24));
        let hours = Math.floor((timeLeft % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
        let minutes = Math.floor((timeLeft % (1000 * 60 * 60)) / (1000 * 60));
        let seconds = Math.floor((timeLeft % (1000 * 60)) / 1000);

        hours = hours < 10 ? "0" + hours : hours;
        minutes = minutes < 10 ? "0" + minutes : minutes;
        seconds = seconds < 10 ? "0" + seconds : seconds;

        callback(days, hours, minutes, seconds, false);
    }, 1000);
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
                if (response.redirected) {
                    window.location = response.url;
                    return;
                }
                if (!response.ok) {

                    if (response.status === 302 || response.status === 301) {
                        window.location = response.url;
                        return;
                    }

                    return response.text().then(text => {
                        console.log(text)
                        throw new Error('\tserver responded with an error:' + response.status);
                    });
                }
                if (response.headers.get("Content-Length") === "0" || !response.headers.get("Content-Type").includes("application/json")) {
                    return {};
                }
                return response.json();
            })
            .then(data => {
                resolve(data);
            })
            .catch(error => {
                console.log(error);
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

    try {
        const response = await fetch(url, requestOptions);
        if (response.redirected) {
            window.location = response.url;
            return;
        }
        if (!response.ok) {
            if ([301, 302, 303, 307, 308].includes(response.status)) {
                // 如果是重定向响应，获取重定向的 URL 并导航到那里
                window.location = response.url;
                return;
            }
            const text = await response.text();
            console.log(text);
            throw new Error('Server responded with an error: ' + response.status);
        }
        if (response.headers.get("Content-Length") === "0" || !response.headers.get("Content-Type").includes("application/json")) {
            return {};
        }

        return await response.json();
    } catch (error) {
        console.error('Error during fetch:', error);
        throw error;
    }
}

const __globalMetaMaskNetworkParam = new Map([
    [toHex(42161), {
        chainId: toHex(42161),
        chainName: 'Arbitrum One',
        nativeCurrency: {
            name: 'ETH',
            symbol: 'ETH',
            decimals: 18
        },
        rpcUrls: ['https://1rpc.io/arb'],
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
    spinner.style.borderTop = '6px solid #4EA0F2';
    spinner.style.borderRadius = '50%';
    spinner.style.width = '40px';
    spinner.style.height = '40px';
    spinner.style.animation = 'spin 2s ease-in-out infinite';

    // Create the loading message
    const message = document.createElement('div');
    message.id = 'loading-message';
    message.style.marginTop = '10px';
    message.textContent = ' ';

    // Add text in the center of the container
    const loadingText = document.createElement('div');
    loadingText.textContent = ' ';
    loadingText.style.position = 'absolute';
    loadingText.style.top = '48.75%';
    loadingText.style.left = '50%';
    loadingText.style.transform = 'translate(-50%, -50%)';
    loadingText.style.fontSize = '10px'; // Adjust font size as needed
    loadingText.style.color = '#4EA0F2'; // Ensure text color is visible

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
let isShowingTips = false;
function showWaiting(message, timeout) {
    if (isShowingTips){
        changeLoadingTips(message);
        return;
    }

    isShowingTips = true;
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
    if (!isShowingTips){
        return;
    }
    isShowingTips = false;
    const modal = document.getElementById('loading-modal');
    if (modal) {
        document.body.removeChild(modal);
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

    static cacheTwBasicInfo(obj) {
        if (!obj.id) {
            throw new Error("invalid twitter basic info")
        }
        localStorage.setItem(lclDbKeyForTwitterUserData(obj.id), JSON.stringify(obj));
        return obj;
    }
}

class NJUserBasicInfo {

    constructor(address, eth_addr, create_at, tw_id, update_at,
                tweet_count, vote_count, be_voted_count,is_elder) {
        this.address = address;
        this.eth_addr = eth_addr;
        this.create_at = create_at;
        this.tw_id = tw_id;
        this.update_at = update_at;
        this.tweet_count = tweet_count;
        this.vote_count = vote_count;
        this.be_voted_count = be_voted_count;
        this.is_elder = is_elder;
    }


    static loadNjBasic(ethAddr) {
        const storedData = localStorage.getItem(DbKeyForNjUserData(ethAddr.toLowerCase()))
        if (!storedData) {
            return null
        }
        const nuObj = JSON.parse(storedData);
        return new NJUserBasicInfo(nuObj.address, nuObj.eth_addr, nuObj.create_at, nuObj.tw_id,
            nuObj.update_at, nuObj.tweet_count, nuObj.vote_count, nuObj.be_voted_count,nuObj.is_elder);
    }

    static cacheNJUsrObj(obj) {
        if (!obj.eth_addr) {
            throw new Error("invalid twitter basic info")
        }

        localStorage.setItem(DbKeyForNjUserData(obj.eth_addr.toLowerCase()), JSON.stringify(obj));
    }
}


const DLevel = Object.freeze({
    Tips: 1, Warning: 2, Error: 3, Success: 4
});

function createDialogElement(imageSrc) {
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
        <div style="background: white; padding: 24px 32px; border-radius: 8px; text-align: center;">
             <img src="${imageSrc}" alt="dialog image" style="max-width: 100%; height: auto; margin-bottom: 2px">
            <p id="dialog-message" style="margin-top: 0; margin-bottom: 16px">Message</p>
            <div style="display: flex; flex-direction: row; justify-content: center">
                <button id="dialog-close" style="margin:0 8px; padding: 8px 42px; border-radius: 12px; background-color: transparent; border: 1px solid rgb(221, 221, 222);font-size: 14px; color: #7E7F82">关闭</button>
                <button id="dialog-confirm" style="margin:0 8px; padding: 8px 42px; border-radius: 12px; border: none;font-size: 14px; color: #ffffff; background-color: #4EA0F2">确定</button>
            </div>
        </div>
        `;
    return dialog;
}

function showDialog(type, msg, confirmCB, cancelCB) {
    hideLoading();
    let imageSrc;
    switch (type) {
        case DLevel.Tips:
            imageSrc = "/assets/file/info-img.png";
            break;
        case DLevel.Error:
            imageSrc = "/assets/file/error-img.png";
            break;
        case DLevel.Warning:
            imageSrc = "/assets/file/warning-img.png";
            break;
        case DLevel.Success:
            imageSrc = "/assets/file/success-img.png";
            break;
    }

    const dialog = createDialogElement(imageSrc);
    document.body.appendChild(dialog);
    const dialogMessage = document.getElementById('dialog-message');

    const dialogCloseButton = document.getElementById('dialog-close');
    const dialogConfirmButton = document.getElementById('dialog-confirm');

    dialogMessage.textContent = msg;

    dialogCloseButton.addEventListener('click', function () {
        document.body.removeChild(dialog);
        if (cancelCB) {
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

let metamaskObj = null;

async function checkMetaMaskEnvironment(callback) {

    if (typeof window.ethereum === 'undefined') {
        window.location.href = "/signIn";
        return
    }

    metamaskObj = window.ethereum;
    metamaskObj.on('accountsChanged', metamaskAccountChanged);
    metamaskObj.on('chainChanged', function (chainID) {
        checkCurrentChainID(chainID, callback)
    });
    const chainID = await metamaskObj.request({method: 'eth_chainId'});

    await checkCurrentChainID(chainID, callback);
}

function metamaskAccountChanged(accounts) {
    if (accounts.length === 0) {
        console.log('MetaMask账户已断开连接，请重新连接');
        window.location.href = "/signOut";
        return;
    }
    window.location.href = "/signOut";
}

async function checkCurrentChainID(chainId, callback) {
    if (__globalTargetChainNetworkID === chainId) {
        const provider = new ethers.providers.Web3Provider(metamaskObj);
        if (callback) {
            await callback(provider);
        }
        return;
    }

    showDialog(DLevel.Tips, "switch to arbitrum", switchToWorkChain, function () {
        window.location.href = "/signOut";
    });
}

async function switchToWorkChain() {
    const result = await switchChain(__globalTargetChainNetworkID);
    if (result.needAdd) {
        await addChain(__globalTargetChainNetworkID);
    }
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
            showDialog(DLevel.Error, "Failed switching to Arbitrum network");
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
        showDialog(DLevel.Error, "Add to network failed: " + addError.toString());
    }
}

let confirmCallback = null;

function openVoteModal(callback) {
    const modal = document.getElementById("vote-no-chose-modal");
    modal.style.display = "block";
    const voteCount = document.getElementById("voteCountInChoseModal");
    voteCount.value = 1;
    confirmCallback = callback;
}

function confirmVoteModal() {
    if (confirmCallback) {
        const voteCount = document.getElementById("voteCountInChoseModal").value;
        const shareOnTwitter = document.getElementById("shareOnTwitter").checked;
        confirmCallback(voteCount, shareOnTwitter);
    }
    closeVoteModal();
}

function closeVoteModal() {
    const modal = document.getElementById("vote-no-chose-modal");
    modal.style.display = "none";
}

function increaseVote() {
    const voteCount = document.getElementById("voteCountInChoseModal");
    voteCount.value = parseInt(voteCount.value) + 1;
}

function decreaseVote() {
    const voteCountElement = document.getElementById("voteCountInChoseModal");
    const newVoteCount = Math.max(1, parseInt(voteCountElement.value) - 1);
    voteCountElement.value = newVoteCount.toString();
}

async function __shareVoteToTweet(create_time, vote_count,slogan) {
    await PostToSrvByJson("/shareVoteAction", {
        create_time: create_time,
        vote_count: Number(vote_count),
        slogan:slogan,
    });
}

function checkMetamaskErr(err) {
    console.error("Transaction error: ", err);
    hideLoading();

    if (err.code === 4001) {
        return null;
    }

    if (err.code === 4100){
        showDialog(DLevel.Warning, "open metamask first");
        return
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
    if (code.includes("insufficient funds")) {
        showDialog(DLevel.Warning, "insufficient funds");
        return
    }
    showDialog(DLevel.Warning, code);
    return code;
}

function __incomeWithdrawHistory(address) {
    let targetUrl = __globalMetaMaskNetworkParam.get(__globalTargetChainNetworkID).blockExplorerUrls[0];
    targetUrl += '/address/' + address;
    targetUrl += '#internaltx';
    window.open(targetUrl);
}

function createThumbnail(originalImageSrc, maxWidth, maxHeight) {
    return new Promise((resolve, reject) => {
        const img = new Image();
        img.src = originalImageSrc;

        img.onload = function() {
            const originalWidth = img.width;
            const originalHeight = img.height;
            const aspectRatio = originalWidth / originalHeight;

            let targetWidth, targetHeight;
            if (originalWidth / originalHeight > maxWidth / maxHeight) {
                targetWidth = maxWidth;
                targetHeight = maxWidth / aspectRatio;
            } else {
                targetHeight = maxHeight;
                targetWidth = maxHeight * aspectRatio;
            }
            const thumbnailDataUrl = __createCanvas(img, targetWidth,targetHeight);
            resolve(thumbnailDataUrl);
        };

        img.onerror = function() {
            reject(new Error('Could not load image'));
        };
    });
}
function __createCanvas(img,targetWidth,targetHeight){
    const canvas = document.createElement('canvas');
    const ctx = canvas.getContext('2d');

    canvas.width = targetWidth;
    canvas.height = targetHeight;

    ctx.drawImage(img, 0, 0, targetWidth, targetHeight);

    return  canvas.toDataURL('image/png');
}

function compressImageByFile(file, quality) {
    return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onload = function (e) {
            const img = new Image();
            img.onload = function () {
                const canvas = document.createElement('canvas');
                const ctx = canvas.getContext('2d');
                canvas.width = img.width;
                canvas.height = img.height;
                ctx.drawImage(img, 0, 0, img.width, img.height);

                // 压缩图片
                canvas.toBlob(function (blob) {
                    // 获取压缩后的图片Blob
                    const compressedFile = new File([blob], file.name, {type: "image/jpeg", lastModified: Date.now()});
                    // 使用resolve返回压缩后的图片文件和Blob
                    resolve({
                        compressedFile: compressedFile,
                        blob: blob
                    });
                }, 'image/jpeg', quality);
            };
            img.onerror = function () {
                reject(new Error('Could not load image'));
            };
            img.src = e.target.result;
        };
        reader.onerror = function () {
            reject(new Error('Could not read file'));
        };
        reader.readAsDataURL(file);
    });
}

function blobToBase64(blob) {
    return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onload = () => {
            // 当读取完成，reader.result 包含了Base64编码的数据URL
            resolve(reader.result);
        };
        reader.onerror = (error) => {
            reject(error);
        };
        reader.readAsDataURL(blob); // 开始将blob转换为data URL
    });
}

function blobToImage(blob) {
    return new Promise((resolve, reject) => {
        const img = new Image();
        const url = URL.createObjectURL(blob);
        img.onload = () => {
            URL.revokeObjectURL(url);
            resolve(img);
        };
        img.onerror = () => {
            URL.revokeObjectURL(url);
            reject(new Error('Failed to load image from blob'));
        };
        img.src = url;
    });
}


function compressBlob(image, quality) {
    return new Promise((resolve, reject) => {
        const canvas = document.createElement('canvas');
        const ctx = canvas.getContext('2d');
        canvas.width = image.width;
        canvas.height = image.height;
        ctx.drawImage(image, 0, 0, image.width, image.height);
        canvas.toBlob( (blob) => {
            resolve(blob);
        }, 'image/jpeg', quality);
    });
}

function readFileAsBlob(file) {
    return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onload = (event) => {
            const img = new Image();
            img.onload = () => {
                resolve(img);
            };
            img.onerror = (error) => {
                reject(error);
            };
            img.src = event.target.result;
        };
        reader.onerror = (error) => {
            reject(error);
        };
        reader.readAsDataURL(file);
    });
}


const __defaultLogo = '/assets/file/logo.png';
const maxTextLenPerImg = 1000;
const maxImgPerTweet = 4;
const defaultTextLenForTweet  = 100;
const MaxRawImgSize = (1<<20) ;
const MaxThumbnailSize = (1<<18);
const CompressQuality = 0.75;