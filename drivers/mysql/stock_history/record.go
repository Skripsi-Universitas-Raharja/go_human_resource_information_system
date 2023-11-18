package stockhistory

import (
	stockhistory "backend-golang/businesses/stock_history"
	"backend-golang/drivers/mysql/stocks"

	"time"

	"gorm.io/gorm"
)

type StockHistory struct {
	ID                  uint           `json:"id" gorm:"primaryKey"`
	CreatedAt           time.Time      `json:"created_at"`
	DeletedAt           gorm.DeletedAt `json:"deleted_at"`
	Stock_Location      string         `json:"stock_location"`
	Stock_Code          string         `json:"stock_code"`
	Stock_Name          string         `json:"stock_name"`
	Stock_Unit          string         `json:"stock_unit"`
	Stock_In            int            `json:"stock_in"`
	Stock_Out           int            `json:"stock_out"`
	Stock_Total         int            `json:"stock_total,omitempty"`
	StockInTransactions stocks.Stock   `json:"-" gorm:"foreignKey:StockID"`
	StockID             uint           `json:"stock_id"`
}

func (rec *StockHistory) ToDomain() stockhistory.Domain {
	return stockhistory.Domain{
		ID:             rec.ID,
		CreatedAt:      rec.CreatedAt,
		DeletedAt:      rec.DeletedAt,
		Stock_Location: rec.Stock_Location,
		Stock_Code:     rec.Stock_Code,
		Stock_Name:     rec.Stock_Name,
		Stock_Unit:     rec.Stock_Unit,
		Stock_In:       rec.Stock_In,
		Stock_Out:      rec.Stock_Out,
		Stock_Total:    rec.Stock_Total,
		StockID:        rec.StockID,
	}
}

func FromDomain(domain *stockhistory.Domain) *StockHistory {
	return &StockHistory{
		ID:             domain.ID,
		CreatedAt:      domain.CreatedAt,
		DeletedAt:      domain.DeletedAt,
		Stock_Location: domain.Stock_Location,
		Stock_Code:     domain.Stock_Code,
		Stock_Name:     domain.Stock_Name,
		Stock_Unit:     domain.Stock_Unit,
		Stock_In:       domain.Stock_In,
		Stock_Out:      domain.Stock_Out,
		Stock_Total:    domain.Stock_Total,
		StockID:        domain.StockID,
	}
}
