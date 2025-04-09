
绑定普通函数
```c++
#include <boost/bind.hpp>
#include <boost/function.hpp>


void tt(int a, bool b){}

int main(){
    boost::function<void(int,bool)> aa = boost::bind(tt, boost::placeholders::_1, boost::placeholders::_2);
    return 0;
}
```

绑定类中的函数
```c++
#include <boost/bind.hpp>
#include <boost/function.hpp>

class A{
public:
    void test(int a){};
};

int main(){
    A a;
    boost::function<void(int)> aa = boost::bind(&A::test, &a, boost::placeholders::_1);
    return 0;
}
```