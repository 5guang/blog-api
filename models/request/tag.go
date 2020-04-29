package request

type ReqAddTag struct {
	Head Head `json:"head" validate:"required"`
	Body ReqAddTagBody `json:"body" validate:"required"`
}

type ReqDelTag struct {
	Head Head `json:"head" validate:"required"`
	Body ReqDelTagBody `json:"body" validate:"required"`
}

type ReqUpdateTag struct {
	Head Head `json:"head" validate:"required"`
	Body ReqUpdateTagBody `json:"body" validate:"required"`
}

type ReqDelTagBody struct {
	ID int `json:"id" validate:"required,gte=1"`
}

type ReqUpdateTagBody struct {
	ID int `json:"id" validate:"gte=1"`
	Name string `json:"name" validate:"required,min=2,max=20"`
	Updater string `json:"updater" validate:"required,min=2,max=50"`
	State int `json:"state" validate:"oneof=0 1"`

}

type ReqAddTagBody struct {
	Creator string `json:"creator" validate:"required"`
	State int `json:"state" validate:"oneof=0 1"`
	Name string `json:"name" validate:"required"`
}

