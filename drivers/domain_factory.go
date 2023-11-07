package drivers

import (
	userDomain "backend-golang/businesses/users"
	userDB "backend-golang/drivers/mysql/users"

	profileDomain "backend-golang/businesses/profiles"
	profileDB "backend-golang/drivers/mysql/profiles"

	stockDomain "backend-golang/businesses/stocks"
	stockDB "backend-golang/drivers/mysql/stocks"

	stockhistoryDomain "backend-golang/businesses/stock_history"
	stockhistoryDB "backend-golang/drivers/mysql/stock_history"

	stockoutDomain "backend-golang/businesses/stock_outs"
	stockoutDB "backend-golang/drivers/mysql/stock_outs"

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

func NewStockHistoryRepository(conn *gorm.DB) stockhistoryDomain.Repository {
	return stockhistoryDB.NewMySQLRepository(conn)
}

func NewStockOutRepository(conn *gorm.DB) stockoutDomain.Repository {
	return stockoutDB.NewMySQLRepository(conn)
}
