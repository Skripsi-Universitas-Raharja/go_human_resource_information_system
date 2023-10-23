package response

import (
	"backend-golang/businesses/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Email     string         `json:"email"`
	Name      string         `json:"name"`
	Password  string         `json:"password"`
	Nip       string         `json:"nip"`
	Division  string         `json:"division"`
	Role      string         `json:"role"`
}

func FromDomain(domain users.Domain) User {
	return User{
		ID:        domain.ID,
		Email:     domain.Email,
		Name:      domain.Name,
		Nip:       domain.Nip,
		Division:  domain.Division,
		Role:      domain.Role,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
