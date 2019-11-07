# easy
easycome的一个应用

## 生成消息协议
使用protobuf来进行与其他服务器间的通信，自动生成消息的id，服务端和客户端的相关文件等

windows环境下
在data文件下执行gen.bat生成client，pb文件的相关文件，其中protoc-gen-go和protoc
是protobuf的安装的文件，proto文件夹下存放消息体的源文件，而protoc-gen-msgcode是插件
文件，插件源码在easycome中protoc-gen-msg的。

linux环境下
编译相应的protoc-gen-msg，protoc-gen-go和protoc文件，gen.sh等，后续会进行更新

## 安装和使用etcd
[etcd地址]https://github.com/etcd-io/etcd