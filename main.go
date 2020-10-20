package main

import (
	"demo_go/logger"
	//_ "demo_go/routers"
	//_ "github.com/lib/pq"
	"sync"
)

func init() {
	// 启动http server
	//beego.BConfig.AppName = config.AppConf.AppName
	//beego.BConfig.RunMode = config.AppConf.RunMode
	//beego.BConfig.CopyRequestBody = config.AppConf.CopyRequestBody
	//beego.BConfig.Listen.HTTPPort = config.AppConf.HttpPort
	//beego.BConfig.WebConfig.EnableDocs = config.AppConf.EnableDocs
	//beego.BConfig.WebConfig.AutoRender = config.AppConf.AutoRender
	//models.InitDB()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		//grpcStart()
		logger.Debug("grpc server start ")
		wg.Done()
	}()
	wg.Wait()
	logger.Debug("ser done")
}

//func grpcStart() {
//	lis, err := net.Listen("tcp", config.AppConf.GrpcListen)
//	if err != nil {
//		panic(err)
//	}
//	s := grpc.NewServer()
//	logger.Debugf("lis:%v grpc-port(%v) s:%v", lis,  config.AppConf.GrpcListen, s)
//	pb.RegisterDeveloperServiceServer(s, &controllers.DeveloperServiceController{})
//	reflection.Register(s)
//	if err := s.Serve(lis); err != nil {
//		panic(err)
//	}
//}