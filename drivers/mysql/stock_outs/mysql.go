package stockouts

import (
	stockouts "backend-golang/businesses/stock_outs"
	"context"
	"errors"
	"fmt"

	_dbStockHistory "backend-golang/drivers/mysql/stock_history"
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

func (sr *stockOutRepository) GetAll(ctx context.Context) ([]stockouts.Domain, error) {
	var records []StockOut
	if err := sr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []stockouts.Domain{}

	for _, category := range records {
		// Dapatkan Stock dari tabel stok menggunakan StockID
		var stock _dbStocks.Stock
		if err := sr.conn.WithContext(ctx).First(&stock, "id = ?", category.StockID).Error; err != nil {
			return nil, err
		}

		domain := category.ToDomain()
		// Set Stock_Code dari Stock ke Domain
		domain.Stock_Code = stock.Stock_Code
		categories = append(categories, domain)
	}

	return categories, nil
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
	record.Stock_Unit = stock.Stock_Unit
	// record.StockInTransactions.Stock_Total = stock.Stock_Total
	record.Stock_Total = stock.Stock_Total

	fmt.Println("Record StockTotal:", record.Stock_Total)
	fmt.Println("Record StockCode:", record.Stock_Code)
	fmt.Println("StockCode:", stock.Stock_Code)

	// Cek apakah catatan sudah ada di tabel StockHistory
	var stockHistory _dbStockHistory.StockHistory
	err := sr.conn.WithContext(ctx).First(&stockHistory, "stock_id = ?", record.StockID).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return stockouts.Domain{}, err
	}

	// Buat baru jika tidak ditemukan
	stockHistory = _dbStockHistory.StockHistory{
		// Inisialisasi nilai-nilai yang diperlukan
		StockID:        record.StockID,
		Stock_Location: record.Stock_Location,
		Stock_Code:     record.Stock_Code,
		Stock_Name:     record.Stock_Name,
		Stock_Unit:     record.Stock_Unit,
		Stock_Out:      record.Stock_Out,
		Stock_Total:    record.Stock_Total,
		// Tambahkan nilai-nilai lain yang diperlukan
	}

	if err := sr.conn.WithContext(ctx).Save(&stockHistory).Error; err != nil {
		return stockouts.Domain{}, err
	}

	// Simpan total stok yang diperbarui di tabel riwayat stok in
	if err := sr.conn.WithContext(ctx).Save(&record).Error; err != nil {
		return stockouts.Domain{}, err
	}

	return record.ToDomain(), nil

}

func (sr *stockOutRepository) ExportToExcel(ctx context.Context) ([]stockouts.Domain, error) {
	var records []StockOut
	if err := sr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []stockouts.Domain{}

	for _, category := range records {
		// Dapatkan Stock dari tabel stok menggunakan StockID
		var stock _dbStocks.Stock
		if err := sr.conn.WithContext(ctx).First(&stock, "id = ?", category.StockID).Error; err != nil {
			return nil, err
		}

		domain := category.ToDomain()
		// Set Stock_Code dari Stock ke Domain
		domain.Stock_Code = stock.Stock_Code
		categories = append(categories, domain)
	}

	return categories, nil
}
