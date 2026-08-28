
|目录|
|----|
|[1-标准输入输出](#1-标准输入输出)|
|[2-rand](#2-随机数rand)|
|[3-String](#3-string用法)|


# 1 标准输入/输出

使用标准输入输出需要 `use std::io`

```rust
use std::io;

fn main() {
    println!("请输入数字：");
    let mut gues = String::new();

    /*read_line这个函数会将终端输入的一行信息以追加的方式写入gues变量中，
    如果希望每次写入时gues中没有之前的数据，可以使用gues.clear()函数清除
    gues中的数据，使之变为一个空字符串。
    
    
    */
    io::stdin().read_line(&mut gues).expect("Failed to read line");
    print!("{}", gues);
}
```


# 2 随机数rand

使用随机数需要`use rand::Rng;`并在Cargo.toml的依赖[section]中添加随机数crate---`rand = "0.9.5"`

```rust
use rand::Rng;

fn main() {
    /*..=是一个范围表达式， 表示区间[1,10], 包含两端的值*/
    let randNumber1 = rand::thread_rng().gen_range(1..= 10);

    /*..是一个范围表达式，表示区间[1,10), 包含1，但不包含10*/
    let randNumber2 = rand::thread_rng().gen_range(1 .. 10);
}
```

# 3 String用法

```rust
let mut str = String::new();

// 删除str中的所有空白字符
str = str.trim().to_string();
```


