# easy
easycome的一个应用


## 安装和使用etcd
[etcd地址]https://github.com/etcd-io/etcd
windows下下载etcd的windows的版本，修改etcdstart.sh路径并直接运行启动

## 生成消息协议
使用protobuf来进行与其他服务器间的通信，自动生成消息的id，服务端和客户端的相关文件等

windows环境下
在data文件下执行gen.bat生成client，pb文件的相关文件，其中protoc-gen-go和protoc
是protobuf的安装的文件，proto文件夹下存放消息体的源文件，而protoc-gen-msgcode是插件
文件，插件源码在easycome/protoc-gen-msg的。

linux环境下
编译相应的protoc-gen-msg，源码在easycome/protoc-gen-msg中，
下载protoc-gen-go和protoc文件

## 服务器配置
在server.config中可以配置相关的服务器端口

## 编译项目
初次运行需要go mod 相关包，当前目录下执行go mod tidy
之后运行build.sh

## 启动整个服务器进程组
执行start.sh开启活动进程，在log下的文件可以查看几个进程是否启动成功
或者执行ps aux|grep _test 查看是否进程启动
示例中启动了两个agent服务器和两个game服务器还有unique和中心服各一个