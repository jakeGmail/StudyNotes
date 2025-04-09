[toc]

# 1 排版概念
## 1.1 字体文件
首先，在<font color=red>FreeType</font>中的基本字体单位是Face，例如simkai.tff通过加载，变为一个<font color=red>FT_Face</font>句柄。而广义上的字体可能是多个Face的组合，例如"Palatino"字体，包含"Palatino常规”.“Palatino斜体"两个不同的Face，它们是分离的文件。所以我们约定术语<font color=red>字体(font)</font>为单个Face，而一套字体包含多个文件我们称为<font color=red>字体集合(font collection)</font>。

## 1.2 字符图像和字符表
字符的图像称为<font color=red>字形（glyphs）</font>，而一个字符可以有多个字形。而一个字形，也可用于多个字符（不同的字符，写法可能一样）。

我们可以只关注两个概念：
- 一个字体文件包含多个字形，每一个都可以储存为位图、矢量、或其他任何方案。而我们通过字形索引访问。

- 字体文件包含一到多个表，称为字符表（character maps）。它可以将字符编码（ASCII、Unicode、GB2312、BIG5等）转为字形索引。（通过字形索引就能获取到字形，便可获得字符的图像）

## 1.3 字符和字体指标
每个字形图像都含有各种<font color=red>指标（Metrics ）</font>，这些指标描述了在呈现文本时如何放置和管理。指标包含<font color=red>字形位置、光标前进、文本布局</font>。

可拓展格式还包含全局指标，以字体单位，描述同一Face所有字形的属性。例如最大字形边框、字体的上升、下降和文本高度。

不可拓展的格式，也包含一些指标。仅适用于一组给定字符的尺寸、分辨率。一般以像素为单位。

## 1.4 字形轮廓
Freetype不是通过像素来存储字形，而是通过字符的形状，我们称之为轮廓（outlines）。它使用点为单位，以下公式计算转换到像素单位：

pixel_size = point_size * resolution / 72

其中resolution为分辨率，以dpi（每英寸点数）为单位。

存储在文件内的数据称为主轮廓，以点作为单位。在转换为位图时，需要进行缩放，这个步骤需要进行<font color=red>网格拟合（grid-fitting）</font>得到图像，它有几种不同的算法，不过我们简单理解一下概念即可

## 1.4 字形指标
### 1.4.1 基线、笔和布局
<font color=red>基线（baseline）</font>是一条假象的线，比如作业本上的横线，使我们方便对齐位置。它可以是横的，也可以是竖的。而笔尖位于基线上的一点，用于定位字形。

水平布局：字形在基线之上（有可能超过基线下方，比如字母q），通过向左或向右增加笔的位置来定位字形。

两个连续笔尖位置（下图线上的小黑方块）的距离与字形有关，称为<font color=red>步进宽度（advance width）</font>。它始终是正数，即使阿拉伯语是从右往左写的（我们排版的时候再进行处理）。

另外笔的位置始终在基线上。
![](img/freetype_1.png)

而垂直布局基线在字形的中央，如下图所示
![](img/freetype_2.png)

### 1.4.2 排版指标和边界框
为字体所有字形定义的各种Face指标：
- Ascent：从基线到最高轮廓点的距离。正值，因为Y轴向上
- Descent：从基线到最低轮廓点的距离。负值，但有些字体是正值
- Linegap(行距)：必须放置在两行文本之间的距离。
  两条基线之间的距离（标准行间距）：linespace = ascent - descent + linegap
- <font color=red>边界框（bounding box）</font>：由xMin、yMin、xMax、yMax表示的包围盒，能够包含所有字形。简写为“bbox”。
- Internal leading：用于传统排版，计算公式为：internal leading = ascent - descent - EM_size
- External leading：行距的别名

（注意这里的大小均是<font color=red>字点单位</font>，通过face可以访问，与我们设置的像素大小无关。要获得指定字体像素大小相关的数据，需要先调用<font color=red>FT_Set_Char_Size</font>，然后再通过ft_face->size->metrics获得<font color=red>FT_Size_Metrics</font>。）

### 1.4.3 方位与步进
每个字形有属于自己的<font color=red>方位（bearing）</font>与<font color=red>步进（advance）</font>。实际值和布局有关，水平和垂直布局是不同的值。
- 左侧方位：笔尖到字形左侧bbox的水平距离。通常水平布局才存在。在FreeType中叫<font color=red>bearingX</font>，简称“lsb”。
- 顶侧方位：基线到字形bbox顶部的垂直距离。通常水平布局为正，垂直布局为负。在FreeType中叫<font color=red>bearingY</font>。
- 步进宽度：渲染自身后，笔尖应该偏移的水平距离（从右向左则是减它）。垂直布局它始终为0。在FreeType中叫<font color=red>advanceX</font>。
- 步进高度：渲染自身后，笔尖应该减少的垂直距离（它为正值，因为Y轴向上，而写字是向下）。水平布局为0。在FreeType中叫<font color=red>advanceY</font>。
- 字形宽度：glyph width = bbox.xMax - bbox.xMin
- 字形高度：glyph height = bbox.yMax - bbox.yMin
- 右侧方位：步进到bbox右侧的距离，仅用于水平布局，一般为正值。缩写为“rsb”。

右侧方位：步进到bbox右侧的距离，仅用于水平布局，一般为正值。缩写为“rsb”。
![](img/freetype_3.png)
水平布局

![](img/freetype_4.png)
垂直布局

### 1.4.4 网格拟合的效果
网格拟合为了使字形的控制点与像素对齐，可能会修改调整字符图像的尺寸，从而影响字形指标。

### 1.4.5 文本宽度与边界框
字形的<font color=red>对齐（origin）点</font>即是笔尖在基线的位置。此对齐点通常不在字形的bbox上。而步进宽度与字形宽度也不是一回事。
对于整个字符串来说：
- 整个字符串包围盒不包含<font color=red>文本光标</font>，并且它也不会在角上。
- 字符串的步进宽度与包围盒无关。特别的是，前后存在空格、制表符。
- 类似字距调整等附加处理，会使整体尺寸与单个字形指标无关。

# 2 freetype安装
# Ubuntu安装
```shell
sudo apt install libfreetype6-dev
```
在编译的时候要链接库`freetype`, 添加寻找头文件路径`freetype2`

# 3 API
## 3.1 加载/释放字体资源
### 3.1.1 初始化freetype字体库
```c++
FT_EXPORT( FT_Error )
FT_Init_FreeType( FT_Library  *alibrary )
```
**参数**
- <font color=red>alibrary</font>: free type字体库指针。

**返回值**：
错误信息，如果不为0，则说明有错误发生

**使用示例**:
```c++
FT_Library mFtLibrary;
if(FT_Init_FreeType(&mFtLibrary)){
    LOGE("freetype init failed");
    mFtLibrary = nullptr;
}
```

**释放字体库资源API**:
```c
FT_EXPORT( FT_Error )
FT_Done_FreeType( FT_Library  library );
```


### 3.1.2 加载字体
```c++
FT_EXPORT( FT_Error )
FT_New_Face( FT_Library   library,
              const char*  filepathname,
              FT_Long      face_index,
              FT_Face     *aface );
```
**参数**:
- <font color=red>alibrary</font>:free type字体库指针。
- <font color=red>filepathname</font>： ttf文件路径
- <font color=red>face_index</font>：字体偏移量
- <font color=red>aface</font>： 传入参数，FT_Face指针，传入将其赋值。

**使用示例**:
```c++
FT_Face face;
if(FT_New_Face(mFtLibrary, ttf_path, 0, &face)){
    LOGE("font %s load failed", ttf_path);
    return false;
}
```

**释放字体资源API**：
```c
FT_EXPORT( FT_Error )
FT_Done_Face( FT_Face  face );
```

### 3.1.3 设置字体大小--接口1
```c++
FT_EXPORT( FT_Error )
FT_Set_Char_Size( FT_Face     face,
                  FT_F26Dot6  char_width,
                  FT_F26Dot6  char_height,
                  FT_UInt     horz_resolution,
                  FT_UInt     vert_resolution );
```

**参数**:
- <font color=red>face</font>: 加载的字体资源
- <font color=red>char_width</font>: 字符宽度，单位是1/64个点，如果传入参数的值单位是像素，需要进行 `char_width<<6` 操作
- <font color=red>char_height</font>: 字符高度，单位是1/64个点，如果传入参数的值单位是像素，需要进行 `char_width<<6` 操作, 如果是0，表示与char_width参数相同
- <font color=red>horz_resolution</font>: 水平方向上的屏幕dpi（1英寸下像素点个数）
- <font color=red>vert_resolution</font>: 垂直方向上的屏幕dpi（1英寸下像素点个数）

**使用示例**：
```c++
FT_Set_Char_Size(face, 36<<6, 36<<6, 96, 96);
```

### 3.1.4 设置字体大小--接口2
```c++
FT_EXPORT( FT_Error )
FT_Set_Pixel_Sizes( FT_Face  face,
                    FT_UInt  pixel_width,
                    FT_UInt  pixel_height );
```

**参数**:
- <font color=red>face</font>: 加载的字体资源
- <font color=red>pixel_width</font>: 字体宽度，单位像素
- <font color=red>pixel_height</font>: 字体高度，单位像素

**使用示例**：
```c++
FT_Set_Pixel_Sizes(face, 36, 36);
```





