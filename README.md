### goJson

go语言解析Json

### 如何使用  
# Importing

    import github.com/hljdxjiang/gojson

### Documentation

### 代码思路

#1、将Json数据类型分为Jsonobject和JsonArray 两种类型

#2、JsonObject 为多个Jsonobjectitem组成

#3、JsonArray 为多个JsonVar组成

#4、Jsonobjectitem为 string,JsonVar两部分组成

#5、JsonVar的内容为interface类型，可以存放多种数据

### 已知问题
#1、所有基本类型的数据Encode的时候，代码均被转成string。
    如bool变为“bool” 10变为了“10”
#2、目前所有功能都已经测试，没有发现大的bug.
代码没有进行效率测试。也没有考虑竞争死锁的问题。

### 联系方式
以上代码参考了java工程的fastJson的实现及go-samplejson的部分代码。
如有发现问题或有好的建议。环境大家与我联系。

