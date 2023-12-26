let twitterUserObj = null;

function setupTwitterElem(twInfo) {
    if (!twInfo) {
        twitterUserObj = null;
        return;
    }
    const isVerifiedLabel = document.getElementById("basic-twitter-verified");
    const twNameLabel = document.getElementById('basic-twitter-name')
    twitterUserObj = twInfo;
    twNameLabel.innerText = twInfo.name;
    if (!twInfo.verified) {
        isVerifiedLabel.innerText = "Premium False";
    } else {
        isVerifiedLabel.innerText = "Premium True";
    }
    if (twInfo.profile_image_url) {
        document.getElementById('user-twitter-logo').src = twInfo.profile_image_url;
    }
}

async function loadTwitterInfo(twitterID, needCache,forceSync) {
    if (!forceSync){
        forceSync = false;
    }

    try {
        if (needCache) {
            let tw_data = TwitterBasicInfo.loadTwBasicInfo(twitterID)
            if (tw_data) {
                return tw_data;
            }
        }
        const response = await GetToSrvByJson("/queryTwBasicById?forceSync="+forceSync);
        if (!response.ok) {
            console.log("query twitter basic info failed")
            return null;
        }

        const text = await response.text();
        console.log(text);
        return TwitterBasicInfo.cacheTwBasicInfo(text);
    } catch (err) {
        console.log("queryTwBasicById err:", err)
        return null;
    }
}

function refreshTwitterInfo() {
    loadTwitterInfo(ninjaUserObj.tw_id, false,true).then(twInfo => {
        setupTwitterElem(twInfo);
    })
}