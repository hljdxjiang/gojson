# goJson

go语言解析Json

# 更新日志

2017-07-17

修复Int bool Float等类型Encode的时候输出为string的问题

# 如何使用  
## Importing

    import github.com/hljdxjiang/gojson

## Documentation

    暂未整理文档


# 代码思路

#1、将Json数据类型分为Jsonobject和JsonArray 两种类型

#2、JsonObject 为多个Jsonobjectitem组成

#3、JsonArray 为多个JsonVar组成

#4、Jsonobjectitem为 string,JsonVar两部分组成

#5、JsonVar的内容为interface类型，可以存放多种数据

# 已知问题
暂无

# 联系方式
 以上代码参考了java工程的fastJson的实现及go-samplejson的部分代码。
 
 如有发现问题或有好的建议。欢迎大家与我联系。
 
 作者邮箱 hljdxjiang@163.com
 
# 版权声明

代码已开源，供所有golang开发者使用，代码中如有侵权，请联系作者。以上未经大量项目使用验证。如有大的bug,导致生产问题，作者不负有任何责任。
