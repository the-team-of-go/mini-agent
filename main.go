package main

import (
	"agent/internal/pkg/core"
	"agent/internal/pkg/setting"
	"agent/service"
	"fmt"
	"log"
	"os"
	"strconv"
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

	if len(os.Args) > 1 {
		machineId, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Panic("bad machine id: ", err)
			return
		}
		setting.ReportSetting.MachineId = int32(machineId)
		log.Printf("set machine id to: %v", machineId)
	}

	wg.Add(1)
	go func() {
		for {
			info := core.GetMachineInfo()
			err := service.ReportOnce(info)
			if err != nil {
				log.Println("report failed", err)
			}

			// 每隔Duration读取一次机器信息
			time.Sleep(time.Duration(setting.ReportSetting.Duration) * time.Second)
		}
	}()
	//go service.StartServer()
	wg.Wait()
}
