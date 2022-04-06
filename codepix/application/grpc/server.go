package grpc

import (
	"fmt"
	"github.com/henriqueholtz/fullcycle-7/application/grpc/pb"
	"github.com/henriqueholtz/fullcycle-7/application/usecase"
	"github.com/henriqueholtz/fullcycle-7/infraestructure/repository"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: pixRepository}
	pixGrpcService := NewPixGrpcService(pixUseCase)

	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

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

func NewPixGrpcService(usecase usecase.PixUseCase) *PixGrpcService {
	return &PixGrpcService{
		PixUseCase: usecase.PixUseCase{},
	}
}
