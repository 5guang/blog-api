package v1

import (
	"blog/models/request"
	"blog/models/response"
	"blog/pkg/app"
	"blog/pkg/e"
	"blog/service/article_service"
	"blog/service/tag_service"
	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
	appG := app.Gin{C: c}

	var (
		reqGetArticle request.ReqGetArticle
		articleList   []response.ResArticle
		err           error
	)

	 errCode := app.BindAndValid(c, &reqGetArticle)
	if errCode != e.SUCCESS {
		appG.Response( errCode, nil)
		return
	}
	articleService := &article_service.ArticleServiceModel{
		CreatedBy: reqGetArticle.Body.Creator,
		State:     reqGetArticle.Body.State,
		ID:        reqGetArticle.Body.ID,
		PageNum:   reqGetArticle.Body.PageNum,
		PageSize:  reqGetArticle.Body.PageSize,
	}
	articleList, err = articleService.GetArticles()
	if err != nil {
		appG.Response( e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	resData := response.ResArticleData{
		ArticleList: articleList,
		Count:       len(articleList),
	}
	appG.Response( errCode, resData)
}

func UpdateArticle(c *gin.Context) {
	appG := app.Gin{c}
	var (
		rUa request.ReqUpdateArticle
		err error
	)
	 errCode := app.BindAndValid(c, &rUa)
	if errCode != e.SUCCESS {
		appG.Response( errCode, nil)
		return
	}
	articleService := &article_service.ArticleServiceModel{
		Title:    rUa.Body.Title,
		Desc:     rUa.Body.Desc,
		Content:  rUa.Body.Content,
		State:    rUa.Body.State,
		UpdateBy: rUa.Body.UpdateBy,
		ID:       rUa.Body.ArticleId,
	}
	err = articleService.UpdateArticle()
	if err != nil {
		appG.Response( e.ERROR, nil)
		return
	}
	appG.Response( e.SUCCESS, nil)

}

func AddArticle(c *gin.Context) {
	appG := app.Gin{c}

	var (
		reqAddArticle request.ReqAddArticle
		tagList       []article_service.TagList
	)

	 errCode := app.BindAndValid(c, &reqAddArticle)

	if errCode != e.SUCCESS {
		appG.Response( errCode, nil)
		return
	}
	for _, tag := range reqAddArticle.Body.TagList {
		tmp := article_service.TagList{
			ID: tag,
		}
		tagList = append(tagList, tmp)
	}
	artService := &article_service.ArticleServiceModel{
		TagList:   tagList,
		Title:     reqAddArticle.Body.Title,
		Desc:      reqAddArticle.Body.Desc,
		Content:   reqAddArticle.Body.Content,
		CreatedBy: reqAddArticle.Body.Creator,
		State:     reqAddArticle.Body.State,
	}
	err, errCode := artService.AddArticle()
	if err != nil {
		appG.Response( errCode, nil)
		return
	}

	appG.Response( errCode, nil)

}

func DelArticle(c *gin.Context) {
	appG := app.Gin{c}
	var (
		rDa request.ReqArticleById
	)

	 errCode := app.BindAndValid(c, &rDa)
	if errCode != e.SUCCESS {
		appG.Response( errCode, nil)
		return
	}
	articleService := &article_service.ArticleServiceModel{ID: rDa.Body.ID}
	err := articleService.DelArticle()
	if err != nil {
		appG.Response( e.ERROR_DELETE_ARTICLE_FAIL, nil)
		return
	}
	appG.Response( e.SUCCESS, nil)
}

func GetArticlesByTagId(c *gin.Context) {
	appG := app.Gin{C: c}

	var (
		t           request.ReqArticleById
		articleList []response.ResArticle
		err         error
	)

	 errCode := app.BindAndValid(c, &t)
	if errCode != e.SUCCESS {
		appG.Response( errCode, nil)
		return
	}
	articleService := &article_service.ArticleServiceModel{
		TagId: t.Body.ID,
	}
	articleList, err = articleService.GetArticlesByTagId()
	if err != nil {
		appG.Response( e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	resData := response.ResArticleData{
		ArticleList: articleList,
		Count:       len(articleList),
	}
	appG.Response( errCode, resData)
}

// 通过文章id获取文章
func GetArticleByArticleId(c *gin.Context) {
	appG := app.Gin{C: c}

	var (
		t           request.ReqArticleById
		articleList []response.ResArticle
		err         error
		resTags     []response.ResTag
	)

	 errCode := app.BindAndValid(c, &t)
	if errCode != e.SUCCESS {
		appG.Response( errCode, nil)
		return
	}
	articleService := &article_service.ArticleServiceModel{
		ID: t.Body.ID,
	}
	tagService := &tag_service.Tag{ArticleId: t.Body.ID}
	// 通过id获取当前文章的所有tag
	resTags, err = tagService.GetTagsByArticleId()
	if err != nil {
		appG.Response( e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}
	// 获取当前文章
	articleList, err = articleService.GetArticle()
	if err != nil {
		appG.Response( e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}
	resData := response.ResArticleData{
		ArticleList: articleList,
		Count:       len(articleList),
		Tags:        resTags,
	}
	appG.Response( errCode, resData)
}
