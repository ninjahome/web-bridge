const __globalTweetMemCache = new Map()
const __globalTweetMemCacheByHash = new Map()

window.onscroll = function () {
    throttle(contentScroll, 200);
}

let throttleTimer;

function throttle(callback, time) {
    if (throttleTimer) return;

    throttleTimer = setTimeout(() => {
        callback();
        clearTimeout(throttleTimer);
        throttleTimer = null;
    }, time);
}

function contentScroll() {
    let cacheObj;
    let uiCallback;

    switch (curScrollContentID) {
        case 0:
            cacheObj = cachedGlobalTweets;
            uiCallback = loadOlderTweetsForHomePage;
            break;
        case 1:
            cacheObj = cachedUserTweets;
            uiCallback = loadOlderMostVotedTweet;
            break;
        case 12:
            cacheObj = cachedTopVotedKolUser;
            uiCallback = loadOlderMostVotedKol;
            break;
        case 13:
            cacheObj = cachedTopVoterUser;
            uiCallback = loadOlderMostVoter;
            break;
        case 2:
            cacheObj = cachedUserTweets;
            uiCallback = olderPostedTweets;
            break;
        case 22:
            cacheObj = cachedUserVotedTweets;
            uiCallback = olderVotedTweets;
            break;
        case 51:
            cacheObj = cachedNinjaUserPostedTweets;
            uiCallback = olderNinjaUsrPostedTweets;
            break;
        case 52:
            cacheObj = cachedNinjaUserVotedTweets;
            uiCallback = olderNinjaUsrVotedTweets;
            break;
        default:
            return;
    }

    if (!cacheObj.canLoadMoreOldData()) {
        return;
    }

    const documentHeight = Math.max(document.body.scrollHeight, document.documentElement.scrollHeight);
    if (window.innerHeight + window.scrollY < documentHeight - 100) {
        return;
    }

    cacheObj.isLoading = true;
    uiCallback().then(() => {
    }).finally(() => {
        cacheObj.isLoading = false;
    });
}

function clearCachedData() {
    databaseDeleteTable(__constCachedItem).then(() => {
        localStorage.clear();
        sessionStorage.clear();
        window.location.href = "/signIn";
    }).catch((error) => {
        console.error('Failed to clear cached data:', error);
        window.location.href = "/signIn";
    });
}

async function showHoverCard(event, twitterObj, web3ID, offset) {

    const hoverCard = document.getElementById('hover-card');
    const rect = event.currentTarget.getBoundingClientRect();

    const njUsrInfo = await loadNJUserInfoFromSrv(web3ID);

    if (twitterObj) {
        document.getElementById('hover-avatar').src = twitterObj.profile_image_url;
        document.getElementById('hover-name').textContent = twitterObj.name;
        document.getElementById('hover-user-name').textContent = '@' + twitterObj.username;
        if (njUsrInfo.is_elder){
            hoverCard.querySelector(".elderFlagOnAvatar").style.display = 'block';
        }else{
            hoverCard.querySelector(".elderFlagOnAvatar").style.display = 'none';
        }

    } else {
        document.getElementById('hover-name').textContent = web3ID;
    }

    let x = 0;
    let y = 0;
    if (offset) {
        x = offset.X;
        y = offset.Y
    }

    hoverCard.style.display = 'block';
    hoverCard.style.left = `${rect.left + x}px`;
    hoverCard.style.top = `${rect.bottom + window.scrollY + y}px`;

    if (!njUsrInfo) {
        console.log("failed to load web3 user:", web3ID);
        return;
    }
    const detailBtn = document.getElementById('show-details-button');
    detailBtn.onclick = () => {
        hoverCard.style.display = 'none';
        showUserProfile(njUsrInfo);
    };
    detailBtn.textContent = i18next.t("show-details-button");
    document.getElementById('hover-tweet-count').textContent = njUsrInfo.tweet_count;
    document.getElementById('hover-vote-count').textContent = njUsrInfo.vote_count;
    document.getElementById('hover-voted-count').textContent = njUsrInfo.be_voted_count;
}

function hideHoverCard(obj) {
    const hoverCard = document.getElementById('hover-card');
    setTimeout(() => {
        if (!hoverCard.matches(':hover') && !obj.matches(':hover')) {
            hoverCard.style.display = 'none';
        }
    }, 300);
}

function cachedToMem(tweetArray, cacheObj) {
    if (!tweetArray) {
        return;
    }
    tweetArray.map(tweet => {
        __globalTweetMemCache.set(tweet.create_time, tweet);
        __globalTweetMemCacheByHash.set(tweet.prefixed_hash, tweet);
        if (tweet.create_time < cacheObj.latestID || cacheObj.latestID === 0) {
            cacheObj.latestID = tweet.create_time;
        }
        cacheObj.CachedItem.push(tweet);
    });
}

async function TweetsQuery(param, newest, cacheObj) {
    try {
        param.start_id = newest ? 0 : cacheObj.latestID;
        if (newest) {
            cacheObj.latestID = 0;
        }
        const tweetArray = await PostToSrvByJson("/tweetQuery", param);

        cacheObj.moreOldTweets = newest || (tweetArray && tweetArray.length !== 0);

        cachedToMem(tweetArray, cacheObj);

        return cacheObj.CachedItem.length > 0;
    } catch (err) {
        console.log(err);
        throw new Error(err);
    }
}

async function __setOnlyHeader(tweetHeader, twitter_id, web3ID) {
    const twitterObj = await TwitterBasicInfo.loadTwBasicInfo(twitter_id);
    const njUsrInfo = await loadNJUserInfoFromSrv(web3ID);

    if (twitterObj) {
        tweetHeader.querySelector('.twitterAvatar').src = twitterObj.profile_image_url;
        if (njUsrInfo && njUsrInfo.is_elder) {
            tweetHeader.querySelector('.elderFlagOnAvatar').style.display = 'block';
        }else{
            tweetHeader.querySelector('.elderFlagOnAvatar').style.display = 'none';
        }
        tweetHeader.querySelector('.twitterName').textContent = twitterObj.name;
        tweetHeader.querySelector('.twitterUserName').textContent = '@' + twitterObj.username;
        return twitterObj;
    }

    const newObj = await loadTwitterUserInfoFromSrv(twitter_id, true)
    if (!newObj) {
        console.log("failed load twitter user info");
        return null;
    }
    tweetHeader.querySelector('.twitterAvatar').src = newObj.profile_image_url;
    // console.log(njUsrInfo);
    if (njUsrInfo && njUsrInfo.is_elder) {
        tweetHeader.querySelector('.elderFlagOnAvatar').style.display = 'block';
    }else{
        tweetHeader.querySelector('.elderFlagOnAvatar').style.display = 'none';
    }
    tweetHeader.querySelector('.twitterName').textContent = newObj.name;
    tweetHeader.querySelector('.twitterUserName').textContent = '@' + newObj.username;

    return newObj;
}

async function showImgRaw(event) {
    event.stopPropagation();
    try {
        showWaiting("loading.....");
        const hash = this.getAttribute('data-hash');
        const obj = await loadTweetImgRaw(hash);
        if (!obj) {
            showDialog(DLevel.Warning, "failed to load raw image");
            return;
        }
        const imgDiv = document.querySelector('.tweet-image-raw')
        imgDiv.style.display = 'block';
        imgDiv.querySelector('.tweet-image-detail').src = obj.raw_data;
        imgDiv.querySelector('.tweet-image-hash').innerText = obj.hash;
    } catch (e) {
        showDialog(DLevel.Error, e.toString());
    } finally {
        hideLoading();
    }
}

function CloseImgDetail() {
    document.querySelector('.tweet-image-raw').style.display = 'none';
}

async function loadTweetImgRaw(hash) {
    let obj = await ImageRawData.load(hash)
    if (obj && obj.raw_data) {
        return obj;
    }

    const response = await GetToSrvByJson("/tweetImgRaw?img_hash=" + hash);
    if (obj) {
        obj.raw_data = response.raw;
    } else {
        obj = new ImageRawData(response.hash, response.raw);
    }
    ImageRawData.sycToDb(obj);
    return obj;
}

async function loadTweetImgThumb(hash) {
    let obj = await ImageRawData.load(hash)
    if (obj && obj.thumb_nail) {
        // console.log("no need to query--->", hash)
        return obj;
    }

    const response = await GetToSrvByJson("/tweetImgThumb?img_hash=" + hash);
    if (obj) {
        obj.thumb_nail = response.raw
    } else {
        obj = new ImageRawData(response.hash, null, response.raw);
    }
    ImageRawData.sycToDb(obj);
    // console.log('found from server =>', obj)
    return obj;
}

async function procTweetTxt(text) {
    let txt = text.replace(/\t/g, "&nbsp;&nbsp;&nbsp;&nbsp;").replace(/\n/g, "<br>").replace(/ /g, '&nbsp;')
    const regex = /<dessage-img>(.*?)<\/dessage-img>/g;
    const result = txt.replace(regex, (match, imgHash) => {
        // console.log(imgHash);
        const cleanedStr = imgHash.replace(/\s+/g, '');
        const images = cleanedStr.split(delimiter);
        // console.log(images);
        const imgManagerDiv = document.getElementById('image-in-tweet-template').cloneNode(true);
        imgManagerDiv.id = '';
        imgManagerDiv.removeAttribute('id');
        imgManagerDiv.style.display = 'grid';

        for (let i = 0; i < images.length; i++) {
            const imgHash = images[i];
            const imgDiv = imgManagerDiv.querySelector('.image-item-in-tweet').cloneNode(true)
            imgDiv.style.display = 'block';
            imgDiv.removeAttribute('id');
            const imgElm = imgDiv.querySelector('.image-src-to-show');
            imgElm.setAttribute('data-hash', imgHash);
            loadTweetImgThumb(imgHash).then(imgObj => {
                if (!imgObj) {
                    console.log("failed to load thumb img=>", imgHash);
                    return
                }
                const selector = `[data-hash="${imgHash}"]`;
                const element = document.querySelectorAll(selector);
                element.forEach(elm=>{
                    elm.src = imgObj.thumb_nail;
                    elm.onclick = showImgRaw;
                })
            });
            imgManagerDiv.appendChild(imgDiv);
        }
        return imgManagerDiv.outerHTML;
    });
    return result;
}

async function setupCommonTweetHeader(tweetHeader, tweet, overlap) {
    tweetHeader.querySelector('.tweetCreateTime').textContent = formatTime(tweet.create_time);
    const twitterObj = await __setOnlyHeader(tweetHeader, tweet.twitter_id, tweet.web3_id);
    const contentArea = tweetHeader.querySelector('.tweet-content');
    const txt = await procTweetTxt(tweet.text);
    contentArea.innerHTML = DOMPurify.sanitize(txt);
    const wrappedHeader = tweetHeader.querySelector('.tweet-header');

    // fulfillTweetImages(tweet, tweetHeader);

    if (overlap) {
        wrappedHeader.addEventListener('mouseenter', (event) => showHoverCard(event, twitterObj, tweet.web3_id));
        wrappedHeader.addEventListener('mouseleave', () => hideHoverCard(wrappedHeader));
    }
    return contentArea;
}

function showSettingBtn(show) {
    const settingDiv = document.getElementById("user-tooltip")
    if (show) {
        settingDiv.style.visibility = 'visible';
    } else {
        settingDiv.style.visibility = 'hidden';
    }
}

async function refreshTwitterInfo() {
    showWaiting("syncing twitter meta");
    await loadTwitterUserInfoFromSrv(ninjaUserObj.tw_id, false, true);
    await setupUserBasicInfoInSetting();
    hideLoading();
}

function quitFromService() {
    fetch("/signOut", {method: 'GET'}).then(() => {
        window.location.href = "/signIn";
    }).catch(err => {
        console.log(err)
        window.location.href = "/signIn";
    })
}

function showTweetDetailInfo() {
    const div = document.getElementById('tweet-detail-info')
    if (div.style.display === 'none') {
        div.style.display = 'flex';
        document.getElementById('tweet-text-info-img').classList.remove('tweet-text-info-img');
    } else {
        document.getElementById('tweet-text-info-img').classList.add('tweet-text-info-img');
        div.style.display = 'none';
    }
}

async function showTweetDetail(parentEleID, tweet) {
    const detail = document.querySelector('#tweet-detail');
    detail.style.display = 'block';
    document.getElementById('tweet-post-on-top').style.display = 'none';
    const parentNode = document.getElementById(parentEleID);
    if (!parentNode) {
        return;
    }
    parentNode.style.display = 'none';

    detail.querySelector('.tweetCreateTime').textContent = formatTime(tweet.create_time);
    await __setOnlyHeader(detail, tweet.twitter_id, tweet.web3_id);
    const txt = await procTweetTxt(tweet.text);
    detail.querySelector('.tweet-text').innerHTML = DOMPurify.sanitize(txt);

    // fulfillTweetImages(tweet, detail);

    detail.querySelector('.back-button').onclick = () => {
        parentNode.style.display = 'block';
        detail.style.display = 'none';
        document.getElementById('tweet-post-on-top').style.display = 'block';
        detail.querySelector('.team-members').innerHTML = '';
        detail.querySelector('.team-members').style.display = 'none';
    }
    detail.querySelector('.tweet-web3_id').textContent = tweet.web3_id;
    detail.querySelector('.tweet-prefixed-hash').textContent = tweet.prefixed_hash;
    detail.querySelector('.tweet-signature').textContent = tweet.signature;
    detail.querySelector('.tweet-vote-number').textContent = tweet.vote_count;

    await __showVoteButton(detail, tweet, function (newVote) {
        detail.querySelector('.tweet-vote-number').textContent = newVote.vote_count;
    });
}

async function __showVoteButton(tweetCard, tweet, callback) {
    const voteBtn = tweetCard.querySelector('.tweet-action-vote');
    if (!voteContractMeta) {
        await initVoteContractMeta();
    }
    // tweetCard.querySelector('.tweet-action-vote-val').textContent = voteContractMeta.votePriceInEth;
    voteBtn.onclick = () => voteToTheTweet(tweet, callback);
}

async function __updateVoteNumberForTweet(tweetObj, newVote) {

    let tweetCard = document.getElementById("tweet-card-for-vote-" + tweetObj.create_time)
    if (tweetCard) {
        tweetCard.querySelector('.total-vote-number').textContent = tweetObj.vote_count;
        if (newVote) {
            const userVoteCounter = tweetCard.querySelector('.user-vote-number');
            userVoteCounter.textContent = newVote.user_vote_count;
            cachedVoteStatusForUser.set(newVote.create_time, newVote.user_vote_count);
        }
    }

    tweetCard = document.getElementById("tweet-card-for-user-" + tweetObj.create_time)
    if (tweetCard) {
        tweetCard.querySelector('.vote-number').textContent = tweetObj.vote_count;
    }

    tweetCard = document.getElementById("tweet-card-for-home-" + tweetObj.create_time)
    if (tweetCard) {
        tweetCard.querySelector('.vote-number').textContent = tweetObj.vote_count;
    }

    tweetCard = document.getElementById("tweet-card-for-njusr-vote-" + tweetObj.create_time)
    if (tweetCard) {
        tweetCard.querySelector('.total-vote-count').textContent = tweetObj.vote_count;
    }

    tweetCard = document.getElementById("tweet-card-for-njusr-post-" + tweetObj.create_time)
    if (tweetCard) {
        tweetCard.querySelector('.total-vote-count').textContent = tweetObj.vote_count;
    }
}

async function voteToTheTweet(obj, callback) {

    if (Number(obj.payment_status) !== TXStatus.Success) {
        showDialog(DLevel.Warning, "can't vote to unpaid tweet")
        return;
    }

    openVoteModal(function (voteCount, shareToTweet) {
        procTweetVotePayment(voteCount, obj, async function (create_time, vote_count, txHash) {
            const newVote = await updateVoteStatusToSrv(create_time, vote_count, txHash);
            obj.vote_count = newVote.vote_count;
            __updateVoteNumberForTweet(obj, newVote).then(() => {
            });

            if (shareToTweet && ninjaUserObj.tw_id) {
                const slogan = i18next.t('slogan_1') + gameContractMeta.totalBonus + " ETH. " + i18next.t('voter-slogan');
                await  __shareVoteToTweet(create_time, vote_count, slogan);
            }
            reloadSelfNjData().then(() => {
            });
            loadUserPointsInfos().then(r=>{
                console.log("load user points success")
            });

            if (callback) {
                callback(newVote);
            }
        });
    });
}

async function updateVoteStatusToSrv(create_time, vote_count, txHash) {
    return await PostToSrvByJson("/updateTweetVoteStatus", {
        create_time: create_time,
        vote_count: Number(vote_count),
        tx_hash: txHash,
    });
}

async function withdrawAction(contract) {
    try {
        showWaiting('prepare transaction')
        const txResponse = await contract.withdraw("0x00", true);
        changeLoadingTips("transaction packaging:" + txResponse.hash);

        const txReceipt = await txResponse.wait();
        showDialog(DLevel.Success, "Transaction: " + txReceipt.status ? "success" : "failed");
    } catch (err) {
        checkMetamaskErr(err);
    } finally {
        hideLoading();
    }
}

async function showTargetTweetDetail() {
    if (!targetTweet || !targetTweet.create_time) {
        return;
    }

    await showTweetDetail('tweets-park', targetTweet);

    const protocol = window.location.protocol;
    const host = window.location.host;
    const rootUrl = protocol + "//" + host;
    const newUrl = rootUrl + '/main';

    history.pushState(null, '', newUrl);
}

async function reloadSelfNjData() {
    let nj_data;
    try {
        nj_data = await GetToSrvByJson("/refreshNjUser");
        ninjaUserObj = nj_data;
        await setupUserBasicInfoInSetting();
    } catch (err) {
        console.log(err)
        showDialog(DLevel.Warning, "reload session failed:" + err.toString())
    }
}

function showTmpTips(msg) {
    const alert = document.getElementById("temporary-alert-popup");
    alert.className = "temporary-alert-popup show";
    alert.querySelector('.tips-content-msg').innerText = msg;
    setTimeout(function () {
        alert.className = alert.className.replace("show", "");
    }, 3000);
}