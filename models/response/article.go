package response

type ResArticle struct {
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Content     string `json:"content"`
	ArticleId   int    `json:"articleId"`
	Creator     string `json:"creator"`
	State       int    `json:"state"`
	CreatedTime string `json:"createdTime"`
}
type ResArticleData struct {
	ArticleList []ResArticle `json:"articleList"`
	Count       int          `json:"count"`
	Tags        []ResTag     `json:"tags"`
}
