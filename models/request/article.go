package request

// 添加文章
type ReqAddArticle struct {
	Head Head              `json:"head" validate:"required"`
	Body ReqAddArticleBody `json:"body" validate:"required"`
}
type ReqGetArticle struct {
	Head Head              `json:"head" validate:"required"`
	Body ReqGetArticleBody `json:"body" validate:"required"`
}

// 获取文章
type ReqGetArticleBody struct {
	Creator  string `json:"creator" validate:"omitempty,min=2,max=50"`
	ID       int    `json:"id"`
	State    int    `json:"state" validate:"omitempty,oneof=0 1"`
	PageSize int    `json:"pageSize" validate:"omitempty,min=1"`
	PageNum  int    `json:"pageNum" validate:"omitempty,min=1"`
}

// 更新文章
type ReqUpdateArticle struct {
	Head Head                 `json:"head" validate:"required"`
	Body ReqUpdateArticleBody `json:"body" validate:"required"`
}

type ReqArticleById struct {
	Head Head                  `json:"head" validate:"required"`
	Body ReqArticleByTagIdBody `json:"body" validate:"required"`
}

type ReqArticleByTagIdBody struct {
	ID int `json:"id" validate:"required"`
}

type ReqUpdateArticleBody struct {
	ArticleId int    `json:"articleId" validate:"required"`
	Title     string `json:"title" validate:"omitempty,min=2,max=50"`
	Desc      string `json:"desc" validate:"omitempty,min=2,max=50"`
	UpdateBy  string `json:"updateBy" validate:"omitempty,min=2,max=20"`
	State     int    `json:"state" validate:"omitempty,oneof=0 1"`
	Content   string `json:"content" validate:"omitempty,min=20"`
}

type ReqAddArticleBody struct {
	TagList []int  `json:"tagList" validate:"required,min=1"`
	Title   string `json:"title" validate:"required,min=2,max=50"`
	Desc    string `json:"desc" validate:"omitempty,min=2,max=50"`
	Creator string `json:"creator" validate:"required,min=2,max=20"`
	State   int    `json:"state" validate:"oneof=0 1"`
	Content string `json:"content" validate:"required,min=20"`
}

//type ReqAddArticleBodyTagList struct {
//	//ReqAddTagBody
//	TagId int `json:"tagId" validate:"required"`
//}
