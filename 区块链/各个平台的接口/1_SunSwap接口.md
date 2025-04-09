[toc]

# 0 SunSwap网址
https://sunswap.com/?lang=zh-CN#/home

# 1 获取交易对的价格信息

**请求**：
```
GET: https://openapi.sun.io/v2/allpairs?ver=3&page_size=100&page_num=1&orderBy=price
```

**参数**:
  - <font color=red>page_size</font>: 数字，分页查询的页码
  - <font color=red>page_num</font>: 数字，每一页的记录条数
  - <font color=red>orderBy</font>: 字符串，以指定的字段排序（默认降序）
  - <font color=red>desc</font>: bool, true:以降序排列， false:以升序排列

**响应示例**：
```json
{
    "data": {
        "TN3W4H6rK2ce4vX9YnFQHwKENnHjoxb3m9_TNUC9Qb1rRpS5CbWLmNMxXBjyFoydXjWFR": {
            "base_decimal": "6",
            "quote_name": "Bitcoin",
            "base_name": "Wrapped TRX",
            "quote_symbol": "BTC",
            "base_id": "TNUC9Qb1rRpS5CbWLmNMxXBjyFoydXjWFR",
            "price": "557658.638820747800222502",
            "quote_volume": "81981",
            "base_volume": "463319974",
            "quote_id": "TN3W4H6rK2ce4vX9YnFQHwKENnHjoxb3m9",
            "base_symbol": "WTRX",
            "quote_decimal": "8"
        }
    }
}
```

# 2 获取市场的交易对成交信息

**请求**：

```
https://abc.endjgfsv.link/swap/v2/exchanges/scan?pageNo=1&pageSize=10
```

**参数**:
 - <font color=red>pageNo</font>: 数字，分页查询的页码
 - <font color=red>pageSize</font>: 数字，每一页的记录条数

**响应示例**：

```json
{
    "code": 0,
    "data": {
        "totalPages": 2029,
        "totalCount": "20290",
        "list": [
            {
                "ver": 1,
                "address": "TYukBQZ2XXCcRCReAUguyXncCWNY9CEiDQ",
                "volume14d": "30145284.095357593",
                "tokenSymbol": "JST",
                "isValid": 1,
                "tokenName": "JUST GOV",
                "volume24hrs": "370591.1261914438",
                "txId": "72990c1ab7700c6c04d3c0f9a8387b1d7d79d4ae51d76db23f84d33bd75b6071",
                "liquidity": "35579143.286247",
                "volume7d": "22100409.10353019",
                "type": "1",
                "tokenAddress": "TCFLL5dx5ZJdKnWuesXxi1VPwjLVmWZZy9",
                "tokenLogoUrl": "https://static.tronscan.org/production/logo/just_icon.png",
                "tokenDecimal": 18,
                "id": 6,
                "fees24hrs": "1111.7733785743314"
            }
        ]
    }
}
```

