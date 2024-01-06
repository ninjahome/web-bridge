package blockchain

type Conf struct {
	TweeTVoteContractAddress string `json:"tweet_vote_contract_address"`
	GameContract             string `json:"game_contract"`
	KolKeyContractAddress    string `json:"kol_key_contract_address"`
	InfuraUrl                string `json:"infura_url"`
	GameTimeInMinute         int    `json:"game_time_in_minute,omitempty"`
	TxCheckerInSeconds       int    `json:"tx_checker_in_seconds,omitempty"`
	ChainID                  int64  `json:"chain_id,omitempty"`
}

func (c *Conf) String() string {
	s := "\n------block chain config------"
	s += "\ntweet vote:" + c.TweeTVoteContractAddress
	s += "\ngame:" + c.GameContract
	s += "\nkol key:" + c.KolKeyContractAddress
	s += "\ninfura url:" + c.InfuraUrl
	s += "\n--------------------------"
	return s
}
