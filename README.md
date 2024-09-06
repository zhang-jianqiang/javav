# java版本管理
## 使用

自己测试使用的windows系统切换java版本工具

**1、配置JAVA_HOME环境变量**

JAVA_HOME作为java软链的目录

将JAVA_HOME/bin目录添加到path环境变量

**2、安装javav**

go版本1.22

源码安装

```shell
go isntall github.com/zhang-jianqiang/javav@latest
```

二进制安装

下载二进制并将二进制所在目录添加到path环境变量

**3、配置**

配置存放多个java版本的目录

```shell
javav config D:/javastudy/java
```

列出可用的java版本

```shell
javav ls
```

切换java版本

```shell
javav use jdk-21.0.1
```
