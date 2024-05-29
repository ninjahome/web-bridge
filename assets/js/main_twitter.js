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
        showDialog(DLevel.Error, "failed to load tweets");
    }
}

async function loadTweetsForHomePage() {
    const tweetsDiv = document.getElementById('tweets-park');
    tweetsDiv.style.display = 'block';
    document.getElementById('tweet-post-on-top').style.display = 'block';
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
            let tw_data = await TwitterBasicInfo.loadTwBasicInfo(twitterID)
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
        } else {
            showMoreBtn.textContent = i18next.t('tweet-show-more');
            showMoreBtn.style.display = 'block';
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

async function convertTweetContentToImg(formattedContent, format = 'image/jpeg', quality = 1.0) {
    try {
        const target = document.getElementById('hidden-tweet-for-img');
        const content = document.getElementById('hidden-tweet-txt');

        content.innerText = formattedContent;
        const canvas = await html2canvas(target, {
            dpi: 300, // Set a higher DPI
            scale: 2  // Increase the scaling level
        });
        const imgURL = canvas.toDataURL(format, quality);
        content.innerText = "";
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
    return compositedTxt;
}

async function convertContentToImages(formattedContent, imageData) {
    let tmpSplitStr = formattedContent;
    for (let i = 0; i < maxImgPerTweet; i++) {
        // const substr = safeSubstring(tmpSplitStr, maxTweetLenPerPage);
        const substr = tmpSplitStr.substring(0, maxTweetLenPerPage);
        // console.log("substring len:", substr.length, substr);
        const txtImg = await convertTweetContentToImg(substr);
        imageData.push(new ImageRawData("converted-" + i, txtImg, ""));
        tmpSplitStr = tmpSplitStr.substring(substr.length);
        // console.log("last content len:", tmpSplitStr.length);
        if (substr.length < maxTweetLenPerPage) {
            break
        }
    }
}

function parseTweetContent(parentDiv) {
    const allTweetDiv = parentDiv.querySelectorAll(".tweets-content-txt-area");

    const txtList = Array.from(allTweetDiv)
        .map(div => {
            const validElm = div.firstChild;
            if (!validElm) {
                return null;
            }
            let textContent = validElm.textContent;
            if (!textContent) {
                return null;
            }
            textContent = textContent.replace(/\u200B/g, '').trim();
            // console.log(`Original text: '${textContent}'`);
            return textContent;
        })
        .filter(txt => {
            const isValid = (txt && txt.length > 0);
            // console.log(`Filter result for '${txt}': ${isValid}`);
            return isValid;
        });

    const formattedTxt = txtList.join('\n');
    const images = parentDiv.querySelectorAll("#twImagePreview img");
    if (txtList.length === 0) {
        showDialog(DLevel.Warning, "content too short")
        return null;
    }

    if (txtList.length > MaxTweetsPerPost + 1) {
        showDialog(DLevel.Warning, "content too long")
        return null;
    }

    if (images.length > maxImgPerTweet) {
        showDialog(DLevel.Warning, "too many images to post");
        return null;
    }
    const imageData = Array.from(images).map(img => {
        const thumbnail = img.src
        const raw = img.getAttribute('data-raw');
        const hash = img.getAttribute('data-hash');
        return new ImageRawData(hash, raw, thumbnail);
    });

    return {formattedTxt, txtList, imageData}
}

function initSloganTxt(nj_tw_id) {
    return "\n" + i18next.t('slogan_1')
        + gameContractMeta.totalBonus
        + i18next.t('slogan_2')
        + "https://" + window.location.hostname + "/buyRights?NjTID="
        + nj_tw_id;
}

async function procTweetContent(tweetContent, slogan) {

    let compositedTxt = tweetContent.formattedTxt + slogan;
    let result = twttr.txt.parseTweet(compositedTxt);
    if (result.valid === true) {
        return compositedTxt;
    }

    const sloganLen = twttr.txt.getTweetLength(slogan);
    if (maxImgPerTweet === tweetContent.imageData.length) {
        showDialog(DLevel.Warning, "tweet content should be short than" + (defaultTextLenForTweet - sloganLen) + " if you have 4 images");
        return null;
    }

    const lastValidTxtLen = maxTweetLenPerPage * (maxImgPerTweet - tweetContent.imageData.length);
    if (tweetContent.formattedTxt.length >= lastValidTxtLen) {
        showDialog(DLevel.Warning, "max tweet content length is:" + lastValidTxtLen + " if you have " + tweetContent.imageData.length + " images");
        return null;
    }

    await convertContentToImages(tweetContent.formattedTxt, tweetContent.imageData)

    return getCompositedTxt(tweetContent.formattedTxt, slogan, sloganLen);
}

async function preparePostMsg(parentDiv) {
    const tweetContent = parseTweetContent(parentDiv);
    if (!tweetContent || tweetContent.txtList.length === 0) {
        return null;
    }

    const nj_tw_id = (new Date()).getTime();
    const slogan = initSloganTxt(nj_tw_id);
    const lastIdx = tweetContent.txtList.length - 1;
    const lastStr = tweetContent.txtList[lastIdx];

    let result = twttr.txt.parseTweet(lastStr + slogan);
    if (result.valid === true) {
        tweetContent.txtList[lastIdx] += slogan;
    } else {
        tweetContent.txtList.push(slogan);
    }

    const tweet = new TweetContentToPost(tweetContent.formattedTxt,
        tweetContent.txtList, nj_tw_id, ninjaUserObj.eth_addr, ninjaUserObj.tw_id);
    const message = JSON.stringify(tweet)

    const signature = await window.ethereum.request({
        method: 'personal_sign', params: [tweetContent.formattedTxt, ninjaUserObj.eth_addr],
    });

    if (!signature) {
        showDialog(DLevel.Warning, "empty signature")
        return null;
    }

    return new SignDataForPost(message, signature, JSON.stringify(tweetContent.imageData));
}

function updatePaymentStatusToSrv(tweet, tx_hash) {
    return PostToSrvByJson("/updateTweetPaymentStatus", {
        create_time: tweet.create_time,
        status: tweet.payment_status,
        hash: tweet.prefixed_hash,
        tx_hash: tx_hash,
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
        showWaiting("preparing dessage")
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

        await procPaymentForPostedTweet(basicTweet, async function (newTweet, txHash) {
            await updatePaymentStatusToSrv(newTweet, txHash);
            if (curScrollContentID === 0) {
                __loadTweetsAtHomePage(true).then(() => {
                });
            } else if (curScrollContentID === 2) {
                __loadTweetAtUserPost(true, ninjaUserObj.eth_addr).then(() => {
                });
            }
        });

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

    initTweetArea('modal-split-tweet-content');

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
    const splitArea = parentDiv.querySelector(".split-tweet-content");
    splitArea.innerHTML = '';
    __globalTweetEditorCount = 0;
    newSplitEditor(splitArea);
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
        this.setAttribute('data-more', 'true');
        this.innerText = i18next.t('tweet-show-more');
    }
}

function loadImgFromLocal() {
    if (!ninjaUserObj.tw_id) {
        showDialog(DLevel.Warning, "bind twitter first", bindingTwitter);
        return;
    }
    const tweetItem = this.closest('.tweet-split-item');
    const images = tweetItem.querySelectorAll(".img-wrapper-container img");
    if (images.length >= maxImgPerTweet) {
        showDialog(DLevel.Tips, "max " + maxImgPerTweet + " images allowed")
        return;
    }
    tweetItem.querySelector('.tweet-file-input').click();
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

    if (files.length > validLen) {
        showDialog(DLevel.Tips, "max " + maxImgPerTweet + " images allowed")
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
            if (rawBase64Str.length > MaxRawImgSize) {
                rawBase64Str = await adjustImageToApproxTargetBase64Length(blob, MaxRawImgSize);
            }

            img.setAttribute('data-raw', rawBase64Str);
            const hash = ethers.hashMessage(rawBase64Str);
            img.setAttribute('data-hash', hash);

            if (blob.src.length > MaxThumbnailSize) {
                img.src = await adjustImageToApproxTargetBase64Length(blob, MaxThumbnailSize);
            } else {
                img.src = blob.src;
            }
            imagePreviewDiv.appendChild(imgWrapper);
        });
    });

    parentDiv.querySelector('.tweet-file-input').value = '';
}

function initTweetArea(divID) {
    const tweetManager = document.getElementById(divID);
    __globalTweetEditorCount = 0;
    tweetManager.innerHTML = '';
    newSplitEditor(tweetManager);
}

function addNextSplitEditor(btn) {
    const tweetManager = btn.closest(".split-tweet-content");
    const siblingNode = btn.closest(".tweet-split-item");
    newSplitEditor(tweetManager, siblingNode);
}

let __globalTweetEditorCount = 0;

function newSplitEditor(tweetManager, siblingNode) {
    if (__globalTweetEditorCount >= MaxTweetsPerPost) {
        showDialog(DLevel.Warning, "too much tweets");
        return;
    }
    const tweetEditorTemplate = document.getElementById("tweet-split-item-template");
    const newEditor = tweetEditorTemplate.cloneNode(true);
    newEditor.style.display = 'block';
    newEditor.id = 'tweet-area-' + __globalTweetEditorCount;
    const editableDiv = newEditor.querySelector('.tweets-content-txt-area');
    editableDiv.innerHTML = '';
    editableDiv.addEventListener('compositionstart', () => {
        isComposing = true;
    });
    editableDiv.addEventListener('compositionend', () => {
            isComposing = false;
            checkTweetLength(editableDiv);
        }
    );
    editableDiv.addEventListener('input', () => {
        if (!ninjaUserObj.tw_id) {
            showDialog(DLevel.Warning, "bind twitter first", bindingTwitter);
            return;
        }
        if (isComposing) {
            return;
        }
        checkTweetLength(editableDiv);
    });
    editableDiv.addEventListener('keydown', handleEnter);

    __globalTweetEditorCount++;
    if (siblingNode) {
        siblingNode.insertAdjacentElement('afterend', newEditor);
    } else {
        tweetManager.appendChild(newEditor);
    }
    // setCursorToStart(newEditor);
}

let isComposing = false;

function delCurrentEditor(btn) {
    if (__globalTweetEditorCount === 1) {
        return;
    }
    const parentDiv = btn.closest('.tweet-split-item');
    parentDiv.remove();
    __globalTweetEditorCount--;
}

function checkSelection() {
    if (window.getSelection) {
        const selection = window.getSelection().toString();
        console.log("Selected text: ", selection);
    }
}

function checkTweetLength(div, isNewLine = false) {
    const tweetTxt = div.innerText;
    const parsedText = twttr.txt.parseTweet(tweetTxt);

    let validText = tweetTxt.substring(0, parsedText.validRangeEnd + 1);
    let excessText = tweetTxt.substring(parsedText.validRangeEnd + 1);
    let restore = saveCaretPosition(div);  // 先保存光标位置

    while (div.firstChild) {
        div.removeChild(div.firstChild);
    }

    let validTextNode = document.createTextNode(validText);
    div.appendChild(validTextNode);

    if (excessText) {
        let newExcess = document.createElement('span');
        newExcess.className = 'tweet-over-flow-red';
        newExcess.innerText = excessText;
        div.appendChild(newExcess);
    }

    restore(isNewLine);

    const parentDiv = div.closest('.tweet-split-item');
    parentDiv.querySelector('.tweet-length-valid').innerText = 280 - parsedText.weightedLength;
}

function handlePaste(event) {
    event.preventDefault();  // 阻止默认粘贴行为
    const clipboardData = event.clipboardData || window.clipboardData;  // 获取剪贴板对象
    const text = clipboardData.getData('text/plain');  // 从剪贴板获取纯文本内容
    // 插入文本到光标位置
    insertTextAtCursor(text);
    checkTweetLength(event.target);
}

function insertTextAtCursor(text) {
    const selection = window.getSelection();
    if (!selection.rangeCount) return;  // 如果没有选区，则不执行任何操作

    const range = selection.getRangeAt(0);
    range.deleteContents();  // 删除选中内容

    // 插入文本并处理换行符
    const lines = text.split('\n');
    const fragment = document.createDocumentFragment();
    lines.forEach((line, index) => {
        if (index > 0) fragment.appendChild(document.createElement('br'));
        fragment.appendChild(document.createTextNode(line));
    });

    range.insertNode(fragment);  // 插入修改后的文本

    // 移动光标到文本末尾
    range.collapse(false);
    selection.removeAllRanges();  // 清除现有的选区
    selection.addRange(range);  // 添加新的范围
}

function saveCaretPosition(context) {
    let selection = window.getSelection();
    if (selection.rangeCount === 0) return () => {
    };

    let activeRange = selection.getRangeAt(0);
    let range = document.createRange();
    range.setStart(context, 0);
    range.setEnd(activeRange.startContainer, activeRange.startOffset);
    let length = range.toString().length;

    return function restore(isNewLine) {
        selection.removeAllRanges();
        let range = document.createRange();
        let nodeStack = [context], node;
        let remainingLength = length;
        if (isNewLine) {
            remainingLength++;
        }

        while (node = nodeStack.pop()) {
            if (node.nodeType === 3) { // 文本节点
                if (remainingLength <= node.length) {
                    range.setStart(node, remainingLength);
                    break;
                } else {
                    remainingLength -= node.length;
                }
            } else {
                let i = node.childNodes.length;
                while (i--) {
                    nodeStack.push(node.childNodes[i]);
                }
            }
        }

        if (range.startContainer && range.startContainer.parentNode) {
            range.collapse(true);
            selection.addRange(range);
            // console.log('Caret position restored:', {node: range.startContainer, offset: range.startOffset});
        } else {
            console.warn('Range not in document, cannot restore caret position');
        }
    };
}

function handleEnter(event) {
    if (event.key === 'Enter') {
        event.preventDefault(); // 阻止默认的回车效果

        const selection = window.getSelection();
        if (!selection.rangeCount) return; // 如果没有选区，则不执行后续操作

        const range = selection.getRangeAt(0);
        // console.log('Current range start:', range.startContainer, range.startOffset);
        range.deleteContents();

        const br = document.createElement('br');
        range.insertNode(document.createTextNode('\u200B'));
        range.insertNode(document.createElement('br'));
        selection.removeAllRanges();
        selection.addRange(range);

        // console.log('New cursor position set after <br>:', range.startContainer, range.startOffset);
        checkTweetLength(event.target, true);
    }
}