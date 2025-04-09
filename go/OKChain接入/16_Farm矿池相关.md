[toc]

# 1 查询全部farm矿池信息
```go
func (fc farmClient) QueryPools() (farmPools []types.FarmPool, err error) 
```

**返回值**:
|返回值|描述|
|------|----|
|[][types.FarmPool](100_各个类型定义.md#17-typesfarmpool类型)|矿池信息|

# 2 获取指定矿池的信息
```go
func (fc farmClient) QueryPool(poolName string) (farmPool types.FarmPool, err error) 
```

**参数**:
|参数名称|参数类型|描述|
|-------|--------|----|
|poolName|string|Farm矿池的名称|

**返回值**:
|返回值|描述|
|------|----|
|farmPool|[types.FarmPool](100_各个类型定义.md#17-typesfarmpool类型)|矿池信息|

# 3 查询指定账户加入的Farm矿池
```go
func (fc farmClient) QueryAccount(accAddrStr string) (poolNames []string, err error) 
```

**参数**:
|参数名称|参数类型|描述|
|-------|--------|----|
|accAddrStr|string|地址|

**返回值**：
|返回值|类型|描述|
|poolNames|[]string|加入的矿池的名称列表|

# 4 查询指定Farm矿池锁定的全部地址信息
```go
func (fc farmClient) QueryAccountsLockedTo(poolName string) (accAddrs []sdk.AccAddress, err error) 
```

# 5 查询指定账户在farm矿池中被锁定的信息
```go
func (fc farmClient) QueryLockInfo(poolName, accAddrStr string) (lockInfo types.LockInfo, err error)
```

# 6 查询指定账户在Farm矿池中的收益信息
```go
func (fc farmClient) QueryEarnings(poolName, accAddrStr string) (types.Earnings, error)
```

**参数**:
|参数名称|参数类型|描述|
|-------|--------|----|
|poolName|string|矿池名称|
|accAddrStr|string|账户地址|

**返回值**:
- **types.Earnings**:
```go
type Earnings struct {
    // 查询收益情况的区块高度，即AmountYielded收益信息是在这个区块高度下的收益
    TargetBlockHeight int64        `json:"target_block_height"`

    // 被锁定的代币数量, eg: 9.611947156751230204ammswap_okt_usdt-25a
    AmountLocked      sdk.SysCoin  `json:"amount_locked"`

    // 获取的收益，eg: 6414.866021891899898769okt
    AmountYielded     sdk.SysCoins `json:"amount_yielded"`
}
```