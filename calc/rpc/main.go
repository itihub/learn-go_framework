package main

import "fmt"

func Add(a, b int) int {
	total := a + b
	return total
}

type Company struct {
	Name    string
	Address string
}

type Employee struct {
	Name    string
	company Company
}

type PrintResult struct {
	Info string
	Err  error
}

func RpcPrintln(employee Employee) PrintResult {
	// rpc中的第二个点传输协、数据编码协议
	// http1.x http2.0协议
	// http协议底层使用的也是 tcp, http 现在主流的是 1.X 这种协议有性能问题 一次性 一旦结果返回连接就断开
	// 1. 直接自己基 tcp/udp协议去封装一层协议 myhttp, 没有通用性, http2.0 既有 http 的特性也有长连接的特性 grpc
	/*
		客户端
			1. 建立连接 tcp/http
			2. 将 employee 对象序列化成 json 字符 - 序列化
			3. 发送 json 字符串 - 调用成功后实际上你接收到的是一个二进制的数据
			4. 等待服务器发送结果
			5. 将服务器返回的数据解析成 PrintResult 对象 - 反序列化
		服务端
			1. 监听网络端口 80
			2. 读取数据 - 二进制的 json 数据
			3. 对数据进行反序列化 Employee 对象
			4. 开始处理业务逻辑
			5. 将处理的结果 PrintResult 序列化成 json 二进制数据 - 序列化
			6. 将数据返回
		序列化和反序列化是可以选择的, 不一定要采用 json、xml、proto、msgpack

	*/

	return PrintResult{}
}

func main() {
	//现在我们想把Add函数变成一个远程的函数调用,也就意味着需要把Add的数放在远程展务器上去运行
	/*
		我们原本的电商系统,这里地方有一段逻辑,这个逻辑是扣减库存,但是库存服务是一个独立的系统, reduce,那如何调用
		一定会牵到网络, 做成一个web服务(gin、beego、net/httpserver)
		1. 这个函数的调用参数如何传递-json(json是一种数据格式的协议)/xml/proto/msgpack -编码协议
			现在网络调用有两个端-客户端 应该干嘛? 将数据传输到gin
			gin-服务端,服务端负责解析数据
	*/
	//fmt.Println(Add(1, 2))

	// 将这个打印的工作放在另一台服务器，我需要将本地的内存对象 struct  这样不行，可行的方式就是将 struct 序列化成 json 对象并进行传输
	fmt.Println(Employee{
		Name: "bobby",
		company: Company{
			Name:    "China",
			Address: "Beijing",
		},
	})
	// 远程的服务需要将二进制对象反解成 struct 对象
	// 搞这么麻烦，直接全部使用 json 去格式化不香吗？这种做法在浏览器和gin服务之间可以，但是如果你是一个大型的分布式系统 就变得难以维护
}
