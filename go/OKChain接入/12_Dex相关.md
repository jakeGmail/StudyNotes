[toc]

</br>
API实现代码位于`exchain-go-sdk/module/dex`

DEX是去中心化交易所

# 1 查询某个地址（通常是交易对的创建者）的产品列表。

```go
QueryProducts(ownerAddr string, page, perPage int) ([]types.TokenPair, error)
```
**作用**：
查询某个地址（通常是交易对的创建者）的产品列表

**参数**
|参数名称|参数类型|描述|
|-------|-------|----|
|ownerAddr|string|交易对所有者的地址。|
|page|int|查询的页面数，通常用于分页处理.|
|perPage|int|每页显示的项数，通常与页面参数一起使用。|

**返回值**:
|返回值|描述|
|------|----|
|[[]types.TokenPair](100_各个类型定义.md#10-typestokenpair类型)|TokenPair 对象的切片。每个 TokenPair 表示一个交易对，包含了相关的信息如交易对的名称、产品名称、里面含有的代币等。|

**示例代码**：
```go
package okb

import (
	gosdk "github.com/okex/exchain-go-sdk"
	"log"
)

const (
	url  = "https://exchaintesttmrpc.okex.org"
	addr = "0x187c99017074132ee2b6e8c7904c37349163db16"
)

var okClient gosdk.Client

func init() {
	clientConfig, err := gosdk.NewClientConfig(url, "exchain-65", gosdk.BroadcastBlock,
		"0.000001okt", 5000,
		1.1, "0.0000000001okt")
	if err != nil {
		log.Fatal("chain config create failed:", err.Error())
	}
	okClient = gosdk.NewClient(clientConfig)
}

func OKExChainTest() {
	tokenPair, err := okClient.Dex().QueryProducts(addr, 10, 10)
	if err != nil {
		log.Fatal("query products failed:", err.Error())
	}
	log.Println("tokenPair=", tokenPair)
}

```

# 2 在OKExChain上注册DEX操作员

用于在OKExChain上注册DEX操作员。
```
RegisterDexOperator(fromInfo keys.Info, passWd, handleFeeAddrStr, website, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
```

**参数**:
|参数名称|类型|描述|
|-------|----|----|
|fromInfo|[keys.Info](100_各个类型定义.md#1-keysinfo类型)|创建者的钱包信息，包括私钥等信息，该信息用来签名交易。| 
|passWd|string|创建者钱包的密码，用于解锁账户。|
|handleFeeAddrStr|string|处理交易费用的地址。区块链交易会有一定的手续费，这个地址就是接收这些手续费的地方。|
|website|string| 操作员的官方网站地址，供其他用户查阅信息。|
|memo|string|可选参数，允许附加一些描述性信息或备注。|
|accNum|uint64|账户号|
|seqNum|uint64|交易序列号|

**返回值**：
|返回值|描述|
|------|----|
|[sdk.TxResponse](100_各个类型定义.md#7-sdktxresponse类型)|注册DEX操作员事务相响应息|

**示例代码**：
```go
package okb

import (
	"github.com/okex/exchain-go-sdk/utils"
	"log"
)

func DexTest() {
	keyInfo, _, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
	if err != nil {
		log.Fatalln("create account failed:", err.Error())
	}
	account, err := okClient.Auth().QueryAccount(keyInfo.GetAddress().String())
	if err != nil {
		log.Fatalln("query account failed:", err.Error())
	}
	accountNumber := account.GetAccountNumber()
	seqNumber := account.GetSequence()

	token, err := okClient.Dex().RegisterDexOperator(keyInfo, passWd, addr187, "https://www.test.com",
		"test register", accountNumber, seqNumber)
	seqNumber++
	if err != nil {
		log.Println("register dex operator failed:", err.Error())
	}
	log.Println("token=", token)
}
```

# 3 编辑DEX操作员的信息
```go
EditDexOperator(fromInfo keys.Info, passWd, handleFeeAddrStr, website, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
```

**参数**:
|参数名称|参数类型|描述|
|-------|--------|----|
|fromInfo|[keys.Info](100_各个类型定义.md#1-keysinfo类型)|dex操作员账户密钥信息|
|passWd|string|钱包密码|
|handleFeeAddrStr|string|新的处理交易费用的地址。|
|website|string|新的网站地址|
|memo|string|本次编辑的备注信息|
|accNum|uint64|账户号|
|seqNum|uint64|交易序列号|

**返回值**：
|返回值|描述|
|------|----|
|[sdk.TxResponse](100_各个类型定义.md#7-sdktxresponse类型)|编辑DEX操作员信息的响应信息|

**示例的代码**：
```go
// 查询账户信息
account, err := okClient.Auth().QueryAccount(keyInfo.GetAddress().String())
if err != nil {
	log.Fatalln("query account failed:", err.Error())
}
accountNumber := account.GetAccountNumber()
seqNumber := account.GetSequence()

// 编辑DEX信息
txRes, err := okClient.Dex().EditDexOperator(keyInfo, passWd, addr187,
	"https://changedWed.com", "edit", accountNumber, seqNumber)
if err != nil {
	log.Fatalln("edit Dex operator failed:", err.Error())
}
fmt.Println(txRes)
```

# 4 在去中心化交易所（DEX）里列出（List）一个新的交易对
```
func (dc dexClient) List(fromInfo keys.Info, passWd, baseAsset, quoteAsset, initPriceStr, memo string, accNum, seqNum uint64) (resp sdk.TxResponse, err error)
```

**参数**：
|参数名称|类型|描述|
|-------|----|----|
|fromInfo|[keys.Info](100_各个类型定义.md#1-keysinfo类型)|DEX操作员的密钥信息|
|passWd|string|钱包密码|
|baseAsset|string|基础代币符号|
|quoteAsset|string|报价资产的标识符，交易对中用于定价的货币。|
|initPriceStr|string|初始价格，表示一个baseAsset期望价值initPriceStr个quoteAsset。例如"100.0"|
|memo|string|创建交易对时的备注信息|
|accNum|uint64|账户号|
|seqNum|uint64|交易序列号|

**示例代码**：
```go
accountNumber := account.GetAccountNumber()
seqNumber := account.GetSequence()

// 列出一个交易对，okt/bbq, 表示1个okt的价格是100个bbq
txRes, err := okClient.Dex().List(keyInfo, passWd, "okt",
	"bbq", "100", "list",accountNumber, seqNumber)
if err != nil {
	log.Fatalln("List Dex operator failed:", err.Error())
}
fmt.Println(txRes)
```

# 5 向交易对中存入指定数量的okt
Deposit okt to a specific product
```go
Deposit(fromInfo keys.Info, passWd, product, amountStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
```

**参数**:
|参数名称|类型|描述|
|-------|----|----|
|fromInfo|[keys.Info](100_各个类型定义.md#1-keysinfo类型)|存入代币的账户密钥信息，用于签名|
|passWd|string|钱包密码|
|product|string|交易对名称|
|amountStr|string|存入代币的数量，例如"1.2okt"|
|memo|string|操作的备注信息|
|accNum|uint64|账户号|
|seqNum|uint64|交易序列号|

**示例代码**：
```go
accountNumber := account.GetAccountNumber()
seqNumber := account.GetSequence()

txRes, err := okClient.Dex().Deposit(keyInfo, passWd, "okt_btc",
	"1.20okt", "okt_btc deposit", accountNumber, seqNumber)
if err != nil {
	log.Fatalln("Deposit Dex operator failed:", err.Error())
}
fmt.Println(txRes)
```

# 6 从交易对中取出指定数量的okt
```go
func (dc dexClient) Withdraw(fromInfo keys.Info, passWd, product, amountStr, memo string, accNum, seqNum uint64) (resp sdk.TxResponse, err error)
```

**参数**:
|参数名称|类型|描述|
|-------|----|----|
|fromInfo|[keys.Info](100_各个类型定义.md#1-keysinfo类型)|取出代币的账户密钥信息，用于签名|
|passWd|string|钱包密码|
|product|string|交易对名称|
|amountStr|string|存入代币的数量，例如"1.2okt"|
|memo|string|操作的备注信息|
|accNum|uint64|账户号|
|seqNum|uint64|交易序列号|

**示例代码**:
```go
accountNumber := account.GetAccountNumber()
seqNumber := account.GetSequence()

txRes, err := okClient.Dex().Withdraw(keyInfo, passWd, "okt_btc",
	"1.20btc", "okt_btc deposit", accountNumber, seqNumber)
if err != nil {
	log.Fatalln("Withdraw Dex operator failed:", err.Error())
}
fmt.Println(txRes)
```

# 7 改变产品的所有者
```go
func (dc dexClient) TransferOwnership(fromInfo keys.Info, passWd, product, toAddrStr, memo string, accNum, seqNum uint64) (resp sdk.TxResponse, err error) 
```

**参数**:
|参数名称|类型|描述|
|-------|----|----|
|fromInfo|[keys.Info](100_各个类型定义.md#1-keysinfo类型)|原产品所有者的账户密钥信息，用于签名|
|passWd|string|钱包密码|
|product|string|产品名称|
|toAddrStr|string|新的产品所有者的地址|
|memo|string|操作的备注信息|
|accNum|uint64|账户号|
|seqNum|uint64|交易序列号|


**示例代码**：
```go
accountNumber := account.GetAccountNumber()
seqNumber := account.GetSequence()

// 将产品的拥有者改为addr187
txRes, err := okClient.Dex().TransferOwnership(keyInfo,passWd,"okt_btc",addr187,
	"transfer onwer", accountNumber, seqNumber)
if err != nil {
	log.Fatalln("TransferOwnership Dex product failed:", err.Error())
}
fmt.Println(txRes)
```

# 8 所有权确认
区块链场景中，所有权确认通常是指确认特定账户拥有特定资产或令牌的所有权。这在多个场景，包括交易前的准备阶段或资产转移中，是一个重要步骤。在使用分布式交易所（DEX）等去中心化金融（DeFi）服务时，所有权确认很可能是一个关键步骤，以确保交易双方都是他们所声称的资产的合法所有者。
```go
func (dc dexClient) ConfirmOwnership(fromInfo keys.Info, passWd, product, memo string, accNum, seqNum uint64) (resp sdk.TxResponse, err error)
```

**参数**:
|参数名称|类型|描述|
|--------|---|----|
|fromInfo|[keys.Info](100_各个类型定义.md#1-keysinfo类型)|确认账户的私钥信息|
|passWd|string|钱包密码|
|product|string|确认所有权的产品名称|
|memo|string|确认所有权的备注信息|
|accNum|uint64|账户号|
|seqNum|uint64|交易序列号|

**示例代码**：
```go
func DexTest() {
	keyInfo, _, err := utils.CreateAccountWithMnemo(mnemonic187, name, passWd)
	if err != nil {
		log.Fatalln("create account failed:", err.Error())
	}
	account, err := okClient.Auth().QueryAccount(keyInfo.GetAddress().String())
	if err != nil {
		log.Fatalln("query account failed:", err.Error())
	}
	accountNumber := account.GetAccountNumber()
	seqNumber := account.GetSequence()

	txRes, err := okClient.Dex().ConfirmOwnership(keyInfo, passWd, "okt_btc",
		"transfer onwer", accountNumber, seqNumber)
	if err != nil {
		log.Fatalln("ConfirmOwnership Dex product failed:", err.Error())
	}
	fmt.Println(txRes)
}
```

