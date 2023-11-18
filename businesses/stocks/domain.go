package stocks

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
	Stock_QRCode   string
	Stock_Name     string
	Stock_Unit     string
	Stock_Total    int
}
type Usecase interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	DownloadBarcodeByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, categoryDomain *Domain) (Domain, error)
	// StockIn(ctx context.Context, categoryDomain *Domain, id string) (Domain, error)
	// StockOut(ctx context.Context, categoryDomain *Domain, id string) (Domain, error)

	// Register(ctx context.Context, userDomain *Domain) (Domain, error)
	// Login(ctx context.Context, userDomain *Domain) (string, error)
	// DeleteUser(ctx context.Context, id string) (error)
}

type Repository interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, categoryDomain *Domain) (Domain, error)
	DownloadBarcodeByID(ctx context.Context, id string) (Domain, error)
	// StockIn(ctx context.Context, categoryDomain *Domain, id string) (Domain, error)
	// StockOut(ctx context.Context, categoryDomain *Domain, id string) (Domain, error)

	// Register(ctx context.Context, userDomain *Domain) (Domain, error)
	// GetByEmail(ctx context.Context, userDomain *Domain) (Domain, error)
	// DeleteCustomer(ctx context.Context, id string) (error)
}
