package server

import (
	"context"

	"hdck007.com/order-service/internal/dto"
	"hdck007.com/order-service/internal/services"
	"hdck007.com/order-service/pb/delivery"
)

type DeliveryGRPCServer struct {
	delivery.UnimplementedDeliveryServiceServer
	deliveryService *services.DeliveryService
}

func NewDeliveryGRPCServer(deliveryService *services.DeliveryService) *DeliveryGRPCServer {
	return &DeliveryGRPCServer{
		deliveryService: deliveryService,
	}
}

func (s *DeliveryGRPCServer) CreateDelivery(ctx context.Context, req *delivery.CreateDeliveryRequest) (*delivery.CreateDeliveryResponse, error) {
	createdDelivery := s.deliveryService.CreateDelivery(
		dto.CreateDeliveryRequest{
			OrderId: uint(req.OrderId),
			RiderId: uint(req.RiderId),
		},
	)

	deliveryResponse := delivery.CreateDeliveryResponse{
		Delivery: &delivery.Delivery{
			Id:      uint32(createdDelivery.ID),
			OrderId: uint32(createdDelivery.OrderId),
			RiderId: uint32(createdDelivery.RiderId),
			Status:  createdDelivery.Status,
		},
	}

	return &deliveryResponse, nil
}

func (s *DeliveryGRPCServer) UpdateDeliveryStatus(ctx context.Context, req *delivery.UpdateDeliveryStatusRequest) (*delivery.UpdateDeliveryStatusResponse, error) {
	updatedDelivery := s.deliveryService.UpdateDeliveryStatus(
		uint(req.DeliveryId),
		req.DeliveryStatus,
	)

	deliveryResponse := delivery.UpdateDeliveryStatusResponse{
		Delivery: &delivery.Delivery{
			Id:      uint32(updatedDelivery.ID),
			OrderId: uint32(updatedDelivery.OrderId),
			RiderId: uint32(updatedDelivery.RiderId),
			Status:  updatedDelivery.Status,
		},
	}

	return &deliveryResponse, nil
}
