package response

import (
	stockins "backend-golang/businesses/stock_ins"

	"time"

	"gorm.io/gorm"
)

type StockIn struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
	Stock_Location string         `json:"stock_location"`
	Stock_Code     string `json:"stock_code"`
	Stock_Name     string `json:"stock_name"`
	Stock_Unit     string `json:"stock_unit"`
	Stock_In       int    `json:"stock_in"`
	Stock_Total    int    `json:"stock_total"`
	StockID        uint   `json:"stock_id"`
}

func FromDomain(domain stockins.Domain) StockIn {
	return StockIn{
		ID:             domain.ID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
		Stock_Location: domain.Stock_Location,
		Stock_Code: domain.Stock_Code,
		Stock_Name:     domain.Stock_Name,
		Stock_Unit:     domain.Stock_Unit,
		Stock_In:       domain.Stock_In,
		Stock_Total:    domain.Stock_Total,
		StockID:        domain.StockID,
	}
}
