package stocks

import (
	"backend-golang/businesses/stocks"
	"context"

	"gorm.io/gorm"
)

type stockRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) stocks.Repository {
	return &stockRepository{
		conn: conn,
	}
}

func (ur *stockRepository) GetByID(ctx context.Context, id string) (stocks.Domain, error) {
	var stock Stock

	if err := ur.conn.WithContext(ctx).First(&stock, "id = ?", id).Error; err != nil {
		return stocks.Domain{}, err
	}

	return stock.ToDomain(), nil

}

func (cr *stockRepository) Create(ctx context.Context, stockDomain *stocks.Domain) (stocks.Domain, error) {
	record := FromDomain(stockDomain)

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return stocks.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return stocks.Domain{}, err
	}

	return record.ToDomain(), nil

}

func (cr *stockRepository) StockIn(ctx context.Context, stockDomain *stocks.Domain, id string) (stocks.Domain, error) {
	stock, err := cr.GetByID(ctx, id)

	if err != nil {
		return stocks.Domain{}, err
	}

	updateStock := FromDomain(&stock)

	updateStock.Stock_Total += stockDomain.Stock_In

	if err := cr.conn.WithContext(ctx).Save(&updateStock).Error; err != nil {
		return stocks.Domain{}, err
	}

	return updateStock.ToDomain(), nil
}

func (cr *stockRepository) StockOut(ctx context.Context, stockDomain *stocks.Domain, id string) (stocks.Domain, error) {
	stock, err := cr.GetByID(ctx, id)

	if err != nil {
		return stocks.Domain{}, err
	}

	updateStock := FromDomain(&stock)

	updateStock.Stock_Total -= stockDomain.Stock_Out

	if err := cr.conn.WithContext(ctx).Save(&updateStock).Error; err != nil {
		return stocks.Domain{}, err
	}

	return updateStock.ToDomain(), nil
}
