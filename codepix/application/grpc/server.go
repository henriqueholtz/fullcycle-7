package grpc

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/reflection"
	// "github.com/henriqueholtz/fullcycle-7/application/grpc/pb"
	"log"
	"net"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Can't start grpc server, err")
	}
	log.Printf("gRPC server has been started on port %d", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Can't start grpc server (err 2)", err)
	}
}
