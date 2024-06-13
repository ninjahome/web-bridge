package database

import "fmt"

const (
	DefaultElderThreshold       = 100
	DefaultBonusRateForReferred = 0.2 //20%
	DefaultBonusForReferrer     = 100
	DefaultPointsForOneRound    = 100.0
)

var __dbConf *FileStoreConf

type FileStoreConf struct {
	ProjectID               string  `json:"project_id"`
	DatabaseID              string  `json:"database_id"`
	KeyFilePath             string  `json:"key_file_path"`
	TweetsPageSize          int     `json:"tweets_page_size"`
	LocalRun                bool    `json:"local_run"`
	PointForPost            float32 `json:"point_for_post"`
	PointForVote            float32 `json:"point_for_vote"`
	PointForBeVote          float32 `json:"point_for_be_vote"`
	ElderNoFirstGot         int     `json:"elder_no_first_got"`
	BonusForReferer         float32 `json:"bonus_for_referer"`
	BonusRateForReferred    float32 `json:"bonus_rate_for_referred"`
	RewardPointsForOneRound float32 `json:"reward_points_for_one_round"`
}

func (c *FileStoreConf) String() string {
	s := "\n------file store config------"
	s += "\nlocal run:" + fmt.Sprintf("%t", c.LocalRun)
	s += "\nproject id:" + c.ProjectID
	s += "\ndatabase id:" + c.DatabaseID
	s += "\nkey path :" + c.KeyFilePath
	s += "\ntweet page size :" + fmt.Sprintf("%d", c.TweetsPageSize)
	s += "\npoint for vote :" + fmt.Sprintf("%d", c.PointForVote)
	s += "\npoint for be voted :" + fmt.Sprintf("%d", c.PointForBeVote)
	s += "\nelder threshold :" + fmt.Sprintf("%d", c.ElderNoFirstGot)
	s += "\nbonus for referrer :" + fmt.Sprintf("%f", c.BonusForReferer)
	s += "\nbonus rate when referred :" + fmt.Sprintf("%f", c.BonusRateForReferred)
	s += "\npoint bonus for one round:" + fmt.Sprintf("%.2f", c.RewardPointsForOneRound)
	s += "\n--------------------------"
	return s
}

func InitConf(c *FileStoreConf) {

	__dbConf = c

	if c.ElderNoFirstGot == 0 {
		c.ElderNoFirstGot = DefaultElderThreshold
	}
	if c.BonusRateForReferred <= 0.0 {
		c.BonusRateForReferred = DefaultBonusRateForReferred
	}
	if c.BonusForReferer <= 0.0 {
		c.BonusForReferer = DefaultBonusForReferrer
	}

	if c.RewardPointsForOneRound == 0 {
		c.RewardPointsForOneRound = DefaultPointsForOneRound
	}

	_ = DbInst()
}
