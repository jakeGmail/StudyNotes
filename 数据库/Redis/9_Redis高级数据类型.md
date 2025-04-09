[toc]
# 1 Redis的高级数据类型
Redis的高级数据类型有3种
- Bitmaps
- HyperLogLog
- GEO
# 2 Bitmaps数据类型
Bitmaps数据类型是可以对数据的位进行操作的。通过一个数据的不同位的值来记录不同的数据信息。
例如记录一个学校的学生学号与性别的对应关系（假设学号从0开始，学校有1000个学生）。那么用Bitmaps数据类型的话可以设计为数据的第一个bit位代表学号为0的学生，而这个bit位的值1代表男生，0代表女生。像这样存储这个学校的学号与性别对应关系的数据就只需要1000bit大小的数据，大大节省了存储空间。
**Bitmaps的特点**：
1. 由于Bitmaps是通过一个比特位来存储信息的，因此比较适合具有2个状态的信息，例如性别、开关、真假等。
2. 由于信息是存储在bit位上，因此获取和修改相对比较麻烦，属于拿时间换空间。
3. Bitmaps数据类型其实是一个string

## 2.1 Bitmaps类型的基础操作
### 2.1.1 获取key对应偏移量上的bit值
<table><tr><td bgcolor="#87CEFA"></br>

```shell
# 获取key对应偏移量上的bit值
getbit key offset
```
</td></tr></table>

**示例**：</br>
<table><tr><td bgcolor=black>
<font color=white>
127.0.0.1:6379> set name jake</br>
OK</br>
127.0.0.1:6379> getbit name 0</br>
(integer) 0</br>
127.0.0.1:6379> getbit name 1</br>
(integer) 1</br>
</font>
</td></tr></table>

**注意**：
1. 如果在超出key的长度的offset取值，也会获取为0. 即在一个bit长度为10的key1中，即使执行```getbit key1 100```也会返回0.

### 2.1.2 设置指定key上对应偏移量的bit值
<table><tr><td bgcolor="#87CEFA"></br>

```shell
# 设置指定key上对应偏移量的bit值,value只能是0或1
setbit key offset value
```
</td></tr></table>

**示例**：
将bitInfo的第0位设置为1
<table><tr><td bgcolor=black>
<font color=white>
127.0.0.1:6379> setbit data 0 1</br>
(integer) 0</br>
127.0.0.1:6379> getbit data 0</br>
(integer) 1</br>
</font>
</td></tr></table>

**注意**:
1. 如果对于一个key的很大的offset上设置了值，也会设置成功，但offset之前的值会自动补充为0. 例如执行```set bits 99999 1```,会在bits的99999索引位设置位1，索引为99999之前的位自动设置位0. 
对此，如果我们需要设置的索引位很长的话，建议先将索引减去一个固定的值在添加到Bitmaps中。

## 2.2 Bitmaps扩展操作
### 2.2.1 或、非、与、异或操作
<table><tr><td bgcolor="#87CEFA"></br>

```shell
# 对key进行或、非、与、异或操作
# 将key1,key2...进行op操作后，将结果存入destkey中
# 其中<op>操作可以是：and、or、not、xor
bitop <op> destKey key1 key2 ...
```
</td></tr></table>

**示例**：
<table><tr><td bgcolor=Black>
<font color=white>
127.0.0.1:6379> set data1 jake</br>
OK</br>
127.0.0.1:6379> set data2 terry</br>
OK</br>
127.0.0.1:6379> bitop or destData data1 data2</br>
(integer) 5</br>
"~e{wy"</br>
</font>
</td></td><table>

### 2.2.2 统计指定key中1的数量
<table><tr><td bgcolor="#87CEFA"></br>


```shell
# 统计key中[start, end]范围内的bit为1的数量
# 如果start和count都不写就是全部数据进行统计
bitcount key1 [start end]
```
</td></tr></table>

**示例**：
<table><tr><td bgcolor=Black>
<font color=white>
127.0.0.1:6379> bitcount destData 1 100</br>
(integer) 21</br>
</font>
</td></td><table>

# 3 HyperLogLog数据类型
HyperLogLog能用于统计不重复数据的数量。
例子：
如果我们需要将统计一个系统中有多少个用户，一种做法是将每个用户的id存入set中，再通过scard命令获取set中元素的个数。但为了效率提升，还可以使用Bitmaps来存储------将每个用户的id作为Bitmaps的下标，用户存在就在对应的下标的bit位置为1，最后通过bitcount命令就可以获取有多少个用户了。
虽然使用Bitmaps做这些要比set效率高。但是如果用户有几千万个，那么一个几千万bit的字符串也是很大的。对此我们可以通过HyperLogLog来完成这样的功能。
HyperLogLog使用了HyperLogLog算法来做不重复数据的统计，大大降低的存储量。
## 3.1 HyperLogLog相关命令
### 3.1.1 添加数据
<table><tr><td bgcolor="#87CEFA"></br>

```shell
# 向key中添加数据
pfadd key element1 element2...
```
</td></tr></table>

**示例**：
<table><tr><td bgcolor=Black>
<font color=white>
127.0.0.1:6379> pfadd hyper a s d f</br>
(integer) 1</br>
127.0.0.1:6379> pfadd hyper a</br>
(integer) 0</br>
</font>
</td></td><table>

### 3.1.2 统计数据
<table><tr><td bgcolor="#87CEFA"></br>

```shell
# 统计key1,key2...合并后的不重复数据个数，key1,key2...合并的时候也是重复数据只保留一个
pfcount key1 [key2...]
```
</td></tr></table>

**示例**：
<table><tr><td bgcolor=Black>
<font color=white>
127.0.0.1:6379> pfcount hyper hyper2</br>
(integer) 8</br>
</font>
</td></td><table>

### 3.1.3 合并数据
<table><tr><td bgcolor="#87CEFA"></br>

```shell
# 合并key1,key2...到destKey中，重复数据只保留1个
pfmerge destKey key1 [key2,key3...]
```
</td></tr></table>

**示例**：
<table><tr><td bgcolor=Black>
<font color=white>
127.0.0.1:6379>  pfadd hyper1 a b c</br>
(integer) 1</br>
127.0.0.1:6379> pfadd hyper2 1 2 3 4</br>
(integer) 1</br>
127.0.0.1:6379> pfmerge destHyper hyper1 hyper2</br>
OK</br>
127.0.0.1:6379> pfcount destHyper</br>
(integer) 7</br>
</font>
</td></td><table>

## 3.2 HyperLogLog注意事项
- HyperlogLog是用于做基数(集合中不重复数据的个数)统计的，不是集合，不保存数据，只记录数量而不记录具体数据。
 - 核心是技术估算算法，最终数值存在一定误差。基数估计的结果是一个带有0.81%标准错误的近似值。 
 - HyperLogLog消耗空间极小，每个HyperLogLog key占用了12Kb的内存用于标记基数。
 - pfadd命令不是一次性分配存储空间12Kb，它是一点一点加的，也就是说存储消耗上限是12Kb
 - 使用pfmerge命令后，destKey的存储空间变为12K, 即使合并的源HyperLogLog key占用空间之和没有12Kb

 # 4 GEO类型
GEO类型主要用于计算两个位置(经纬度)之间的距离，这只计算水平距离，不算高度。
## 4.1 GEO基本操作
### 4.1.1 添加、获取坐标点
<table><tr><td bgcolor="#87CEFA"></br>

```shell
# 将一个名为member的坐标点添加到key中,经纬度是(longitude,latitude)
geoadd key longitude latitude member [longitude1 latitude1 member1 ...]

# 获取key中坐标点member1,member2...的经纬度
geopos key member1 member2...
```
</td></tr></table>

**示例**：
<table><tr><td bgcolor=Black>
<font color=white>
127.0.0.1:6379> geoadd map 12.4 22.1 pos1 -1.3 56 pos2</br>
(integer) 2</br>
127.0.0.1:6379> geopos map pos1 pos2</br>
1) 1) "12.39999979734420776"</br>
&emsp;2) "22.09999932656907617"</br>
1) 1) "-1.2999996542930603"</br>
&emsp;2) "55.9999988613003552"</br>
</font>
</td></tr></table>

### 4.1.2 计算坐标点的距离
<table><tr><td bgcolor="#87CEFA"></br>

```shell
# 获取key中的两个点member1,member2之间的水平距离.
#默认单位米，如果要转化为其他单位，需要在命令最后加上对应的长度单位（m, km, ft(英尺), mi(英里)）来替换[unit参数]

geodist key member1 member2 [unit]
```
</td></tr></table>

**示例**：
<table><tr><td bgcolor=Black>
<font color=white>
127.0.0.1:6379> geodist map pos1 pos2</br>
"3935828.1157"</br>
127.0.0.1:6379> geodist map pos1 pos2 km</br>
"3935.8281"</br>
127.0.0.1:6379> geodist map pos1 pos2 ft</br>
</font>
</td></tr></table>

### 4.1.3 求范围内的点的信息
**根据坐标求范围内的数据**
<table><tr><td bgcolor="#87CEFA"></br>

```shell
# 求位于 经纬度(longitude,latitude)的点半径redius（单位是m|km|ft|mi）范围内的点
# 如果还需要显示与圆心的距离，就加上withdist
# 如果需要显示点的hash值信息，就添加withhash
georadius key longitude latitude redius m|km|ft|mi [withdist] [withhash] [count count]
```
</td></tr></table>

**示例**：
<table><tr><td bgcolor=Black>
<font color=white>
127.0.0.1:6379> georadius map 0.0 1.2 3000 km withdist</br>
1) 1) "pos1"</br>
&emsp; 2) "2684.2777"</br>
127.0.0.1:6379> georadius map 0.0 1.2 3000 km withdist withhash</br>
1) 1) "pos1"</br>
&emsp; 2) "2684.2777"</br>
&emsp; 3) (integer) 3456981065411951</br>
</font>
</td></tr></table>

**根据点求范围内的数据**
<table><tr><td bgcolor="#87CEFA"></br>

```shell
# 求在 以key中的member为圆心，半径radius(单位是m|km|ft|mi)的圆 范围内的点的信息
# 如果还需要显示与圆心的距离，就加上withdist
# 如果需要显示点的hash值信息，就添加withhash
georadiusbymember key member raduis m|km|ft|mi [withdist] [withhash] [count count]
```
</td></tr></table>

**示例**：
<table><tr><td bgcolor=Black>
<font color=white>
127.0.0.1:6379> georadius map 123 21 1000 km withdist withhash</br>
1) 1) "post1"</br>
&nbsp&nbsp&nbsp 2) "124.0859"</br>
&nbsp&nbsp&nbsp 3) (integer) 4049197069141317</font>
</td></td><table>

### 4.1.4 获取对应点的hash值
<table><tr><td bgcolor="#87CEFA"></br>

```shell
# 获取key中membr1,member2...的hash值
geohash key member1 member2...
```
</td></tr></table>

**示例**：
<table><tr><td bgcolor=Black>
<font color=white>
127.0.0.1:6379> geohash map pos1 pos2</br>
1) "wezterptb40"</br>
2) "s067by61mk0"</br>
</font>
</td></td><table>
