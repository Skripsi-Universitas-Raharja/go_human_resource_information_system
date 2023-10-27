package users

import (
	"backend-golang/businesses/users"
	"backend-golang/drivers/mysql/profiles"

	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Email     string         `json:"email" gorm:"unique"`
	Name      string         `json:"name"`
	Password  string         `json:"password"`
	Nip       string         `json:"nip"`
	Division  string         `json:"division"`
	Role      string         `json:"role"`
	ProfileID uint           `json:"-" gorm:"uniqueIndex"`
	Profile   profiles.Profile
}

func (rec *User) ToDomain() users.Domain {
	return users.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		Email:     rec.Email,
		Name:      rec.Name,
		Password:  rec.Password,
		Nip:       rec.Nip,
		Division:  rec.Division,
		Role:      rec.Role,
		ProfileID: rec.ProfileID,
		Profile:   rec.Profile.ToDomain(),
	}
}

func FromDomain(domain *users.Domain) *User {
	return &User{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Email:     domain.Email,
		Name:      domain.Name,
		Password:  domain.Password,
		Nip:       domain.Nip,
		Division:  domain.Division,
		Role:      domain.Role,
		ProfileID: domain.ProfileID,
		Profile:   *profiles.FromDomain(&domain.Profile),
	}
}
