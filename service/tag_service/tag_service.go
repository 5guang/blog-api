package tag_service

import (
	"blog/dao"
	"blog/models/response"
	"blog/pkg/goredis"
	"blog/pkg/logging"
	"blog/pkg/util"
	"blog/service/cache_service"
	"encoding/json"
)

type Tag struct {
	Name      string
	CreatedBy string
	UpdatedBy string
	DeletedOn int
	State     int
	ID        int
	PageNum   int
	PageSize  int
	ArticleId int
}

//通过name判断tag是否存在
func (t *Tag) ExistByName() (bool, error) {
	return dao.ExistTagByName(t.Name)
}

//通过id判断tag是否存在
func (t *Tag) ExistByID() (bool, error) {
	return dao.ExistTagByID(t.ID)
}

// 获取tag总数
func (t *Tag) Count() (int, error) {
	return dao.GetTagTotal(t.getMaps())
}

// 获取所有文章
func (t *Tag) GetAll() ([]response.ResTag, error) {
	var (
		tags               []dao.Tag
		err                error
		cacheTags, resTags []response.ResTag
	)
	// 先去redis中去查找 如果没有再去调用数据库
	bol, key := cache_service.IsCacheTagsExist()
	if bol == true {
		data, err := goredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheTags)
			return cacheTags, nil
		}
	}
	tags, err = dao.GetTags(t.PageSize, t.PageNum, t.getMaps())
	if err != nil {
		return nil, err
	}
	resTags = util.DaoTag2ResTag(tags)
	goredis.Set(key, resTags, 3600)
	return resTags, nil
}

// 通过articleID获取所有的tag
func (t *Tag) GetTagsByArticleId() ([]response.ResTag, error) {
	var (
		tags    []dao.Tag
		err     error
		resTags []response.ResTag
	)
	cacheService := cache_service.CacheTag{
		ArticleId: t.ArticleId,
	}
	key := cacheService.GetCacheTagKey()
	isExist := cacheService.IsCacheExist(key)
	// 判断缓存是否存在，存在直接使用缓存
	if isExist == true {
		data, err := cacheService.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			err = json.Unmarshal(data, &resTags)
			if err != nil {
				return nil, err
			}
			return resTags, nil
		}
	}
	tags, err = dao.GetTagsByArticleId(t.ArticleId)
	if err != nil {
		return nil, err
	}
	resTags = util.DaoTag2ResTag(tags)
	cacheService.Set(key, resTags)
	return resTags, nil
}

// 添加一个tag
func (t *Tag) Add() error {
	// 每次更新数据库前需要删除缓存 确保缓存和数据库同步
	key := cache_service.GetCacheTagsKey()
	_, err := goredis.Delete(key)
	if err != nil {
		logging.Info(err)
	}
	return dao.AddTag(t.Name, t.State, t.CreatedBy)
}

// 删除一个tag
func (t *Tag) DelTag() error {
	// 每次删除前需要删除缓存 确保缓存和数据库同步
	key := cache_service.GetCacheTagsKey()
	_, err := goredis.Delete(key)
	if err != nil {
		logging.Info(err)
	}
	return dao.DeleteTag(t.ID)
}

// 更新tag
func (t *Tag) UpdateTag() error {
	// 清除通过tagId查找文章的缓存
	{
		cacheService := cache_service.CacheArticle{
			TagId: t.ID,
		}
		key := cacheService.GetCacheArticleKey()
		if _, err := cacheService.Del(key); err != nil {
			logging.Info(err)
		}
	}
	// 每次更新前需要删除缓存 确保缓存和数据库同步
	{
		allKey := cache_service.GetCacheTagsKey()
		_, err := goredis.Delete(allKey)
		if err != nil {
			logging.Info(err)
		}
	}
	var tag = dao.Tag{
		DataBaseModel: dao.DataBaseModel{
			UpdatedBy: t.UpdatedBy,
			State:     t.State,
		},
		Name: t.Name,
	}
	return dao.UpdateTag(t.ID, tag)
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	// 软删除 DeletedOn=0表示删除了此条数据
	maps["deleted_on"] = 0

	if t.Name != "" {
		maps["name"] = t.Name
	}
	if t.State >= 0 {
		maps["state"] = t.State
	}

	return maps
}
