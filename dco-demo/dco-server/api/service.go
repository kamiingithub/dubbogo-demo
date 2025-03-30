package api

import (
	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/protocol/triple/triple_protocol"
	"dubbo.apache.org/dubbo-go/v3/server"
	"golang.org/x/net/context"
)

// DcoCreativeServiceImpl 实现 DcoCreativeService 接口
type DcoCreativeServiceImpl struct {
}

// GetDynamicCreative 实现接口方法
func (s *DcoCreativeServiceImpl) GetDynamicCreative(ctx context.Context, req *DcoCreativeRequest) (*DcoCreativeResponse, error) {
	resultPtrs := []*Result{
		{
			Request: Request{
				Realtime:   false,
				ProductId:  req.RequestList[0].ProductId,
				CreativeId: req.RequestList[0].CreativeId,
				BuType:     req.RequestList[0].BuType,
			},
			Resources: []Resource{
				{
					Uri: "https://dimg04.fws.qa.nt.ctripcorp.com/images/0zg4v1200002v94jy9E7F.jpg",
					Attrs: map[string]string{
						"LEFT_ANGLE_BRACKET":  "&lt;",
						"text_width@139273":   "1625",
						"RIGHT_ANGLE_BRACKET": "&gt;",
						"star":                "4",
						"_transparent":        "false",
						"v_civ":               "b679f5fbead7eb5f1a86ebc55df3f0ff",
						"_civ":                "b679f5fbead7eb5f1a86ebc55df3f0ff",
						"Name":                "北京和平里宾馆11(Beijing Peace HotelBeijing Peace Hotel)",
					},
					Timestamp: 1741334102846,
				},
			},
		},
	}

	// 将 []*api.Result 转换为 []api.Result
	results := make([]Result, len(resultPtrs))
	for i, ptr := range resultPtrs {
		results[i] = *ptr
	}

	resp := &DcoCreativeResponse{
		Results: results,
		Status: Status{
			Success: true,
			ErrCode: 0,
			Message: "ok",
		},
	}
	return resp, nil
}

func SetProviderService(srv common.RPCService) {
	dubbo.SetProviderServiceWithInfo(srv, &DcoService_ServiceInfo)
}

var DcoService_ServiceInfo = server.ServiceInfo{
	InterfaceName: "dco.DcoService",
	ServiceType:   (*DcoCreativeService)(nil),
	Methods: []server.MethodInfo{
		{
			Name: "GetDynamicCreative",
			Type: constant.CallUnary,
			ReqInitFunc: func() interface{} {
				return new(DcoCreativeRequest)
			},
			MethodFunc: func(ctx context.Context, args []interface{}, handler interface{}) (interface{}, error) {
				req := args[0].(*DcoCreativeRequest)
				res, err := handler.(DcoCreativeService).GetDynamicCreative(ctx, req)
				if err != nil {
					return nil, err
				}
				return triple_protocol.NewResponse(res), nil
			},
		},
	},
}
