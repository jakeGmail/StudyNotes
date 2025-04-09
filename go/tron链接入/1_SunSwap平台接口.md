# 1 波场见证节点信息获取

**请求URL**:
```url
GET  https://apilist.tronscanapi.com/api/vote/witness
```

**响应结果示例**

```json
{
    "total": 411,
    "totalVotes": 44297151283,
    "fastestRise": {
        "lastRanking": 16,
        "realTimeRanking": 10,
        "address": "TDpt9adA6QidL1B1sy3D8NC717C6L5JxFo", 
        "name": "Chain Cloud",
        "url": "chaincloud",
        "hasPage": false,
        "lastCycleVotes": 1623256721,
        "realTimeVotes": 1624830346,
        "changeVotes": 1573625,
        "brokerage": 0,
        "votesPercentage": 3.6644720348483455,
        "lastCycleVotesPercentage": 3.664541780526695,
        "change_cycle": 6,
        "witnessType": 1,
        "annualizedRate": "4.180365",
        "producedTotal": 695698,
        "producedEfficiency": 99.99353209446592,
        "blockReward": 11131168,
        "version": 29,
        "totalOutOfTimeTrans": 2441,
        "lastWeekOutOfTimeTrans": 0,
        "changedBrokerage": false,
        "lowestBrokerage": 100
    },
    "maxVotesRise": {
        "lastRanking": 18,
        "realTimeRanking": 18,
        "address": "TJBtdYunmQkeK5KninwgcjuK1RPDhyUWBZ",
        "name": "JD Investment",
        "url": "JDinvestment",
        "hasPage": false,
        "lastCycleVotes": 1607064025,
        "realTimeVotes": 1614186687,
        "changeVotes": 7122662,
        "brokerage": 0,
        "votesPercentage": 3.6279173230192483,
        "lastCycleVotesPercentage": 3.627986372953941,
        "change_cycle": 0,
        "witnessType": 1,
        "annualizedRate": "4.182818",
        "producedTotal": 1214858,
        "producedEfficiency": 99.94759336731126,
        "blockReward": 19437728,
        "version": 29,
        "totalOutOfTimeTrans": 8736,
        "lastWeekOutOfTimeTrans": 0,
        "changedBrokerage": false,
        "lowestBrokerage": 100
    },
    "data": [
        {
            "lastRanking": 1,
            "realTimeRanking": 1,
            "address": "TLyqzVGLV1srkB7dToTAEqgDSfPtXRJZYH",  // 验证节点地址
            "name": "Binance Staking",   // 验证节点名称
            "url": "https://www.binance.com/en/staking",
            "hasPage": false,
            "lastCycleVotes": 4243747149,
            "realTimeVotes": 4243730221,  // 投票数
            "changeVotes": -16928,
            "brokerage": 20,  // 分成比例 = （100-brokerage/100）
            "votesPercentage": 9.580180725139837,
            "lastCycleVotesPercentage": 9.58036306414994,
            "change_cycle": 0,
            "witnessType": 1,
            "annualizedRate": "3.155004",  // 年化收益
            "producedTotal": 1726935,
            "producedEfficiency": 99.48137548691255,
            "blockReward": 28248624,
            "version": 29,
            "totalOutOfTimeTrans": 16405,
            "lastWeekOutOfTimeTrans": 0,
            "changedBrokerage": false,
            "lowestBrokerage": 100
        }
        ...
    ]
}
```

# 2 获取token对TRX的交易对信息
