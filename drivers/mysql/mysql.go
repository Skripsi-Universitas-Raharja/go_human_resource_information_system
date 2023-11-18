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

func SeedStocketail(db *gorm.DB) error {
	stocksData := []stocks.Stock{
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", Stock_Total: 365},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "Fiber Optik", Stock_Unit: "dus", Stock_Total: 365},
	}

	var record stocks.Stock
	_ = db.First(&record)

	if record.ID != 0 {
		log.Printf("stock detail already exists\n")
	} else {
		for _, stock := range stocksData {
			result := db.Create(&stock)
			if result.Error != nil {
				return result.Error
			}
		}
		log.Printf("%d stock detail created\n", len(stocksData))
	}

	return nil
}

func SeedStockOutsDetail(db *gorm.DB) {
	stockOutsData := []stockouts.StockOut{
		// type StockOut struct {
		// 	ID                  uint           `json:"id" gorm:"primaryKey"`
		// 	CreatedAt           time.Time      `json:"created_at"`
		// 	DeletedAt           gorm.DeletedAt `json:"deleted_at"`
		// 	Stock_Location      string         `json:"stock_location"`
		// 	Stock_Code          string         `json:"stock_code"`
		// 	Stock_Name          string         `json:"stock_name"`
		// 	Stock_Unit          string         `json:"stock_unit"`
		// 	Stock_Out           int            `json:"stock_out"`
		// 	Stock_Total         int            `json:"stock_total,omitempty"`
		// 	StockInTransactions stocks.Stock   `json:"-" gorm:"foreignKey:StockID"`
		// 	StockID             uint           `json:"stock_id"`
		// }
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
		{Stock_Location: "Tangerang", Stock_Code: "12345", Stock_Name: "RJ 45", Stock_Unit: "dus", StockID: 2, Stock_Out: 1},
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		// Panggil fungsi SeedStock di dalam transaksi
		if err := SeedStocketail(tx); err != nil {
			return err
		}

		var record stockouts.StockOut
		_ = db.First(&record)

		if record.ID != 0 {
			log.Printf("stock detail already exists\n")
		} else {
			for _, stock := range stockOutsData {
				result := db.Create(&stock)
				if result.Error != nil {
					return result.Error
				}
			}
			log.Printf("%d stock detail created\n", len(stockOutsData))
		}

		return nil
	})
	if err != nil {
		log.Fatalf("failed to seed stock detail: %s\n", err.Error())
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
