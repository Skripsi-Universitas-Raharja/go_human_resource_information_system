package request

import (
	"backend-golang/businesses/stocks"

	"github.com/go-playground/validator/v10"
)

type Stock struct {
	Stock_Location string `json:"stock_location" validate:"required"`
	Stock_Code     string `json:"stock_code" validate:"required"`
	Stock_Name     string `json:"stock_name" validate:"required"`
	Stock_QRCode   string `json:"stock_qrcode"`
	Stock_Unit     string `json:"stock_unit" validate:"required"`
	Stock_Total    int    `json:"stock_total"`
}

func (req *Stock) ToDomain() *stocks.Domain {
	return &stocks.Domain{
		Stock_Location: req.Stock_Location,
		Stock_Code:     req.Stock_Code,
		Stock_QRCode:   req.Stock_QRCode,
		Stock_Name:     req.Stock_Name,
		Stock_Unit:     req.Stock_Unit,
		Stock_Total:    req.Stock_Total,
	}
}

func (req *Stock) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
