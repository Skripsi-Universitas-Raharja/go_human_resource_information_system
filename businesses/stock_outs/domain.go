package stockouts

import (
	"context"
	"time"

	"gorm.io/gorm"

	"backend-golang/businesses/stocks"
)

type Domain struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	Stock_Location string
	Stock_Name     string
	Unit           string
	Stock_Out      float64
	Stock_Total    float64
	StockID        uint
	Stock          stocks.Domain
}
type Usecase interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	StockOut(ctx context.Context, categoryDomain *Domain) (Domain, error)
}

type Repository interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	StockOut(ctx context.Context, categoryDomain *Domain) (Domain, error)
}
