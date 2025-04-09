[toc]
# 1 FMath介绍
Fmath是Unreal平台相关的数学库

# 2 FMath中的方法
## 2.1 Clamp
```c++
template< class T > 
static T Clamp( const T X, const T Min, const T Max )
{
    return X<Min ? Min : X<Max ? X : Max;
}
```
**作用**：夹值，如果X在[Min， Max]范围内，则返回X。 如果X小于Min则返回Min。如果X大于Max则返回Max。使得Clamp的返回值始终在[Min, Max]中。

## 2.2 
```c++
int32 RandRange(int32 Min, int32 Max);
int64 RandRange(int64 Min, int64 Max);
float RandRange(float InMin, float InMax)
```
**作用**：随机返回[Min, Max]区间中的值