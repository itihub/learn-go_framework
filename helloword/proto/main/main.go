package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"learngoframework/helloword/proto"
)

type Hello struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Courses []string `json:"courses"`
}

// 测试序列化效果
func main() {

	// json 与 protobuf 对比

	jsonStruct := Hello{
		Name: "bobby",
		Age: 18,
		Courses: []string{"go", "gin", "微服务"},
	}
	jsonRsp, _ := json.Marshal(jsonStruct) // json进行序列化
	fmt.Println(string(jsonRsp))
	fmt.Println(len(jsonRsp))


	req := helloword.HelloRequest {
		Name: "bobby",
		Age: 18,
		Courses: []string{"go", "gin", "微服务"},
	}
	rsp, _ := proto.Marshal(&req) // protobuf进行序列化
	fmt.Println(string(rsp))
	fmt.Println(len(rsp))

	newReq := helloword.HelloRequest{}
	_  = proto.Unmarshal(rsp, &newReq) // protobuf进行反序列化
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses)

}
