package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	AppName    string
	PageSize   int
	LogPath    string
	SqlDir     string
	DailyDir   string
	LogFileExt string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type         string
	User         string
	Password     string
	Host         string
	Name         string
	TablePrefix  string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifeTime  int
}

var DatabaseSetting = &Database{}
var MetaDatabaseSetting = &Database{}

var (
	Cfg *ini.File
)

func Init() {
	var err error
	Cfg, err = ini.Load("env.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/env.ini': %v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	err = Cfg.Section("meat_database").MapTo(MetaDatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}

	err = Cfg.Section("100_database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}
}
