package stockouts

import (
	stockouts "backend-golang/businesses/stock_outs"
	"context"

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

func (cr *stockOutRepository) StockOut(ctx context.Context, stockOutDomain *stockouts.Domain) (stockouts.Domain, error) {
	record := FromDomain(stockOutDomain)

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return stockouts.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return stockouts.Domain{}, err
	}

	return record.ToDomain(), nil

}
