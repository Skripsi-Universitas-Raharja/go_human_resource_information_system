package request

import (
	stockins "backend-golang/businesses/stock_ins"

	"github.com/go-playground/validator/v10"
)

type StockIn struct {
	Stock_Location string `json:"stock_location"`
	Stock_Code     string `json:"stock_code"`
	Stock_Name     string `json:"stock_name"`
	Stock_Unit     string `json:"stock_unit"`
	Stock_In       int    `json:"stock_in"`
	Stock_Total    int    `json:"stock_total"`
	StockID        uint   `json:"stock_id"`
}

func (req *StockIn) ToDomain() *stockins.Domain {
	return &stockins.Domain{
		Stock_Location: req.Stock_Location,
		Stock_Code:     req.Stock_Code,
		Stock_Name:     req.Stock_Name,
		Stock_Unit:     req.Stock_Unit,
		Stock_In:       req.Stock_In,
		Stock_Total:    req.Stock_Total,
		StockID:        req.StockID,
	}
}

func (req *StockIn) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
