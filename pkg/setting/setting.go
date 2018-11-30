package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	// Cfg load app.ini
	Cfg *ini.File
	// BaseConf 基础相关配置
	BaseConf *BaseConfig
	// ServerConf server相关配置
	ServerConf *ServerConfig
	// AppConf app相关配置
	AppConf *AppConfig
	// DBConf database相关配置
	DBConf *DBConfig
)

// BaseConfig conf/app.ini中的基础配置信息
type BaseConfig struct {
	RunMode string
}

// ServerConfig conf/app.ini中的server配置
type ServerConfig struct {
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// AppConfig conf/app.ini中的app配置
type AppConfig struct {
	PageSize  int
	JwtSecret string
}

// DBConfig conf/app.ini中的database配置
type DBConfig struct {
	DBType      string
	DBName      string
	UserName    string
	Password    string
	DBHost      string
	TablePrefix string
}

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadDB()
}

// LoadBase LoadBase
func LoadBase() {
	BaseConf = &BaseConfig{
		RunMode: Cfg.Section("").Key("RUN_MODE").MustString("debug"),
	}
}

// LoadServer LoadServer
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	ServerConf = &ServerConfig{
		HTTPPort:     sec.Key("HTTP_PORT").MustInt(8000),
		ReadTimeout:  time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second,
		WriteTimeout: time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second,
	}
}

// LoadApp LoadApp
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	AppConf = &AppConfig{
		JwtSecret: sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)"),
		PageSize:  sec.Key("PAGE_SIZE").MustInt(10),
	}
}

// LoadDB LoadDB
func LoadDB() {
	sec, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	DBConf = &DBConfig{
		DBType:      sec.Key("TYPE").String(),
		DBName:      sec.Key("NAME").String(),
		UserName:    sec.Key("USER").String(),
		Password:    sec.Key("PASSWORD").String(),
		DBHost:      sec.Key("HOST").String(),
		TablePrefix: sec.Key("TABLE_PREFIX").String(),
	}
}
