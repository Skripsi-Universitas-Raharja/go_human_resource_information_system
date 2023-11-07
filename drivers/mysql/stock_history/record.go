package stockhistory

import (
	stockhistory "backend-golang/businesses/stock_history"
	// "backend-golang/drivers/mysql/stocks"

	"time"

	"gorm.io/gorm"
)

type StockHistory struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Stock_Location string         `json:"stock_location"`
	Stock_Name     string         `json:"stock_name"`
	Unit           string         `json:"unit"`
	Stock_Out      float64        `json:"stock_out"`
	Stock_Total    float64        `json:"stock_total"`
	StockID        uint           `json:"stock_id"`
}

func (rec *StockHistory) ToDomain() stockhistory.Domain {
	return stockhistory.Domain{
		ID:             rec.ID,
		CreatedAt:      rec.CreatedAt,
		UpdatedAt:      rec.UpdatedAt,
		DeletedAt:      rec.DeletedAt,
		Stock_Location: rec.Stock_Location,
		Stock_Name:     rec.Stock_Name,
		Unit:           rec.Unit,
		Stock_Out:      rec.Stock_Out,
		Stock_Total:    rec.Stock_Total,
		StockID:        rec.StockID,
	}
}

func FromDomain(domain *stockhistory.Domain) *StockHistory {
	return &StockHistory{
		ID:             domain.ID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
		Stock_Location: domain.Stock_Location,
		Stock_Name:     domain.Stock_Name,
		Unit:           domain.Unit,
		Stock_Out:      domain.Stock_Out,
		Stock_Total:    domain.Stock_Total,
		StockID:        domain.StockID,
	}
}
