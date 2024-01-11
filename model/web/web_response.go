package web

type WebResponse struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type WebResponseCount struct {
	Code      int         `json:"code"`
	Status    bool        `json:"status"`
	TotalData int         `json:"total_data"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}
