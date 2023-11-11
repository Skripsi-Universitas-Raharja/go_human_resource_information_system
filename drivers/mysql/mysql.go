package mysql_driver

import (
	"backend-golang/drivers/mysql/profiles"
	stockhistory "backend-golang/drivers/mysql/stock_history"
	stockins "backend-golang/drivers/mysql/stock_ins"
	stockouts "backend-golang/drivers/mysql/stock_outs"

	"backend-golang/drivers/mysql/stocks"
	"backend-golang/drivers/mysql/users"

	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func (config *DBConfig) InitDB() *gorm.DB {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when creating a connection to the database: %s\n", err)
	}

	log.Println("connected to the database")

	return db
}

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&users.User{}, &profiles.Profile{}, &stocks.Stock{}, &stockins.StockIn{}, &stockouts.StockOut{}, &stockhistory.StockHistory{})

	if err != nil {
		log.Fatalf("failed to perform database migration: %s\n", err)
	}
}

func CloseDB(db *gorm.DB) error {
	database, err := db.DB()

	if err != nil {
		log.Printf("error when getting the database instance: %v", err)
		return err
	}

	if err := database.Close(); err != nil {
		log.Printf("error when closing the database connection: %v", err)
		return err
	}

	log.Println("database connection is closed")

	return nil
}
