[toc]

# 1 查询产品的蜡烛数据
```go
func (ac ammswapClient) QuerySwapTokenPairs() (exchanges []types.SwapTokenPair, err error)
```

**返回值**：
|返回值|描述|
|--------|----|
|types.SwapTokenPair|交易对信息|


**使用示例**:

```go
package okb

import (
	"fmt"
	"log"
)

func AmmSwapTest() {
    tokenpairs, err := OkClient.AmmSwap().QuerySwapTokenPairs()
    if err != nil {
        log.Fatal("QuerySwapTokenPairs failed")
    }
    for _, tokenPair := range tokenpairs {
        fmt.Println("tokenPair Name=", tokenPair.TokenPairName())
        fmt.Println("content", tokenPair.String())
        fmt.Println("============================")
    }
}
```

