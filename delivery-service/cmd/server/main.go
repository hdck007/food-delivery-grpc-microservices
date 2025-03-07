package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"hdck007.com/order-service/internal/database"
	"hdck007.com/order-service/internal/repositories"
	"hdck007.com/order-service/internal/server"
	"hdck007.com/order-service/internal/services"
	"hdck007.com/order-service/pb/delivery"
)

func Start() {
	dbEngine, err := database.InitDatabaseEngine()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	deliveryRepository := repositories.NewDeliveryRepository(dbEngine)
	deliveryService := services.NewDeliveryService(deliveryRepository)

	grpcServer := server.NewDeliveryGRPCServer(deliveryService)

	s := grpc.NewServer()
	delivery.RegisterDeliveryServiceServer(s, grpcServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
