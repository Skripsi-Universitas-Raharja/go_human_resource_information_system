package stockins

import (
	stockins "backend-golang/businesses/stock_ins"
	"context"
	"errors"
	"fmt"

	_dbStockHistory "backend-golang/drivers/mysql/stock_history"
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

func (sr *stockInRepository) GetAll(ctx context.Context) ([]stockins.Domain, error) {
	var records []StockIn
	if err := sr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []stockins.Domain{}

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

	// Cek apakah catatan sudah ada di tabel StockHistory
	var stockHistory _dbStockHistory.StockHistory
	err := sr.conn.WithContext(ctx).First(&stockHistory, "stock_id = ?", record.StockID).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return stockins.Domain{}, err
	}

	// Perbarui total stok pada rekam riwayat stok
	record.Stock_Location = stock.Stock_Location
	record.Stock_Code = stock.Stock_Code
	record.Stock_Name = stock.Stock_Name
	record.Stock_Unit = stock.Stock_Unit
	record.Stock_Total = stock.Stock_Total

	fmt.Println("Record StockTotal:", record.Stock_Total)

	// Buat baru jika tidak ditemukan
	stockHistory = _dbStockHistory.StockHistory{
		// Inisialisasi nilai-nilai yang diperlukan
		StockID:        record.StockID,
		Stock_Location: record.Stock_Location,
		Stock_Code:     record.Stock_Code,
		Stock_Name:     record.Stock_Name,
		Stock_Unit:     record.Stock_Unit,
		Stock_In:       record.Stock_In,
		Stock_Total:    record.Stock_Total,
		// Tambahkan nilai-nilai lain yang diperlukan
	}

	if err := sr.conn.WithContext(ctx).Save(&stockHistory).Error; err != nil {
		return stockins.Domain{}, err
	}

	// Simpan total stok yang diperbarui di tabel riwayat stok in
	if err := sr.conn.WithContext(ctx).Save(&record).Error; err != nil {
		return stockins.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (sr *stockInRepository) ExportToExcel(ctx context.Context) ([]stockins.Domain, error) {
	var records []StockIn
	if err := sr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []stockins.Domain{}

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
