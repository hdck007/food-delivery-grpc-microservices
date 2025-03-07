package services

import (
	"hdck007.com/order-service/internal/dto"
	"hdck007.com/order-service/internal/models"
	"hdck007.com/order-service/internal/repositories"
)

type RestaurantService struct {
	repository repositories.IRestaurantRepository
}

func NewRestaurantService(restaurantRepository repositories.IRestaurantRepository) *RestaurantService {
	return &RestaurantService{
		repository: restaurantRepository,
	}
}

func (restaurantService *RestaurantService) CreateRestaurant(createRestaurantQuery dto.CreateRestaurantRequest) *models.Restaurant {
	restaurant := &models.Restaurant{
		Name: createRestaurantQuery.Name,
	}

	restaurantService.repository.Create(restaurant)

	return restaurant
}

func (restaurantService *RestaurantService) GetRestaurantById(restaurantId uint) *models.Restaurant {
	return restaurantService.repository.GetRestaurantById(restaurantId)
}

func (restaurantService *RestaurantService) GetRestaurantsByName(name string) []models.Restaurant {
	return restaurantService.repository.GetRestaurantsByName(name)
}

func (restaurantService *RestaurantService) GetRestaurants() []models.Restaurant {
	return restaurantService.repository.GetRestaurants()
}
