[toc]

# 1 glfwInit
**定义**: int glfwInit(void)
**作用**：初始化GLFW
**注意**：需要在调用其他glfw API之前调用本函数

# 2 glfwWindowHint 
**定义**: void glfwWindowHint(int hint, int value)
**作用**：设置glfw的参数
**参数**：
  - hint: 选项的名称，我们可以从很多以GLFW_开头的枚举值中选择
  - value: 设置这个选项的值

**使用示例**：
```c
glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, 3); // 设置glfw的主版本
glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, 3); // 设置glfw次版本
glfwWindowHint(GLFW_OPENGL_PROFILE, GLFW_OPENGL_CORE_PROFILE); // 设置渲染模式为核心模式
```

**附加**:
|hint参数可选值|value参数的默认值|对应value参数的值|
|--------------|---------------|-----------------|
|GLFW_RESIZABLE(设置窗口大小是否可以改变)|	GLFW_TRUE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_VISIBLE(设置窗口是否可见)|	GLFW_TRUE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_DECORATED(设置窗口是否带有装饰，像边框/最小化/最大化/退出按钮)|	GLFW_TRUE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_FOCUSED(窗口创建时是否获得输入焦点，即打字可以输入到该窗口)|	GLFW_TRUE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_AUTO_ICONIFY|	GLFW_TRUE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_FLOATING|	GLFW_FALSE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_MAXIMIZED|	GLFW_FALSE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_CENTER_CURSOR|	GLFW_TRUE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_TRANSPARENT_FRAMEBUFFER|	GLFW_FALSE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_FOCUS_ON_SHOW|	GLFW_TRUE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_SCALE_TO_MONITOR|	GLFW_FALSE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_SCALE_FRAMEBUFFER|	GLFW_TRUE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_MOUSE_PASSTHROUGH|	GLFW_FALSE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_POSITION_X|	GLFW_ANY_POSITION|	Any valid screen x-coordinate or GLFW_ANY_POSITION|
|GLFW_POSITION_Y|	GLFW_ANY_POSITION|	Any valid screen y-coordinate or GLFW_ANY_POSITION|
|GLFW_RED_BITS|	8|	0 to or INT_MAXGLFW_DONT_CARE|
|GLFW_GREEN_BITS|	8|	0 to or INT_MAXGLFW_DONT_CARE|
|GLFW_BLUE_BITS|	8|	0 to or INT_MAXGLFW_DONT_CARE|
|GLFW_ALPHA_BITS|	8|	0 to or INT_MAXGLFW_DONT_CARE|
|GLFW_DEPTH_BITS|	24|	0 to or INT_MAXGLFW_DONT_CARE|
|GLFW_STENCIL_BITS|	8|	0 to or INT_MAXGLFW_DONT_CARE|
|GLFW_ACCUM_RED_BITS|	0|	0 to or INT_MAXGLFW_DONT_CARE|
|GLFW_ACCUM_GREEN_BITS|	0|	0 to or INT_MAXGLFW_DONT_CARE|
|GLFW_ACCUM_BLUE_BITS|	0|	0 to or INT_MAXGLFW_DONT_CARE|
|GLFW_ACCUM_ALPHA_BITS|	0|	0 to or INT_MAXGLFW_DONT_CARE|
|GLFW_AUX_BUFFERS|	0|	0 to or INT_MAXGLFW_DONT_CARE|
|GLFW_SAMPLES|	0|	0 to or INT_MAXGLFW_DONT_CARE|
|GLFW_REFRESH_RATE|	GLFW_DONT_CARE|	0 to or INT_MAXGLFW_DONT_CARE|
|GLFW_STEREO|	GLFW_FALSE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_SRGB_CAPABLE|	GLFW_FALSE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_DOUBLEBUFFER|	GLFW_TRUE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_CLIENT_API|	GLFW_OPENGL_API|	GLFW_OPENGL_API, or GLFW_OPENGL_ES_APIGLFW_NO_API|
|GLFW_CONTEXT_CREATION_API|	GLFW_NATIVE_CONTEXT_API	GLFW_NATIVE_CONTEXT_API, or GLFW_EGL_CONTEXT_APIGLFW_OSMESA_CONTEXT_API|
|GLFW_CONTEXT_VERSION_MAJOR（设置主版本号）|	1|	Any valid major version number of the chosen client API|
|GLFW_CONTEXT_VERSION_MINOR（设置次版本号）|	0|	Any valid minor version number of the chosen client API|
|GLFW_CONTEXT_ROBUSTNESS|	GLFW_NO_ROBUSTNESS|	GLFW_NO_ROBUSTNESS, or GLFW_NO_RESET_NOTIFICATIONGLFW_LOSE_CONTEXT_ON_RESET
|GLFW_CONTEXT_RELEASE_BEHAVIOR|	GLFW_ANY_RELEASE_BEHAVIOR|	GLFW_ANY_RELEASE_BEHAVIOR, or GLFW_RELEASE_BEHAVIOR_FLUSHGLFW_RELEASE_BEHAVIOR_NONE|
|GLFW_OPENGL_FORWARD_COMPAT|	GLFW_FALSE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_CONTEXT_DEBUG|	GLFW_FALSE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_OPENGL_PROFILE（设置渲染模式）|	GLFW_OPENGL_ANY_PROFILE|	GLFW_OPENGL_ANY_PROFILE, or GLFW_OPENGL_COMPAT_PROFILE, GLFW_OPENGL_CORE_PROFILE(核心模式)|
|GLFW_WIN32_KEYBOARD_MENU|	GLFW_FALSE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_WIN32_SHOWDEFAULT|	GLFW_FALSE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_COCOA_FRAME_NAME|	""	|A UTF-8 encoded frame autosave name|
|GLFW_COCOA_GRAPHICS_SWITCHING|	GLFW_FALSE|	GLFW_TRUE or GLFW_FALSE|
|GLFW_WAYLAND_APP_ID|	""|	An ASCII encoded Wayland name app_id|
|GLFW_X11_CLASS_NAME|	""|	An ASCII encoded class name WM_CLASS|
|GLFW_X11_INSTANCE_NAME|	""|	An ASCII encoded instance name WM_CLASS|

# 3 glfwCreateWindow
**作用**: 创建窗口
**原型**：GLFWwindow* glfwCreateWindow(int width, int height, const char* title, GLFWmonitor* monitor, GLFWwindow* share)
**参数**：
  - width： 窗口的宽度，单位为像素
  - height: 窗口的高度，单位像素
  - title： 窗口的标题
  - monitor：显示器，指定在哪个监视器上创建窗口。如果传入的是 NULL，则会创建一个窗口模式的窗口（即窗口可以在桌面上拖动的普通窗口）。
  - share：用于指定一个已经存在的窗口（GLFWwindow*），与新创建的窗口共享 OpenGL 上下文资源，如纹理、着色器等。如果传入 NULL，新创建的窗口将不与任何现有窗口共享资源。

# 4 glfwWindowShouldClose
**作用**：判断窗口是否应该关闭
**原型**：int glfwWindowShouldClose(GLFWwindow* window)
**参数**：
  - window：创建的窗口
**返回值**：是否应该关闭窗口，0:否 1:是

# 5 glfwMakeContextCurrent
**作用**：设置当前线程的opengl上下文，设置了上下文到当前线程后，在该线程下调用opengl api才会生效
**原型**：void glfwMakeContextCurrent(GLFWwindow* window)
**参数**：
  - window：将指定窗口的opengl上下文设置到当前线程

# 5 glfwSetWindowShouldClose
**作用**：设置窗口是否应该关闭
**原型**：void glfwSetWindowShouldClose(GLFWwindow* window, int value)
**参数**：
  - window：创建的窗口
  - value: 设置窗口是否应该关闭，GLFW_FALSE:不应该关闭， GLFW_TRUE:应该关闭

# 6 glfwSwapBuffers
**作用**：交换绘制缓冲区，在绘制完一帧后，交换绘制缓冲区，用于将绘制的图形显示在显示屏上
**原型**：void glfwSwapBuffers(GLFWwindow* window)
**参数**：
  - window：创建的窗口

# 7 glfwPollEvents
**作用**：在应用程序的主循环中定期调用，以处理事件队列中的所有待处理事件，使得应用程序能够对用户输入（例如键盘、鼠标等设备的输入事件）和窗口事件做出响应。
**原型**：void glfwPollEvents(void)

# 8 glfwGetKey
**作用**：获取键盘事件
**原型**：int glfwGetKey(GLFWwindow* window, int key)
**参数**：
  - window: 创建的窗口
  - key：检测key所对应的键盘按键是否有事件发生，例如GLFW_KEY_SPACE表示空格键
**返回值**: 返回对应的键盘按键的状态，例如GLFW_PRESS表示键盘或鼠标按下，GLFW_RELEASE：键盘或鼠标释放
**示例**：
```c++
if (glfwGetKey(window, GLFW_KEY_SPACE) == GLFW_PRESS) { // 监听空格按键
    ...
}
```

# 9 glfwTerminate
**作用**：销毁opengl上下文的数据，包含创建的opengl各个对象等
**原型**： void glfwTerminate(void)
