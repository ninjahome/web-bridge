window.onscroll = function () {
    throttle(() => loadMoreTweets(), 200);
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

let confirmCallback = null;

function openVoteModal(callback) {
    const modal = document.getElementById("vote-no-chose-modal");
    modal.style.display = "block";
    confirmCallback = callback;
}

function confirmVoteModal() {
    if (confirmCallback) {
        const voteCount = document.getElementById("voteCount").value;
        confirmCallback(voteCount);
    }
    closeVoteModal();
}

function closeVoteModal() {
    const modal = document.getElementById("vote-no-chose-modal");
    modal.style.display = "none";
}

function increaseVote() {
    const voteCount = document.getElementById("voteCount");
    voteCount.value = parseInt(voteCount.value) + 1;
}

function decreaseVote() {
    const voteCountElement = document.getElementById("voteCount");
    const newVoteCount = Math.max(1, parseInt(voteCountElement.value) - 1);
    voteCountElement.value = newVoteCount.toString();
}

function clearCachedData() {
    localStorage.clear();
    sessionStorage.clear();
    window.location.href = "/signIn";
}


function showHoverCard() {
    const tweetCard = this.closest('.tweet-card');
    const obj = JSON.parse(tweetCard.dataset.rawObj);

    const hoverCard = document.getElementById('hover-card');
    const rect = this.getBoundingClientRect();
    const avatar = this.querySelector('img').src;
    const name = this.querySelector('.name').textContent;
    const tweetCount = '0'; // obj.tweet_no;
    const voteCount = '0'; // obj.vote_count;

    // 设置悬浮卡片内容
    document.getElementById('hover-avatar').src = avatar;
    document.getElementById('hover-name').textContent = name;
    document.getElementById('hover-tweet-count').textContent = tweetCount;
    document.getElementById('hover-vote-count').textContent = voteCount;

    // 设置悬浮卡片的位置
    hoverCard.style.display = 'block';
    hoverCard.style.left = `${rect.left}px`;
    hoverCard.style.top = `${rect.bottom + window.scrollY}px`;
}

function hideHoverCard(obj) {
    // console.log(obj);
    if(obj){
        obj.style.display = 'none';
        return;
    }
    // 检查鼠标是否在 hover-card 或 tweet-header 上
    const hoverCard = document.getElementById('hover-card');
    setTimeout(() => {
        if (!hoverCard.matches(':hover') && !this.matches(':hover')) {
            hoverCard.style.display = 'none';
        }
    }, 300);
}