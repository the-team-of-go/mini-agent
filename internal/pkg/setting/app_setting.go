package setting

import (
	"gopkg.in/ini.v1"
	"log"
)

type RpcServer struct {
	Host string
	Port string
}

var RpcSrverSetting = &RpcServer{}

type Report struct {
	MachineId int32
	Duration  int
}

var ReportSetting = &Report{}

var cfg *ini.File

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalln(section, "项目配置注入失败！", err)
	}
}

func SetUp() {
	// 读取项目配置
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalln("读取项目配置失败")
	}
	mapTo("rpc-server", &RpcSrverSetting)
	mapTo("report", &ReportSetting)
}
