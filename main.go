package main

import (
	"agent/internal/pkg/core"
	"agent/internal/pkg/setting"
	"agent/service"
	"fmt"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func Init() {
	defer fmt.Println("settings read finished!")
	setting.SetUp()
}

func main() {
	Init()
	wg.Add(2)
	go func() {
		for {
			info := core.GetMachineInfo()
			service.ReportOnce(info)
			// 每隔Duration读取一次机器信息
			time.Sleep(time.Duration(setting.ReportSetting.Duration) * time.Second)
		}
	}()
	//go service.StartServer()
	wg.Wait()
}
