package setting

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	Cfg *ini.File

	RunMode string
	MysqlHost string
	MysqlUser string
	MysqlPassword string
	MysqlDb string

	JwtSecret string

	Port int
	ReadTimeout int
	WriteTimeout int
)

func init()  {
	var err error
	Cfg, err = ini.Load("conf/api.conf")
	if err != nil{
		log.Println(err)
	}

	LoadSql()
	LoadServer()
	LoadMode()
	LoadJwt()
}

func LoadSql()  {
	sec, _ := Cfg.GetSection("mysql")
	MysqlHost = sec.Key("host").MustString("")
	MysqlUser = sec.Key("user").MustString("")
	MysqlPassword = sec.Key("password").MustString("")
	MysqlDb = sec.Key("db").MustString("")
}

func LoadServer()  {
	sec, _ := Cfg.GetSection("server")
	Port = sec.Key("port").MustInt()
	ReadTimeout = sec.Key("port").MustInt()
	WriteTimeout = sec.Key("port").MustInt()
}

func LoadMode()  {
	RunMode = Cfg.Section("").Key("run_mode").String()
}

func LoadJwt()  {
	JwtSecret = Cfg.Section("api").Key("jwt_secret").String()
}