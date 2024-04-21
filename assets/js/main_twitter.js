const cachedGlobalTweets = new MemCachedTweets();

function bindingTwitter() {
    showWaiting("prepare to bind twitter");
    window.location.href = "/signUpByTwitter";
}

async function __loadTweetsAtHomePage(newest) {
    try {
        const param = new TweetQueryParam(0, "", []);
        const needUpdateUI = await TweetsQuery(param, newest, cachedGlobalTweets);
        if (needUpdateUI) {
            await fillTweetParkAtHomePage(newest);
            cachedGlobalTweets.CachedItem = [];
        }
    } catch (err) {
        console.log(err);
        showDialog(DLevel.Error, err.toString());
    }
}

async function loadTweetsForHomePage() {
    const tweetsDiv = document.getElementById('tweets-park');
    tweetsDiv.style.display = 'block';
    showWaiting("loading.....");
    await __loadTweetsAtHomePage(true);
    hideLoading();
}

async function loadOlderTweetsForHomePage() {
    if (cachedGlobalTweets.latestID === 0) {
        console.log("no need to load older data");
        return;
    }
    return __loadTweetsAtHomePage(false)
}

async function loadTwitterUserInfoFromSrv(twitterID, useCache, syncFromTwitter) {
    try {
        if (syncFromTwitter) {
            useCache = false;
        }
        if (useCache) {
            let tw_data = TwitterBasicInfo.loadTwBasicInfo(twitterID)
            if (tw_data) {
                return tw_data;
            }
        }

        const response = await GetToSrvByJson("/queryTwBasicById?twitterID=" + twitterID + "&&forceSync=" + syncFromTwitter);
        if (!response) {
            return null;
        }
        return TwitterBasicInfo.cacheTwBasicInfo(response);
    } catch (err) {
        console.log("queryTwBasicById err:", err)
        return null;
    }
}

async function __fillNormalTweet(clear, parkID, data, templateId, cardID, overlap, detailType, callback) {
    const tweetsPark = document.getElementById(parkID);
    if (clear) {
        tweetsPark.innerHTML = '';
    }

    for (const tweet of data) {

        const tweetCard = document.getElementById(templateId).cloneNode(true);
        tweetCard.style.display = '';
        tweetCard.id = cardID + tweet.create_time;
        tweetCard.dataset.createTime = tweet.create_time;
        const tweetHeader = document.getElementById('tweet-header-template').cloneNode(true);
        tweetHeader.style.display = '';
        tweetHeader.id = "";


        const sibling = tweetCard.querySelector('.tweet-footer')
        const contentArea = await setupCommonTweetHeader(tweetHeader, tweet, overlap);

        if (TweetDetailSource.NoNeed !== detailType) {
            contentArea.onclick = () => showTweetDetail(parkID, tweet)
        }

        tweetCard.insertBefore(tweetHeader, sibling);

        if (callback) {
            callback(tweetCard, tweetHeader, tweet)
        }

        tweetsPark.appendChild(tweetCard);

        const showMoreBtn = tweetCard.querySelector('.show-more');
        if (contentArea.scrollHeight <= contentArea.clientHeight) {
            showMoreBtn.style.display = 'none';
            sibling.style.marginTop = '8px';
        } else {
            showMoreBtn.textContent = i18next.t('tweet-show-more');
            showMoreBtn.style.display = 'block';
            sibling.style.marginTop = '-12px';
        }
    }
}

async function fillTweetParkAtHomePage(clear) {

    return __fillNormalTweet(clear, 'tweets-park',
        cachedGlobalTweets.CachedItem, 'tweetTemplateForHomePage',
        "tweet-card-for-home-", true, TweetDetailSource.HomePage,
        function (tweetCard, tweetHeader, tweet) {
            tweetCard.querySelector('.vote-number').textContent = tweet.vote_count;
            __showVoteButton(tweetCard, tweet);
        });
}


let currentTargetIdx = 1

async function checkAtTarget() {
    const tweetsContentTxtArea = document.getElementById("tweets-content-txt-area");
    const text = tweetsContentTxtArea.innerText;
    const usrName = findAtTarget(text)
    if (!usrName) {
        return;
    }
    console.log(usrName);

    const data = await GetToSrvByJson('/searchTwitterUsr?q=' + usrName);
    if (!data) {
        return;
    }
    console.log(data);
}

function findAtTarget(text) {
    const regex = /(?:^|\s)@(\w+)/g;
    let match;
    if ((match = regex.exec(text)) !== null) {
        console.log(`Found mention: ${match[currentTargetIdx]}`);
        return match[currentTargetIdx];
    }
    return null;
}

async function convertTweetContentToImg(formattedContent) {
    try {
        const target = document.getElementById('hidden-tweet-txt');

        target.innerText = formattedContent;
        const canvas = await html2canvas(target);
        const imgURL = canvas.toDataURL("image/png");
        // document.getElementById('outputImage').src = imgURL;//debug infos
        target.innerText = "";
        return imgURL;
    } catch (e) {
        console.log(e);
    }
}

function getCompositedTxt(formattedContent, slogan, sloganLen) {
    let prefix = safeSubstring(formattedContent, defaultTextLenForTweet - 4 - sloganLen) + "...";
    let compositedTxt = prefix + slogan;
    const twtLen = twttr.txt.getTweetLength(compositedTxt);
    if (twtLen > defaultTextLenForTweet) {
        prefix = safeSubstring(formattedContent, (defaultTextLenForTweet - 4 - sloganLen) / 2) + "...";
        compositedTxt = prefix + slogan;
    }
    console.log(compositedTxt, twttr.txt.getTweetLength(compositedTxt));
    return compositedTxt;
}

async function preparePostMsg(parentDiv) {
    const contentHtml = parentDiv.querySelector(".tweets-content-txt-area").innerText.trim();
    const formattedContent = contentHtml
        .replace(/<br\s*[\/]?>/gi, "\n") // 将 <br> 标签转换为换行符
        .replace(/<\/?p>/gi, "\n") // 将 <p> 标签转换为换行符
        .replace(/<[^>]+>/g, ''); // 移除所有其他HTML标签

    if (formattedContent < 4 && images.length === 0) {
        showDialog(DLevel.Warning, "content too short")
        return null;
    }

    const images = parentDiv.querySelectorAll("#twImagePreview img");
    if (images.length > maxImgPerTweet) {
        showDialog(DLevel.Warning, "too many images to post");
        return null;
    }

    const nj_tw_id = (new Date()).getTime();
    const domainName = "https://" + window.location.hostname;
    const slogan = "\r\n" + i18next.t('slogan_1') + gameContractMeta.totalBonus + i18next.t('slogan_2') + domainName + "/buyRights?NjTID=" + nj_tw_id;
    const sloganLen = twttr.txt.getTweetLength(slogan);

    let compositedTxt = formattedContent + slogan;
    let result = twttr.txt.parseTweet(compositedTxt);
    console.log("Weighted Length:", result.weightedLength);

    const imageData = Array.from(images).map(img => {
        const thumbnail = img.src
        const raw = img.getAttribute('data-raw');
        const hash = img.getAttribute('data-hash');
        return new ImageRawData(hash, raw, thumbnail);
    });

    if (result.valid === false) {
        if (maxImgPerTweet === images.length) {
            showDialog(DLevel.Warning, "tweet content should be short than" + (defaultTextLenForTweet - sloganLen) + " if you have 4 images");
            return null;
        }

        const lastValidTxtLen = maxTweetLenPerPage * (maxImgPerTweet - images.length);
        console.log("simple len :", formattedContent.length);
        if (formattedContent.length >= lastValidTxtLen) {
            showDialog(DLevel.Warning, "max tweet content length is:" + lastValidTxtLen + "if you have " + images.length + " images");
            return null;
        }

        compositedTxt = getCompositedTxt(formattedContent, slogan, sloganLen);
        let tmpSplitStr = formattedContent;
        for (let i = 0; i < maxImgPerTweet; i++) {

            const substr = tmpSplitStr.substring(0, maxTweetLenPerPage);//safeSubstring(tmpSplitStr, maxTweetLenPerPage);
            console.log(substr, substr.length);
            const txtImg = await convertTweetContentToImg(substr);
            imageData.push(new ImageRawData("converted-" + i, txtImg, ""));
            tmpSplitStr = tmpSplitStr.substring(substr.length);
            console.log(tmpSplitStr, tmpSplitStr.length);
            if (substr.length < maxTweetLenPerPage){
                break
            }
        }
    }


    const tweet = new TweetContentToPost(formattedContent,
        nj_tw_id, ninjaUserObj.eth_addr, ninjaUserObj.tw_id, compositedTxt);
    const message = JSON.stringify(tweet)


    const signature = await window.ethereum.request({
        method: 'personal_sign', params: [message, ninjaUserObj.eth_addr],
    });

    if (!signature) {
        showDialog(DLevel.Warning, "empty signature")
        return null;
    }

    return new SignDataForPost(message, signature, JSON.stringify(imageData));
}

function updatePaymentStatusToSrv(tweet) {
    return PostToSrvByJson("/updateTweetPaymentStatus", {
        create_time: tweet.create_time,
        status: tweet.payment_status,
        hash: tweet.prefixed_hash
    }).then(r => {
        console.log(r);
    })
}

async function postTweetWithPayment(parentID) {

    if (!ninjaUserObj.tw_id) {
        showDialog(DLevel.Warning, "bind twitter first", bindingTwitter);
        return;
    }

    try {
        const parentDiv = document.querySelector(parentID)

        const tweetObj = await preparePostMsg(parentDiv);
        if (!tweetObj) {
            return;
        }
        closePostTweetDiv();
        showWaiting("posting to twitter");
        const basicTweet = await PostToSrvByJson("/postTweet", tweetObj);
        if (!basicTweet) {
            return;
        }
        clearDraftTweetContent(parentDiv);
        const paySuccess = await procPaymentForPostedTweet(basicTweet);
        if (!paySuccess) {
            return;
        }
        showWaiting("updating tweet status")
        await updatePaymentStatusToSrv(basicTweet)

        if (curScrollContentID === 0) {
            __loadTweetsAtHomePage(true).then(() => {
            });
        } else if (curScrollContentID === 2) {
            __loadTweetAtUserPost(true, ninjaUserObj.eth_addr).then(() => {
            });
        }
    } catch (err) {
        checkMetamaskErr(err);
    } finally {
        hideLoading();
    }
}

async function showPostTweetDiv() {
    if (!voteContractMeta) {
        showDialog(DLevel.Warning, "please change metamask to arbitrum network")
        return;
    }

    if (!ninjaUserObj.tw_id) {
        showDialog(DLevel.Warning, "bind twitter first", bindingTwitter);
        return;
    }

    const modal = document.querySelector('.modal-for-tweet-post');
    modal.style.display = 'block';
    document.getElementById('modal-overlay').style.display = 'block';


    const postBtn = document.getElementById("tweet-post-with-eth-btn-txt-2");
    postBtn.innerText = i18next.t('btn-tittle-post-tweet') + "(" + voteContractMeta.votePriceInEth + " eth)"
}

function closePostTweetDiv() {
    const modal = document.querySelector('.modal-for-tweet-post');
    modal.style.display = 'none';
    document.getElementById('modal-overlay').style.display = 'none';
}

function clearDraftTweetContent(parentDiv) {
    parentDiv.querySelector(".tweets-content-txt-area").innerHTML = '';
    parentDiv.querySelector(".img-wrapper-container").innerHTML = '';
    parentDiv.querySelector(".img-wrapper-container").style.display = 'none'
}

function showFullTweetContent() {

    const tweetCard = this.closest('.tweet-card');
    const tweetContent = tweetCard.querySelector('.tweet-content');
    const isMore = this.getAttribute('data-more') === 'true';

    if (isMore) {
        tweetContent.style.display = 'block';
        tweetContent.classList.remove('tweet-content-collapsed');
        tweetCard.style.maxHeight = 'none';
        this.innerText = i18next.t('tweet-show-less');
        this.setAttribute('data-more', 'false');
    } else {
        tweetContent.style.display = '-webkit-box';
        tweetContent.classList.add('tweet-content-collapsed');
        tweetCard.style.maxHeight = '400px';
        this.setAttribute('data-more', 'true');
        this.innerText = i18next.t('tweet-show-more');
    }
}

function loadImgFromLocal(parentId) {

    if (!ninjaUserObj.tw_id) {
        showDialog(DLevel.Warning, "bind twitter first", bindingTwitter);
        return;
    }

    const parentDiv = document.querySelector(parentId);
    const images = parentDiv.querySelectorAll("#twImagePreview img");
    if (images.length >= maxImgPerTweet) {
        showDialog(DLevel.Tips, "max " + maxImgPerTweet + " images allowed")
        return;
    }
    parentDiv.querySelector('.tweet-file-input').click();
}

function previewImage(parentId) {
    const parentDiv = document.querySelector(parentId);
    let files = parentDiv.querySelector('.tweet-file-input').files;
    const imagePreviewDiv = parentDiv.querySelector('.img-wrapper-container');
    imagePreviewDiv.style.display = 'block';
    const images = parentDiv.querySelectorAll("#twImagePreview img");
    const validLen = maxImgPerTweet - images.length;
    if (validLen <= 0) {
        return;
    }

    files = Array.from(files).slice(0, validLen);
    files.forEach(file => {
        const imgWrapper = parentDiv.querySelector('.img-wrapper').cloneNode(true);
        imgWrapper.style.display = 'block';
        imgWrapper.id = "";
        const img = imgWrapper.querySelector('.img-preview');
        const deleteBtn = imgWrapper.querySelector('.delete-btn');
        deleteBtn.onclick = function () {
            imagePreviewDiv.removeChild(imgWrapper);
        };

        readFileAsBlob(file).then(async blob => {
            let rawBase64Str = blob.src;
            // console.log("blob size:", rawBase64Str.length);
            if (rawBase64Str.length > MaxRawImgSize) {
                let quality = MaxRawImgSize / rawBase64Str.length * 0.75;
                if (quality > CompressQuality) {
                    quality = CompressQuality;
                }
                const compressedBlob = await compressBlob(blob, quality);
                rawBase64Str = await blobToBase64(compressedBlob);
                // console.log('image Base64 String:', rawBase64Str.length, compressedBlob.size);
            }

            img.setAttribute('data-raw', rawBase64Str);
            const hash = ethers.hashMessage(rawBase64Str);
            // console.log(hash);
            img.setAttribute('data-hash', hash);

            if (blob.src.length > MaxThumbnailSize) {
                let quality = MaxThumbnailSize / blob.src.length * 0.75;
                if (quality > CompressQuality) {
                    quality = CompressQuality;
                }
                const thumbNailBlob = await compressBlob(blob, quality);
                img.src = await blobToBase64(thumbNailBlob);
                // console.log("thumbNail size:", img.src.length, thumbNailBlob.size);
            } else {
                img.src = blob.src;
            }
            imagePreviewDiv.appendChild(imgWrapper);
        });
    });

    parentDiv.querySelector('.tweet-file-input').value = '';
}