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
	Stock_Code     string         `json:"stock_code"`
	Stock_QRCode   string         `json:"stock_qrcode"`
	Stock_Name     string         `json:"stock_name"`
	Stock_Unit     string         `json:"stock_unit"`
	Stock_Total    int            `json:"stock_total"`
}

func FromDomain(domain stocks.Domain) Stock {
	return Stock{
		ID:             domain.ID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
		Stock_Location: domain.Stock_Location,
		Stock_QRCode:   domain.Stock_QRCode,
		Stock_Code:     domain.Stock_Code,
		Stock_Name:     domain.Stock_Name,
		Stock_Unit:     domain.Stock_Unit,
		Stock_Total:    domain.Stock_Total,
	}
}
