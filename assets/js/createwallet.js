function generateMnemonic() {
    // 示例助记词
    const mnemonic = ethers.Wallet.createRandom().mnemonic.phrase;
    const words = mnemonic.split(" ");
    const wallet = ethers.Wallet.fromMnemonic(mnemonic);
    console.log(wallet);

    const table = document.getElementById("mnemonicTable");
    table.innerHTML = ""; // 清空表格

    let row = table.insertRow(); // 初始创建第一行
    words.forEach((word, index) => {
        if (index % 3 === 0 && index !== 0) { // 每3个单词换行，但跳过第一次（因为第一行已经创建）
            row = table.insertRow();
        }
        let cell = row.insertCell();
        cell.innerHTML = `${index + 1}. ${word}`;
    });
}
