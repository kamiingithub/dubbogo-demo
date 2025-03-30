package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"encoding/json"
	"fmt"
	"github.com/dubbogo/gost/log/logger"
	"log"
)

type Request struct {
	Name string
}

type Response struct {
	NickName string
}

// 定义服务结构体
var u = new(struct {
	GetUser     func(ctx context.Context, name string) (string, error)
	GetUserJson func(ctx context.Context, req []byte) ([]byte, error)
})

func main() {
	//u := new(UserProvider)
	// 注册服务消费者
	config.SetConsumerService(u)
	if err := config.Load(); err != nil {
		panic(err)
	}

	// 调用服务方法
	//resp, err := u.GetUser(context.Background(), "John")
	//if err != nil {
	//	logger.Errorf("Error calling GetUser: %v", err)
	//	return
	//}
	request := &Request{Name: "Kami"}
	req, err2 := json.Marshal(request)
	if err2 != nil {
		fmt.Printf("marchal error", err2)
	}
	resp, err := u.GetUserJson(context.Background(), req)
	if err != nil {
		logger.Errorf("Error calling GetUserJson: %v", err)
		return
	}
	var response Response
	if err := json.Unmarshal(resp, &response); err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}
	fmt.Println("Response from server:", response.NickName)
}
