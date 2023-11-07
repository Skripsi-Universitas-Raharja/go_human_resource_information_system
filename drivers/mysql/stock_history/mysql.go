package stockhistory

import (
	stockhistory "backend-golang/businesses/stock_history"
	"context"

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
