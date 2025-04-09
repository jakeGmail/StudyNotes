[toc]

# 1 uniswap-v2接口的schema
```shell
{
    "data": {
        "__schema": {
            "types": [
                {
                    "name": "Aggregation_interval",
                    "fields": null
                },
                {
                    "name": "BigDecimal",
                    "fields": null
                },
                {
                    "name": "BigInt",
                    "fields": null
                },
                {
                    "name": "BlockChangedFilter",
                    "fields": null
                },
                {
                    "name": "Block_height",
                    "fields": null
                },
                {
                    "name": "Boolean",
                    "fields": null
                },
                {
                    "name": "Bundle",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "ethPrice"
                        }
                    ]
                },
                {
                    "name": "Bundle_filter",
                    "fields": null
                },
                {
                    "name": "Bundle_orderBy",
                    "fields": null
                },
                {
                    "name": "Burn",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "transaction"
                        },
                        {
                            "name": "timestamp"
                        },
                        {
                            "name": "pair"
                        },
                        {
                            "name": "liquidity"
                        },
                        {
                            "name": "sender"
                        },
                        {
                            "name": "amount0"
                        },
                        {
                            "name": "amount1"
                        },
                        {
                            "name": "to"
                        },
                        {
                            "name": "logIndex"
                        },
                        {
                            "name": "amountUSD"
                        },
                        {
                            "name": "needsComplete"
                        },
                        {
                            "name": "feeTo"
                        },
                        {
                            "name": "feeLiquidity"
                        }
                    ]
                },
                {
                    "name": "Burn_filter",
                    "fields": null
                },
                {
                    "name": "Burn_orderBy",
                    "fields": null
                },
                {
                    "name": "Bytes",
                    "fields": null
                },
                {
                    "name": "Float",
                    "fields": null
                },
                {
                    "name": "ID",
                    "fields": null
                },
                {
                    "name": "Int",
                    "fields": null
                },
                {
                    "name": "Int8",
                    "fields": null
                },
                {
                    "name": "LiquidityPosition",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "user"
                        },
                        {
                            "name": "pair"
                        },
                        {
                            "name": "liquidityTokenBalance"
                        }
                    ]
                },
                {
                    "name": "LiquidityPositionSnapshot",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "liquidityPosition"
                        },
                        {
                            "name": "timestamp"
                        },
                        {
                            "name": "block"
                        },
                        {
                            "name": "user"
                        },
                        {
                            "name": "pair"
                        },
                        {
                            "name": "token0PriceUSD"
                        },
                        {
                            "name": "token1PriceUSD"
                        },
                        {
                            "name": "reserve0"
                        },
                        {
                            "name": "reserve1"
                        },
                        {
                            "name": "reserveUSD"
                        },
                        {
                            "name": "liquidityTokenTotalSupply"
                        },
                        {
                            "name": "liquidityTokenBalance"
                        }
                    ]
                },
                {
                    "name": "LiquidityPositionSnapshot_filter",
                    "fields": null
                },
                {
                    "name": "LiquidityPositionSnapshot_orderBy",
                    "fields": null
                },
                {
                    "name": "LiquidityPosition_filter",
                    "fields": null
                },
                {
                    "name": "LiquidityPosition_orderBy",
                    "fields": null
                },
                {
                    "name": "Mint",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "transaction"
                        },
                        {
                            "name": "timestamp"
                        },
                        {
                            "name": "pair"
                        },
                        {
                            "name": "to"
                        },
                        {
                            "name": "liquidity"
                        },
                        {
                            "name": "sender"
                        },
                        {
                            "name": "amount0"
                        },
                        {
                            "name": "amount1"
                        },
                        {
                            "name": "logIndex"
                        },
                        {
                            "name": "amountUSD"
                        },
                        {
                            "name": "feeTo"
                        },
                        {
                            "name": "feeLiquidity"
                        }
                    ]
                },
                {
                    "name": "Mint_filter",
                    "fields": null
                },
                {
                    "name": "Mint_orderBy",
                    "fields": null
                },
                {
                    "name": "OrderDirection",
                    "fields": null
                },
                {
                    "name": "Pair",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "token0"
                        },
                        {
                            "name": "token1"
                        },
                        {
                            "name": "reserve0"
                        },
                        {
                            "name": "reserve1"
                        },
                        {
                            "name": "totalSupply"
                        },
                        {
                            "name": "reserveETH"
                        },
                        {
                            "name": "reserveUSD"
                        },
                        {
                            "name": "trackedReserveETH"
                        },
                        {
                            "name": "token0Price"
                        },
                        {
                            "name": "token1Price"
                        },
                        {
                            "name": "volumeToken0"
                        },
                        {
                            "name": "volumeToken1"
                        },
                        {
                            "name": "volumeUSD"
                        },
                        {
                            "name": "untrackedVolumeUSD"
                        },
                        {
                            "name": "txCount"
                        },
                        {
                            "name": "createdAtTimestamp"
                        },
                        {
                            "name": "createdAtBlockNumber"
                        },
                        {
                            "name": "liquidityProviderCount"
                        },
                        {
                            "name": "pairHourData"
                        },
                        {
                            "name": "liquidityPositions"
                        },
                        {
                            "name": "liquidityPositionSnapshots"
                        },
                        {
                            "name": "mints"
                        },
                        {
                            "name": "burns"
                        },
                        {
                            "name": "swaps"
                        }
                    ]
                },
                {
                    "name": "PairDayData",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "date"
                        },
                        {
                            "name": "pairAddress"
                        },
                        {
                            "name": "token0"
                        },
                        {
                            "name": "token1"
                        },
                        {
                            "name": "reserve0"
                        },
                        {
                            "name": "reserve1"
                        },
                        {
                            "name": "totalSupply"
                        },
                        {
                            "name": "reserveUSD"
                        },
                        {
                            "name": "dailyVolumeToken0"
                        },
                        {
                            "name": "dailyVolumeToken1"
                        },
                        {
                            "name": "dailyVolumeUSD"
                        },
                        {
                            "name": "dailyTxns"
                        }
                    ]
                },
                {
                    "name": "PairDayData_filter",
                    "fields": null
                },
                {
                    "name": "PairDayData_orderBy",
                    "fields": null
                },
                {
                    "name": "PairHourData",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "hourStartUnix"
                        },
                        {
                            "name": "pair"
                        },
                        {
                            "name": "reserve0"
                        },
                        {
                            "name": "reserve1"
                        },
                        {
                            "name": "totalSupply"
                        },
                        {
                            "name": "reserveUSD"
                        },
                        {
                            "name": "hourlyVolumeToken0"
                        },
                        {
                            "name": "hourlyVolumeToken1"
                        },
                        {
                            "name": "hourlyVolumeUSD"
                        },
                        {
                            "name": "hourlyTxns"
                        }
                    ]
                },
                {
                    "name": "PairHourData_filter",
                    "fields": null
                },
                {
                    "name": "PairHourData_orderBy",
                    "fields": null
                },
                {
                    "name": "Pair_filter",
                    "fields": null
                },
                {
                    "name": "Pair_orderBy",
                    "fields": null
                },
                {
                    "name": "Query",
                    "fields": [
                        {
                            "name": "uniswapFactory"
                        },
                        {
                            "name": "uniswapFactories"
                        },
                        {
                            "name": "token"
                        },
                        {
                            "name": "tokens"
                        },
                        {
                            "name": "pair"
                        },
                        {
                            "name": "pairs"
                        },
                        {
                            "name": "user"
                        },
                        {
                            "name": "users"
                        },
                        {
                            "name": "liquidityPosition"
                        },
                        {
                            "name": "liquidityPositions"
                        },
                        {
                            "name": "liquidityPositionSnapshot"
                        },
                        {
                            "name": "liquidityPositionSnapshots"
                        },
                        {
                            "name": "transaction"
                        },
                        {
                            "name": "transactions"
                        },
                        {
                            "name": "mint"
                        },
                        {
                            "name": "mints"
                        },
                        {
                            "name": "burn"
                        },
                        {
                            "name": "burns"
                        },
                        {
                            "name": "swap"
                        },
                        {
                            "name": "swaps"
                        },
                        {
                            "name": "bundle"
                        },
                        {
                            "name": "bundles"
                        },
                        {
                            "name": "uniswapDayData"
                        },
                        {
                            "name": "uniswapDayDatas"
                        },
                        {
                            "name": "pairHourData"
                        },
                        {
                            "name": "pairHourDatas"
                        },
                        {
                            "name": "pairDayData"
                        },
                        {
                            "name": "pairDayDatas"
                        },
                        {
                            "name": "tokenDayData"
                        },
                        {
                            "name": "tokenDayDatas"
                        },
                        {
                            "name": "_meta"
                        }
                    ]
                },
                {
                    "name": "String",
                    "fields": null
                },
                {
                    "name": "Subscription",
                    "fields": [
                        {
                            "name": "uniswapFactory"
                        },
                        {
                            "name": "uniswapFactories"
                        },
                        {
                            "name": "token"
                        },
                        {
                            "name": "tokens"
                        },
                        {
                            "name": "pair"
                        },
                        {
                            "name": "pairs"
                        },
                        {
                            "name": "user"
                        },
                        {
                            "name": "users"
                        },
                        {
                            "name": "liquidityPosition"
                        },
                        {
                            "name": "liquidityPositions"
                        },
                        {
                            "name": "liquidityPositionSnapshot"
                        },
                        {
                            "name": "liquidityPositionSnapshots"
                        },
                        {
                            "name": "transaction"
                        },
                        {
                            "name": "transactions"
                        },
                        {
                            "name": "mint"
                        },
                        {
                            "name": "mints"
                        },
                        {
                            "name": "burn"
                        },
                        {
                            "name": "burns"
                        },
                        {
                            "name": "swap"
                        },
                        {
                            "name": "swaps"
                        },
                        {
                            "name": "bundle"
                        },
                        {
                            "name": "bundles"
                        },
                        {
                            "name": "uniswapDayData"
                        },
                        {
                            "name": "uniswapDayDatas"
                        },
                        {
                            "name": "pairHourData"
                        },
                        {
                            "name": "pairHourDatas"
                        },
                        {
                            "name": "pairDayData"
                        },
                        {
                            "name": "pairDayDatas"
                        },
                        {
                            "name": "tokenDayData"
                        },
                        {
                            "name": "tokenDayDatas"
                        },
                        {
                            "name": "_meta"
                        }
                    ]
                },
                {
                    "name": "Swap",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "transaction"
                        },
                        {
                            "name": "timestamp"
                        },
                        {
                            "name": "pair"
                        },
                        {
                            "name": "sender"
                        },
                        {
                            "name": "from"
                        },
                        {
                            "name": "amount0In"
                        },
                        {
                            "name": "amount1In"
                        },
                        {
                            "name": "amount0Out"
                        },
                        {
                            "name": "amount1Out"
                        },
                        {
                            "name": "to"
                        },
                        {
                            "name": "logIndex"
                        },
                        {
                            "name": "amountUSD"
                        }
                    ]
                },
                {
                    "name": "Swap_filter",
                    "fields": null
                },
                {
                    "name": "Swap_orderBy",
                    "fields": null
                },
                {
                    "name": "Token",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "symbol"
                        },
                        {
                            "name": "name"
                        },
                        {
                            "name": "decimals"
                        },
                        {
                            "name": "totalSupply"
                        },
                        {
                            "name": "tradeVolume"
                        },
                        {
                            "name": "tradeVolumeUSD"
                        },
                        {
                            "name": "untrackedVolumeUSD"
                        },
                        {
                            "name": "txCount"
                        },
                        {
                            "name": "totalLiquidity"
                        },
                        {
                            "name": "derivedETH"
                        },
                        {
                            "name": "tokenDayData"
                        },
                        {
                            "name": "pairDayDataBase"
                        },
                        {
                            "name": "pairDayDataQuote"
                        },
                        {
                            "name": "pairBase"
                        },
                        {
                            "name": "pairQuote"
                        }
                    ]
                },
                {
                    "name": "TokenDayData",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "date"
                        },
                        {
                            "name": "token"
                        },
                        {
                            "name": "dailyVolumeToken"
                        },
                        {
                            "name": "dailyVolumeETH"
                        },
                        {
                            "name": "dailyVolumeUSD"
                        },
                        {
                            "name": "dailyTxns"
                        },
                        {
                            "name": "totalLiquidityToken"
                        },
                        {
                            "name": "totalLiquidityETH"
                        },
                        {
                            "name": "totalLiquidityUSD"
                        },
                        {
                            "name": "priceUSD"
                        }
                    ]
                },
                {
                    "name": "TokenDayData_filter",
                    "fields": null
                },
                {
                    "name": "TokenDayData_orderBy",
                    "fields": null
                },
                {
                    "name": "Token_filter",
                    "fields": null
                },
                {
                    "name": "Token_orderBy",
                    "fields": null
                },
                {
                    "name": "Transaction",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "blockNumber"
                        },
                        {
                            "name": "timestamp"
                        },
                        {
                            "name": "mints"
                        },
                        {
                            "name": "burns"
                        },
                        {
                            "name": "swaps"
                        }
                    ]
                },
                {
                    "name": "Transaction_filter",
                    "fields": null
                },
                {
                    "name": "Transaction_orderBy",
                    "fields": null
                },
                {
                    "name": "UniswapDayData",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "date"
                        },
                        {
                            "name": "dailyVolumeETH"
                        },
                        {
                            "name": "dailyVolumeUSD"
                        },
                        {
                            "name": "dailyVolumeUntracked"
                        },
                        {
                            "name": "totalVolumeETH"
                        },
                        {
                            "name": "totalLiquidityETH"
                        },
                        {
                            "name": "totalVolumeUSD"
                        },
                        {
                            "name": "totalLiquidityUSD"
                        },
                        {
                            "name": "txCount"
                        }
                    ]
                },
                {
                    "name": "UniswapDayData_filter",
                    "fields": null
                },
                {
                    "name": "UniswapDayData_orderBy",
                    "fields": null
                },
                {
                    "name": "UniswapFactory",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "pairCount"
                        },
                        {
                            "name": "totalVolumeUSD"
                        },
                        {
                            "name": "totalVolumeETH"
                        },
                        {
                            "name": "untrackedVolumeUSD"
                        },
                        {
                            "name": "totalLiquidityUSD"
                        },
                        {
                            "name": "totalLiquidityETH"
                        },
                        {
                            "name": "txCount"
                        }
                    ]
                },
                {
                    "name": "UniswapFactory_filter",
                    "fields": null
                },
                {
                    "name": "UniswapFactory_orderBy",
                    "fields": null
                },
                {
                    "name": "User",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "liquidityPositions"
                        },
                        {
                            "name": "usdSwapped"
                        }
                    ]
                },
                {
                    "name": "User_filter",
                    "fields": null
                },
                {
                    "name": "User_orderBy",
                    "fields": null
                },
                {
                    "name": "_Block_",
                    "fields": [
                        {
                            "name": "hash"
                        },
                        {
                            "name": "number"
                        },
                        {
                            "name": "timestamp"
                        },
                        {
                            "name": "parentHash"
                        }
                    ]
                },
                {
                    "name": "_Meta_",
                    "fields": [
                        {
                            "name": "block"
                        },
                        {
                            "name": "deployment"
                        },
                        {
                            "name": "hasIndexingErrors"
                        }
                    ]
                },
                {
                    "name": "_SubgraphErrorPolicy_",
                    "fields": null
                }
            ]
        }
    }
}
```


