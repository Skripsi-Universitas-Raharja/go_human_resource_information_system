package response

import (
	"backend-golang/businesses/profiles"
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name      string         `json:"name"`
	Nip       string         `json:"nip"`
	Division  string         `json:"division"`
}

func FromDomain(domain profiles.Domain) Profile {
	return Profile{
		ID:        domain.ID,
		Name:      domain.Name,
		Nip:       domain.Nip,
		Division:  domain.Division,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}
