[toc]

# 1 glClearColor
**作用**：设置用于刷新像素的颜色，在绘制缓冲区的数据在界面展示完成之后，需要把绘制缓冲区的数据清除，这个函数用于指定使用什么颜色来刷新绘制缓冲区的颜色。
最终的原始是三原色的混合颜色
**原型**：void glClearColor(GLfloat red, GLfloat green, GLfloat blue, GLfloat alpha)
**参数**：
  - red：红色分量的值，范围[0, 1]，值越大颜色越红
  - green：绿色分量的值，范围[0, 1]，值越大颜色越绿
  - blue：蓝色分量的值，范围[0, 1]，值越大颜色越蓝
  - alpha：透明度分量的值，范围[0, 1]， 0的完全透明，1是完全不透明

# 2 glClear
**作用**：清除指定缓冲区的数据
**原型**：void glClear(GLbitfield mask)
**参数**：
  - mask：指定清除的缓冲区，
    GL_COLOR_BUFFER_BIT：清除颜色缓冲区，会将颜色缓冲区中的每个像素都设置为通过 glClearColor 指定的颜色。通常用于清屏，将屏幕背景设置为某种颜色。
    GL_DEPTH_BUFFER_BIT：清除深度缓冲区，glClear 会将深度缓冲区中的每个像素都设置为通过 glClearDepth 指定的深度值。通常用于在开始新的渲染帧时重置深度信息。
    GL_STENCIL_BUFFER_BIT：清除模板缓冲区,glClear 会将模板缓冲区中的每个像素都设置为通过 glClearStencil 指定的值。用于重置模板测试相关的信息。
    这三个缓冲区类型可以通过或符号来运算，用于表示同时清除指定的缓冲区
**示例**：
```c++
glClear(GL_COLOR_BUFFER_BIT|GL_DEPTH_BUFFER_BIT|GL_STENCIL_BUFFER_BIT);
```

# 3 glClearDepth
**作用**：用于指定在使用 glClear 函数清除深度缓冲区时，深度缓冲区将被设置为的指定的深度值。
**原型**：void glClearDepth(GLdouble depth)
**参数**：
  - depth: 这是一个指定深度缓冲区清除值的参数。它是一个双精度浮点数，通常范围在 [0.0, 1.0] 之间。0.0 表示最接近观察者的位置，而 1.0 表示最远的位置。
**使用场景**：在3D图形渲染中，每一帧开始时通常需要清除深度缓冲区，以确保不会有前一帧的深度信息残留，干扰当前帧的渲染。

# 4 glGenBuffers
**作用**：生成一个或多个缓冲区对象，这些缓冲区对象通常用于存储顶点数据、索引数据或其他与渲染相关的数据。
虽然 glGenBuffers 生成了缓冲区对象的名称，但它们尚未与任何数据或具体的缓冲区类型关联。要实际使用这些缓冲区对象，通常需要通过 glBindBuffer 绑定到特定的缓冲区目标（如 GL_ARRAY_BUFFER 或 GL_ELEMENT_ARRAY_BUFFER），然后再使用其他函数（如 glBufferData）来填充数据。
**原型**：void glGenBuffers(GLsizei n, GLuint *buffers)
**参数**：
  - n: 需要生成的缓冲区对象的个数
  - buffers：指向一个数组的指针，该数组将存储生成的缓冲区对象的标识符。

# 5 glGenVertexArrays
**作用**：生成一个或多个顶点数组对象（VAO）。VAO 是一个容器对象，它存储与顶点属性相关的状态（如顶点缓冲对象、顶点属性指针等）。通过使用 VAO，你可以将所有与顶点数组相关的状态绑定到一个对象上，从而简化后续的渲染过程。
**原型**：void glGenVertexArrays(GLsizei n, GLuint *arrays)
**参数**：
  - **n**: 指定要生成的 VAO 的数量。
  - **arrays**: 一个 GLuint 类型的数组，用于存储生成的 VAO 的 ID。
**示例**：
```c++
unsigned int VBO, VAO;
    glGenVertexArrays(1, &VAO); // 生成顶点数组对象
    glGenBuffers(1, &VBO); // 生成顶点缓冲对象
    glBindVertexArray(VAO); // 绑定生成的顶点数组对象, 将其设置为当前的VAO

    glBindBuffer(GL_ARRAY_BUFFER, VBO); // 绑定顶点缓冲对象到顶点缓冲区
    glBufferData(GL_ARRAY_BUFFER, sizeof(vertices), vertices, GL_STATIC_DRAW); // 向顶点缓冲区所绑定的顶点缓冲对象中添加数据

    glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 3 * sizeof(float), (void*)0); // 
    glEnableVertexAttribArray(0);
```

# 6 glBindVertexArray
**作用**：用于绑定一个顶点数组对象（VAO）。通过绑定 VAO，OpenGL 将其设为当前的 VAO，并且所有与顶点数组相关的操作（如设置顶点属性、绑定顶点缓冲对象等）都会影响到这个绑定的 VAO。
**原型**：void glBindVertexArray(GLuint array)
**参数**：
  - **array**: 需要绑定的 VAO 的 ID。如果传入 0，则表示解绑当前的 VAO，这意味着后续的顶点数组操作不再影响任何 VAO。

# 6 glBindBuffer
**作用**：将缓冲区对象绑定到指定的缓冲区目标
**原型**：void glBindBuffer(GLenum target, GLuint buffer)
**参数**：
  - target：缓冲区目标
    - **GL_ARRAY_BUFFER**：顶点缓冲区
    - **GL_ELEMENT_ARRAY_BUFFER**：
  - buffer：创建的缓冲区对象的句柄, 其值为`glGenBuffers`方法创建的缓冲区对象。如果为0，表示解绑指定的缓冲区目标
**示例**：
```c++
GLuint VBO;
glGenBuffers(1, &VBO); // 创建一个顶点缓冲对象
glBindBuffer(GL_ARRAY_BUFFER, VBO); // 将顶点缓冲对象绑定到GL_ARRAY_BUFFER目标上。从这一刻起，我们使用的任何（在GL_ARRAY_BUFFER目标上的）缓冲调用都会用来配置当前绑定的缓冲(VBO)
```

**拓展**：
- VAO 和 VBO 的关系
绑定关系：VAO 本身不存储顶点数据，而是存储与顶点数据如何组织和解释相关的状态信息。VBO 存储实际的顶点数据。当一个 VBO 被绑定到 VAO 时，VAO 会记录这个 VBO 以及与其相关的顶点属性设置。

状态保存：通过 VAO，OpenGL 可以记住哪些 VBO 被绑定，以及如何解释这些 VBO 中的数据。这意味着当你绑定一个 VAO 时，OpenGL 会自动设置相关的 VBO 和顶点属性指针。

简化流程：使用 VAO 可以大大简化渲染代码。每次绘制物体时，只需要绑定相应的 VAO，OpenGL 就会自动配置所有与顶点相关的状态，而不需要每次重新绑定 VBO 并设置顶点属性指针。

# 7 glBufferData
**作用**：向指定的缓冲区目标所绑定的缓冲区对象中分配内存并填充数据。这个函数通常用于上传顶点数据、索引数据等到 GPU 以供后续渲染使用。
**原型**：void glBufferData(GLenum target, GLsizeiptr size, const void *data, GLenum usage)
**参数**：
  - target：指定目标缓冲区对象的类型。
    - **GL_ARRAY_BUFFER**：表示顶点缓冲区对象，用于存储顶点属性数据。
    - **GL_ELEMENT_ARRAY_BUFFER**：表示元素缓冲区对象，用于存储顶点索引数据。
  - size：填充数据的大小，单位：字节
  - data： 填充的数据
  - usage：指定数据的使用模式，告诉 OpenGL 如何使用这段数据。
    - **GL_STATIC_DRAW**：表示数据不会经常改变，主要用于从应用程序到 GPU 的一次性传输（静态数据）。
    - **GL_DYNAMIC_DRAW**：表示数据会经常改变，主要用于频繁从应用程序到 GPU 的传输（动态数据）。
    - **GL_STREAM_DRAW**：表示数据几乎每次绘制时都会改变，主要用于从应用程序到 GPU 的频繁更新（流数据）。

    如果，比如说一个缓冲中的数据将频繁被改变，那么使用的类型就是GL_DYNAMIC_DRAW或GL_STREAM_DRAW，这样就能确保显卡把数据放在能够高速写入的内存部分。

# 8 glVertexAttribPointer
**作用**: 指定顶点属性数组的信息，并将其与当前绑定的顶点缓冲对象（VBO）关联起来。该函数是顶点属性配置的重要步骤，它告诉 OpenGL 如何解释存储在 VBO 中的顶点数据。
**原型**: void glVertexAttribPointer(GLuint index, GLint size, GLenum type, GLboolean normalized, GLsizei stride, const void *pointer)
**参数**: 
  - **index**: 指定顶点属性的索引（即着色器中的位置）
  - **size**: 指定顶点属性的分量个数。有效值为 1 到 4 之间。例如，如果每个顶点的颜色由 R, G, B 三个分量组成，那么 size 就是 3。
  - **type**: 指定数组中每个分量的数据类型。
    - GL_FLOAT: 每个分量是一个 GLfloat 类型（浮点数）。
    - GL_INT:  每个分量是一个 GLint 类型（整数）。
    - GL_UNSIGNED_BYTE: 每个分量是一个 GLubyte 类型（无符号字节）等。
  - **normalized**: 指定当数据类型为整型时，是否将其映射到 [0, 1]（无符号类型）或 [-1, 1]（有符号类型）的浮点范围。如果设置为 GL_TRUE，数据会被归一化；如果为 GL_FALSE，数据将保持原样。
  - **stride**: 指定连续顶点属性之间的字节偏移量。简单来说，就是每个顶点的总字节数（包括所有属性）。如果属性是紧密排列的，可以设置为 0，OpenGL 会自动计算正确的步幅。
  - **pointer**: 指定数组中第一个顶点属性的偏移量（通常是相对于数组起始位置的字节偏移）。对于存储在 VBO 中的顶点数据，这个值通常是 0 或者一个字节偏移量。
**示例**：
```c++
glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, 3 * sizeof(float), (void*)0);


static const GLfloat vertexTexture[]{
    0.0f, 0.0f,
    0.0f, 1.0f,
    1.0f, 0.0f,
    1.0f, 1.0f, 
};
glVertexAttribPointer(textureIndex, 2, GL_FLOAT, GL_FALSE, 2*sizeof(GL_FLOAT), vertexTexture);
```


# 8 glCreateShader
**作用**：创建着色器
**原型**： GLuint glCreateShader(GLenum type)
**参数**：
  - type: 指定创建的着色器类型
    - **GL_VERTEX_SHADER**: 创建顶点着色器
    - **GL_FRAGMENT_SHADER**：创建片段着色器
**返回值**：返回创建的shader的句柄

# 9 glShaderSource
**作用**: 向着色器代码设置到创建的shader中
**原型**: void glShaderSource(GLuint shader, GLsizei count, const GLchar *const*string, const GLint *length)
**参数**：
  - **shader**: glCreateShader方法创建着色器程序
  - **count**: 着色器代码的字符串数量。也就是说，你可以通过这个参数指定有多少个字符串片段需要传递给着色器
  - **string**: 着色器源代码
  - **length**: 一个指向长度数组的指针，每个元素对应 string 数组中每个字符串的长度。如果你传递 NULL，表示每个字符串都是以 '\0' 结尾的，函数会自动计算每个字符串的长度。

**示例1**：
```c++
const char* vertexShaderSource = "your GLSL vertex shader code here";
GLuint vertexShader = glCreateShader(GL_VERTEX_SHADER);
glShaderSource(vertexShader, 1, &vertexShaderSource, NULL);
```

**示例2**：
```c++
const char* shaderPart1 = "void main() { gl_Position = vec4(0.0);";
const char* shaderPart2 = " }";

// 每个片段的长度
GLint lengths[] = { strlen(shaderPart1), strlen(shaderPart2) };

// 创建并编译着色器
GLuint vertexShader = glCreateShader(GL_VERTEX_SHADER);
const char* shaderSources[] = { shaderPart1, shaderPart2 };

// 源代码由2个字符串组成，这两个字符串的长度信息存储在lengths数组中，第一个元素是一个字符串的长度，第二个元素是第二个字符串的长度
glShaderSource(vertexShader, 2, shaderSources, lengths);
```

# 11 glCompileShader
**作用**： 编译着色器
**原型**: void glCompileShader(GLuint shader)
**参数**：
  - **shader**: glCreateShader方法创建的shader对象的句柄

# 12 glGetShaderiv
**作用**: 用于查询着色器对象状态的一个函数。它可以获取与指定着色器对象相关的各种信息，如编译状态、着色器源码的长度等
**原型**: void glGetShaderiv(GLuint shader, GLenum pname, GLint *params)
**参数**：
  - **shader**: shader对象
  - **pname**：指定要查询的着色器属性
    - GL_SHADER_TYPE： 获取着色器类型（param的值会设置为 GL_VERTEX_SHADER 或 GL_FRAGMENT_SHADER
    - GL_COMPILE_STATUS: 获取着色器的编译状态（param的值会设置为 GL_TRUE 或 GL_FALSE）
    - GL_INFO_LOG_LENGTH: 获取编译日志的长度。
    - GL_SHADER_SOURCE_LENGTH： 获取着色器源码的长度。
  - **params**: 传出参数，用于获取查询的shader属性信息
**示例**：
```c++
unsigned int vertexShader = glCreateShader(GL_VERTEX_SHADER);
glShaderSource(vertexShader, 1, &vertexShaderSource, NULL);
glCompileShader(vertexShader);
// check for shader compile errors
GLint success;
char infoLog[512];
glGetShaderiv(vertexShader, GL_COMPILE_STATUS, &success);
```

# 13 glGetShaderInfoLog
**作用**： 获取着色器编译日志
**原型**：void glGetShaderInfoLog(GLuint shader, GLsizei bufSize, GLsizei *length, GLchar *infoLog)
**参数**：
  - **shader**: 着色器句柄
  - **bufSize**: 指定infoLog缓存的最大大小（读取日志的最大长度）
  - **length**: 用于存储实际写入 infoLog 的字符串长度（不包括空终止符）。如果不需要该长度信息，可以传递 NULL
  - **infoLog**: 指向一个字符数组的指针，该数组用于存储函数返回的信息日志。如果编译或链接过程中没有生成信息日志，返回的 infoLog 将是一个空字符串。
**示例**:
```c++
GLint success;
glGetShaderiv(vertexShader, GL_COMPILE_STATUS, &success);
if (!success)
{
    glGetShaderInfoLog(vertexShader, 512, NULL, infoLog);
    std::cout << "ERROR::SHADER::VERTEX::COMPILATION_FAILED\n" << infoLog << std::endl;
}
```

# 14 glCreateProgram
**作用**：创建着色器程序对象
**原型**：GLuint glCreateProgram()

# 15 glAttachShader
**作用**：将着色器对象添加到程序对象中
**原型**：void glAttachShader(GLuint program, GLuint shader)
**参数**：
  - **program**: 程序对象
  - **shader**: 需要被添加到程序上的着色器对象

# 16 glLinkProgram
**作用**: 链接程序，将添加到程序上的顶点和片段着色器一起组成一个渲染管线程序
**原型**： void glLinkProgram(GLuint program)
**参数**：
  - **program**: 需要被链接的程序对象

# 17 glGetProgramiv
**作用**: 获取程序对象的相关信息
**原型**： void glGetProgramiv(GLuint program, GLenum pname, GLint *params)
**参数**：
  - **program**： 需要检查的程序对象
  - **pname**: 指明需要检查的信息类型
    - GL_LINK_STATUS： 查询程序是否成功链接， params为1表示成功，0表示失败
    - GL_VALIDATE_STATUS： 查询程序是否经过验证
    - GL_INFO_LOG_LENGTH：查询链接或验证操作的日志长度
    - GL_ATTACHED_SHADERS：查询附加到该程序的着色器对象的数量
    - GL_ACTIVE_UNIFORMS： 查询活动的 uniform 变量的数量。
    - GL_ACTIVE_ATTRIBUTES： 查询活动的属性变量的数量。 
  - **params**： 传出参数，用于存储pname参数指定的消息类型的查询结果

# 18 glGetProgramInfoLog
**作用**： 查询链接程序中的日志
**原型**：void glGetProgramInfoLog(GLuint program, GLsizei bufSize, GLsizei *length, GLchar *infoLog)
**参数**：
  - **program**： 需要查看日志的程序对象
  - **bufSize**: 指定查询日志内容的最大大小
  - **length**: 传出参数，函数将返回实际写入 infoLog 的字符数。如果你不关心这个值，可以传递 NULL。
  - **infoLog**: 用于存储获取到的日志信息。你需要为这个指针分配足够的空间来容纳日志内容。
**示例**
```c++
char infoLog[512];
glGetProgramInfoLog(shaderProgram, 512, NULL, infoLog);
```

# 19 glDeleteShader
**参数**： 删除着色器对象
**原型**： void glDeleteShader(GLuint shader)
**参数**： 
  - **shader**: 需要删除的着色器对象

# 20 glDeleteProgram
**参数**： 删除程序对象
**原型**： void glDeleteProgram(GLuint program)
**参数**： 
  - **program**：需要删除的程序对象

# 21 glDrawElements
**作用**：绘制图元
**原型**：void glDrawArrays(GLenum mode, GLint first, GLsizei count)
**参数**：
  - **mode**：根据当前绑定的顶点数据绘制图元（如点、线、三角形等）
    - GL_TRIANGLES: 绘制三角形图元
    - GL_POINTS：绘制点
    - GL_LINES: 绘制独立的线段
    - GL_LINE_STRIP: 绘制连续的线段（连接每对相邻的顶点）。
    - GL_LINE_LOOP: 绘制一个线环（类似于 GL_LINE_STRIP，但最后一个顶点与第一个顶点相连）。
    - GL_TRIANGLE_STRIP: 绘制连续的三角形条带（每个新顶点与前两个顶点组成一个三角形）。例如对于顶点数组[a,b,c,d], 获取2个三角形(a,b,c)和(b,c,d)
    - GL_TRIANGLE_FAN: 绘制一个三角形扇形。
  - **first**: 指定从顶点数组的第几个顶点开始绘制。这个索引是基于当前绑定的 VBO 和设置的顶点属性指针来解析的。
  - **count**: 指定要绘制的顶点数量。
