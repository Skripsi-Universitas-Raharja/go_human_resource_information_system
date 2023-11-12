package stockouts

import (
	stockouts "backend-golang/businesses/stock_outs"
	"context"
	"fmt"

	_dbStocks "backend-golang/drivers/mysql/stocks"

	"gorm.io/gorm"
)

type stockOutRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) stockouts.Repository {
	return &stockOutRepository{
		conn: conn,
	}
}

func (ur *stockOutRepository) GetByID(ctx context.Context, id string) (stockouts.Domain, error) {
	var stockout StockOut

	if err := ur.conn.WithContext(ctx).First(&stockout, "id = ?", id).Error; err != nil {
		return stockouts.Domain{}, err
	}

	return stockout.ToDomain(), nil

}

func (sr *stockOutRepository) StockOut(ctx context.Context, stockOutDomain *stockouts.Domain) (stockouts.Domain, error) {
	record := FromDomain(stockOutDomain)

	// Buat rekam stok pada riwayat
	result := sr.conn.WithContext(ctx).Create(&record)
	if err := result.Error; err != nil {
		return stockouts.Domain{}, err
	}

	// Perbarui total stok di tabel stok
	var stock _dbStocks.Stock

	if err := sr.conn.WithContext(ctx).First(&stock, "id = ?", record.StockID).Error; err != nil {
		return stockouts.Domain{}, err
	}

	stock.Stock_Total -= record.Stock_Out
	fmt.Println("StockTotal after addition:", stock.Stock_Total)

	// Simpan total stok yang diperbarui di tabel stok
	if err := sr.conn.WithContext(ctx).Save(&stock).Error; err != nil {
		return stockouts.Domain{}, err
	}

	// Perbarui total stok pada rekam riwayat stok
	record.Stock_Location = stock.Stock_Location
	record.Stock_Code = stock.Stock_Code
	record.Stock_Name = stock.Stock_Name
	record.Stock_Unit = stock.Unit
	// record.StockInTransactions.Stock_Total = stock.Stock_Total
	record.Stock_Total = stock.Stock_Total

	fmt.Println("Record StockTotal:", record.Stock_Total)

	// Simpan total stok yang diperbarui di tabel riwayat stok in
	if err := sr.conn.WithContext(ctx).Save(&record).Error; err != nil {
		return stockouts.Domain{}, err
	}

	return record.ToDomain(), nil

}
