package web


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

