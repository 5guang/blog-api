package v1

import (
	"blog/models/request"
	"blog/models/response"
	"blog/pkg/app"
	"blog/pkg/e"
	"blog/pkg/setting"
	"blog/pkg/util"
	"blog/service/tag_service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetTags(c *gin.Context) {
	appG := app.Gin{c}
	name := c.Query("name")
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}
	tagService := tag_service.Tag{
		Name:     name,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	tags, err := tagService.GetAll()

	if err != nil {
		appG.Response( e.ERROR_GET_TAGS_FAIL, nil)
		return
	}
	count := len(tags)
	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_TAG_FAIL, nil)
	//	return
	//}
	resTagData := response.ResTagData{
		TagList: tags,
		Count:   count,
	}
	appG.Response( e.SUCCESS, resTagData)
}

func GetTagsByArticleId(c *gin.Context) {
	appG := app.Gin{C: c}

	var (
		t    request.ReqArticleById
		tags []response.ResTag
		err  error
	)

	 errCode := app.BindAndValid(c, &t)
	if errCode != e.SUCCESS {
		appG.Response( errCode, nil)
		return
	}
	tagService := &tag_service.Tag{
		ArticleId: t.Body.ID,
	}
	tags, err = tagService.GetTagsByArticleId()
	if err != nil {
		appG.Response( e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	resData := response.ResTagData{
		TagList: tags,
		Count:   len(tags),
	}
	appG.Response( errCode, resData)
}

func AddTag(c *gin.Context) {
	var (
		appG      = app.Gin{C: c}
		reqAddTag request.ReqAddTag
		err       error
	)
	 errCode := app.BindAndValid(c, &reqAddTag)
	if errCode != e.SUCCESS {
		appG.Response( errCode, nil)
		return
	}
	tagService := tag_service.Tag{
		Name:      reqAddTag.Body.Name,
		CreatedBy: reqAddTag.Body.Creator,
		State:     reqAddTag.Body.State,
	}
	exists, err := tagService.ExistByName()
	if err != nil {
		appG.Response( e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if exists {
		appG.Response( e.ERROR_EXIST_TAG, nil)
		return
	}
	err = tagService.Add()
	if err != nil {
		appG.Response( e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response( e.SUCCESS, nil)
}
func DelTag(c *gin.Context) {
	var (
		appG      = app.Gin{C: c}
		reqDelTag request.ReqDelTag
		err       error
	)
	 errCode := app.BindAndValid(c, &reqDelTag)
	if errCode != e.SUCCESS {
		appG.Response( errCode, nil)
		return
	}
	tagService := tag_service.Tag{
		ID: reqDelTag.Body.ID,
	}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response( e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if !exists {
		appG.Response( e.ERROR_NOT_EXIST_TAG, nil)
		return
	}
	err = tagService.DelTag()
	if err != nil {
		appG.Response( e.ERROR_DELETE_TAG_FAIL, nil)
		return
	}

	appG.Response( e.SUCCESS, nil)
}
func UpdateTag(c *gin.Context) {
	var (
		appG         = app.Gin{C: c}
		reqUpdateTag request.ReqUpdateTag
		err          error
	)
	 errCode := app.BindAndValid(c, &reqUpdateTag)
	if errCode != e.SUCCESS {
		appG.Response( errCode, nil)
		return
	}
	tagService := tag_service.Tag{
		ID:        reqUpdateTag.Body.ID,
		State:     reqUpdateTag.Body.State,
		Name:      reqUpdateTag.Body.Name,
		UpdatedBy: reqUpdateTag.Body.Updater,
	}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response( e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}
	if !exists {
		appG.Response( e.ERROR_NOT_EXIST_TAG, nil)
		return
	}
	err = tagService.UpdateTag()
	if err != nil {
		appG.Response( e.ERROR_EDIT_TAG_FAIL, nil)
		return
	}

	appG.Response( e.SUCCESS, nil)
}
