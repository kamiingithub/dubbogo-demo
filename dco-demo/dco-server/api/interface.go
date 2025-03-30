package api

import "golang.org/x/net/context"

// DcoCreativeService 定义服务接口
type DcoCreativeService interface {
	GetDynamicCreative(ctx context.Context, req *DcoCreativeRequest) (*DcoCreativeResponse, error)
}
