package response



type CommonResponse struct {
	Head Head `json:"head"`
	Body Body `json:"body"`
}

type Head struct {
	Token string `json:"token"`
}

type Status struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
}
// Response 基础序列化器
type Body struct {
	Status Status         `json:"status"`
	Data interface{} `json:"data,omitempty"`
}
