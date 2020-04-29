package cache_service

import (
	"blog/pkg/e"
	"blog/pkg/goredis"
	"strconv"
	"strings"
)

type CacheTag struct {
	ID        int
	Name      string
	State     int
	ArticleId int
	PageNum   int
	PageSize  int
}

func (t *CacheTag) GetCacheTagKey() string {
	keys := []string{
		e.CACHE_TAG,
		"LIST",
	}
	if t.Name != "" {
		keys = append(keys, t.Name)
	}
	if t.State >= 0 {
		keys = append(keys, "STATE", strconv.Itoa(t.State))
	}
	if t.PageNum > 0 {
		keys = append(keys, "PAGE_NUM", strconv.Itoa(t.PageNum))
	}
	if t.ArticleId > 0 {
		keys = append(keys, "ARTICLE_ID", strconv.Itoa(t.ArticleId))
	}
	if t.PageSize > 0 {
		keys = append(keys, "PAGE_SIZE", strconv.Itoa(t.PageSize))
	}
	return strings.Join(keys, "_")
}

func GetCacheTagsKey() string {
	keys := []string{
		e.CACHE_TAG,
		"LIST",
		"ALL",
	}
	return strings.Join(keys, "_")
}

// 判断获取总条数的缓存key是否存在缓存中
func IsCacheTagsExist() (bool, string) {
	key := GetCacheTagsKey()
	return goredis.IsExist(key), key
}

// 判断缓存是否存在
func (c *CacheTag) IsCacheExist(key string) bool {
	return goredis.IsExist(key)
}

// 获取缓存
func (c *CacheTag) Get(key string) ([]byte, error) {
	return goredis.Get(key)
}

// 设置缓存 默认缓存过期时间为3小时
func (c *CacheTag) Set(key string, data interface{}) error {
	return goredis.Set(key, data, 10800)
}

// 删除指定缓存
func (c *CacheTag) Del(key string) (bool, error) {
	return goredis.Delete(key)
}

// 清空缓存
func (c *CacheTag) Clear() (bool, error) {
	return goredis.Clear()
}
