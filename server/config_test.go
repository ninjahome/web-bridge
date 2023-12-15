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
	}

	bts, _ := json.MarshalIndent(cfg, "", "\t")
	_ = os.WriteFile("config.json.sample", bts, 0644)
}
