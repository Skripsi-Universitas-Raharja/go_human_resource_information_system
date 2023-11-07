package stocks

import (
	"backend-golang/businesses/stocks"
	// stockhistory "backend-golang/drivers/mysql/stock_history"
	stockouts "backend-golang/drivers/mysql/stock_outs"

	"time"

	"gorm.io/gorm"
)

type Stock struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Stock_Location string         `json:"stock_location"`
	Stock_Code     string         `json:"stock_code"`
	Stock_QRCode   string         `json:"stock_qrcode"`
	Stock_Name     string         `json:"stock_name"`
	Unit           string         `json:"unit"`
	Stock_In       int            `json:"stock_in"`
	// Stock_Out      float64
	Stock_Total int                  `json:"stock_total"`
	Stock_Out   []stockouts.StockOut `json:"-" gorm:"foreignKey:StockID"`
	// sembunyikan dulu agar dapat berjalan
	// 	STRAINT `fk_stocks_stock_history` FOREIGN KEY (`stock_id`) REFERENCES `stocks`(`id`)
	// 2023/11/07 01:49:26 failed to perform database migration: Error 1452 (23000): Cannot add or update a child row: a foreign key constraint fails (`hrisdb`.`#sql-683c_f`, CONSTRAINT `fk_stocks_stock_history` FOREIGN KEY (`stock_id`) REFERENCES `stocks` (`id`))
	// Stock_History []stockhistory.StockHistory `json:"-" gorm:"foreignKey:StockID"`
}

func (rec *Stock) ToDomain() stocks.Domain {
	return stocks.Domain{
		ID:             rec.ID,
		CreatedAt:      rec.CreatedAt,
		UpdatedAt:      rec.UpdatedAt,
		DeletedAt:      rec.DeletedAt,
		Stock_Code:     rec.Stock_Code,
		Stock_Location: rec.Stock_Location,
		Stock_Name:     rec.Stock_Name,
		Unit:           rec.Unit,
		Stock_In:       rec.Stock_In,
		Stock_Total:    rec.Stock_Total,
	}
}

func FromDomain(domain *stocks.Domain) *Stock {
	return &Stock{
		ID:             domain.ID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		DeletedAt:      domain.DeletedAt,
		Stock_Location: domain.Stock_Location,
		Stock_Code:     domain.Stock_Code,
		Stock_Name:     domain.Stock_Name,
		Unit:           domain.Unit,
		Stock_In:       domain.Stock_In,
		Stock_Total:    domain.Stock_Total,
	}
}
