# 1 SurfaceFlinger介绍
SurfaceFlinger是Android系统与本地窗口进行通信的协议。绘制界面的应用是SurfaceFlinger的客户端，在客户端绘制好后，会将绘制的buffer发送带到SurfaceFlinger服务中，由SurfaceFlinger服务对各个客户端的buffer进行合成然后发送到屏幕。
SurfaceFlinger 接受缓冲区，对它们进行合成，然后发送到屏幕。WindowManager 为 SurfaceFlinger 提供缓冲区和窗口元数据，而 SurfaceFlinger 可使用这些信息将 Surface 合成到屏幕。

SurfaceFlinger 可通过两种方式接受缓冲区：通过 BufferQueue 和 SurfaceControl，或通过 ASurfaceControl。

**BufferQueue 和 SurfaceControl方式**： 当应用进入前台时，它会从 WindowManager 请求缓冲区。然后，WindowManager 会从 SurfaceFlinger 请求层。层是 surface（包含 BufferQueue）和 SurfaceControl（包含显示框架等层元数据）的组合。SurfaceFlinger 创建层并将其发送至 WindowManager。然后，WindowManager 将 Surface 发送至应用，但会保留 SurfaceControl 来操控应用在屏幕上的外观。

**ASurfaceControl方式**： Android 10 新增了 ASurfaceControl，这是 SurfaceFlinger 接受缓冲区的另一种方式。ASurfaceControl 将 Surface 和 SurfaceControl 组合到一个事务包中，该包会被发送至 SurfaceFlinger。ASurfaceControl 与层相关联，应用可通过 ASurfaceTransactions 更新该层。然后，应用可通过回调（用于传递包含锁定时间、获取时间等信息的 ASurfaceTransactionStats）获取有关 ASurfaceTransactions 的信息。

下表包含有关 ASurfaceControl 及其相关组件的更多详细信息。

|组件|	说明|
|----|-----|
|ASurfaceControl|	对 SurfaceControl 进行封装，并让应用能够创建与屏幕上的各层相对应的 SurfaceControl。可作为 ANativeWindow 的一个子级或者另一个 ASurfaceControl 的子级创建。|
|ASurfaceTransaction|	对事务进行包装，以使客户端能够修改层的描述性属性（比如几何图形），并将经过更新的缓冲区发送至 SurfaceFlinger。|
|ASurfaceTransactionStats|	通过预先注册的回调将有关已显示事务的信息（比如锁定时间、获取时间和上一个释放栅栏）发送至应用。|

虽然应用可以随时提交缓冲区，但 SurfaceFlinger 仅能在屏幕处于两次刷新之间时唤醒，以接受缓冲区，这会因设备而异。这样可以最大限度地减少内存使用量，并避免屏幕上出现可见的撕裂现象（如果显示内容在刷新期间更新，则会出现此现象）。

在屏幕处于两次刷新之间时，屏幕会向 SurfaceFlinger 发送 VSYNC 信号。VSYNC 信号表明可对屏幕进行刷新而不会产生撕裂。当 SurfaceFlinger 接收到 VSYNC 信号后，SurfaceFlinger 会遍历其层列表，以查找新的缓冲区。如果 SurfaceFlinger 找到新的缓冲区，SurfaceFlinger 会获取缓冲区；否则，SurfaceFlinger 会继续使用上一次获取的那个缓冲区。SurfaceFlinger 必须始终显示内容，因此它会保留一个缓冲区。如果在某个层上没有提交缓冲区，则该层会被忽略。

SurfaceFlinger 在收集可见层的所有缓冲区之后，便会询问硬件混合渲染器 (HWC) 应如何进行合成。如果 HWC 将层合成类型标记为客户端合成，则 SurfaceFlinger 将合成这些层。然后，SurfaceFlinger 会将输出缓冲区传递给 HWC。

# 2 SurfaceFlinger原理
