package drivers

import (
	userDomain "backend-golang/businesses/users"
	userDB "backend-golang/drivers/mysql/users"

	profileDomain "backend-golang/businesses/profiles"
	profileDB "backend-golang/drivers/mysql/profiles"

	stockDomain "backend-golang/businesses/stocks"
	stockDB "backend-golang/drivers/mysql/stocks"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}

func NewProfileRepository(conn *gorm.DB) profileDomain.Repository {
	return profileDB.NewMySQLRepository(conn)
}

func NewStockRepository(conn *gorm.DB) stockDomain.Repository {
	return stockDB.NewMySQLRepository(conn)
}
