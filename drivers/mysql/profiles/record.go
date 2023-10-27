package profiles

import (
	"backend-golang/businesses/profiles"
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name"`
	Nip       string         `json:"nip"`
	Division  string         `json:"division"`
	ImagePath string         `json:"image_path"`
}

func (record *Profile) ToDomain() profiles.Domain {
	return profiles.Domain{
		ID:        record.ID,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		DeletedAt: record.DeletedAt,
		Name:      record.Name,
		Nip:       record.Nip,
		Division:  record.Division,
		ImagePath: record.ImagePath,
	}
}

func FromDomain(domain *profiles.Domain) *Profile {
	return &Profile{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Name:      domain.Name,
		Nip:       domain.Nip,
		Division:  domain.Division,
		ImagePath: domain.ImagePath,
	}
}
