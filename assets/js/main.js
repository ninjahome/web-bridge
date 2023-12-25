


function checkSystemEnvironment() {

    if (typeof window.ethereum === 'undefined') {
        window.location.href = "/signIn";
        return
    }
    metamaskObj = window.ethereum;
    metamaskObj.on('accountsChanged', metamaskAccountChanged);
    metamaskObj.on('chainChanged', metamaskChainChanged);
    metamaskObj.request({method: 'eth_chainId'}).then(chainID => {
        metamaskChainChanged(chainID).then(r=>{});
    })
}

async function metamaskChainChanged(chainId) {
    const chainBtn = document.getElementById('change-chain-id-button')
    const chainBalance = document.getElementById('basic-web3-balance')
    if (__globalTargetChainNetworkID === chainId) {
        chainBtn.style.display = 'none';
        chainBalance.style.display = 'inline-block';
        chainBalance.innerText = await metamaskAccBalance()
        return;
    }
    chainBalance.style.display = 'none';
    chainBtn.style.display = 'inline-block';
}

async function metamaskAccBalance() {
    const balance = await metamaskObj.request({
        method: 'eth_getBalance',
        params: [ninjaUserObj.eth_addr, 'latest'],
    });
    // console.log('eth balance:', balance);
    if (balance === "0x0") {
        return "0.00 eth";
    }
    if (balance < 10 ** 12) {
        return "< 0.000001 eth"
    }
    return parseFloat((balance / 10 ** 18).toFixed(6)) + ' eth';
}

function metamaskAccountChanged(accounts) {
    if (accounts.length === 0) {
        window.location.href = "/signOut";
        return;
    }
    window.location.href = "/signOut";
}

function setupBasicInfo() {
    const twBtn = document.getElementById('sign-up-by-twitter-button')
    const twNameLabel = document.getElementById('basic-twitter-name')
    const isVerifiedLabel = document.getElementById("basic-twitter-verified");
    document.getElementById('basic-web3-id').innerText = ninjaUserObj.eth_addr;
    if (!ninjaUserObj.tw_id) {
        twNameLabel.style.display = 'none';
        twBtn.style.display = 'inline-block';
    } else {
        twBtn.style.display = 'none';
        twNameLabel.style.display = 'inline-block';
        loadTwitterInfo(ninjaUserObj.tw_id).then(twInfo => {
            if (!twInfo) {
                twitterUserObj = null;
                return;
            }
            twitterUserObj = twInfo;
            twNameLabel.innerText = twInfo.name;
            if (!twInfo.verified){
                isVerifiedLabel.innerText = "Premium False";
            }else{
                isVerifiedLabel.innerText = "Premium True";
            }
            if (twInfo.profile_image_url) {
                document.getElementById('user-twitter-logo').src = twInfo.profile_image_url;
            }
        })
    }
}

function switchToWorkChain() {
    metamaskObj.request({
        method: 'wallet_switchEthereumChain',
        params: [{chainId: __globalTargetChainNetworkID}],
    }).catch((switchError) => {
        if (switchError.code !== 4902) {
            showDialog("error", "failed switching to arbitrum network");
            return;
        }

        const arbParam = __globalMetaMaskNetworkParam.get(__globalTargetChainNetworkID);
        metamaskObj.request({
            method: 'wallet_addEthereumChain',
            params: [arbParam],
        }).then(r => {
            console.log(r);
        }).catch(err => {
            showDialog("error", "add to network failed:" + err.toString());
        });
    });
}

function signUpByTwitter() {
    window.location.href = "/signUpByTwitter";
}

async function loadTwitterInfo(twitterID) {
    try {
        let tw_data = TwitterBasicInfo.loadTwBasicInfo(twitterID)
        if (tw_data) {
            return tw_data;
        }

        const response = await GetToSrvByJson("/queryTwBasicById");
        if (!response.ok) {
            console.log("query twitter basic info failed")
            return null;
        }

        const text = await response.text();
        console.log(text);
        return TwitterBasicInfo.cacheTwBasicInfo(text);
    } catch (err) {
        console.log("queryTwBasicById err:", err)
        showDialog("error", err.toString())
    }
}

function quitFromService() {
    fetch("/signOut", {method: 'GET'}).then(r => {
        window.location.href = "/signIn";
    }).catch(err => {
        console.log(err)
        window.location.href = "/signIn";
    })
}

class TweetContent {
    constructor(tweet_content, createAt, web3Id, twitterID, tweet_id, signature) {
        this.text = tweet_content;
        this.create_time = createAt;
        this.web3_id = web3Id;
        this.twitter_id = twitterID;
        this.tweet_id = tweet_id;
        this.signature = signature;
    }
}

async function postTweet() {
    try {
        const content = document.getElementById("tweets-content").value.trim();
        if (!content) {
            showDialog("tips", "content can't be empty")
            return;
        }

        const twitterID = ninjaUserObj.tw_id;
        if (!twitterID) {
            showDialog("tips", "bind your twitter first")
            return;
        }
        if (!metamaskObj) {
            window.location.href = "/signIn";
            return;
        }
        const web3Id = ninjaUserObj.eth_addr;
        const tweet = new TweetContent(content, (new Date()).getTime(), web3Id, twitterID);
        const message = JSON.stringify(tweet);
        const signature = await metamaskObj.request({
            method: 'personal_sign',
            params: [message, web3Id],
        })

        const obj = new SignDataForPost(message, signature, null)

        PostToSrvByJson("/postTweet", obj).then(resp => {
            console.log(resp);
            const refreshedTweet = JSON.parse(resp)
            document.getElementById("tweets-content").value='';
            showDialog("success","post success");

        }).catch(err => {
            console.log(err);
            showDialog("error",err.toString())
        })
    } catch (err) {
        showDialog("error",err.toString())
    }

    function LoadTweets() {
        const lastTweetId = 0;
        fetchData('/allNinjaTweets?lastTwId='+lastTweetId).then(tweets => {
            const tweetsPark = document.querySelector('.tweets-park');
            tweetsPark.innerHTML = '';

            tweets.forEach(tweet => {
                // 创建tweet-card元素
                const tweetCard = document.createElement('div');
                tweetCard.classList.add('tweet-card');

                // 头部：头像、用户名、时间
                const header = document.createElement('div');
                header.classList.add('tweet-header');
                header.innerHTML = `
                <img src="${tweet.avatar}" alt="Avatar">
                <span class="name">${tweet.name}</span>
                <span class="username">@${tweet.username}</span>
                <span class="time">${tweet.time}</span>
            `;

                // 内容
                const content = document.createElement('div');
                content.classList.add('tweet-content');
                content.textContent = tweet.content;

                // 底部：按钮和信息
                const footer = document.createElement('div');
                footer.classList.add('tweet-footer');
                footer.innerHTML = `
                <div class="tweet-action">
                    <button>$10U赞</button>
                    <span>${tweet.likes} 赞</span>
                </div>
                <div class="tweet-action">
                    <button>$10U踩</button>
                    <span>${tweet.dislikes} 踩</span>
                </div>
                <div class="tweet-info">彩票奖池200U</div>
                <div class="tweet-info">中奖者: ${tweet.winner}</div>
            `;

                // 将元素添加到tweet-card
                tweetCard.appendChild(header);
                tweetCard.appendChild(content);
                tweetCard.appendChild(footer);

                // 将tweet-card添加到tweets-park
                tweetsPark.appendChild(tweetCard);
            });
        });
    }

}

