package main

import (
	"github.com/henriqueholtz/fullcycle-7/application/grpc"
	"github.com/henriqueholtz/fullcycle-7/infraestructure/db"
	"github.com/jinzhu/gorm"
	"os"
)

var database *gorm.DB

func main() {
	print("Starting... ")
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
