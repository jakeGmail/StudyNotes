[toc]

# 0 区块浏览器
测试链：https://testnet.bscscan.com/

# 1 Biance接口在线文档
https://www.binance.com/zh-CN/binance-api

# 2 Binance Http请求base
https://api.binance.com
https://api-gcp.binance.com
https://api1.binance.com
https://api2.binance.com
https://api3.binance.com
https://api4.binance.com
后续如果使用http请求 Binance平台数据可以使用这些url base

# 3 接口调用限制

**Http限制**：
- 上述列表的最后4个接口 (api1-api4) 可能会提供更好的性能，但其稳定性略为逊色
- 所有接口的响应都是 JSON 格式
- 响应中如有数组，数组元素以时间升序排列，越早的数据越提前
- 所有时间、时间戳均为UNIX时间，单位为**毫秒**
- 对于仅发送公开市场数据的 API，您可以使用 base URL https://data-api.binance.vision 。
    GET /api/v3/aggTrades
    GET /api/v3/avgPrice
    GET /api/v3/depth
    GET /api/v3/exchangeInfo
    GET /api/v3/klines
    GET /api/v3/ping
    GET /api/v3/ticker
    GET /api/v3/ticker/24hr
    GET /api/v3/ticker/bookTicker
    GET /api/v3/ticker/price
    GET /api/v3/time
    GET /api/v3/trades
    GET /api/v3/uiKlines

- 收到429时，您有责任停止发送请求，不得滥用API。收到429后仍然继续违反访问限制，会被封禁IP，并收到418错误码
频繁违反限制，封禁时间会逐渐延长，从最短2分钟到最长3天。

**web socket限制**： 
- 单个连接最多可以订阅 1024 个Streams。
- 每IP地址、每5分钟最多可以发送300次连接请求
- Websocket服务器每秒最多接受5个消息。消息包括:
    PING帧
    PONG帧
    JSON格式的消息, 比如订阅, 断开订阅.
- 如果用户发送的消息超过限制，连接会被断开连接。反复被断开连接的IP有可能被服务器屏蔽。
- 每IP地址、每5分钟最多可以发送300次连接请求。

# k 线数据获取
K线间隔:

s -> 秒; m -> 分钟; h -> 小时; d -> 天; w -> 周; M -> 月
1s
1m
3m 
5m  
15m  
30m 
1h 
2h
4h
6h
8h
12h
1d
3d
1w
1M