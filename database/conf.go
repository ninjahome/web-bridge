package database

import "fmt"

var __dbConf *FileStoreConf

type FileStoreConf struct {
	ProjectID      string `json:"project_id"`
	DatabaseID     string `json:"database_id"`
	KeyFilePath    string `json:"key_file_path"`
	TweetsPageSize int    `json:"tweets_page_size"`
	LocalRun       bool   `json:"local_run"`
	PointForPost   int    `json:"point_for_post"`
	PointForVote   int    `json:"point_for_vote"`
	PointForBeVote int    `json:"point_for_be_vote"`
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
	s += "\n--------------------------"
	return s
}

func InitConf(c *FileStoreConf) {
	__dbConf = c
	_ = DbInst()

}
