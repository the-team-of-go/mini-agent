# agent
## 上报协议
```proto
message ReportReq {
  int32  MachineId = 1;
  double CpuPercent = 2;
  double MemPercent = 3;
  double DiskPercent = 4;
  int64  TimeStamp = 5;
}

message ReportRsp {
  int32 code = 1;
  string msg = 2;
}

service Report {
  rpc Send(ReportReq) returns (ReportRsp) {}
}
```

目前上报频率为`10s`，可手动配置
