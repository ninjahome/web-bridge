
let curScrollContentID = 0;
document.addEventListener("DOMContentLoaded", initMainPage);

function showGameRule() {
    const rule_div = document.getElementById('gameRuleInfo')
    if (rule_div.style.display === 'none') {
        rule_div.style.display = 'block';
    } else {
        rule_div.style.display = 'none';
    }
}

async function initNjBasic() {
    if (!ninjaUserObj) {
        window.location.href = "/signIn";
        return;
    }

    NJUserBasicInfo.cacheNJUsrObj(ninjaUserObj);

    let tw_data = await TwitterBasicInfo.loadTwBasicInfo(ninjaUserObj.tw_id)
    if (tw_data) {
        document.getElementById("dessage-log-to-nj-log").src = tw_data.profile_image_url;
        if (ninjaUserObj.is_elder) {
            document.getElementById("dessage-is-elder-flag").style.display = 'block';
        } else {
            document.getElementById("dessage-is-elder-flag").style.display = 'none';
        }
    }
    loadUserPointsInfos().then();
}

async function initMainPage() {
    translatePage();
    try {
        showWaiting("loading from block chain");
        initTweetArea('split-tweet-content');
        await initDatabase();
        await initNjBasic();
        await checkMetaMaskEnvironment(initBlockChainContract);
        await loadTweetsForHomePage();
        await showTargetTweetDetail();
    } catch (err) {
        console.log(err);
        showDialog(DLevel.Warning, err.toString());
    } finally {
        hideLoading();
    }
}

async function menuOnClicked(curMenuId) {
    curScrollContentID = curMenuId;
    document.getElementById('hover-card').style.display = 'none';
    document.querySelectorAll('.menu-item-left-area').forEach(i => {
        i.classList.remove('active');
        const icon = i.querySelector('.menu-icon');
        icon.src = i.getAttribute('data-icon-inactive');
    });
    this.classList.add('active');
    const activeIcon = this.querySelector('.menu-icon');
    activeIcon.src = this.getAttribute('data-icon-active');

    const contentId = this.getAttribute('data-content');
    document.querySelectorAll('.content-in-middle-area').forEach(c => c.classList.remove('active'));
    document.getElementById(contentId).classList.add('active');
    document.getElementById("tweet-detail").style.display = 'none';
    document.getElementById("nj-user-profile").style.display = 'none';

    if (contentId === "middle-div-user-tweets") {
        await loadTweetsUserPosted();
    } else if (contentId === "middle-div-home") {
        await loadTweetsForHomePage()
    } else if (contentId === "middle-div-leaderboard") {
        await initTopPage();
    } else if (contentId === "middle-div-about-user") {
        await initAboutUserPage();
    }
}

function openLotteryMainPage() {
    window.location.href = "/lotteryGame";
}

async function setupUserBasicInfoInSetting() {

    document.getElementById('tweet-no-has-income').innerText = ninjaUserObj.tweet_count;
    document.getElementById("basic-web3-id").innerText = ninjaUserObj.eth_addr;
    document.getElementById("referral-Code").innerText = ninjaUserObj.eth_addr.slice(-6);
    document.getElementById("dessage-my-vote").innerText = ninjaUserObj.vote_count;
    document.getElementById("dessage-be-vote").innerText = ninjaUserObj.be_voted_count;

    if (!ninjaUserObj.tw_id) {
        return;
    }
    const twitterInfo = await loadTwitterUserInfoFromSrv(ninjaUserObj.tw_id, true, false);
    if (!twitterInfo) {
        console.log("no such twitter user");
        return;
    }
    document.getElementById("user-twitter-avatar").src = twitterInfo.profile_image_url;
    document.getElementById("user-setting-name").innerText = twitterInfo.name;
    document.getElementById("user-setting-user-name").innerText = twitterInfo.username;
}

async function initAboutUserPage() {
    try {

        showWaiting("syncing from block chain");

        await reloadGameBalance();
        await reloadTweetBalance();

        const obj = await loadNJUserInfoFromSrv(ninjaUserObj.eth_addr, false);
        if (!obj) {
            console.log("can't be empty")
            return;
        }
        ninjaUserObj = obj;
        await setupUserBasicInfoInSetting();
        loadUserPointsInfos().then(r=>{
            console.log("load user points success")
        });
    } catch (err) {
        checkMetamaskErr(err);
    } finally {
        hideLoading();
    }
}

function translatePage() {
    i18next
        .use(i18nextBrowserLanguageDetector)
        .init({
            fallbackLng: 'en',
            debug: false,
            resources: {
                en: {
                    translation: {
                        "title": "Dessage",
                        "left-menu-item-home": "Home",
                        "left-menu-item-top": "Leaderboard",
                        "left-menu-item-my-tweet": "MyTweet",
                        "left-menu-item-my-setting": "Setting",
                        "btn-tittle-post-tweet": "Post Tweet",
                        "btn-post-tweet-on-left-menu": "Post",
                        "cancelButton": "Cancel",
                        "confirmButton": "Confirm",
                        "btn-tittle-share-when-vote": "Share to Twitter",
                        "vote-dialog-vote-no": "Votes",
                        "vote-to-tweet-dialog-tittle": "Vote On Tweet",
                        "tweet-show-less": "Less",
                        "tweet-show-more": "More",
                        "tittle-top-vote": "Most Voted",
                        "tittle-top-best-writer": "Top Creators Leaderboard",
                        "tittle-top-best-voter": "Whale Honor Roll",
                        "show-details-button": "Show More >>",
                        "hover-tittle-tweet-no": "Tweet No",
                        "hover-tittle-vote-no": "Voted No",
                        "hover-tittle-voted-no": "Vote No",
                        "label-vote-no": "Votes",
                        "btn-tittle-vote-no": "Vote",
                        "tweet-detail-creator": "Creator Web3 ID",
                        "tweet-detail-hash": "Tweet Hash",
                        "tweet-detail-signature": "Tweet Signature",
                        "tweet-detail-vote-no": "Vote No",
                        "tweet-detail-vote": "Vote",
                        "back": "Back",
                        "twee-switch-area-his-tweet": "His/Her Posts",
                        "twee-switch-area-his-vote": "His/Her Votes",
                        "nj-user-tweet-his-vote-no": "His/Her Votes",
                        "user-tweet-btn-my-tweet": "My Tweet",
                        "user-tweet-btn-my-vote": "My Vote",
                        "user-tweet-btn-goon-pay": "Pay It",
                        "user-tweet-btn-delete": "Delete",
                        "user-tweet-label-total-vote": "Total Vote",
                        "user-tweet-label-my-vote": "My Vote",
                        "user-setting-clear": "Clean Cache",
                        "user-setting-quit": "Sign Out",
                        "user-setting-sync-twitter": "Sync Twitter",
                        "user-setting-points": "Points",
                        "user-setting-my-voted": "Number of Votes",
                        "user-setting-be-voted": "Number of Votes Received",
                        "user-setting-tweet-income": "Tweet Revenue",
                        "user-setting-tweet-no": "Total Tweets",
                        "withdraw-history": "Withdrawal Records",
                        "withdraw": "Withdraw",
                        "user-setting-game-income": "Game Revenue",
                        "user-setting-lottery-income-tittle": "Lottery Revenue",
                        "user-setting-tweet-no-suffix": "Tweet(s)",
                        "lottery-round-tittle-1": "Phase ",
                        "lottery-round-tittle-2": "",
                        "lottery-tittle-tips": "Current Prize Pool",
                        "countdown-result": "Drawing in Progress",
                        "countdown-tips-days": "D(s)",
                        "countdown-tips-sub-tittle": "Next Draw Time",
                        "lottery-info-history-bonus": "Total Historical Prize",
                        "lottery-info-current-tickets-no": "Total Votes This Phase",
                        "lottery-stats-unit-2": " ",
                        "lottery-info-btn-mine": "My Tickets",
                        "lottery-info-tittle": "Game Rule",
                        "lottery-info-content": "This game is a lottery game, with draws every 48 hours that anyone can participate in.",
                        "lottery-info-content-1": "The method of purchasing lottery tickets",
                        "lottery-info-content-2": "1.Reward The Tweet",
                        "lottery-info-content-3": "30% of any tips on your tweets will go into the lottery prize pool, 30% into the points prize pool, 30% to the tweet author, and 10% to Dessage protocol income.",
                        "lottery-info-content-4": "2. Place a separate bet.",
                        "lottery-info-content-5": "You can also place individual bets on the 'My Tickets' page, with 45% of the bet amount going into the lottery prize pool, 45% into the points prize pool, and 10% belonging to Dessage protocol income.",
                        "lottery-info-content-6": "During Dessage's Chapter 1 phase, it will continue to accumulate, with the amount being public. Upon entering Chapter 2, it will be distributed proportionally according to all users' points situation. The ways users can earn points are as follows:",
                        "lottery-info-content-7": "1. Each vote cast earns 2 points.",
                        "lottery-info-content-8": "2. Each tweet posted earns 2 points.",
                        "lottery-info-content-9": "3. When a user's tweet is voted on by other users, the creator can earn points at a rate of 1 point per vote.",
                        "lottery-info-content-10": "4.At the end of Chapter 1, the voters of the top 10 tweets with the highest total votes will receive additional points rewards based on the chronological order of their votes. The top 10% of voters by voting time under each tweet will be rewarded with points ten times their number of votes, the next 10%-20% will receive nine times the points, and so on.",
                        "lottery-info-content-11": "5. At the end of Chapter 1, the top 10 users on the Whale Leaderboard can receive an additional 50% bonus to their points.",
                        "lottery-info-content-13": "*Points are not only usable for participating in the distribution of the points prize pool, but they can also be used to obtain other rights and interests in Dessage in the future.",
                        "lottery-info-content-14": "To post a tweet participating in the game, you need to bind your Arbitrum address and Twitter account. Then, by clicking to tweet on this platform, your tweet will be synchronously forwarded to your Twitter account. Other users can then vote and tip under your tweet, allowing you to earn creator income.",
                        "lottery-info-content-15": "6. During Chapter 1, the first 10 KOLs to reach 100 votes can receive the Elder Medal. When any tweet posted by these 10 KOLs receives votes, the Elder can earn twice the points for votes compared to ordinary creators.",
                        "lottery-info-content-16": "7.Referral System: Users who log in to the system using a referral code can receive a 100-point reward. This reward is not given all at once but accumulates as the user earns other points until reaching 100 points. Additionally, the referrer will receive a 20% bonus of the points earned by the referred user.",
                        "lottery-info-content-17": "8.Points Mining: The system distributes 100 points every 8 hours. The system takes a snapshot of the points every 8 hours. The user's mining reward is equal to the proportion of the user's points to the total points at the time of the snapshot, and the 100 points are divided based on that proportion.",
                        "bonus-for-points-tittle": "Points Prize Pool",
                        "voter-slogan": "Come to this blockchain social network to grab lottery tickets!",
                        "slogan_1": "Current Prize Pool: ",
                        "tweet-detail-in-details": "Tweet Info",
                        "slogan_2": "ETH，Vote on this tweet and participate in the prize pool. link:",
                        "vote-price-in-modal-uint": "Vote",
                        "hidden-appoint-tips": "Powered By Dessage",
                        "point-bonus-annual-tittle": "Annualized:",
                        "referral-code-btn-txt": "Referral Code",
                        "referral-code-tips": "Click to copy the referral code",
                        "user-setting-bonus": "Unclaimed reward",
                    }
                },
                zh: {
                    translation: {
                        "title": "Dessage",
                        "left-menu-item-home": "主页",
                        "left-menu-item-top": "排行榜",
                        "left-menu-item-my-tweet": "推文",
                        "left-menu-item-my-setting": "我的",
                        "btn-tittle-post-tweet": "发布推文",
                        "btn-post-tweet-on-left-menu": "发推",
                        "cancelButton": "取消",
                        "confirmButton": "确定",
                        "btn-tittle-share-when-vote": "转发到推特",
                        "vote-dialog-vote-no": "票数",
                        "vote-to-tweet-dialog-tittle": "推文投票",
                        "tweet-show-less": "更少",
                        "tweet-show-more": "更多",
                        "tittle-top-vote": "总得票排行榜",
                        "tittle-top-best-writer": "最佳创作者排行榜",
                        "tittle-top-best-voter": "鲸鱼荣誉榜",
                        "show-details-button": "查看详情 >>",
                        "hover-tittle-tweet-no": "推文总数",
                        "hover-tittle-vote-no": "被打赏数",
                        "hover-tittle-voted-no": "打赏数",
                        "label-vote-no": "次投票",
                        "btn-tittle-vote-no": "投票",
                        "tweet-detail-creator": "作者Web3 ID",
                        "tweet-detail-hash": "推文哈希",
                        "tweet-detail-signature": "推文数字签名",
                        "back": "返回",
                        "user-tweet-btn-my-tweet": "我的推文",
                        "user-tweet-btn-my-vote": "我的投票",
                        "user-tweet-btn-goon-pay": "继续支付",
                        "user-tweet-btn-delete": "删除",
                        "user-tweet-label-total-vote": "总票数",
                        "user-tweet-label-my-vote": "我的票数",
                        "user-setting-clear": "清除缓存",
                        "user-setting-quit": "退出",
                        "user-setting-sync-twitter": "同步推特",
                        "user-setting-points": "积分",
                        "user-setting-my-voted": "投票数",
                        "twee-switch-area-his-tweet": "他发布推文",
                        "twee-switch-area-his-vote": "他投票推文",
                        "nj-user-tweet-his-vote-no": "他投票推文",
                        "user-setting-be-voted": "被投票数",
                        "user-setting-tweet-income": "推文收入(eth)",
                        "user-setting-tweet-no": "推文总数",
                        "withdraw-history": "提现记录",
                        "withdraw": "提现",
                        "user-setting-game-income": "游戏收入",
                        "user-setting-lottery-income-tittle": "彩票收入(eth)",
                        "user-setting-tweet-no-suffix": "篇",
                        "lottery-round-tittle-1": "第 ",
                        "lottery-round-tittle-2": "期",
                        "lottery-tittle-tips": "本期奖池",
                        "countdown-result": "开奖中",
                        "countdown-tips-days": "天",
                        "countdown-tips-sub-tittle": "开奖时间",
                        "lottery-info-history-bonus": "历史累计奖金(eth)",
                        "lottery-info-current-tickets-no": "本期总票数",
                        "lottery-stats-unit-2": "张",
                        "lottery-info-btn-mine": "我的彩票",
                        "lottery-info-tittle": "玩法规则",
                        "lottery-info-content": "本游戏是一个彩票游戏，每48小时开奖一次任何人都可参与",
                        "lottery-info-content-1": "购买彩票方式",
                        "lottery-info-content-2": "1.推文下打赏",
                        "lottery-info-content-3": "你的任何推文打赏都会有30%进入彩票奖池，30%进入积分奖池，30%归属推文作者，10%归属Dessage协议收入",
                        "lottery-info-content-4": "2.单独投注",
                        "lottery-info-content-5": "你也可以在“我的彩票”页面进行单独投注，单独投注的投注额的45%进入彩票奖池，45%进入积分奖池，10%归属Dessage协议收入",
                        "lottery-info-content-6": "在Dessage的chapter 1阶段持续累积，数额公开，在进入chapter 2阶段时将按照所有用户的积分情况等比例分配。用户获得积分的途径如下",
                        "lottery-info-content-7": "1.每投票一票可获得2分",
                        "lottery-info-content-8": "2.每发推一次可获得2分",
                        "lottery-info-content-9": "3.用户发布的推文被其他用户投票时，创作者可按1分每票获得积分",
                        "lottery-info-content-10": "4.chapter 1结束时，总得票数排名前10的推文,其投票用户可依据投票时间先后获得额外积分奖励。每篇推文下投票时间排名前10%的用户奖励投票数10倍的积分，前10%-20%的可获得额外9倍的积分，以此类推。",
                        "lottery-info-content-11": "5.chapter 1结束时，大户排行榜前10位的用户可获得额外积分加成50%",
                        "lottery-info-content-13": "*积分不仅仅可用于参与积分奖池的分配，后续还能获得Dessage的其他权益",
                        "lottery-info-content-14": "要发布参与游戏的推文，需绑定你的arbitrum地址和推特账号，然后在本平台点击发推即可，推文会同步转发至你的推特账号，其他用户即可在你的推文下投票打赏，你可获得创作者收益",
                        "lottery-info-content-15": "6.chapter 1进行中，获得投票数率先达到100的前10位创作者可以获得元老勋章，这10位创作者的所有推文被投票时，元老可以获得2倍于普通创作的被投票积分。",
                        "lottery-info-content-16": "7.推荐制度:使用推荐码登录本系统的用户，可以获得100积分奖励，该积分不是一次性奖励，而是在获取其它积分的时候，奖励同样的积分并累积到100为止。同时推荐人可以在被推荐人获取积分时，获取该积分的20%的积分奖励。",
                        "lottery-info-content-17": "8.积分挖矿:系统每8小时发放100积分，系统会每8小时取一次积分快照，用户挖矿奖励等于快照时用户的积分占总积分的比例，根据该比例瓜分100积分。",
                        "bonus-for-points-tittle": "积分奖池",
                        "slogan_1": "当前彩池: ",
                        "tweet-detail-in-details": "推文信息",
                        "slogan_2": "ETH，给本推文投票并参与彩池 链接:",
                        "vote-price-in-modal-uint": "票",
                        "hidden-appoint-tips": "来自 Dessage",
                        "point-bonus-annual-tittle": "年化积分:",
                        "referral-code-btn-txt": "推荐码",
                        "referral-code-tips": "点击复制推荐码",
                        "user-setting-bonus": "未兑奖励",

                    }
                }
            }
        }, function (err, t) {
            updateContent();
        });
}

function updateContent() {
    document.title = i18next.t('title');

    document.getElementById('left-menu-item-home').textContent = i18next.t('left-menu-item-home');
    document.getElementById('left-menu-item-top').textContent = i18next.t('left-menu-item-top');
    document.getElementById('left-menu-item-my-tweet').textContent = i18next.t('left-menu-item-my-tweet');
    document.getElementById('left-menu-item-my-setting').textContent = i18next.t('left-menu-item-my-setting');
    document.getElementById('btn-post-tweet-on-left-menu').textContent = i18next.t('btn-post-tweet-on-left-menu');
    document.getElementById('cancelButton').textContent = i18next.t('cancelButton');
    document.getElementById('confirmButton').textContent = i18next.t('confirmButton');
    document.getElementById('btn-tittle-share-when-vote').textContent = i18next.t('btn-tittle-share-when-vote');
    document.getElementById('vote-dialog-vote-no').textContent = i18next.t('vote-dialog-vote-no');
    document.getElementById('vote-to-tweet-dialog-tittle').textContent = i18next.t('vote-to-tweet-dialog-tittle');

    document.getElementById('tittle-top-vote').textContent = i18next.t('tittle-top-vote');
    document.getElementById('tittle-top-best-writer').textContent = i18next.t('tittle-top-best-writer');
    document.getElementById('tittle-top-best-voter').textContent = i18next.t('tittle-top-best-voter');

    document.getElementById('hover-tittle-tweet-no').textContent = i18next.t('hover-tittle-tweet-no');
    document.getElementById('hover-tittle-vote-no').textContent = i18next.t('hover-tittle-vote-no');
    document.getElementById('hover-tittle-voted-no').textContent = i18next.t('hover-tittle-voted-no');

    document.getElementById('label-vote-no-home').textContent = i18next.t('label-vote-no');
    document.getElementById('btn-tittle-vote-no-home').textContent = i18next.t('btn-tittle-vote-no');

    document.getElementById('label-vote-no-top').textContent = i18next.t('label-vote-no');
    document.getElementById('btn-tittle-vote-no-top').textContent = i18next.t('btn-tittle-vote-no');

    document.getElementById('tweet-detail-creator').textContent = i18next.t('tweet-detail-creator');
    document.getElementById('tweet-detail-hash').textContent = i18next.t('tweet-detail-hash');
    document.getElementById('tweet-detail-signature').textContent = i18next.t('tweet-detail-signature');

    document.getElementById('tweet-detail-vote-no').textContent = i18next.t('label-vote-no');
    document.getElementById('tweet-detail-vote').textContent = i18next.t('btn-tittle-vote-no');
    document.getElementById('back-button-tweet-tittle').textContent = i18next.t('back');

    document.getElementById('user-tweet-btn-my-tweet').textContent = i18next.t('user-tweet-btn-my-tweet');
    document.getElementById('user-tweet-btn-my-vote').textContent = i18next.t('user-tweet-btn-my-vote');
    document.getElementById('user-tweet-label-vote-no').textContent = i18next.t('label-vote-no');
    document.getElementById('user-tweet-btn-delete').textContent = i18next.t('user-tweet-btn-delete');
    document.getElementById('user-tweet-label-total-vote').textContent = i18next.t('user-tweet-label-total-vote');
    document.getElementById('user-tweet-label-my-vote').textContent = i18next.t('user-tweet-label-my-vote');
    document.getElementById('user-tweet-btn-vote').textContent = i18next.t('btn-tittle-vote-no');

    document.getElementById('user-setting-quit').textContent = i18next.t('user-setting-quit');
    document.getElementById('user-setting-clear').textContent = i18next.t('user-setting-clear');
    document.getElementById('user-setting-points').textContent = i18next.t('user-setting-points');
    document.getElementById('user-setting-bonus').textContent = i18next.t('user-setting-bonus');
    document.getElementById('user-setting-my-voted').textContent = i18next.t('user-setting-my-voted');
    document.getElementById('user-setting-be-voted').textContent = i18next.t('user-setting-be-voted');
    document.getElementById('user-setting-tweet-no').textContent = i18next.t('user-setting-tweet-no');
    document.getElementById('user-setting-tweet-income-tittle').textContent = i18next.t('user-setting-tweet-income');
    document.getElementById('user-setting-withdraw-history').textContent = i18next.t('withdraw-history');
    document.getElementById('user-setting-withdraw').textContent = i18next.t('withdraw');
    document.getElementById('user-setting-lottery-income-tittle').textContent = i18next.t('user-setting-lottery-income-tittle');
    document.getElementById('user-setting-withdraw-history2').textContent = i18next.t('withdraw-history');
    document.getElementById('user-setting-withdraw2').textContent = i18next.t('withdraw');
    document.getElementById('user-setting-tweet-no-suffix').textContent = i18next.t('user-setting-tweet-no-suffix');
    document.getElementById('user-setting-tweet-no-suffix').textContent = i18next.t('user-setting-tweet-no-suffix');

    document.getElementById('twee-switch-area-his-tweet').textContent = i18next.t('twee-switch-area-his-tweet');
    document.getElementById('twee-switch-area-his-vote').textContent = i18next.t('twee-switch-area-his-vote');
    document.getElementById('nj-user-tweet-vote-no').textContent = i18next.t('user-setting-my-voted');
    document.getElementById('nj-user-tweet-his-vote-no').textContent = i18next.t('nj-user-tweet-his-vote-no');

    document.getElementById('lottery-round-tittle-1').textContent = i18next.t('lottery-round-tittle-1');
    document.getElementById('lottery-round-tittle-2').textContent = i18next.t('lottery-round-tittle-2');
    document.getElementById("countdown-result").textContent = i18next.t('countdown-result');
    document.getElementById("countdown-tips-days").textContent = i18next.t('countdown-tips-days');
    document.getElementById("countdown-tips-sub-tittle").textContent = i18next.t('countdown-tips-sub-tittle');
    document.getElementById("lottery-info-history-bonus").textContent = i18next.t('lottery-info-history-bonus');
    document.getElementById("lottery-info-current-tickets-no").textContent = i18next.t('lottery-info-current-tickets-no');
    document.getElementById("lottery-info-btn-mine").textContent = i18next.t('lottery-info-btn-mine');

    document.getElementById("lottery-info-tittle").textContent = i18next.t('lottery-info-tittle');
    document.getElementById("lottery-gameRule").textContent = i18next.t('lottery-info-tittle');
    document.getElementById("bonus-for-points-tittle").textContent = i18next.t('bonus-for-points-tittle');
    document.getElementById("bonus-for-points-tittle-2").textContent = i18next.t('bonus-for-points-tittle');
    document.getElementById("lottery-info-content").textContent = i18next.t('lottery-info-content');
    document.getElementById("lottery-info-content-1").textContent = i18next.t('lottery-info-content-1');
    document.getElementById("lottery-info-content-2").textContent = i18next.t('lottery-info-content-2');
    document.getElementById("lottery-info-content-3").textContent = i18next.t('lottery-info-content-3');
    document.getElementById("lottery-info-content-4").textContent = i18next.t('lottery-info-content-4');
    document.getElementById("lottery-info-content-5").textContent = i18next.t('lottery-info-content-5');
    document.getElementById("lottery-info-content-6").textContent = i18next.t('lottery-info-content-6');
    document.getElementById("lottery-info-content-7").textContent = i18next.t('lottery-info-content-7');
    document.getElementById("lottery-info-content-8").textContent = i18next.t('lottery-info-content-8');
    document.getElementById("lottery-info-content-9").textContent = i18next.t('lottery-info-content-9');
    document.getElementById("lottery-info-content-10").textContent = i18next.t('lottery-info-content-10');
    document.getElementById("lottery-info-content-11").textContent = i18next.t('lottery-info-content-11');
    document.getElementById("lottery-info-content-13").textContent = i18next.t('lottery-info-content-13');
    document.getElementById("lottery-info-content-14").textContent = i18next.t('lottery-info-content-14');
    document.getElementById("lottery-info-content-15").textContent = i18next.t('lottery-info-content-15');
    document.getElementById("lottery-info-content-16").textContent = i18next.t('lottery-info-content-16');
    document.getElementById("lottery-info-content-17").textContent = i18next.t('lottery-info-content-17');

    document.getElementById("user-tweet-btn-goon-pay").textContent = i18next.t('user-tweet-btn-goon-pay');
    document.getElementById("tweet-post-with-eth-btn-txt-1").textContent = i18next.t('btn-tittle-post-tweet');
    document.getElementById("tweet-post-with-eth-btn-txt-2").textContent = i18next.t('btn-tittle-post-tweet');

    document.getElementById("tweet-detail-in-details").textContent = i18next.t('tweet-detail-in-details');
    document.getElementById("vote-price-in-modal-uint").textContent = i18next.t('vote-price-in-modal-uint');

    document.getElementById('point-bonus-unit').textContent = i18next.t('user-setting-points');
    document.getElementById('point-bonus-annual-tittle').textContent = i18next.t('point-bonus-annual-tittle');
    document.getElementById('referral-code-btn-txt').textContent = i18next.t('referral-code-btn-txt');
    document.getElementById('referral-code-tips').textContent = i18next.t('referral-code-tips');

}