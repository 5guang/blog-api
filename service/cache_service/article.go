package cache_service

import (
	"blog/pkg/e"
	"blog/pkg/goredis"
	"strconv"
	"strings"
)

type CacheArticle struct {
	ID        int
	TagId     int
	CreatedBy string
	State     int
	PageSize  int
	PageNum   int
}

// 获取所有文章的缓存key
//func (c *CacheArticle)GetCacheArticlesKey() string {
//	keys := []string{
//		e.CACHE_ARTICLE,
//		"LIST",
//		"ALL",
//	}
//	return strings.Join(keys,"_")
//}

func (c *CacheArticle) GetCacheArticleKey() string {
	keys := []string{
		e.CACHE_ARTICLE,
		"LIST",
	}
	if c.ID > 0 {
		keys = append(keys, "ARTICLE_ID", strconv.Itoa(c.ID))
	}
	if c.PageNum > 0 {
		keys = append(keys, "PAGE_NUM", strconv.Itoa(c.PageNum))
	}
	if c.PageSize > 0 {
		keys = append(keys, "PAGE_SIZE", strconv.Itoa(c.PageSize))
	}
	if c.TagId > 0 {
		keys = append(keys, "TAG_ID", strconv.Itoa(c.TagId))
	}
	if c.CreatedBy != "" {
		keys = append(keys, c.CreatedBy)
	}
	if c.State > 0 {
		keys = append(keys, "STATE", strconv.Itoa(c.State))
	}
	return strings.Join(keys, "_")
}

// 判断缓存是否存在
func (c *CacheArticle) IsCacheExist(key string) bool {
	return goredis.IsExist(key)
}

// 获取缓存
func (c *CacheArticle) Get(key string) ([]byte, error) {
	return goredis.Get(key)
}

// 设置缓存 默认缓存过期时间为3小时
func (c *CacheArticle) Set(key string, data interface{}) error {
	return goredis.Set(key, data, 10800)
}

// 删除指定缓存
func (c *CacheArticle) Del(key string) (bool, error) {
	return goredis.Delete(key)
}

// 清空缓存
func (c *CacheArticle) Clear() (bool, error) {
	return goredis.Clear()
}
