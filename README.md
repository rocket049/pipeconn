# pipeconn 用标准输入输出和管道模拟 io.ReadWriteCloser

用标准输入输出和管道模拟 io.ReadWriteCloser，可以用于编写不依靠网络连接的 rpc server/client。

服务器方建立连接时调用`pipeconn.NewServerPipeConn()`

客户端建立连接时调用`NewClientPipeConn(progPath , args...)`

**`rpc` 目录中的是一个示例程序。**