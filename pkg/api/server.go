package api

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/api/handler"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/pb"
	"google.golang.org/grpc"
)

type ServerHttp struct {
	Engine *gin.Engine
}

func NewGrpcServer(adminHandler *handler.AdminHandler, grpcPort string) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalln("Failed to listen to the GRPC Port", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServiceServer(grpcServer, adminHandler)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Could not serve the GRPC Server: ", err)
	}
}

func NewServerHttp(adminHandler *handler.AdminHandler) *ServerHttp {
	engine := gin.New()

	go NewGrpcServer(adminHandler, "8891")

	engine.Use(gin.Logger())
	return &ServerHttp{
		Engine: engine,
	}
}

func (ser *ServerHttp) Start() {
	ser.Engine.Run(":9999")
}
