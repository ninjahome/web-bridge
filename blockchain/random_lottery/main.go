package main

import (
	"flag"
	"fmt"
	"github.com/ninjahome/web-bridge/util"
)

func testData() {
	passphraseFlag := flag.String("p", "", "Encryption passphrase")
	decryptFlag := flag.Bool("d", false, "Decrypt mode")
	encryptedDataFlag := flag.String("e", "", "Encrypted data for decryption")

	flag.Parse()

	if *decryptFlag == false {
		fmt.Println(util.GenerateRandomData(*passphraseFlag))
		return
	}

	if *passphraseFlag == "" || *encryptedDataFlag == "" {
		fmt.Println("Passphrase and encrypted data are required for decryption")
		flag.Usage()
		return
	}
	fmt.Println(util.DecryptRandomData(*encryptedDataFlag, *passphraseFlag))
	return
}

func main() {
	//walletFile := flag.String("wallet", "", "wallet path")
	//
	//fmt.Print("Enter Password: ")
	//passwordBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	//if err != nil {
	//	fmt.Println("\nError reading password")
	//	return
	//}
	//
	//password := string(passwordBytes)

}
