package stockhistory

import (
	stockhistory "backend-golang/businesses/stock_history"
	"context"

	_dbStocks "backend-golang/drivers/mysql/stocks"

	"gorm.io/gorm"
)

type stockHistoryRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) stockhistory.Repository {
	return &stockHistoryRepository{
		conn: conn,
	}
}

func (sr *stockHistoryRepository) GetAll(ctx context.Context) ([]stockhistory.Domain, error) {
	var records []StockHistory
	if err := sr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []stockhistory.Domain{}

	for _, category := range records {
		// Dapatkan Stock dari tabel stok menggunakan StockID
		var stock _dbStocks.Stock
		if err := sr.conn.WithContext(ctx).First(&stock, "id = ?", category.StockID).Error; err != nil {
			return nil, err
		}

		domain := category.ToDomain()
		// Set Stock_Code dari Stock ke Domain
		// domain.Stock_Code = stock.Stock_Code
		categories = append(categories, domain)
	}

	return categories, nil
}

func (ur *stockHistoryRepository) GetByID(ctx context.Context, id string) (stockhistory.Domain, error) {
	var stockhistorys StockHistory

	if err := ur.conn.WithContext(ctx).First(&stockhistorys, "id = ?", id).Error; err != nil {
		return stockhistory.Domain{}, err
	}

	return stockhistorys.ToDomain(), nil
}

func (cr *stockHistoryRepository) Create(ctx context.Context, stockHistoryDomain *stockhistory.Domain) (stockhistory.Domain, error) {
	record := FromDomain(stockHistoryDomain)

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return stockhistory.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return stockhistory.Domain{}, err
	}

	return record.ToDomain(), nil

}

func (sr *stockHistoryRepository) ExportToExcel(ctx context.Context) ([]stockhistory.Domain, error) {
	var records []StockHistory
	if err := sr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []stockhistory.Domain{}

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
