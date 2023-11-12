package request

import (
	stockouts "backend-golang/businesses/stock_outs"

	"github.com/go-playground/validator/v10"
)

type StockOut struct {
	Stock_Location string `json:"stock_location"`
	Stock_Code     string `json:"stock_code"`
	Stock_Name     string `json:"stock_name"`
	Stock_Unit     string `json:"stock_unit"`
	Stock_Out      int    `json:"stock_out"`
	Stock_Total    int    `json:"stock_total"`
	StockID        uint   `json:"stock_id"`
}

func (req *StockOut) ToDomain() *stockouts.Domain {
	return &stockouts.Domain{
		Stock_Location: req.Stock_Location,
		Stock_Code:     req.Stock_Code,
		Stock_Name:     req.Stock_Name,
		Stock_Unit:     req.Stock_Unit,
		Stock_Out:      req.Stock_Out,
		Stock_Total:    req.Stock_Total,
		StockID:        req.StockID,
	}
}

func (req *StockOut) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
