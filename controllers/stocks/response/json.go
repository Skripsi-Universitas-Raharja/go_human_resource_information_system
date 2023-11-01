package response

import (
	"backend-golang/businesses/stocks"
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	Stock_Location string         `json:"stock_location"`
	Stock_Name     string         `json:"stock_name"`
	Unit           string         `json:"unit"`
	Stock_In       float64
	Stock_Out      float64
	Stock_Total    float64 `json:"stock_total"`
}

func FromDomain(domain stocks.Domain) Stock {
	return Stock{
		ID:             domain.ID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
		Stock_Location: domain.Stock_Location,
		Stock_Name:     domain.Stock_Name,
		Unit:           domain.Unit,
		Stock_In:       domain.Stock_In,
		Stock_Out:      domain.Stock_Out,
		Stock_Total:    domain.Stock_Total,
	}
}
