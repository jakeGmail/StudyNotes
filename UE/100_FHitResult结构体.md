[toc]
# 1 FHitResult介绍
FHitResult结构体用于描述发生碰撞时的碰撞信息

# 2 碰撞位置
在FHitResult中，`Location`变量中，记录发生碰撞时的位置.
```c++
USTRUCT(BlueprintType, meta = (HasNativeBreak = "Engine.GameplayStatics.BreakHitResult", HasNativeMake = "Engine.GameplayStatics.MakeHitResult"))
struct ENGINE_API FHitResult
{
    ...
    UPROPERTY()
    FVector_NetQuantize Location;
    ...
}
```
FVector_NetQuantize继承自FVector。
- `FHitResult.Location.X`: 碰撞位置的X坐标
- `FHitResult.Location.Y`: 碰撞位置的Y坐标
- `FHitResult.Location.Z`: 碰撞位置的Z坐标