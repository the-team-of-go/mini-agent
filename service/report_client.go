package service

import (
	"agent/internal/pkg/core"
	"agent/internal/pkg/setting"
	"agent/service/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

var (
	ReportConn unsafe.Pointer
	ReportLock sync.Mutex
)

func GetReportConn() (*grpc.ClientConn, error) {
	if atomic.LoadPointer(&ReportConn) != nil {
		return (*grpc.ClientConn)(ReportConn), nil
	}

	ReportLock.Lock()
	defer ReportLock.Unlock()

	if atomic.LoadPointer(&ReportConn) != nil {
		return (*grpc.ClientConn)(ReportConn), nil
	}

	conn, err := NewReportConn()
	if err != nil {
		return nil, err
	}

	atomic.StorePointer(&ReportConn, unsafe.Pointer(conn))
	return conn, nil
}

func NewReportConn() (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, net.JoinHostPort(setting.RpcSrverSetting.Host, setting.RpcSrverSetting.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func GetReportClient() (pb.ReportClient, error) {
	conn, err := GetReportConn()
	if err != nil {
		return nil, err
	}
	return pb.NewReportClient(conn), nil
}

func ReportOnce(info *core.MachineInfo) error {
	reportCli, err := GetReportClient()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := reportCli.Send(ctx, &pb.ReportReq{
		MachineId:   setting.ReportSetting.MachineId,
		CpuPercent:  info.CpuPercent,
		MemPercent:  info.MemPercent,
		DiskPercent: info.DiskPercent,
		TimeStamp:   info.TimeStamp,
	})
	if err != nil {
		return err
	}

	log.Println("send one record finished:", r.GetCode(), r.GetMsg())
	return nil
}
