[toc]

# 1 Token相关类型概览
```go
// 令牌类型接口
type Token interface {
    gosdktypes.Module
    TokenTx
    TokenQuery
}

type TokenTx interface {
    Send(fromInfo keys.Info, passWd, toAddrStr, coinsStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
    MultiSend(fromInfo keys.Info, passWd string, transfers []types.TransferUnit, memo string, accNum, seqNum uint64) (
        sdk.TxResponse, error)
    Issue(fromInfo keys.Info, passWd, orgSymbol, wholeName, totalSupply, tokenDesc, memo string, mintable bool, accNum,
        seqNum uint64) (sdk.TxResponse, error)
    Mint(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
    Burn(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
    Edit(fromInfo keys.Info, passWd, symbol, description, wholeName, memo string, isDescEdit, isWholeNameEdit bool, accNum,
        seqNum uint64) (sdk.TxResponse, error)
}

type TokenQuery interface {
    QueryTokenInfo(ownerAddr, symbol string) ([]types.TokenResp, error)
}
```

# 2 获取令牌信息
```go
type TokenQuery interface {
	QueryTokenInfo(ownerAddr, symbol string) ([]types.TokenResp, error)
}
```
**作用**：
来从 OKExChain 查询指定地址拥有者的特定代币信息的。

**参数**:
- <font color=red>ownerAddr</font>: 账户地址信息，可以通过types.Account.GetAddress()方法获取
- <font color=red>symbol</font>:想查询的代币的符号的字符串。OKB是： "okt"

**返回值**:
<font color=red>[]types.TokenResp</font>: 包含了 types.TokenResp 对象的切片，每个对象都代表了一个代币的详细信息。
types.TokenResp类型详见[types.TokenResp类型](100_各个类型定义.md#6-typestokenresp类型)

**示例代码**:
```go
// 生成客户端
config, err := sdk.NewClientConfig(rpcURL, "exchain-65", sdk.BroadcastBlock, "0.001okt", 57138,
		0, "")
if err != nil {
	log.Fatal(err)
}
client := sdk.NewClient(config)

// 查询令牌信息
tokenRes, err := client.Token().QueryTokenInfo("0x187c99017074132ee2b6e8c7904c37349163db16", "okt")
if err != nil {
	fmt.Println("getTokenInfo faled:", err.Error())
} else {
	fmt.Println("token info:", tokenRes)
}

/*运行结果：
token info: [{"description":"OKExChain Native Token",
"symbol":"okt",
"original_symbol":"okt","whole_name":"OKT",
"original_total_supply":"1000000000.000000000000000000",
"type":0,
"owner":"ex1v8segh8mlw297s2ksy6pp6nwxn3el0wmkuqsx2","mintable":true,
"total_supply":"11015029183.875000000000000000"}]
*/
```


# 3 发行令牌
```go
type TokenTx interface {
    Issue(fromInfo keys.Info, passWd, orgSymbol, wholeName, totalSupply, tokenDesc, memo string,
	mintable bool, accNum,
	seqNum uint64) (sdk.TxResponse, error)
    ...
    ...
}
```
**作用**:
在OKChain区块链上发行令牌.
Issue 的主要作用是创建一个新的代币（Token）并为其提供初始供应。这在主权发行代币，如初始币（ICO）等场景中有着重要的作用。你可以通过该功能发行你自己的代币，然后根据你的业务逻辑进行配置和操作。所发布的代币在区块链中是一种完全去中心化的资产，通过这种方式，开发者和企业有能力创建和管理自己的数字资产。
**参数**：
- <font color=red>fromInfo</font>:提供基本的账户信息包括密钥信息。
- <font color=red>passWd</font>:账户的密码凭据，用于正确签署交易。
- <font color=red>orgSymbol</font>:要发行的新令牌的简称或符号，通常是以大写字母表示，例如“BTC”，“ETH”。
- <font color=red>wholeName</font>:代币的全称。
- <font color=red>totalSupply</font>:该代币的总供应量，表示发行的令牌数量。
- <font color=red>tokenDesc</font>:代币的描述，介绍关于这个令牌的基本信息。
- <font color=red>memo</font>:备注信息，可以用于记录一些额外信息。
- <font color=red>mintable</font>:一个布尔值，表明是否可以增发这种令牌。(可增发令牌允许你在初始发行后随时创建更多的令牌）
- <font color=red>accNum</font>: 提供账户的地址。
- <font color=red>seqNum</font>: 提供账户的顺序号。

**返回值**:
- <font color=red>sdk.TxResponse</font>: 如果调用成功，则返回交易的详细信息。

**代码示例**:
```go
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
	mnemonic = "giggle sibling fun arrow elevator spoon blood grocery laugh tortoise culture tool"
	// user's password
	passWd = "12345678"
	// target address
	addr     = "ex1qj5c07sm6jetjz8f509qtrxgh4psxkv3ddyq7u"
	baseCoin = "okt"
)

func TestRun1() {
	config, err := sdk.NewClientConfig(rpcURL, "exchain-65", sdk.BroadcastBlock, "0.01okt", 2000,
		0, "")
	if err != nil {
		log.Fatal(err)
	}
	client := sdk.NewClient(config)

	// 使用给定的助记符和密码生成密钥信息
	info, mnemo, err := utils.CreateAccountWithMnemo(mnemonic, "jake", "123")
	if err != nil {
		fmt.Printf("create account failed:%s\n", err.Error())
		panic("create account failed!")
	}
	fmt.Printf("mnemo=|%s|\n", mnemo)

    // 获取账户信息
	accInfo, err := client.Auth().QueryAccount(info.GetAddress().String())
	if err != nil {
		log.Fatal(err)
	}
    

    // 发行代币
	txResponse, err := client.Token().Issue(info, "123", "usdk", "sudk",
		"100", "statable coins", "nn", true,
		accInfo.GetSequence(), accInfo.GetAccountNumber())
	if err == nil {
		fmt.Println(txResponse)
	} else {
		fmt.Println("Issue error:", err.Error())
	}
}
```

**报错信息解析**：
- <font color=red>failed. build stdTx error: ciphertext decryption failed</font>:表示在解密过程中出现问题，这通常和用错了密钥或者密码有关。
在 OKEx 的区块链操作中，你需要用到你的私钥或密码。这些都是敏感的信息，用于确认你的身份和签署你的交易。如果用错了，或者这些信息被篡改，你的请求就会失败。
所有的密码或密钥在加密和解密的过程中都必须是正确的，否则会导致解密失败。
下面是一些可能的问题和解决方案：
1. fromInfo 或者 passWd 输入错误：请确认你提供的密钥信息和密码都是正确的。
2. 密码或密钥的管理问题：在实际的开发和部署过程中，一定要确保密码或密钥的安全，防止他人恶意使用。
3. 如果你复制和粘贴的密钥，确认没有意外的空格或者其他字符。
4. 如果这个私钥是从其他地方导入的，确认这个私钥是正确的，没有被加密或修改。
5. 最后，你可以切换到一个新的、安全的环境，重新生成一个新的私钥和密码，然后再试一次。
6. 如果问题依旧存在，你可能需要向 OKEx 的开发者社区或者他们的技术支持寻求更深入的帮助。

- <font color=red>out of gas: out of gas in location: ReadFlat; gasWanted: 2000, gasUsed: 2015</font>:这个错误 "out of gas: out of gas in location: ReadFlat; gasWanted: 2000, gasUsed: 2015" 表示在执行该交易时，提供的 Gas 数量不足以完成操作。在这个错误中，你为这个交易指定了 2000 Gas，但实际操作需要 2015 Gas。

在区块链中，每次操作都需要消耗 Gas，Gas的基本单位是`Gwei`。Gas 是一种支付区块链网络计算资源使用费用的方式。Gas 价格会随着网络的拥堵程度、供需关系等因素而波动。

以下是一些可能的原因和解决方案：

1. <font color=blue>提供更高的 Gas 数量</font>：你可以在发起交易时提供更多的Gas。但是，这需要一些试错，因为过多的Gas会被浪费，过少的Gas则会导致交易失败。你可以将 Gas 数量设置为比实际需求稍高的数量。

2. <font color=blue>查询链上 Gas 价格</font>：你可以先查询链上当前的 Gas 价格，然后根据实际情况调整你的 Gas 价格和数量。这样你可以避免因为价格波动导致的交易失败。

3. <font color=blue>检查智能合约代码</font>：如果Gas消耗超出预期，可能存在智能合约代码的问题。检查代码以避免不必要的操作或计算。优化代码可以降低 Gas 消耗，从而减少交易费用。

4. <font color=blue>在网络拥堵时段之外进行操作</font>：在区块链网络拥堵时段，Gas 价格往往会上涨。在非拥堵时段进行交易可能有更低的 Gas 价格。

请尝试上述方法并根据实际情况调整你的操作。希望你的交易在提高 Gas 之后能够成功执行。

- <font color=red>unauthorized: signature verification failed; verify correct account sequence and chain-id, sign msg:... ...</font>:这个错误提示的是在执行方法或交易时发生签名验证失败，导致交易未能成功执行。出现此问题的原因可能有以下几个方面：

1. **账号序列不正确**：区块链交易通常需要一个递增的账号序列（account sequence）来确保交易安全和顺序性。这个错误可能是因为你提供了一个错误的账号序列。通常情况下，区块链会自动为你填写正确的账号序列，但是在某些情况下可能需要手动指定。你需要确保提供正确的账号序列。

2. **chain-id错误**：所有区块链交易都需要指定一个chain-id来确定哪个区块链网络将执行此交易。chain-id错误可能是因为你指定了一个无法匹配正确网络的chain-id值。你需要检查交易中的chain-id是否正确，确保它与目标区块链网络匹配。

3. **私钥错误或丢失**：当发送交易时，你的私钥用于签名以证明你是合法的交易发起者。错误可能是因为你提供的私钥丢失或不正确，导致签名无法验证。请检查您的私钥，并确保正确地执行了签名过程。

要解决这个问题，你需要根据以上可能原因进行排查。请确认你所使用的账号序列、chain-id以及私钥是否正确，以确保成功执行交易。

- <font color=red>not allowed original symbol: not allowed original symbol: xxx</font>: 发行代币的时候，名称命名不符合规范。

# 4 给单个账户转账
```go
type TokenTx interface {
    Send(fromInfo keys.Info, passWd, toAddrStr, coinsStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
    ...
    ...
}
```
**作用**:
发送一笔代币转账交易。
当执行这个 Send() 函数时，它会用指定的发送方账户信息、密码、接收方信息、转账金额等构建一笔转账交易，并对该交易进行本地签名。然后将签名后的交易广播到 OKExChain 网络。函数执行结束后，如果成功，将返回一个包含交易响应信息的sdk.TxResponse 对象；否则返回一个错误信息error。

请注意，在广播交易之前，请确保发送方账户拥有足够的代币和Gas费用，以确保交易成功执行。
**参数**：
- **fromInfo**：发送方的账户信息keys.Info，包括了发送者的地址、公钥和私钥等相关信息。
- **passWd**：发送方的账户密码，该密码用于程序对交易进行本地签名。
- **toAddrStr**：接收方的地址，以字符串形式表示。
- **coinsStr**：转账金额和代币信息，使用字符串描述。例如："10okt"表示发送10个OKT代币。
- **memo**：（可选）交易备注。可供用户在交易中附带额外的信息。
- **accNum**：发送方的账户序号，用于正确识别发送方的信息以保证交易安全。
- **seqNum**：发送方的交易序号，用于防止重复交易和提高交易的安全性。

**返回值**：
- **sdk.TxResponse**：转账反馈。

**示例代码**：
```go
// 创建客户端
config, err := sdk.NewClientConfig(rpcURL, "exchain-65", sdk.BroadcastBlock, "0.001okt", 98138,
		0, "")
if err != nil {
	log.Fatal(err)
}
client := sdk.NewClient(config)

// 使用给定的名称和密码创建账户信息
info, mnemo, err := utils.CreateAccountWithMnemo(mnemonic, name, passWd)
if err != nil {
	fmt.Printf("create account failed:%s\n", err.Error())
	panic("create account failed!")
}

// 进行转账
res, err := client.Token().Send(info, "Jake123", "0xC3C250CD18AC910BE1E7898693968829224AF6B8", "0.02okt", "", account.GetAccountNumber(), account.GetSequence())
if err != nil {
	fmt.Println("send failed:", err.Error())
} else {
	fmt.Println("send response:", res)
}

/* 运行结果
send response: Response:
  Height: 25058079
  TxHash: 7136FC38A50D06FFF396E6B7D284246E390089721D8DCABFAB40C1AEC08CB38B
  Raw Log: [{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"send"},{"key":"sender","value":"ex1rp7fjqtswsfjac4karreqnphxjgk8kckzueq9k"},{"key":"module","value":"token"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"ex1c0p9pngc4jgshc083xrf895g9y3y4a4c5je8dw"},{"key":"sender","value":"ex1rp7fjqtswsfjac4karreqnphxjgk8kckzueq9k"},{"key":"amount","value":"0.020000000000000000okt"}]}]}]
  Logs: [{"msg_index":0,"log":"","events":[{"type":"message","attributes":[{"key":"action","value":"send"},{"key":"sender","value":"ex1rp7fjqtswsfjac4karreqnphxjgk8kckzueq9k"},{"key":"module","value":"token"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"ex1c0p9pngc4jgshc083xrf895g9y3y4a4c5je8dw"},{"key":"sender","value":"ex1rp7fjqtswsfjac4karreqnphxjgk8kckzueq9k"},{"key":"amount","value":"0.020000000000000000okt"}]}]}]
  GasWanted: 98138
  GasUsed: 81231
*/
```

# 5 同时给多个账户转账
```go
type TokenTx interface {
    // 给一个指定账户转账
	Send(fromInfo keys.Info, passWd, toAddrStr, coinsStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)

    // 同时给多个账户转账
	MultiSend(fromInfo keys.Info, passWd string, transfers []types.TransferUnit, memo string, accNum, seqNum uint64) (
		sdk.TxResponse, error)
}
```
**参数**：
|参数名称|参数类型|描述|
|-------|--------|----|
|fromInfo|keys.Info|公密钥相关信息|
|passWd|string|创建账户时设置的密码|
|transfers|[]types.TransferUnit|转账相关信息|
|meno|string|转账的备注|
|accNum|uint64|发送方的账户号码|
|seqNum|uint64|发送方的账户序列号|

**代码示例**：
```go
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
```

# 6 增加发行令牌的总量
```go
type TokenTx interface{
    ...
    Mint(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
    ...
```
**作用**:
增发一定数量的令牌
**参数**：
- <font color=red>fromInfo</font>:增大令牌的账户信息
- <font color=red>passWd</font>: 账户密码（创建账户时自定义）
- <font color=red>coinsStr</font>:增发的令牌信息，格式为"<增发数量> 令牌符号"，例如"1000 okt"表示增发1000枚okt
- <font color=red>memo</font>：增发令牌的备注信息
- <font color=red>accNum</font>：账户号
- <font color=red>seqNum</font>：交易序列号

**示例**：
```go
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
		"0.0001okt", 177138, 0.1, "")
	if err != nil {
		log.Fatal("NewClientConfig failed:", err.Error())
	}

	client = sdk.NewClient(clientConfig)
}

func TestRun2() {
	mnemonic := "charge caution name brain crowd summer angry legal fence champion month yellow"

	// 从助记符生成账户
	keyInfo, mnemo, err := utils.CreateAccountWithMnemo(mnemonic, "jake", "Jake123")
	if err != nil {
		log.Fatal("create account from mnemo failed", err.Error())
	}
	fmt.Println("mno=", mnemo)

    // 获取账户信息
	account, err := client.Auth().QueryAccount(keyInfo.GetAddress().String())
	if err != nil {
		log.Fatal("query account failed:", err.Error())
	}

	// 增发令牌
	res, err := client.Token().Mint(keyInfo, "Jake123", "1 jjk",
		"test issue", account.GetAccountNumber(), account.GetSequence())
	if err != nil {
		log.Fatal("incr token failed:", err.Error())
	}
	fmt.Println(res)
}
```

# 7 减少发行的令牌总量
```go
type TokenTx interface {
    ...
    Burn(fromInfo keys.Info, passWd, coinsStr, memo string, accNum, seqNum uint64) (sdk.TxResponse, error)
    ...
}
```
**参数**:
- <font color=red>fromInfo</font>: 减少发行令牌的账户信息，只有令牌的所有者账户才有权利进行减少发行量
- <font color=red>passWd</font>: 账户密码
- <font color=red>coinsStr</font>: 减少的令牌数量，格式为"<减少数量> 令牌符号"，例如"1000 okt"表示减少1000枚okt
- <font color=red>memo</font>：增发令牌的备注信息
- <font color=red>accNum</font>：账户号
- <font color=red>seqNum</font>：交易序列号

**示例**:
```go
txResponse, err := client.Token().Burn(keyInfo, "Jake123", "100 jjk",
		"dec", account.GetAccountNumber(), account.GetSequence())
if err != nil {
	fmt.Println("Burn token failed:", err.Error())
}
```

# 8 修改令牌信息
```go
type TokenTx interface {
    Edit(fromInfo keys.Info, passWd, symbol, description, wholeName, memo string, isDescEdit, isWholeNameEdit bool, accNum,
         seqNum uint64) (sdk.TxResponse, error)
}
```
**作用**：
它用来编辑已经在OKExChain上发行的令牌的信息。可以用于修改令牌的符号、全称、描述信息
**参数**：
- <font color=red>fromInfo<font>: 令牌所有者的信息
- <font color=red>passWd<font>: 账户密码
- <font color=red>symbol<font>: 需要修改的令牌的符号
- <font color=red>description<font>: 修改后的新的描述信息。
- <font color=red>wholeName<font>: 令牌修改新的全称。
- <font color=red>memo<font>: 备注信息
- <font color=red>isDescEdit<font>:是否更改描述信息
- <font color=red>isWholeNameEdit<font>: 是否更改全称
- <font color=red>accNum<font>: 操作者账号的账户编号。
- <font color=red>seqNum<font>:账号的序列号。

**示例代码**:
```go
txResponse, err := client.Token().Edit(keyInfo, "Jake123", "jjk",
		"jjk coins", "jakejake",
		"change in 2023-11-04", true, true,
		account.GetAccountNumber(), account.GetSequence())
if err != nil {
	log.Fatal("edit token failed:", err.Error())
} else {
	fmt.Println("txResponse=", txResponse)
}
```
