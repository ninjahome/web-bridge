package server

import (
	"encoding/json"
	"os"
	"testing"
)

func TestCreateDefaultConfigFile(t *testing.T) {
	cfg := &Conf{
		&SrvConf{
			DebugMode:   true,
			UseHttps:    false,
			SSLCertFile: "",
			SSLKeyFile:  "",
		},
		&TwitterConf{
			ClientID:     "",
			ClientSecret: "",
		},
		&FileStoreConf{
			ProjectID:   DefaultProjectID,
			KeyFilePath: "dessage-c3b5c95267fb.json",
		},
	}

	bts, _ := json.MarshalIndent(cfg, "", "\t")
	_ = os.WriteFile("../config.sample.json", bts, 0644)
}
