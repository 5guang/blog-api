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
	Body User `json:"body" validate:"required"`
}

type User struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

