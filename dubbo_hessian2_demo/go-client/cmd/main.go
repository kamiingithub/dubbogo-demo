package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"fmt"
	hessian "github.com/apache/dubbo-go-hessian2"
	"github.com/dubbogo/gost/log/logger"
)

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
	// request可以省略
	//hessian.RegisterPOJO(&Request{})
	hessian.RegisterPOJO(&Response{})
}

// 定义服务结构体
var u = new(struct {
	GetUser     func(ctx context.Context, name string) (string, error)
	GetUserJson func(ctx context.Context, req *Request) (*Response, error)
})

func main() {

	//u := new(UserProvider)
	// 注册服务消费者
	config.SetConsumerService(u)
	if err := config.Load(); err != nil {
		panic(err)
	}

	resp1, _ := u.GetUser(context.Background(), "Jack")
	fmt.Println("Response from server:", resp1)

	resp, err := u.GetUserJson(context.Background(), &Request{Name: "Kami"})
	if err != nil {
		logger.Errorf("Error calling GetUserJson: %v", err)
		return
	}
	fmt.Println("Response from server:", resp.NickName)
}
