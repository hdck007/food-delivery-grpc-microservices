package repositories

import (
	"hdck007.com/order-service/internal/database"
	"hdck007.com/order-service/internal/models"
)

type RestaurantRepository struct {
	databaseEngine *database.DatabaseEngine
}

func NewRestaurantRepository(databaseEngine *database.DatabaseEngine) IRestaurantRepository {
	return &RestaurantRepository{
		databaseEngine: databaseEngine,
	}
}

func (restaurantRepository *RestaurantRepository) Create(restaurant *models.Restaurant) {
	restaurantRepository.databaseEngine.Db.Create(&restaurant)
}

func (restaurantRepository *RestaurantRepository) GetRestaurants() []models.Restaurant {
	var restaurants []models.Restaurant
	restaurantRepository.databaseEngine.Db.Find(&restaurants)
	return restaurants
}

func (restaurantRepository *RestaurantRepository) GetRestaurantById(restaurantId uint) *models.Restaurant {
	var restaurant models.Restaurant
	restaurantRepository.databaseEngine.Db.Preload("MenuItems").Find(&restaurant).Where(&models.Restaurant{ID: restaurantId})
	return &restaurant
}

func (restaurantRepository *RestaurantRepository) GetRestaurantsByName(name string) []models.Restaurant {
	var restaurants []models.Restaurant
	restaurantRepository.databaseEngine.Db.Preload("MenuItems").Find(&restaurants).Where("name ilike ?", "%"+name+"%")
	return restaurants
}
