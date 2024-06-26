
function toHex(number) {
    return '0x' + number.toString(16);
}

const __globalContractConf = new Map([
    [toHex(421614), {
        tweetVote: "0x4c9aF3E257eC34Ce5b32816896A84F733B6ADb9D",
        gameLottery: "0x7e1394f0b7C460eE717cea8A7D509379ae757D2e",
        kolKey: "0x8836D288165F6136b2be970DAed0612B37dD0B6A",
        postPrice: "0.005",
        votePrice: "0.005"
    }],
    [toHex(42161), {
        tweetVote: "0x6500Cda46979F1956a46486B1a88768cb425E23a",
        gameLottery: "0x1B85654c3576972Bdd457413E9BEcCc27B52f6a1",
        kolKey: "",
        postPrice: "0.005",
        votePrice: "0.005"
    }]]);

const __globalTargetChainNetworkID = toHex(421614);