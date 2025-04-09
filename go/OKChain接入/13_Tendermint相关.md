[toc]

# 1 查询指定账户的交易信息
```go
QueryTxsByEvents(eventsStr string, page, limit int) (*ctypes.ResultTxSearch, error)
```

**参数**
|参数名称|类型|描述|
|-------|----|----|
|eventsStr|string|查询条件，以 `{eventType}.{eventAttribute}={value}`格式来描述，如果有多个条件，用`&`符号相连，条件之间不留空格。详见**备注**|
|page|int|指定要查询的结果页数|
|limit|int|指定每页显示的结果数量。|

**备注**:
{eventType} 是事件类型，比如 ‘message’, ‘transfer’ 等，
{eventAttribute} 是该事件的属性，比如 ‘action’, ‘sender’, ‘recipient’ 等，
{value} 是指的是否符合属性的值，比如 ‘send’, ‘0xYourPublicKeyHere’ 等。

- `message.action=send&message.sender=ex1rp7fjqtswsfjac4karreqnphxjgk8kckzueq9k`
表示查询ex1rp7fjqtswsfjac4karreqnphxjgk8kckzueq9k账户的发起转账的交易记录

- `message.action=send&transfer.recipient=ex1rp7fjqtswsfjac4karreqnphxjgk8kckzueq9k`表示查询账户`ex1rp7fjqtswsfjac4karreqnphxjgk8kckzueq9k`收到转账的记录

- `message.action=withdraw&unbond.sender=ex1rp7fjqtswsfjac4karreqnphxjgk8kckzueq9k`表示查询对应地址减少质押的记录

- `message.action=deposit&delegate.validator=ex1rp7fjqtswsfjac4karreqnphxjgk8kckzueq9k`表示查询指定账户的增加质押的记录

**返回值**
|返回值|描述|
|------|----|
|[*ctypes.ResultTxSearch](100_各个类型定义.md#11-ctypesresulttxsearch类型)|查询到的交易的记录|

# 2 查询指定高度的区块信息
```go
QueryBlock(height int64) (*types.Block, error)
```

**参数**:
|参数名称|类型|描述|
|-------|----|----|
|height|int64|区块高度|

**返回值**：
|返回值|描述|
|------|----|
|[types.Block](100_各个类型定义.md#14-typesblock类型)|区块信息|

# 3 获取给定区块高度的区块结果
```go
func (tc tendermintClient) QueryBlockResults(height int64) (pBlockResults *types.ResultBlockResults, err error)
```

**参数**:
|参数名称|类型|描述|
|-------|----|----|
|height|int64|区块高度|

**返回值**：
|返回值|描述|
|------|----|
|[types.ResultBlockResults](100_各个类型定义.md#15-typesresultblockresults类型)|区块查询结果信息|


