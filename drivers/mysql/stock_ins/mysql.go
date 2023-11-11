package stockins

import (
	stockins "backend-golang/businesses/stock_ins"
	// "backend-golang/businesses/stocks"
	"context"
	"fmt"

	_dbStocks "backend-golang/drivers/mysql/stocks"

	"gorm.io/gorm"
)

type stockInRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) stockins.Repository {
	return &stockInRepository{
		conn: conn,
	}
}

func (sr *stockInRepository) Create(ctx context.Context, stockInDomain *stockins.Domain) (stockins.Domain, error) {
	record := FromDomain(stockInDomain)

	result := sr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return stockins.Domain{}, err
	}

	err := sr.conn.WithContext(ctx).Last(&record).Error

	if err != nil {
		return stockins.Domain{}, err
	}

	return record.ToDomain(), nil
}

// func (sr *stockInRepository) StockIn(ctx context.Context, stockDomain *stocks.Domain) (stocks.Domain, error) {
// 	record := _dbStocks.FromDomain(stockDomain)

// 	// Buat rekam stok pada riwayat
// 	result := sr.conn.WithContext(ctx).Create(&record)
// 	if err := result.Error; err != nil {
// 		return stocks.Domain{}, err
// 	}

// 	// Perbarui total stok di tabel stok
// 	var stockIn StockIn
// 	if err := sr.conn.WithContext(ctx).First(&stockIn, "id = ?", record.ID).Error; err != nil {
// 		return stocks.Domain{}, err
// 	}

// 	record.Stock_Total += stockIn.Stock_In
// 	fmt.Println("StockTotal after addition:", stockIn.Stock_Total)

// 	// Simpan total stok yang diperbarui di tabel stok
// 	if err := sr.conn.WithContext(ctx).Save(&stockIn).Error; err != nil {
// 		return stocks.Domain{}, err
// 	}

// 	// Perbarui total stok pada rekam riwayat stok
// 	record.Stock_Total = stockIn.Stock_Total
// 	fmt.Println("Record StockTotal:", record.Stock_Total)

// 	// Simpan total stok yang diperbarui di tabel riwayat stok
// 	if err := sr.conn.WithContext(ctx).Save(&record).Error; err != nil {
// 		return stocks.Domain{}, err
// 	}

// 	return record.ToDomain(), nil
// }

func (sr *stockInRepository) StockIn(ctx context.Context, stockInDomain *stockins.Domain) (stockins.Domain, error) {
	record := FromDomain(stockInDomain)

	// Buat rekam stok pada riwayat
	result := sr.conn.WithContext(ctx).Create(&record)
	if err := result.Error; err != nil {
		return stockins.Domain{}, err
	}

	// Perbarui total stok di tabel stok
	var stock _dbStocks.Stock

	if err := sr.conn.WithContext(ctx).First(&stock, "id = ?", record.StockID).Error; err != nil {
		return stockins.Domain{}, err
	}

	stock.Stock_Total += record.Stock_In
	fmt.Println("StockTotal after addition:", stock.Stock_Total)

	// Simpan total stok yang diperbarui di tabel stok
	if err := sr.conn.WithContext(ctx).Save(&stock).Error; err != nil {
		return stockins.Domain{}, err
	}

	// Perbarui total stok pada rekam riwayat stok
	record.Stock_Location = stock.Stock_Location
	record.Stock_Code = stock.Stock_Code
	record.Stock_Name = stock.Stock_Name
	record.Stock_Unit = stock.Unit
	record.Stock_Total = stock.Stock_Total
	fmt.Println("Record StockTotal:", record.Stock_Total)

	// Simpan total stok yang diperbarui di tabel riwayat stok
	if err := sr.conn.WithContext(ctx).Save(&record).Error; err != nil {
		return stockins.Domain{}, err
	}

	return record.ToDomain(), nil
}
