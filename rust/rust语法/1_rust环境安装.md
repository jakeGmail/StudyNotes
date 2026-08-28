
[toc]
# 1 安装Rust

**windows环境安装**

- 官方网址:  https://www.rust-lang.org 进入后点击install进入安装界面。安装好后需要重启电脑

安装好后在终端执行以下命令来确认环境是否安装好

```shell
# rustc是rust编译工具
rustc --version

# Cargo 是 Rust 的构建系统和包管理器
cargo --version

# rustup是用于更新rust的工具
rustup --version
```

- 还需要安装Visual Studio Community， 安装是需要勾选Windows 10/11 SDK和 MSVC，
在Visual Studio Installer中需要安装“使用C++桌面开发”，修改后需要重启电脑

然后执行以下命令

```shell
# 确认 Rust 使用 MSVC
rustup default stable-x86_64-pc-windows-msvc

# 执行完以下命令后应该能够看到host: x86_64-pc-windows-msvc
rustc -vV
```

**linux/macOS环境安装**

- 如果是linux系统，会给一个终端命令，使用终端命令来安装
    ```shell
    curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
    ```
- 查看rust版本
    ```shell
    rustc --version
    ```
- 更新Rust
    ```shell
    rustup update
    ```
- 卸载rust
    ```shell
    rustup self uninstall
    ```
- 在浏览器打开本地rust文档
    安装rust的时候还会下载rust文档
    ```shell
    rustup doc
    ```
# 2 Rust开发工具

- RustRover
  - Rust专门的IDE
  - 下载网址： https://www.jetbrains.com/rust/download/?section=windows
- VSCode
  - Rust插件
- Clion(Intellij Idea系列)
  - Rust插件

# 3 rust常用命令

```shell
# 更新rust, 更新到新发布的rust版本
rustup update

# 卸载rust
rustup self uninstall

# 编译程序，这会生成一个执行文件和一个.pdb文件； pdb文件中包含了调试信息
rustc <main.rc路径>

# 使用cargo创建项目，会自动生成一个项目文件目录
cargo new <项目文件夹名称>

# 自动在一个项目中创建Cargo.toml文件
cargo init

# 编译一个项目，会将编译结果放在./target/debug/<项目名称>这个文件夹下
# Cargo 还会在项目根目录创建一个新文件：Cargo.lock。这个文件会记录项目依赖的精确版本，类似于go.mod
cargo build 

# 编译发布版本
cargo build --release

# 编译并运行一个项目
cargo run

# 快速检查代码是否可以编译，但不产生任何文件
cargo check

# 升级Cargo.toml文件中的版本，但不会升级大版本
# 例如0.8.5可以升级到小于0.9.0的版本，但不会升级为0.9.0及其以上版本
# 如果需要升级到0.9.0及其以上，需要手动修改Cargo.toml中依赖项的版本号
cargo update
```