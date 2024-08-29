package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sale/config"
	"sale/logs"
	"sale/service"
	"sale/storage"
	"sale/storage/mongosh"
	"sale/storage/postgres"
)

func main() {
	mdb, err := mongosh.Connect(context.Background())
	if err != nil {
		panic(err)
	}
	defer mdb.Client().Disconnect(context.Background())

	pdb, err := postgres.ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer pdb.Close()

	storage := storage.NewStoragePro(mdb, pdb)
	conf := config.Load()
	logger := logs.NewLogger()

	fmt.Println("Starting server...")
	lis, err := net.Listen("tcp", conf.Server.SALE_SERVICE)
	if err != nil {
		log.Fatalf("error while listening: %v", err)
	}
	defer lis.Close()

	ServiceProduct := service.NewProductService(logger, storage)

}
