[toc]
# 1 获取Node节点
```c++
// 获取根节点
using namespace kanzi;
kanzi::Node2DSharedPtr root = getRoot(); 
// 获取Node2D
kanzi::Node2DSharedPtr node = root->lookup<kanzi::Node2D>("#nodeName")
// 获取EmptyNode2D
kanzi::EmptyNode2DPtr mCarModle = root->lookup<kanzi::EmptyNode2D>("#nodeName")

/*如果想要获取EmptyNode2D、EmptyNode3D、Node3D，将模板参数
和返回值改为对应的值即可。等号的左边也可以改为其父类的指针*/
```

# 2 获取Node的属性值
想要获取其他的类型的属性值就将模板参数改成对应的类型
```c++
// mCarModle是Node指针
std::string name = mCarModle->getProperty(kanzi::DynamicPropertyType<std::string>("name"));
printf("get car modle name propert=%s\n", name.c_str());
```

# 3 设置Node节点的属性
通过设置Node节点的属性值可以改变Node对应的UI行为、状态
```c++
// 获取域指针
Domain* domain = getDomain();
kanzi::TaskDispatcher* mDispatcher = domain->getTaskDispatcher()

// 设置3D/2D Node的属性设置的回调函数 //
template <typename NodeTypePtr>
static void setPropertytaskInt(NodeTypePtr node, std::string property, int data){
    LOGI("[%s] set |%s| property %s:%d",__FUNCTION__,node->getName().c_str(),property.c_str(),data);
    node->setProperty(kanzi::DynamicPropertyType<int>(property.c_str()), data);
}

// 设置node节点中的attr属性为value
void UINodeManager::setProperty(kanzi::Node* node, std::string attr, int value){
    if(node != NULL){
        mDispatcher->submit(bind(setPropertytaskInt<kanzi::Node*>, node, attr, value));
        return;
    }
    LOGE("[UINodeManager::%s<bool>] node is NULL",__FUNCTION__);
}
```
如果是其他类型的属性如bool,float等，只需把value和data变量的类型修改。像kanzi::Vector2，kanzi::Vector3
```c++
template <typename Node2DPtr>
static void setPropertytaskVector2( Node2DPtr node,
                                    std::string property,
                                    int dataX, int dataY)
{
    LOGI("[%s] set |%s| property %s:(%d,%d)",__FUNCTION__,node->getName().c_str(),property.c_str(),dataX,dataY);
    node->setProperty(kanzi::DynamicPropertyType<kanzi::Vector2>(property.c_str()),
                                                        kanzi::Vector2(dataX, dataY));
}

// 设置kanzi::Vector2类型的属性，kanzi::Vector3类型的类似
void UINodeManager::setProperty(std::string nodeName, std::string attr, int dataX, int dataY){
    kanzi::Node2D* node2d = get2DNodeByName(nodeName);
    if(node2d != NULL){
        mDispatcher->submit(bind(setPropertytaskVector2<kanzi::Node2D*>, node2d, attr, dataX, dataY));
        return;
    }
    LOGE("[%s<int,int>] not found node named %s, it's maybe a 3D resource",__FUNCTION__,nodeName.c_str());
}
```
# 4 监听属性值的变化
```c++
#include <kanzi/core/property/property_type_descriptor.hpp>

// 获取一个node节点
kanzi::EmptyNode2DPtr mCarModle = root->lookup<kanzi::EmptyNode2D>("#nodeName");

// bool类型属性变化回调
void NotificationCallback(PropertyObject& prop, const typename PropertyTypeDescriptor<bool>::Traits::StorageType& value, PropertyNotificationReason reason, void* target){
    char* attrInfo = (char*)target;
    bool visable = prop.getProperty(kanzi::DynamicPropertyType<bool>(NODE_VISIBLE));
    LOGI("target=%s, Visible change to %d, storage=%d",attrInfo, visable, value);
}

// 设置node节点监听的Node.Visible属性,最后一个参数是自定义类型，可以用于传递额外信息，如表示是哪个节点的哪个属性
// 挡对应的属性值变化的时候就会回调NotificationCallback
mCarModle->addPropertyNotificationHandler(kanzi::DynamicPropertyType<bool>("Node.Visible"), NotificationCallback, (void*)"2D Page");
```

# 5 监听Node点击事件
**添加Node监听回调函数流程**
```c++
ScreenSharedPtr screen = getScreen();
Domain* domain = getDomain();

//使用别名获取ClickNode 节点。
NodeSharedPtr clickNode = screen->lookupNode<Node>("#ClickNode");

//创建生成点击消息的输入操纵器。
ClickManipulatorSharedPtr clickManipulator = ClickManipulator::create(domain);

//添加输入操纵器到 ClickNode 节点。
clickNode->addInputManipulator(clickManipulator);

//订阅 ClickNode 节点的 ClickManipulator::ClickMessage 消息。
// ClickManipulator 操纵器在用户点击节点时生成此消息。
clickNode->addMessageHandler(ClickManipulator::ClickMessage, bind(&MyProject::onNodeClicked, this, boost::placeholders::_1));

//订阅 ClickNode 节点的 ClickManipulator::ClickCancelMessage 消息。
// ClickManipulator 操纵器在用户第一次按下
//该节点，然后将指针移动到节点区域外并提起指针时，生成此消息。
clickNode->addMessageHandler(ClickManipulator::ClickCancelMessage, bind(&MyProject::onNodeClickCanceled, this, placeholders::_1));
```

```c++
/*回调函数定义*/
//为 ClickManipulator::ClickMessage 消息定义处理程序，该消息来自
//具有可生成 ClickMessage 消息的输入操纵器的节点。
void onNodeClicked(ClickManipulator::ClickMessageArguments& messageArguments)
{
    //添加处理点击事件的代码。
    // 获取点击的Node
    shared_ptr<Node> node = messageArguments.getSource();

    // 获取点击node的名称(非别名的最后一个名字，如/RootPage/car/aa ==> aa)
    LOGD("click node %s",node->getName().c_str());
}

//为 ClickManipulator::ClickCancelMessage 消息定义处理程序，该消息来自
//具有可生成 ClickCancelMessage 消息的输入操纵器的节点。
void onNodeClickCanceled(ClickManipulator::ClickCancelMessageArguments& messageArguments)
{
    //添加处理点击取消事件的代码。
    // 获取点击的Node
    shared_ptr<Node> node = messageArguments.getSource();

    // 获取点击node的名称(非别名的最后一个名字，如/RootPage/car/aa ==> aa)
    LOGD("click node %s",node->getName().c_str());
}
```

**点击事件消息**
除了ClickMessage和ClickCancelMessage事件，还有可以监听以下点击相关事件，
再设置的时候，只需要替换成对应的类型就可以了
```c++
/// Message type for notifying recognition of click gesture.
/// Click message is generated when the tracked touch is released over the attached node or when the fail manipulator fails (see InputManipulator::requireToFail).
static MessageType<ClickMessageArguments> ClickMessage;
/// Message type for notifying the beginning of a click gesture.
/// There are two scenarios for ClickBeginMessage to be generated:
/// 1) When touch is pressed over the attached node. This is the default behavior. 
/// 2) If the gesture is set to start on hover with the setHoverToBegin function, click gesture begins when the touch arrives in 
/// the node area while being pressed.
static MessageType<ClickBeginMessageArguments> ClickBeginMessage;
/// Message type for notifying cancellation of click gesture.
/// Click gesture is canceled when the tracked touch is released outside the node area, or when the fail manipulator succeeds (see InputManipulator::requireToFail).
static MessageType<ClickCancelMessageArguments> ClickCancelMessage;
/// Message type for notifying the tracked touch entering the node area.
/// Click enter message is generated when touch is pressed over the node area or when the touch returns to the node area after leaving it.
static MessageType<ClickEnterMessageArguments> ClickEnterMessage;
/// Message type for notifying the tracked touch leaving the node area.
/// Click leave message is generated when the pointer leaves the area or when the click gesture is completed or canceled.
static MessageType<ClickLeaveMessageArguments> ClickLeaveMessage;
```
