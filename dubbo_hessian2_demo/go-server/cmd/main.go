package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	_ "dubbo.apache.org/dubbo-go/v3/registry/protocol"
	"fmt"
	hessian "github.com/apache/dubbo-go-hessian2"
)

// 定义服务结构体
type UserProvider struct{}

type Request struct {
	Name string
}

func (r *Request) JavaClassName() string {
	return "com.apache.dubbo.hessian2.Request"
}

type Response struct {
	NickName string
}

func (r *Response) JavaClassName() string {
	return "com.apache.dubbo.hessian2.Response"
}

func init() {
	hessian.RegisterPOJO(&Request{})
	hessian.RegisterPOJO(&Response{})
}

// 实现服务方法
func (u *UserProvider) GetUserJson(ctx context.Context, req *Request) (*Response, error) {
	fmt.Printf("Received request: %+v", req)

	// 构造响应
	resp := &Response{fmt.Sprintf("===SUPER %s===", req.Name)}
	fmt.Println("call server success!")
	return resp, nil
}

// 实现服务方法
func (u *UserProvider) GetUser(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s! This is a JSON response.", name), nil
}

func main() {
	userProvider := &UserProvider{}
	// 注册服务提供者
	config.SetProviderService(userProvider)
	// 自定义序列化器
	//config.SetProviderCustomSerializer("json", func() interface{} {
	//	return &UserProvider{}
	//})
	//dubbo.SetProviderService(userProvider)
	if err := config.Load(); err != nil {
		panic(err)
	}
	select {}
}
