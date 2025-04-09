# 1 读取xml文件
## 1.1 从文件中读取xml
函数原型
```c++
XMLError XMLDocument::LoadFile( const char* filename);
```
**参数**:
- **<font color=green>filename</font>**：xml文件路径

**返回值**：
一个错误类型的枚举，具体值：
```c++
enum XMLError {
    XML_SUCCESS = 0,
    XML_NO_ATTRIBUTE,
    XML_WRONG_ATTRIBUTE_TYPE,
    XML_ERROR_FILE_NOT_FOUND,
    XML_ERROR_FILE_COULD_NOT_BE_OPENED,
    XML_ERROR_FILE_READ_ERROR,
    XML_ERROR_PARSING_ELEMENT,
    XML_ERROR_PARSING_ATTRIBUTE,
    XML_ERROR_PARSING_TEXT,
    XML_ERROR_PARSING_CDATA,
    XML_ERROR_PARSING_COMMENT,
    XML_ERROR_PARSING_DECLARATION,
    XML_ERROR_PARSING_UNKNOWN,
    XML_ERROR_EMPTY_DOCUMENT,
    XML_ERROR_MISMATCHED_ELEMENT,
    XML_ERROR_PARSING,
    XML_CAN_NOT_CONVERT_TEXT,
    XML_NO_TEXT_NODE,
    XML_ELEMENT_DEPTH_EXCEEDED,
    XML_ERROR_COUNT
};
```
## 1.2 获取根节点
函数原型
```c++
XMLElement* XMLDocument::RootElement();
```
**说明：**
对于xml
```xml
<?xml version="1.0" encoding="UTF-8"?>
<Root>
    <Child1></Child1>
</Root>
```
RootElement()获取的是Root这个节点。
## 1.3 获取节点的名称
```c++
const char* XMLElement::Value();

// 对Value()的封装
const char* XMLElement::Name();
```
## 1.4 获取节点属性
```c++
// 获取节点的第一个属性
XMLAttribute* XMLElement::FirstAttribute();

/*获取下一个属性，若没有则返回NULL*/ 
XMLAttribute* XMLAttribute::Next();
```
## 1.5 获取节点的文本信息
```c++
char* XMLElement::GetText();
```
说明：
```xml
<?xml version="1.0" encoding="UTF-8"?>
<Root>hello world</Root>
```
获取结果为：hello world


## 1.6 获取节点的子节点
```c++
// 获取节点的第一个子节点
XMLElement* XMLElement::FirstChildElement( const char* name = 0 )

/*获取当前节点的兄弟节点*/
XMLElement*	XMLElement::NextSiblingElement( const char* name = 0 )
```
代码示例：
```c++
// 获取current节点的第一个子节点
XMLElement* element = current->FirstChildElement();

// 遍历current节点的所有子节点
for (; element != NULL; element = element->NextSiblingElement()) {
    ...
}
```

# 2 创建XML文件

```c++
tinyxml2::XMLDocument doc;
 
//1.添加声明
tinyxml2::XMLDeclaration* declaration = doc.NewDeclaration();
doc.InsertFirstChild(declaration);
 
//2.创建根节点
tinyxml2::XMLElement* root = doc.NewElement("school");
doc.InsertEndChild(root);
 
//3.创建子节点
tinyxml2::XMLElement* childNodeStu = doc.NewElement(“student”);
tinyxml2::XMLElement* childNodeTea = doc.NewElement(“teacher”);
tinyxml2::XMLElement* childNodeTeaGender = doc.NewElement(“gender”);
 
//4.为子节点增加文本内容
tinyxml2::XMLText* contentStu = doc.NewText(“stu”);
childNodeStu->InsertFirstChild(contentStu);
 
tinyxml2::XMLText* contentGender = doc.NewText(“man”);
childNodeTeaGender->InsertFirstChild(contentGender);
 
//5.为子节点增加属性
childNodeStu->SetAttribute("Name", "libai");
 
// 将子节点添加到父节点中作为父节点的子节点
root->InsertEndChild(childNodeStu);//childNodeStu是root的子节点
root->InsertEndChild(childNodeTea);//childNodeTea是root的子节点
childNodeTea->InsertEndChild(childNodeTeaGender);//childNodeTeaGender是childNodeTea的子节点
 
//6.保存xml文件
doc.SaveFile(“school.xml”);
```