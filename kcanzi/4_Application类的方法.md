[toc]
由于我们自己的逻辑处理的类是继承自框架的Application类，因此有必要了解Application类中提供的常见API。
全部详细API位于[SampleCode/Engine/include/kanzi/core.ui/application/application.hpp]()

# 1 处理按键事件
  ```c++
  // 处理按键事件
  void onKeyInputEvent(const KzsInputEventKey* inputData);
  ```
  使用示例：
  ```c++
  void CarApplication::onKeyInputEvent(const KzsInputEventKey* inputData){
        KzsInputKey button = kzsInputEventKeyGetButton(inputData);
        if (button == KZS_KEY_ESC || button == KZS_KEY_Q || button == KZS_KEY_BACKSPACE)
        {
            // 逻辑代码
            ...
        }
    }
  ```

# 2 处理鼠标/触摸事件
```c++
void onPointerInputEvent(const KzsInputEventPointer* inputData);
```

# 退出程序包
  ```c++
  // 退出程序
  void quit();
  ```