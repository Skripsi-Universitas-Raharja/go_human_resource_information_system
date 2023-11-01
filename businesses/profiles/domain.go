package profiles

import (
	// "backend-golang/businesses/users"

	// "github.com/google/wire"

	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
	Name       string
	Nip        string
	Division   string
	Image_Path string
}

type Usecase interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	UpdateProfileUser(ctx context.Context, profileDomain *Domain, id string) (Domain, error)
	UploadProfileImage(ctx context.Context, profileDomain *Domain, avatarPath string, id string) (Domain, string, error)
	// DeleteProfile(ctx context.Context, id string) (error)
}

type Repository interface {
	GetByID(ctx context.Context, id string) (Domain, error)
	UpdateProfileUser(ctx context.Context, profileDomain *Domain, id string) (Domain, error)
	UploadProfileImage(ctx context.Context, profileDomain *Domain, avatarPath string, id string) (Domain, string, error)
	// DeleteCustomer(ctx context.Context, id string) (error)
}
