package web

import (
	"demo_go/controller"
	"github.com/gin-gonic/gin"
)

func GinStart()  {
	r := gin.Default()
	r.GET("/ping", controller.PingHandler)
	r.GET("/pool", controller.PoolHandler)
	r.Run()
}


//func GrpcStart() {
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

//func init() {
// bee run
	// 启动http server
	//beego.BConfig.AppName = config.AppConf.AppName
	//beego.BConfig.RunMode = config.AppConf.RunMode
	//beego.BConfig.CopyRequestBody = config.AppConf.CopyRequestBody
	//beego.BConfig.Listen.HTTPPort = config.AppConf.HttpPort
	//beego.BConfig.WebConfig.EnableDocs = config.AppConf.EnableDocs
	//beego.BConfig.WebConfig.AutoRender = config.AppConf.AutoRender
//}

