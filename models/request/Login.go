package request

type ReqLoginBody struct {
	// validate校验中间用逗号隔开不能有空格 不然会会报错！！！
	Username string `json:"username" validate:"max=20,min=2" `
	Password string `json:"password" validate:"max=256"`
}

