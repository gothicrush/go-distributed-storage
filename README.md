# 单机版云存储

### 架构图

![](https://github.com/gothicrush/go-distributed-storage/blob/master/ReadmeImage/version1.png)

### API接口

* PUT /objects/<object_name>
  * <object_name>中特殊字符必须先转义
  * 不能含有 '/'，否则'/'后的名字被丢弃

  

* GET /objects/<object_name>
  * 当请求资源不存在时，返回找不到资源
  * <object_name>中特殊字符必须先转义
  * 不能含有 '/'，否则'/'后的名字被丢弃
