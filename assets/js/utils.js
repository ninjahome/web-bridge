function showDialog(title, msg, callback) {
    // 加载通用提示框的内容
    const dialogFrame = document.createElement('iframe');
    dialogFrame.src = '/assets/html/dialog.html';
    dialogFrame.style.width = '100%';
    dialogFrame.style.height = '100%';
    dialogFrame.style.border = 'none';
    dialogFrame.style.position = 'fixed';
    dialogFrame.style.top = '0';
    dialogFrame.style.left = '0';
    dialogFrame.style.zIndex = '10000';

    // 当通用提示框加载完成后，设置标题和消息，并显示
    dialogFrame.onload = function () {
        const dialogWindow = dialogFrame.contentWindow;
        const dialogDocument = dialogWindow.document;
        const dialogTitle = dialogDocument.getElementById('dialog-title');
        const dialogMessage = dialogDocument.getElementById('dialog-message');
        const dialogCloseButton = dialogDocument.getElementById('dialog-close');
        const dialogConfirmButton = dialogDocument.getElementById('dialog-confirm');

        dialogTitle.textContent = title;
        dialogMessage.textContent = msg;

        // 关闭按钮的点击事件
        dialogCloseButton.addEventListener('click', function () {
            document.body.removeChild(dialogFrame);
        });

        // 确认按钮的点击事件
        if (callback) {
            dialogConfirmButton.style.display = 'block';
            dialogConfirmButton.addEventListener('click', function () {
                callback();
                document.body.removeChild(dialogFrame);
            });
        } else {
            dialogConfirmButton.style.display = 'none';
        }
    };

    // 在当前页面中添加通用提示框
    document.body.appendChild(dialogFrame);
}

function PostToSrvByJson(url, data) {
    const requestOptions = {
        method: 'POST', // 请求方法
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    };
    return new Promise((resolve, reject) => {
        fetch(url, requestOptions)
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
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

function showWaiting(message, timeout) {
    const loadingFrame = document.createElement('iframe');
    loadingFrame.src = '/assets/html/loading.html';
    loadingFrame.style.width = '100%';
    loadingFrame.style.height = '100%';
    loadingFrame.style.border = 'none';
    loadingFrame.style.position = 'fixed';
    loadingFrame.style.top = '0';
    loadingFrame.style.left = '0';
    loadingFrame.style.zIndex = '10000';

    loadingFrame.onload = function () {
        const loadingWindow = loadingFrame.contentWindow;
        const loadingDocument = loadingWindow.document;
        const loadingMessage = loadingDocument.getElementById('loading-message');
        const loadingContainer = loadingDocument.getElementById('loading-container');

        loadingMessage.textContent = message;
        loadingContainer.style.display = 'flex';

        if (timeout) {
            setTimeout(() => {
                document.body.removeChild(loadingFrame);
            }, timeout * 1000);
        }
    };

    document.body.appendChild(loadingFrame);
}

function hideLoading() {
    const loadingFrame = document.querySelector('iframe[src="/assets/html/loading.html"]');
    if (loadingFrame) {
        document.body.removeChild(loadingFrame);
    }
}
function toHex(number) {
    return '0x' + number.toString(16);
}


const __globalTargetChainNetworkID = toHex(421613);
const __globalMetaMaskNetworkParam = new Map([
    [toHex(42161), {
        chainId:  toHex(42161),
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
        blockExplorerUrls: ['https://sepolia-explorer.arbitrum.io'],
    }]
]);

class SignDataForPost {
    constructor(msg, sig,payload) {
        this.message = msg;
        this.signature = sig;
        this.payload = payload;
    }
}