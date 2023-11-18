package request

import (
	stockhistory "backend-golang/businesses/stock_history"

	"github.com/go-playground/validator/v10"
)

type StockHistory struct {
	Stock_Location string `json:"stock_location"`
	Stock_Code     string `json:"stock_code"`
	Stock_Name     string `json:"stock_name"`
	Stock_Unit     string `json:"stock_unit"`
	Stock_In       int    `json:"stock_in"`
	Stock_Out      int    `json:"stock_out"`
	Stock_Total    int    `json:"stock_total"`
}

func (req *StockHistory) ToDomain() *stockhistory.Domain {
	return &stockhistory.Domain{
		Stock_Location: req.Stock_Location,
		Stock_Code:     req.Stock_Code,
		Stock_Name:     req.Stock_Name,
		Stock_Unit:     req.Stock_Unit,
		Stock_In:       req.Stock_In,
		Stock_Out:      req.Stock_Out,
		Stock_Total:    req.Stock_Total,
	}
}

func (req *StockHistory) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
