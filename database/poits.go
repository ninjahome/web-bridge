package database

type SysPoints struct {
	EthAddr    string `json:"eth_addr" firestore:"eth_addr"`
	ReferBonus int    `json:"refer_bonus" firestore:"refer_bonus"`
}
