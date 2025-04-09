# 2 uniswap-v3接口schema

```
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
                            "name": "ethPriceUSD"
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
                            "name": "pool"
                        },
                        {
                            "name": "token0"
                        },
                        {
                            "name": "token1"
                        },
                        {
                            "name": "timestamp"
                        },
                        {
                            "name": "owner"
                        },
                        {
                            "name": "origin"
                        },
                        {
                            "name": "amount"
                        },
                        {
                            "name": "amount0"
                        },
                        {
                            "name": "amount1"
                        },
                        {
                            "name": "amountUSD"
                        },
                        {
                            "name": "tickLower"
                        },
                        {
                            "name": "tickUpper"
                        },
                        {
                            "name": "logIndex"
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
                    "name": "Collect",
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
                            "name": "pool"
                        },
                        {
                            "name": "owner"
                        },
                        {
                            "name": "amount0"
                        },
                        {
                            "name": "amount1"
                        },
                        {
                            "name": "amountUSD"
                        },
                        {
                            "name": "tickLower"
                        },
                        {
                            "name": "tickUpper"
                        },
                        {
                            "name": "logIndex"
                        }
                    ]
                },
                {
                    "name": "Collect_filter",
                    "fields": null
                },
                {
                    "name": "Collect_orderBy",
                    "fields": null
                },
                {
                    "name": "Factory",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "poolCount"
                        },
                        {
                            "name": "txCount"
                        },
                        {
                            "name": "totalVolumeUSD"
                        },
                        {
                            "name": "totalVolumeETH"
                        },
                        {
                            "name": "totalFeesUSD"
                        },
                        {
                            "name": "totalFeesETH"
                        },
                        {
                            "name": "untrackedVolumeUSD"
                        },
                        {
                            "name": "totalValueLockedUSD"
                        },
                        {
                            "name": "totalValueLockedETH"
                        },
                        {
                            "name": "totalValueLockedUSDUntracked"
                        },
                        {
                            "name": "totalValueLockedETHUntracked"
                        },
                        {
                            "name": "owner"
                        }
                    ]
                },
                {
                    "name": "Factory_filter",
                    "fields": null
                },
                {
                    "name": "Factory_orderBy",
                    "fields": null
                },
                {
                    "name": "Flash",
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
                            "name": "pool"
                        },
                        {
                            "name": "sender"
                        },
                        {
                            "name": "recipient"
                        },
                        {
                            "name": "amount0"
                        },
                        {
                            "name": "amount1"
                        },
                        {
                            "name": "amountUSD"
                        },
                        {
                            "name": "amount0Paid"
                        },
                        {
                            "name": "amount1Paid"
                        },
                        {
                            "name": "logIndex"
                        }
                    ]
                },
                {
                    "name": "Flash_filter",
                    "fields": null
                },
                {
                    "name": "Flash_orderBy",
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
                            "name": "pool"
                        },
                        {
                            "name": "token0"
                        },
                        {
                            "name": "token1"
                        },
                        {
                            "name": "owner"
                        },
                        {
                            "name": "sender"
                        },
                        {
                            "name": "origin"
                        },
                        {
                            "name": "amount"
                        },
                        {
                            "name": "amount0"
                        },
                        {
                            "name": "amount1"
                        },
                        {
                            "name": "amountUSD"
                        },
                        {
                            "name": "tickLower"
                        },
                        {
                            "name": "tickUpper"
                        },
                        {
                            "name": "logIndex"
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
                    "name": "Pool",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "createdAtTimestamp"
                        },
                        {
                            "name": "createdAtBlockNumber"
                        },
                        {
                            "name": "token0"
                        },
                        {
                            "name": "token1"
                        },
                        {
                            "name": "feeTier"
                        },
                        {
                            "name": "liquidity"
                        },
                        {
                            "name": "sqrtPrice"
                        },
                        {
                            "name": "feeGrowthGlobal0X128"
                        },
                        {
                            "name": "feeGrowthGlobal1X128"
                        },
                        {
                            "name": "token0Price"
                        },
                        {
                            "name": "token1Price"
                        },
                        {
                            "name": "tick"
                        },
                        {
                            "name": "observationIndex"
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
                            "name": "feesUSD"
                        },
                        {
                            "name": "txCount"
                        },
                        {
                            "name": "collectedFeesToken0"
                        },
                        {
                            "name": "collectedFeesToken1"
                        },
                        {
                            "name": "collectedFeesUSD"
                        },
                        {
                            "name": "totalValueLockedToken0"
                        },
                        {
                            "name": "totalValueLockedToken1"
                        },
                        {
                            "name": "totalValueLockedETH"
                        },
                        {
                            "name": "totalValueLockedUSD"
                        },
                        {
                            "name": "totalValueLockedUSDUntracked"
                        },
                        {
                            "name": "liquidityProviderCount"
                        },
                        {
                            "name": "poolHourData"
                        },
                        {
                            "name": "poolDayData"
                        },
                        {
                            "name": "mints"
                        },
                        {
                            "name": "burns"
                        },
                        {
                            "name": "swaps"
                        },
                        {
                            "name": "collects"
                        },
                        {
                            "name": "ticks"
                        }
                    ]
                },
                {
                    "name": "PoolDayData",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "date"
                        },
                        {
                            "name": "pool"
                        },
                        {
                            "name": "liquidity"
                        },
                        {
                            "name": "sqrtPrice"
                        },
                        {
                            "name": "token0Price"
                        },
                        {
                            "name": "token1Price"
                        },
                        {
                            "name": "tick"
                        },
                        {
                            "name": "feeGrowthGlobal0X128"
                        },
                        {
                            "name": "feeGrowthGlobal1X128"
                        },
                        {
                            "name": "tvlUSD"
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
                            "name": "feesUSD"
                        },
                        {
                            "name": "txCount"
                        },
                        {
                            "name": "open"
                        },
                        {
                            "name": "high"
                        },
                        {
                            "name": "low"
                        },
                        {
                            "name": "close"
                        }
                    ]
                },
                {
                    "name": "PoolDayData_filter",
                    "fields": null
                },
                {
                    "name": "PoolDayData_orderBy",
                    "fields": null
                },
                {
                    "name": "PoolHourData",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "periodStartUnix"
                        },
                        {
                            "name": "pool"
                        },
                        {
                            "name": "liquidity"
                        },
                        {
                            "name": "sqrtPrice"
                        },
                        {
                            "name": "token0Price"
                        },
                        {
                            "name": "token1Price"
                        },
                        {
                            "name": "tick"
                        },
                        {
                            "name": "feeGrowthGlobal0X128"
                        },
                        {
                            "name": "feeGrowthGlobal1X128"
                        },
                        {
                            "name": "tvlUSD"
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
                            "name": "feesUSD"
                        },
                        {
                            "name": "txCount"
                        },
                        {
                            "name": "open"
                        },
                        {
                            "name": "high"
                        },
                        {
                            "name": "low"
                        },
                        {
                            "name": "close"
                        }
                    ]
                },
                {
                    "name": "PoolHourData_filter",
                    "fields": null
                },
                {
                    "name": "PoolHourData_orderBy",
                    "fields": null
                },
                {
                    "name": "Pool_filter",
                    "fields": null
                },
                {
                    "name": "Pool_orderBy",
                    "fields": null
                },
                {
                    "name": "Position",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "owner"
                        },
                        {
                            "name": "pool"
                        },
                        {
                            "name": "token0"
                        },
                        {
                            "name": "token1"
                        },
                        {
                            "name": "tickLower"
                        },
                        {
                            "name": "tickUpper"
                        },
                        {
                            "name": "liquidity"
                        },
                        {
                            "name": "depositedToken0"
                        },
                        {
                            "name": "depositedToken1"
                        },
                        {
                            "name": "withdrawnToken0"
                        },
                        {
                            "name": "withdrawnToken1"
                        },
                        {
                            "name": "collectedFeesToken0"
                        },
                        {
                            "name": "collectedFeesToken1"
                        },
                        {
                            "name": "transaction"
                        },
                        {
                            "name": "feeGrowthInside0LastX128"
                        },
                        {
                            "name": "feeGrowthInside1LastX128"
                        }
                    ]
                },
                {
                    "name": "PositionSnapshot",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "owner"
                        },
                        {
                            "name": "pool"
                        },
                        {
                            "name": "position"
                        },
                        {
                            "name": "blockNumber"
                        },
                        {
                            "name": "timestamp"
                        },
                        {
                            "name": "liquidity"
                        },
                        {
                            "name": "depositedToken0"
                        },
                        {
                            "name": "depositedToken1"
                        },
                        {
                            "name": "withdrawnToken0"
                        },
                        {
                            "name": "withdrawnToken1"
                        },
                        {
                            "name": "collectedFeesToken0"
                        },
                        {
                            "name": "collectedFeesToken1"
                        },
                        {
                            "name": "transaction"
                        },
                        {
                            "name": "feeGrowthInside0LastX128"
                        },
                        {
                            "name": "feeGrowthInside1LastX128"
                        }
                    ]
                },
                {
                    "name": "PositionSnapshot_filter",
                    "fields": null
                },
                {
                    "name": "PositionSnapshot_orderBy",
                    "fields": null
                },
                {
                    "name": "Position_filter",
                    "fields": null
                },
                {
                    "name": "Position_orderBy",
                    "fields": null
                },
                {
                    "name": "Query",
                    "fields": [
                        {
                            "name": "factory"
                        },
                        {
                            "name": "factories"
                        },
                        {
                            "name": "bundle"
                        },
                        {
                            "name": "bundles"
                        },
                        {
                            "name": "token"
                        },
                        {
                            "name": "tokens"
                        },
                        {
                            "name": "pool"
                        },
                        {
                            "name": "pools"
                        },
                        {
                            "name": "tick"
                        },
                        {
                            "name": "ticks"
                        },
                        {
                            "name": "position"
                        },
                        {
                            "name": "positions"
                        },
                        {
                            "name": "positionSnapshot"
                        },
                        {
                            "name": "positionSnapshots"
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
                            "name": "collect"
                        },
                        {
                            "name": "collects"
                        },
                        {
                            "name": "flash"
                        },
                        {
                            "name": "flashes"
                        },
                        {
                            "name": "uniswapDayData"
                        },
                        {
                            "name": "uniswapDayDatas"
                        },
                        {
                            "name": "poolDayData"
                        },
                        {
                            "name": "poolDayDatas"
                        },
                        {
                            "name": "poolHourData"
                        },
                        {
                            "name": "poolHourDatas"
                        },
                        {
                            "name": "tickHourData"
                        },
                        {
                            "name": "tickHourDatas"
                        },
                        {
                            "name": "tickDayData"
                        },
                        {
                            "name": "tickDayDatas"
                        },
                        {
                            "name": "tokenDayData"
                        },
                        {
                            "name": "tokenDayDatas"
                        },
                        {
                            "name": "tokenHourData"
                        },
                        {
                            "name": "tokenHourDatas"
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
                            "name": "factory"
                        },
                        {
                            "name": "factories"
                        },
                        {
                            "name": "bundle"
                        },
                        {
                            "name": "bundles"
                        },
                        {
                            "name": "token"
                        },
                        {
                            "name": "tokens"
                        },
                        {
                            "name": "pool"
                        },
                        {
                            "name": "pools"
                        },
                        {
                            "name": "tick"
                        },
                        {
                            "name": "ticks"
                        },
                        {
                            "name": "position"
                        },
                        {
                            "name": "positions"
                        },
                        {
                            "name": "positionSnapshot"
                        },
                        {
                            "name": "positionSnapshots"
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
                            "name": "collect"
                        },
                        {
                            "name": "collects"
                        },
                        {
                            "name": "flash"
                        },
                        {
                            "name": "flashes"
                        },
                        {
                            "name": "uniswapDayData"
                        },
                        {
                            "name": "uniswapDayDatas"
                        },
                        {
                            "name": "poolDayData"
                        },
                        {
                            "name": "poolDayDatas"
                        },
                        {
                            "name": "poolHourData"
                        },
                        {
                            "name": "poolHourDatas"
                        },
                        {
                            "name": "tickHourData"
                        },
                        {
                            "name": "tickHourDatas"
                        },
                        {
                            "name": "tickDayData"
                        },
                        {
                            "name": "tickDayDatas"
                        },
                        {
                            "name": "tokenDayData"
                        },
                        {
                            "name": "tokenDayDatas"
                        },
                        {
                            "name": "tokenHourData"
                        },
                        {
                            "name": "tokenHourDatas"
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
                            "name": "pool"
                        },
                        {
                            "name": "token0"
                        },
                        {
                            "name": "token1"
                        },
                        {
                            "name": "sender"
                        },
                        {
                            "name": "recipient"
                        },
                        {
                            "name": "origin"
                        },
                        {
                            "name": "amount0"
                        },
                        {
                            "name": "amount1"
                        },
                        {
                            "name": "amountUSD"
                        },
                        {
                            "name": "sqrtPriceX96"
                        },
                        {
                            "name": "tick"
                        },
                        {
                            "name": "logIndex"
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
                    "name": "Tick",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "poolAddress"
                        },
                        {
                            "name": "tickIdx"
                        },
                        {
                            "name": "pool"
                        },
                        {
                            "name": "liquidityGross"
                        },
                        {
                            "name": "liquidityNet"
                        },
                        {
                            "name": "price0"
                        },
                        {
                            "name": "price1"
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
                            "name": "feesUSD"
                        },
                        {
                            "name": "collectedFeesToken0"
                        },
                        {
                            "name": "collectedFeesToken1"
                        },
                        {
                            "name": "collectedFeesUSD"
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
                            "name": "feeGrowthOutside0X128"
                        },
                        {
                            "name": "feeGrowthOutside1X128"
                        }
                    ]
                },
                {
                    "name": "TickDayData",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "date"
                        },
                        {
                            "name": "pool"
                        },
                        {
                            "name": "tick"
                        },
                        {
                            "name": "liquidityGross"
                        },
                        {
                            "name": "liquidityNet"
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
                            "name": "feesUSD"
                        },
                        {
                            "name": "feeGrowthOutside0X128"
                        },
                        {
                            "name": "feeGrowthOutside1X128"
                        }
                    ]
                },
                {
                    "name": "TickDayData_filter",
                    "fields": null
                },
                {
                    "name": "TickDayData_orderBy",
                    "fields": null
                },
                {
                    "name": "TickHourData",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "periodStartUnix"
                        },
                        {
                            "name": "pool"
                        },
                        {
                            "name": "tick"
                        },
                        {
                            "name": "liquidityGross"
                        },
                        {
                            "name": "liquidityNet"
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
                            "name": "feesUSD"
                        }
                    ]
                },
                {
                    "name": "TickHourData_filter",
                    "fields": null
                },
                {
                    "name": "TickHourData_orderBy",
                    "fields": null
                },
                {
                    "name": "Tick_filter",
                    "fields": null
                },
                {
                    "name": "Tick_orderBy",
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
                            "name": "volume"
                        },
                        {
                            "name": "volumeUSD"
                        },
                        {
                            "name": "untrackedVolumeUSD"
                        },
                        {
                            "name": "feesUSD"
                        },
                        {
                            "name": "txCount"
                        },
                        {
                            "name": "poolCount"
                        },
                        {
                            "name": "totalValueLocked"
                        },
                        {
                            "name": "totalValueLockedUSD"
                        },
                        {
                            "name": "totalValueLockedUSDUntracked"
                        },
                        {
                            "name": "derivedETH"
                        },
                        {
                            "name": "whitelistPools"
                        },
                        {
                            "name": "tokenDayData"
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
                            "name": "volume"
                        },
                        {
                            "name": "volumeUSD"
                        },
                        {
                            "name": "untrackedVolumeUSD"
                        },
                        {
                            "name": "totalValueLocked"
                        },
                        {
                            "name": "totalValueLockedUSD"
                        },
                        {
                            "name": "priceUSD"
                        },
                        {
                            "name": "feesUSD"
                        },
                        {
                            "name": "open"
                        },
                        {
                            "name": "high"
                        },
                        {
                            "name": "low"
                        },
                        {
                            "name": "close"
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
                    "name": "TokenHourData",
                    "fields": [
                        {
                            "name": "id"
                        },
                        {
                            "name": "periodStartUnix"
                        },
                        {
                            "name": "token"
                        },
                        {
                            "name": "volume"
                        },
                        {
                            "name": "volumeUSD"
                        },
                        {
                            "name": "untrackedVolumeUSD"
                        },
                        {
                            "name": "totalValueLocked"
                        },
                        {
                            "name": "totalValueLockedUSD"
                        },
                        {
                            "name": "priceUSD"
                        },
                        {
                            "name": "feesUSD"
                        },
                        {
                            "name": "open"
                        },
                        {
                            "name": "high"
                        },
                        {
                            "name": "low"
                        },
                        {
                            "name": "close"
                        }
                    ]
                },
                {
                    "name": "TokenHourData_filter",
                    "fields": null
                },
                {
                    "name": "TokenHourData_orderBy",
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
                            "name": "gasUsed"
                        },
                        {
                            "name": "gasPrice"
                        },
                        {
                            "name": "mints"
                        },
                        {
                            "name": "burns"
                        },
                        {
                            "name": "swaps"
                        },
                        {
                            "name": "flashed"
                        },
                        {
                            "name": "collects"
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
                            "name": "volumeETH"
                        },
                        {
                            "name": "volumeUSD"
                        },
                        {
                            "name": "volumeUSDUntracked"
                        },
                        {
                            "name": "feesUSD"
                        },
                        {
                            "name": "txCount"
                        },
                        {
                            "name": "tvlUSD"
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