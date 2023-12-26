function checkSystemEnvironment() {

    if (typeof window.ethereum === 'undefined') {
        window.location.href = "/signIn";
        return
    }
    setupMetamask();
}

function setupBasicInfo() {
    const twBtn = document.getElementById('sign-up-by-twitter-button')
    const twNameLabel = document.getElementById('basic-twitter-name')
    document.getElementById('basic-web3-id').innerText = ninjaUserObj.eth_addr;
    if (!ninjaUserObj.tw_id) {
        twNameLabel.style.display = 'none';
        twBtn.style.display = 'inline-block';
    } else {
        twBtn.style.display = 'none';
        twNameLabel.style.display = 'inline-block';
        loadTwitterInfo(ninjaUserObj.tw_id, true).then(twInfo => {
            setupTwitterElem(twInfo);
        })
    }
}

function signUpByTwitter() {
    window.location.href = "/signUpByTwitter";
}

function quitFromService() {
    fetch("/signOut", {method: 'GET'}).then(r => {
        window.location.href = "/signIn";
    }).catch(err => {
        console.log(err)
        window.location.href = "/signIn";
    })
}
