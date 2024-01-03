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
