package request

import (
	stockhistory "backend-golang/businesses/stock_history"

	"github.com/go-playground/validator/v10"
)

type StockHistory struct {
	Stock_Location string  `json:"stock_location" validate:"required"`
	Stock_Name     string  `json:"stock_name" validate:"required"`
	Unit           string  `json:"unit" validate:"required"`
	Stock_Out      float64 `json:"stock_out"`
	Stock_Total    float64 `json:"stock_total"`
}

func (req *StockHistory) ToDomain() *stockhistory.Domain {
	return &stockhistory.Domain{
		Stock_Location: req.Stock_Location,
		Stock_Name:     req.Stock_Name,
		Unit:           req.Unit,
		Stock_Out:      req.Stock_Out,
		Stock_Total:    req.Stock_Total,
	}
}

func (req *StockHistory) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
