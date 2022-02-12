package service

import (
	"agent/internal/pkg/setting"
	"agent/service/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedReportServer
}

func (s *server) Send(ctx context.Context, req *pb.ReportReq) (*pb.ReportRsp, error) {
	cpuPercent := req.GetCpuPercent()
	memPercent := req.GetMemPercent()
	diskPercent := req.GetDiskPercent()
	timestamp := req.GetTimeStamp()
	log.Printf("received one record. cpu: %v%%, mem: %v%%, disk: %v%%, timestamp: %v", cpuPercent, memPercent, diskPercent, timestamp)

	rsp := &pb.ReportRsp{}
	rsp.Code = 200
	rsp.Msg = "OK"
	return rsp, nil
}

func StartServer() {
	port := fmt.Sprintf(":%v", setting.RpcSrverSetting.Port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterReportServer(s, &server{})
	log.Printf("server is listerning at %v...", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
