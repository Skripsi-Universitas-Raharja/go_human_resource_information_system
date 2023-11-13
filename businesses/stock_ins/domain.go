package stockins

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	Stock_Location string
	Stock_Code     string
	Stock_Name     string
	Stock_Unit     string
	Stock_In       int
	Stock_Total    int
	StockID        uint
}
type Usecase interface {
	// GetByID(ctx context.Context, id string) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Create(ctx context.Context, stockInDomain *Domain) (Domain, error)
	StockIn(ctx context.Context, stockInDomain *Domain) (Domain, error)
	ExportToExcel(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	// GetByID(ctx context.Context, id string) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Create(ctx context.Context, stockInDomain *Domain) (Domain, error)
	StockIn(ctx context.Context, stockInDomain *Domain) (Domain, error)
	ExportToExcel(ctx context.Context) ([]Domain, error)
}
