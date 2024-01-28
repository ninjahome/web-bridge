
function toHex(number) {
    return '0x' + number.toString(16);
}

const __globalContractConf = new Map([
    [toHex(421614), {
        tweetVote: "0xbC086f9cF06Bc186a1a6eB619b50302dD347518C",
        gameLottery: "0x1939C2E865A691eff314D25b3005D4aF9df27076",
        kolKey: "0x8836D288165F6136b2be970DAed0612B37dD0B6A",
        postPrice: "0.005",
        votePrice: "0.005"
    }],
    [toHex(42161), {
        tweetVote: "",
        gameLottery: "",
        kolKey: "",
        postPrice: "0.005",
        votePrice: "0.005"
    }]]);

const __globalTargetChainNetworkID = toHex(421614);