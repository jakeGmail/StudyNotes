
# 2 boost::shared_ptr使用
boost::shared_ptr跟std::shared_ptr用法相似
```c++
#include <boost/shared_ptr.hpp>
#include <boost/make_shared.hpp>
#include <boost/scoped_ptr.hpp> // 对应std::unique_ptr
class A{
};
int main() {
    // 创建智能指针
    boost::shared_ptr<A> a0(new A);
    boost::shared_ptr<A> a1 = boost::make_shared<A>();

    // 获取持有的指针
    A* pa = a0.get();

    // 重新设置持有指针
    A* a =  new A;
    a0.reset(a);

    /////////////////////////////////////////////////////

    // 创建scoped_ptr
    boost::scoped_ptr<A> a2(new A);

    // 获取持有的指针
    A* pa = a0.get();

    // 重新设置持有指针
    A* a =  new A;
    a0.reset(a);
    return 0;
}

```