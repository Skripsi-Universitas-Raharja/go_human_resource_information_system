package request

import (
	"backend-golang/businesses/stocks"

	"github.com/go-playground/validator/v10"
)

type Stock struct {
	Stock_Location string  `json:"stock_location" validate:"required"`
	Stock_Name     string  `json:"stock_name" validate:"required"`
	Unit           string  `json:"unit" validate:"required"`
	Stock_In       float64 `json:"stock_in"`
	Stock_Out      float64 `json:"stock_out"`
	Stock_Total    float64 `json:"stock_total"`
}

func (req *Stock) ToDomain() *stocks.Domain {
	return &stocks.Domain{
		Stock_Location: req.Stock_Location,
		Stock_Name:     req.Stock_Name,
		Unit:           req.Unit,
		Stock_In:       req.Stock_In,
		Stock_Out:      req.Stock_Out,
		Stock_Total:    req.Stock_Total,
	}
}

func (req *Stock) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
