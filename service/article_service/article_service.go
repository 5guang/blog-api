package article_service

import (
	"blog/dao"
	"blog/models/response"
	"blog/pkg/e"
	"blog/pkg/logging"
	"blog/pkg/util"
	"blog/service/cache_service"
	"encoding/json"
)

type TagList struct {
	Name      string
	CreatedBy string
	ID        int
	State     int
}

type ArticleServiceModel struct {
	TagList   []TagList
	Title     string
	Desc      string
	Content   string
	CreatedBy string
	State     int
	ID        int
	PageNum   int
	PageSize  int
	TagId     int
	UpdateBy  string
}

func (a *ArticleServiceModel) GetArticles() ([]response.ResArticle, error) {
	var (
		articles   []dao.Article
		resArticle []response.ResArticle
		err        error
	)
	art := dao.Article{
		Model: dao.Model{
			ID: a.ID,
		},
		DataBaseModel: dao.DataBaseModel{
			CreatedBy: a.CreatedBy,
			State:     a.State,
		},
	}
	cacheService := &cache_service.CacheArticle{
		ID:        a.ID,
		CreatedBy: a.CreatedBy,
		State:     a.State,
		PageSize:  a.PageSize,
		PageNum:   a.PageNum,
	}
	key := cacheService.GetCacheArticleKey()
	isExist := cacheService.IsCacheExist(key)
	// 判断缓存是否存在，存在直接使用缓存
	if isExist == true {
		data, err := cacheService.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			err = json.Unmarshal(data, &resArticle)
			if err != nil {
				return nil, err
			}
			return resArticle, nil
		}
	}
	articles, err = dao.GetArticles(a.PageSize, a.PageNum, art)
	if err != nil {
		return nil, err
	}
	resArticle = util.DaoArticle2ResArticle(articles)
	cacheService.Set(key, resArticle)
	return resArticle, nil
}

func (a *ArticleServiceModel) UpdateArticle() error {

	cacheService := &cache_service.CacheArticle{
		ID: a.ID,
	}
	// 每次更新数据库前需要删除缓存 确保缓存和数据库同步
	{
		key := cacheService.GetCacheArticleKey()
		_, err := cacheService.Del(key)
		if err != nil {
			logging.Info(err)
		}
	}
	art := dao.Article{
		Title:   a.Title,
		Desc:    a.Desc,
		Content: a.Content,
		DataBaseModel: dao.DataBaseModel{
			UpdatedBy: a.UpdateBy,
		},
	}
	return dao.UpdateArticle(a.ID, art)
}

func (a *ArticleServiceModel) AddArticle() (error, int) {
	cacheService := &cache_service.CacheArticle{}
	// 每次更新数据库前需要删除缓存 确保缓存和数据库同步
	_, err := cacheService.Clear()
	if err != nil {
		logging.Info(err)
	}
	var tags []dao.Tag
	for _, tag := range a.TagList {
		exist, err := dao.ExistTagByID(tag.ID)
		if err != nil || !exist {
			continue
		}
		tags = append(tags, dao.Tag{
			Model: dao.Model{ID: tag.ID},
			Name:  tag.Name,
			DataBaseModel: dao.DataBaseModel{
				CreatedBy: tag.CreatedBy,
				State:     tag.State,
			},
		})
	}
	art := dao.Article{
		Title:   a.Title,
		Desc:    a.Desc,
		Content: a.Content,
		DataBaseModel: dao.DataBaseModel{
			CreatedBy: a.CreatedBy,
			State:     a.State,
		},
		Tags: tags,
	}
	// artId为插入成功后返回的主键id
	err, _ = dao.AddArticle(art)
	if err != nil {
		return err, e.ERROR_ADD_ARTICLE_FAIL
	}
	return nil, e.SUCCESS
}

func (a *ArticleServiceModel) DelArticle() error {
	cacheService := &cache_service.CacheArticle{}
	// 每次更新数据库前需要删除缓存 确保缓存和数据库同步
	_, err := cacheService.Clear()
	if err != nil {
		logging.Info(err)
	}
	return dao.DelArticle(a.ID)
}

func (a *ArticleServiceModel) GetArticlesByTagId() (resArticle []response.ResArticle, err error) {
	var articles []dao.Article
	cacheService := &cache_service.CacheArticle{
		TagId: a.TagId,
	}
	key := cacheService.GetCacheArticleKey()
	isExist := cacheService.IsCacheExist(key)
	// 判断缓存是否存在，存在直接使用缓存
	if isExist == true {
		data, err := cacheService.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			err = json.Unmarshal(data, &resArticle)
			if err != nil {
				return nil, err
			}
			return resArticle, nil
		}
	}
	articles, err = dao.GetArticlesByTagId(a.TagId)
	if err != nil {
		return nil, err
	}
	resArticle = util.DaoArticle2ResArticle(articles)
	cacheService.Set(key, resArticle)
	return resArticle, nil
}

// 通过articleID获取指定文章
func (a *ArticleServiceModel) GetArticle() (resArticle []response.ResArticle, err error) {
	var (
		articles []dao.Article
	)
	cacheService := &cache_service.CacheArticle{
		ID: a.ID,
	}
	key := cacheService.GetCacheArticleKey()
	isExist := cacheService.IsCacheExist(key)
	// 判断缓存是否存在，存在直接使用缓存
	if isExist == true {
		data, err := cacheService.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			err = json.Unmarshal(data, &resArticle)
			if err != nil {
				return nil, err
			}
			return resArticle, nil
		}
	}
	articles, err = dao.GetArticle(a.ID)

	if err != nil {
		return nil, err
	}
	resArticle = util.DaoArticle2ResArticle(articles)
	cacheService.Set(key, resArticle)
	return resArticle, nil
}
