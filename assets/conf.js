
function toHex(number) {
    return '0x' + number.toString(16);
}

const __globalContractConf = new Map([
    [toHex(421614), {
        tweetVote: "0xaCfAf616D89534793378979CeF5B4a0FEc15A08a",
        gameLottery: "0x398356dc8EDD5dc9eb897A120440cfe6F99253e9",
        kolKey: "0x8836D288165F6136b2be970DAed0612B37dD0B6A",
        postPrice: "0.005",
        votePrice: "0.005"
    }],
    [toHex(42161), {
        tweetVote: "0x6212246b6bE9D814DF4DC348370D590B41EDBB54",
        gameLottery: "0xbeC347D7A0914A1BB669e24204b98CeBa90DD456",
        kolKey: "",
        postPrice: "0.005",
        votePrice: "0.005"
    }]]);

const __globalTargetChainNetworkID = toHex(421614);