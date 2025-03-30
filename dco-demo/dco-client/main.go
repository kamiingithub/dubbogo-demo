package main

import (
	"context"
	"dubbo-demo/dco-client/api"
	"fmt"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

var dcoCreativeService = new(api.DcoCreativeService)

func main() {
	config.SetConsumerService(dcoCreativeService)
	//config.SetConsumerServiceByInterfaceName("api.DcoService", dcoCreativeService)
	if err := config.Load(); err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	req := &api.DcoCreativeRequest{
		RequestList: []api.RequestItem{
			{
				ProductId:  433114,
				CreativeId: 4686809,
				BuType:     "DHTL",
			},
		},
		RequestId: "11",
		OnlyImage: false,
	}

	resp, err := (*dcoCreativeService).GetDynamicCreative(context.Background(), req)
	if err != nil {
		fmt.Printf("failed to call GetDynamicCreative: %v\n", err)
		return
	}

	fmt.Printf("response: %+v\n", resp)
}
