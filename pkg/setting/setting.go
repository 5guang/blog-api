package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret       string
	RuntimeRootPath string
	PageSize        int
	TimeFormat  string
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	SaltSecret string
	Delimiter          string
	StretchingPassword string
	SaltLocalSecret string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var (
	cfg *ini.File
	AppSetting = &App{}
	ServerSetting   = &Server{}
	DataBaseSetting = &Database{}
	RedisSetting    = &Redis{}
)

func Init()  {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'setting/app.ini': %v", err)
	}
	// 读取配置文件
	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DataBaseSetting)
	mapTo("redis", RedisSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second

}

// 映射配置文件
func mapTo(section string, obj interface{})  {
	err := cfg.Section(section).MapTo(obj)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}