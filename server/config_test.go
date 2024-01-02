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
		LogLevel: "debug",
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
		BlockChainConf: &BlockChainConf{
			TweeTVoteContractAddress:  "0x63713037a9E337D7Db5D383070199B948598e0Da",
			GamePluginContractAddress: "0x57F0bbE85f5822911003A8fa425D5595D139FDFe",
			KolKeyContractAddress:     "",
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

	hash, retErr := util.Verify(obj.EthAddr, message, signedMessage)
	fmt.Println("verify result :", retErr, "hash", hash)
}
func TestGenASessionKey(t *testing.T) {
	secretKey := util.RandomBytesInHex(16)
	fmt.Println(secretKey)
}

func TestVerify(t *testing.T) {
	hash, retErr := util.Verify("0x00a7539cc7cc54f08a761175aa678005ef91f4dc",
		`{"text":"const messageHash = ethers.utils.hashMessage(message);","create_time":1703758019240,"web3_id":"0x00a7539cc7cc54f08a761175aa678005ef91f4dc","twitter_id":"1472854871548170246"}`, "0x35759ef0f8749e1ca203d0da78e0048facf0e3fbf26dc953039de1e92484bf5b5ed8e5ec957df3d2b111ec8dd9d45c8000219a6e0af4c8ea9ed9bac80743036c1b")
	fmt.Println("verify result :", retErr, "hash", hash, hash == "0x7beff7fdd64827d2cc82dbbf9525a6f4712db6d9c6944594e669958374c66f22")
}
