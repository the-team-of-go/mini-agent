package core

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"time"
)

type MachineInfo struct {
	CpuPercent  float64 `json:"cpu-percent"`  // Cpu使用率
	MemPercent  float64 `json:"mem-percent"`  // 内存使用率
	DiskPercent float64 `json:"disk-percent"` // 磁盘使用率
	TimeStamp   int64   `json:"timestamp"`
}

func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	return percent[0]
}

func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}

// GetMachineInfo 获取机器cpu、内存使用率数据
func GetMachineInfo() *MachineInfo {
	info := &MachineInfo{}
	info.CpuPercent = GetCpuPercent()
	info.MemPercent = GetMemPercent()
	info.DiskPercent = GetDiskPercent()
	info.TimeStamp = time.Now().UnixMilli()
	return info
}
