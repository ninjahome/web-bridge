package server

import (
	"encoding/json"
	"fmt"
	"github.com/ninjahome/web-bridge/util"
	"os"
	"testing"
)

func TestCreateDefaultConfigFile(t *testing.T) {
	cfg := &SysConf{
		Log:      "debug",
		LocalRun: false,
		HttpPort: "8880",
		UrlHome:  "https://bridge.simplenets.org",
		HttpConf: &HttpConf{
			RefreshContent: true,
			UseHttps:       false,
			SSLCertFile:    "",
			SSLKeyFile:     "",
			SessionKey:     "",
		},
		TwitterConf: &TwitterConf{
			FontPath:       "util/Noto_Sans_SC.ttf",
			FontSize:       28.0,
			ClientID:       "",
			ClientSecret:   "",
			ConsumerKey:    "",
			ConsumerSecret: "",
		},
		FileStoreConf: &FileStoreConf{
			TweetsPageSize: 20,
			ProjectID:      DefaultTwitterProjectID,
			KeyFilePath:    "dessage-c3b5c95267fb.json",
		},
	}

	bts, _ := json.MarshalIndent(cfg, "", "\t")
	_ = os.WriteFile("../config.sample.json", bts, 0644)
}

func TestSignParam(t *testing.T) {
	message := `{"address":"0x2ba4e30628742e55e98e4a5253b510f5f2c60219","signTim":1702880038532}`
	obj := &SignInObj{}

	err := json.Unmarshal([]byte(message), obj)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(obj)
	signedMessage := "0x227caaf479a4fcc890694919d81dfbd3eb234137c93362f71411c7172584bcd1765d2924135fa15934f0f1543f1524a86de15d62b6383d737bd77eec334233ed1b"

	retErr := util.Verify(obj.EthAddr, message, signedMessage)
	fmt.Println("verify result :", retErr)
}
func TestGenASessionKey(t *testing.T) {
	secretKey := util.RandomBytesInHex(16)
	fmt.Println(secretKey)
}
