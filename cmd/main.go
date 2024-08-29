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

	pb "sale/genproto/sale"

	"google.golang.org/grpc"
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
	serviceWishlist := service.NewWishlistService(logger, storage)
	serviceFeedback := service.NewFeedbackService(logger, storage)
	serviceBought := service.NewBoughtService(logger, storage)
	serviceProcess := service.NewProcessService(logger, storage)

	s := grpc.NewServer()

	pb.RegisterProductServer(s, ServiceProduct)
	pb.RegisterWishlistServer(s, serviceWishlist)
	pb.RegisterFeedbackServer(s, serviceFeedback)
	pb.RegisterBoughtServer(s, serviceBought)
	pb.RegisterProcessServer(s, serviceProcess)

	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("error while serving: %v", err)
	}

}
