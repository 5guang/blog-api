package cache_service

import (
	"blog/pkg/e"
	"blog/pkg/goredis"
	"strings"
)

type Auth struct {
	Username string
	Password string
}

func (a *Auth) GetCacheAuthKey() string {
	keys := []string{
		e.CACHE_AUTH,
		"TOKEN",
	}
	if a.Username != "" {
		keys = append(keys, a.Username)
	}
	if a.Password != "" {
		keys = append(keys, a.Password[4:8])
	}
	return strings.Join(keys, "-")
}
// 判断缓存是否存在
func (a *Auth) IsCacheExist(key string) bool {
	return goredis.IsExist(key)
}

// 获取缓存
func (a *Auth) Get(key string) ([]byte, error) {
	return goredis.Get(key)
}

// 设置缓存 默认缓存过期时间为3小时
func (a *Auth) Set(key string, data interface{}) error {
	return goredis.Set(a.GetCacheAuthKey(), data, 10800)
}

// 删除指定缓存
func (a *Auth) Del(key string) (bool, error) {
	return goredis.Delete(key)
}

// 清空缓存
func (a *Auth) Clear() (bool, error) {
	return goredis.Clear()
}
