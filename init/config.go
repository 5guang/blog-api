package init

import (
	"blog/dao"
	"blog/pkg/goredis"
	"blog/pkg/logging"
	"blog/pkg/setting"
)

func Init() {
	setting.Init()
	dao.Init()
	goredis.Init()
	logging.Init()
}
