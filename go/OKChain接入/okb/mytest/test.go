package mytest

import (
	"fmt"
	sdk "github.com/okex/exchain-go-sdk"
	"github.com/okex/exchain-go-sdk/utils"
	"log"
)

const (
	// TODO: link to mainnet of ExChain later
	rpcURL = "https://exchaintesttmrpc.okex.org"
	// user's name
	name = "alice"
	// user's mnemonic
	//mnemonic = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
	mnemonic = "expire lesson shoot glory double spirit enhance prison flip agree spawn model"
	// user's password
	passWd = "Jake123"
	// target address
	addr     = "ex1qj5c07sm6jetjz8f509qtrxgh4psxkv3ddyq7u"
	baseCoin = "okt"
)

func TestRun1() {
	config, err := sdk.NewClientConfig(rpcURL, "exchain-65", sdk.BroadcastBlock, "0.001okt", 98138,
		0, "")
	if err != nil {
		log.Fatal(err)
	}
	client := sdk.NewClient(config)

	// 使用给定的名称和密码创建账户信息
	info, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	if err != nil {
		fmt.Printf("create account failed:%s\n", err.Error())
		panic("create account failed!")
	}

	account, err := client.Auth().QueryAccount(info.GetAddress().String())

	transferUit, err := utils.ParseTransfersStr("0xC3C250CD18AC910BE1E7898693968829224AF6B8 0.01okt")
	if err != nil {
		log.Fatal("parse transfer failed:", err.Error())
	}

	tx_res, err := client.Token().MultiSend(info, "Jake123", transferUit, "transfer test", account.GetAccountNumber(), account.GetSequence())
	if err != nil {
		log.Fatal("multi send failed:", err.Error())
	}
	fmt.Println("multiSend res:", tx_res)
}
