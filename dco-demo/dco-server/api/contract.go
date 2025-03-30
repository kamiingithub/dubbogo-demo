package api

// DcoCreativeRequest 定义请求结构体
type DcoCreativeRequest struct {
	RequestList []RequestItem `json:"requestList"`
	RequestId   string        `json:"requestId"`
	OnlyImage   bool          `json:"onlyImage"`
}

// RequestItem 定义请求项结构体
type RequestItem struct {
	ProductId  int64  `json:"productId"`
	CreativeId int64  `json:"creativeId"`
	BuType     string `json:"buType"`
}

// DcoCreativeResponse 定义响应结构体
type DcoCreativeResponse struct {
	Results []Result `json:"results"`
	Status  Status   `json:"status"`
}

// Result 定义结果结构体
type Result struct {
	Request   Request    `json:"request"`
	Resources []Resource `json:"resources"`
}

// Request 定义内部请求结构体
type Request struct {
	Realtime   bool   `json:"realtime"`
	ProductId  int64  `json:"productId"`
	CreativeId int64  `json:"creativeId"`
	BuType     string `json:"buType"`
}

// Resource 定义资源结构体
type Resource struct {
	Uri       string            `json:"uri"`
	Attrs     map[string]string `json:"attrs"`
	Timestamp int64             `json:"timestamp"`
}

// Status 定义状态结构体
type Status struct {
	Success bool   `json:"success"`
	ErrCode int    `json:"errCode"`
	Message string `json:"message"`
}
