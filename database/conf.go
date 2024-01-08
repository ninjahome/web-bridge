package database

import "fmt"

var __dbConf *FileStoreConf

type FileStoreConf struct {
	ProjectID      string `json:"project_id"`
	DatabaseID     string `json:"database_id"`
	KeyFilePath    string `json:"key_file_path"`
	TweetsPageSize int    `json:"tweets_page_size"`
	LocalRun       bool   `json:"local_run"`
}

func (c *FileStoreConf) String() string {
	s := "\n------file store config------"
	s += "\nproject id:" + c.ProjectID
	s += "\nkey path :" + c.KeyFilePath
	s += "\ntweet page size :" + fmt.Sprintf("%d", c.TweetsPageSize)
	s += "\n--------------------------"
	return s
}
func InitConf(c *FileStoreConf) {
	__dbConf = c
	_ = DbInst()

}
