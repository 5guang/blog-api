package request

type ReqRegister struct {
	Head Head `json:"head" validate:"required"`
	Body ReqRegisterBody `json:"body" validate:"required"`
}

type ReqRegisterBody struct {
	User
	Nickname string `json:"nickname" validate:"required,min=1,max=15"`
	Email string `json:"email" validate:"required,email"`
}

type ReqLogin struct {
	Head Head `json:"head" validate:"required"`
	Body ReqLoginBody `json:"body" validate:"required"`
}

type ReqLoginBody struct {
	User
	// validate校验中间用逗号隔开不能有空格 不然会会报错！！！
	AdminPassword string `json:"adminPassword" validate:"omitempty,max=30"`
}

type User struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

