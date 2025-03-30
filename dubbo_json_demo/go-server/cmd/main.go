package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	_ "dubbo.apache.org/dubbo-go/v3/registry/protocol"
	"encoding/json"
	"fmt"
)

// 定义服务结构体
type UserProvider struct{}

type Request struct {
	Name string
}

type Response struct {
	NickName string
}

// 实现服务方法
func (u *UserProvider) GetUserJson(ctx context.Context, req []byte) ([]byte, error) {
	// 解析请求体
	var request Request
	if err := json.Unmarshal(req, &request); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %v", err)
	}
	fmt.Printf("Received request: %+v", request)

	// 构造响应
	response := &Response{fmt.Sprintf("===SUPER%s===", request.Name)}

	// 序列化响应体
	resp, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %v", err)
	}
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
