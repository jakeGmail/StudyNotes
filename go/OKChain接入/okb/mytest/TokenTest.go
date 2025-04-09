package mytest

import (
	"fmt"
	sdk "github.com/okex/exchain-go-sdk"
	"github.com/okex/exchain-go-sdk/types"
	"github.com/okex/exchain-go-sdk/utils"
	"log"
)

const (
	// TODO: link to mainnet of ExChain later
	rpcURL2 = "https://exchaintesttmrpc.okex.org"
)

var err error
var clientConfig types.ClientConfig
var client sdk.Client

func init() {
	clientConfig, err = sdk.NewClientConfig(rpcURL2, "exchain-65", sdk.BroadcastBlock,
		"0.001okt", 57138, 0, "")
	if err != nil {
		log.Fatal("NewClientConfig failed:", err.Error())
	}

	client = sdk.NewClient(clientConfig)
}

func TestRun2() {
	mnemonic := "charge caution name brain crowd summer angry legal fence champion month yellow"

	privateKey, err := utils.GeneratePrivateKeyFromMnemo(mnemonic)
	if err != nil {
		log.Fatal("generate private key failed:", err.Error())
	} else {
		fmt.Println("privateKey=", privateKey)
	}

	keyInfo, err := utils.CreateAccountWithPrivateKey(privateKey, "jake", "123")
	if err != nil {
		log.Fatal("create account failed:", err.Error())
	} else {
		fmt.Println("keyInfo.PubAddr=", keyInfo.GetPubKey().Address().String())
		fmt.Println("keyInfo.PubAddr=", keyInfo.GetPubKey().Address())
		fmt.Println("keyInfo.Addr=", keyInfo.GetAddress().String())
	}

	account, err := client.Auth().QueryAccount(keyInfo.GetAddress().String())
	if err != nil {
		log.Fatal("query account failed:", err.Error())
	} else {
		fmt.Println("account=", account)
	}
}
