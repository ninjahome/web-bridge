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

function openLotteryModal(lotteryNumbers, teamNumbers) {
    const tbody = document.querySelector('#lottery-table tbody');
    tbody.innerHTML = '';

    lotteryNumbers.forEach((number, index) => {
        const tr = document.createElement('tr');
        tr.innerHTML = `<td>${number}</td><td>${teamNumbers[index]}</td>`;
        tbody.appendChild(tr);
    });

    document.getElementById('lottery-detail-modal').style.display = 'block';
}

function closeLotteryModal() {
    document.getElementById('lottery-detail-modal').style.display = 'none';
}

function clearCachedData() {
    localStorage.clear();
    sessionStorage.clear();
    window.location.href = "/signIn";
}