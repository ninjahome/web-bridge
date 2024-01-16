SHELL=PATH='$(PATH)' /bin/sh

.PHONY: all
all: abi
	abigen --abi eth_abis/TweetVote.abi --pkg ethapi --type TweetVote --out ./ethapi/tweet_vote.go
	abigen --abi eth_abis/TweetLotteryGame.abi --pkg ethapi --type TweetLotteryGame --out ./ethapi/game_lottery.go
	abigen --abi eth_abis/KolKeys.abi --pkg ethapi --type KolKeys --out ./ethapi/kol_keys.go
	abigen --abi eth_abis/KolKeySimple.abi --pkg ethapi --type KolKeySimple --out ./ethapi/kol_key_simple.go

abi:
	solc --abi --overwrite *.sol -o ./eth_abis

clean:
	rm -f *.abi
	rm -f *.go
