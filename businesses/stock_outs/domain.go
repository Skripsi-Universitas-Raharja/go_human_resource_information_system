package stockouts

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID             uint
	CreatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	Stock_Location string
	Stock_Code     string
	Stock_Name     string
	Stock_Unit     string
	Stock_Out      int
	Stock_Total    int
	StockID        uint
}
type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	StockOut(ctx context.Context, stockOutDomain *Domain) (Domain, error)
	ExportToExcel(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id string) (Domain, error)
	StockOut(ctx context.Context, stockOutDomain *Domain) (Domain, error)
	ExportToExcel(ctx context.Context) ([]Domain, error)
}
