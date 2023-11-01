package users

import (
	"backend-golang/businesses/profiles"

	"context"
	"time"

	// "github.com/google/wire"
	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string
	Email     string
	Password  string
	Nip       string
	Division  string
	Role      string
	ProfileID uint
	Profile   profiles.Domain
}
type Usecase interface {
	Register(ctx context.Context, userDomain *Domain) (Domain, error)
	Login(ctx context.Context, userDomain *Domain) (string, error)

	// UpdateProfileUser(ctx context.Context, userDomain *Domain, id string) (Domain, error)
	// UploadProfileImage(ctx context.Context, userDomain *Domain, avatarPath string, id string) (Domain, string, error)
	// DeleteUser(ctx context.Context, id string) (error)
}

type Repository interface {
	Register(ctx context.Context, userDomain *Domain) (Domain, error)
	GetByEmail(ctx context.Context, userDomain *Domain) (Domain, error)

	// UpdateProfileUser(ctx context.Context, userDomain *Domain, id string) (Domain, error)
	// UploadProfileImage(ctx context.Context, userDomain *Domain, avatarPath string, id string) (Domain, string, error)
	// DeleteCustomer(ctx context.Context, id string) (error)
}
