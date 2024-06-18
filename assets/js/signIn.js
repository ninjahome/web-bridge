window.addEventListener('resize', function () {
    const container = document.querySelector('.container');
    const containerRightDistance = window.innerWidth - container.getBoundingClientRect().right;
    const minRightDistance = 300;
    if (containerRightDistance <= minRightDistance) {
        container.style.right = minRightDistance + 'px';
    } else {
        container.style.right = '20vw';
    }
});

class SignInData {
    constructor(ethAddr, referrer_code, signTim, refererID) {
        this.eth_addr = ethAddr;
        this.referrer_code = referrer_code;
        this.referrer_id = refererID
        if (!signTim) {
            this.sign_time = (new Date()).getTime();
        } else {
            this.sign_time = signTim;
        }
    }
}

function showError(txt) {
    if (!txt || txt.length === 0) {
        document.getElementById("error-txt").textContent = '';
        document.getElementById("error-div").style.display = 'none';
        return;
    }
    document.getElementById("error-txt").textContent = txt;
    document.getElementById("error-div").style.display = 'block';
}

async function showSignBtn(web3ID) {
    document.querySelectorAll(".logo-container").forEach(item => item.classList.remove("active"));
    const web3IdTips = document.getElementById("web3-id-tips");
    const signInDiv = document.getElementById("sign-in-div");
    const metamaskDiv = document.getElementById("metamask-div");
    const web3Id = document.getElementById("web3-id");
    const refererDiv = document.querySelector(".referral-code-popup");

    if (!web3ID) {
        web3IdTips.style.display = 'block';
        signInDiv.style.display = 'none';
        metamaskDiv.style.display = 'block';
        web3Id.style.display = 'none';
        refererDiv.style.display = 'none';
        showError(null);
        return;
    }
    web3IdTips.style.display = 'none';
    signInDiv.style.display = 'block';
    metamaskDiv.style.display = 'none';
    web3Id.style.display = 'block';
    web3Id.textContent = web3ID;
    document.getElementById("logo-container-arbitrum").classList.add("active")
    document.getElementById("logo-container-eth").classList.add("active")

    const savedUserInfo = await loadNJUserInfoFromSrv(web3ID);
    if (!savedUserInfo) {
        refererDiv.style.display = 'block';
        return;
    }

    if (!savedUserInfo.referrer_code) {
        refererDiv.style.display = 'block';
    } else {
        refererDiv.style.display = 'none';
    }

    let tw_data = await TwitterBasicInfo.loadTwBasicInfo(savedUserInfo.tw_id);
    if (!tw_data) {
        return;
    }

    document.getElementById("logo-container-twitter").classList.add("active")
}

document.addEventListener("DOMContentLoaded", initPageElem);

async function initPageElem() {
    await initDatabase();

    if (typeof window.ethereum === 'undefined') {
        await showSignBtn(null);
        return;
    }

    window.ethereum.on('accountsChanged', async function (accounts) {
        showError(null);
        if (accounts.length > 0) {
            await showSignBtn(accounts[0]);
            return;
        }
        await showSignBtn(null);
    });

    try {
        const accounts = await window.ethereum.request({method: 'eth_accounts'})
        if (accounts.length === 0) {
            await showSignBtn(null);
            return
        }
        await showSignBtn(accounts[0]);
    } catch (err) {
        console.error("Error checking MetaMask status:", err);
        showError(err);
    }
}

function closeReferralCode() {
    document.querySelector(".referral-code-popup").style.display = 'none';
}

async function setupReferralCode() {
    const referralInput = document.getElementById("referral-code-val");
    const val = referralInput.value;
    if (val.length !== referrerCodeLen) {
        showDialog(DLevel.Tips, "referral code length is:" + referrerCodeLen)
        return;
    }
    showWaiting("checking referral code");
    const referrerInfo = await loadNJUserByReferrer(val);
    if (!referrerInfo) {
        showDialog(DLevel.Tips, "no such referrer!")
        referralInput.value = '';
        referralInput.dataset.ethAddr = '';
        hideLoading();
        return;
    }

    hideLoading();
    const web3Id = document.getElementById("web3-id").textContent;
    if (referrerInfo.eth_addr === web3Id) {
        showDialog(DLevel.Tips, "Referrer should not be yourself");
        referralInput.value = '';
        referralInput.dataset.ethAddr = '';
        return;
    }

    document.querySelector(".referral-code-popup").style.display = 'none';
    referralInput.dataset.ethAddr = referrerInfo.eth_addr;
}

async function connectToMetamask() {
    if (!window.ethereum || !window.ethereum.isMetaMask) {
        window.open('https://metamask.io/download/', '_blank');
        return;
    }
    try {
        const accounts = await window.ethereum.request({method: "eth_requestAccounts"})
        // console.log(accounts)
        await showSignBtn(accounts[0]);
        // window.location.reload();
    } catch (error) {
        console.error("Metamask login failed:", error.code, error.message.toString());
        if (error.code === -32002) {
            showError("Please click the metamask plugin and sign in ")
            return;
        }
        showError(error.message);
    }
}

async function signInByEth() {
    showError(null);
    const web3Id = document.getElementById("web3-id").textContent;
    if (web3Id.length < 20) {
        showError("no valid web3 id selected");
        return;
    }
    const referralInput = document.getElementById("referral-code-val");
    const val = referralInput.value;

    const signParam = new SignInData(web3Id)
    if (val.length === referrerCodeLen) {
        signParam.referrer_code = val;
        signParam.referrer_id = referralInput.dataset.ethAddr;
    }
    const message = JSON.stringify(signParam);
    try {
        const signature = await window.ethereum.request({
            method: 'personal_sign',
            params: [message, web3Id],
        })

        showWaiting("signing in......", 15)
        const obj = new SignDataForPost(message, signature, null)
        PostToSrvByJson('/signInByEth', obj).then(response => {
            console.log("sign in by eth success", response);
            hideLoading();
            if (!response) {
                showError("failed to load ninja user data");
                return
            }
            window.location.href = "/main";
        }).catch(err => {
            showError(err.toString());
            hideLoading();
        })
    } catch (err) {
        showError("sign in err:\n" + err.message);
        hideLoading();
    }
}

function showAppList(listID, isHide) {
    if (isHide) {
        document.getElementById(listID).style.display = 'none';
    } else {
        document.getElementById(listID).style.display = 'block';
    }
}
