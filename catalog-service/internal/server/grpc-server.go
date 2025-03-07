package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"hdck007.com/order-service/internal/dto"
	"hdck007.com/order-service/internal/services"
	catalogservice "hdck007.com/order-service/pb/catalog-service"
)

type CatalogGRPCServer struct {
	catalogservice.UnimplementedCatalogServiceServer
	restaurantService *services.RestaurantService
	menuItemService   *services.MenuItemService
}

func NewCatalogGRPCServer(restaurantService *services.RestaurantService, menuItemService *services.MenuItemService) *CatalogGRPCServer {
	return &CatalogGRPCServer{
		restaurantService: restaurantService,
		menuItemService:   menuItemService,
	}
}

func (s *CatalogGRPCServer) CreateRestaurant(ctx context.Context, req *catalogservice.CreateRestaurantRequest) (*catalogservice.Restaurant, error) {
	restaurant := s.restaurantService.CreateRestaurant(dto.CreateRestaurantRequest{
		Name:      req.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	})

	return &catalogservice.Restaurant{
		Id:        uint64(restaurant.ID),
		Name:      restaurant.Name,
		Latitude:  restaurant.Latitude,
		Longitude: restaurant.Longitude,
		MenuItems: []*catalogservice.MenuItem{},
	}, nil
}

func (s *CatalogGRPCServer) GetRestaurantById(ctx context.Context, req *catalogservice.GetRestaurantRequest) (*catalogservice.Restaurant, error) {
	restaurant := s.restaurantService.GetRestaurantById(uint(req.Id))
	if restaurant == nil {
		return nil, nil
	}

	menuItems := s.menuItemService.GetMenuItemsByRestaurantId(uint(req.Id))
	grpcMenuItems := []*catalogservice.MenuItem{}

	for _, item := range menuItems {
		grpcMenuItems = append(grpcMenuItems, &catalogservice.MenuItem{
			Id:           uint64(item.ID),
			Name:         item.Name,
			RestaurantId: uint64(item.RestaurantId),
			Price:        item.Price,
			Stock:        int32(item.Stock),
		})
	}

	return &catalogservice.Restaurant{
		Id:        uint64(restaurant.ID),
		Name:      restaurant.Name,
		Latitude:  restaurant.Latitude,
		Longitude: restaurant.Longitude,
		MenuItems: grpcMenuItems,
	}, nil
}

func (s *CatalogGRPCServer) GetRestaurantsByName(ctx context.Context, req *catalogservice.GetRestaurantsByNameRequest) (*catalogservice.GetRestaurantsResponse, error) {
	restaurants := s.restaurantService.GetRestaurantsByName(req.Name)
	grpcRestaurants := []*catalogservice.Restaurant{}

	for _, restaurant := range restaurants {
		grpcRestaurants = append(grpcRestaurants, &catalogservice.Restaurant{
			Id:        uint64(restaurant.ID),
			Name:      restaurant.Name,
			Latitude:  restaurant.Latitude,
			Longitude: restaurant.Longitude,
			MenuItems: []*catalogservice.MenuItem{},
		})
	}

	return &catalogservice.GetRestaurantsResponse{
		Restaurants: grpcRestaurants,
	}, nil
}

func (s *CatalogGRPCServer) GetRestaurants(ctx context.Context, req *emptypb.Empty) (*catalogservice.GetRestaurantsResponse, error) {
	restaurants := s.restaurantService.GetRestaurants()
	grpcRestaurants := []*catalogservice.Restaurant{}

	for _, restaurant := range restaurants {
		grpcRestaurants = append(grpcRestaurants, &catalogservice.Restaurant{
			Id:        uint64(restaurant.ID),
			Name:      restaurant.Name,
			Latitude:  restaurant.Latitude,
			Longitude: restaurant.Longitude,
			MenuItems: []*catalogservice.MenuItem{},
		})
	}

	return &catalogservice.GetRestaurantsResponse{
		Restaurants: grpcRestaurants,
	}, nil
}

func (s *CatalogGRPCServer) CreateMenuItem(ctx context.Context, req *catalogservice.CreateMenuItemRequest) (*catalogservice.MenuItem, error) {
	menuItem := s.menuItemService.CreateMenuItem(dto.CreateMenuItemRequest{
		Name:         req.Name,
		RestaurantId: uint(req.RestaurantId),
		Price:        req.Price,
		Stock:        int(req.Stock),
	})

	return &catalogservice.MenuItem{
		Id:           uint64(menuItem.ID),
		Name:         menuItem.Name,
		RestaurantId: uint64(menuItem.RestaurantId),
		Price:        menuItem.Price,
		Stock:        int32(menuItem.Stock),
	}, nil
}

func (s *CatalogGRPCServer) GetMenuItemById(ctx context.Context, req *catalogservice.GetMenuItemRequest) (*catalogservice.MenuItem, error) {
	menuItem := s.menuItemService.GetMenuItemById(uint(req.Id))
	if menuItem == nil {
		return nil, nil
	}
	return &catalogservice.MenuItem{
		Id:           uint64(menuItem.ID),
		Name:         menuItem.Name,
		RestaurantId: uint64(menuItem.RestaurantId),
		Price:        menuItem.Price,
		Stock:        int32(menuItem.Stock),
	}, nil
}

func (s *CatalogGRPCServer) GetMenuItemsByRestaurantId(ctx context.Context, req *catalogservice.GetMenuItemsByRestaurantIdRequest) (*catalogservice.GetMenuItemsResponse, error) {
	menuItems := s.menuItemService.GetMenuItemsByRestaurantId(uint(req.RestaurantId))
	grpcMenuItems := []*catalogservice.MenuItem{}

	for _, item := range menuItems {
		grpcMenuItems = append(grpcMenuItems, &catalogservice.MenuItem{
			Id:           uint64(item.ID),
			Name:         item.Name,
			RestaurantId: uint64(item.RestaurantId),
			Price:        item.Price,
			Stock:        int32(item.Stock),
		})
	}

	return &catalogservice.GetMenuItemsResponse{
		MenuItems: grpcMenuItems,
	}, nil
}
