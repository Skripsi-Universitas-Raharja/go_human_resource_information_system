package profiles

import (
	"backend-golang/businesses/profiles"
	"context"

	"gorm.io/gorm"
)

type profileRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) profiles.Repository {
	return &profileRepository{
		conn: conn,
	}
}

func (ur *profileRepository) GetByID(ctx context.Context, id string) (profiles.Domain, error) {
	var profile Profile

	if err := ur.conn.WithContext(ctx).First(&profile, "id = ?", id).Error; err != nil {
		return profiles.Domain{}, err
	}

	return profile.ToDomain(), nil

}

func (ur *profileRepository) UpdateProfileUser(ctx context.Context, profileDomain *profiles.Domain, id string) (profiles.Domain, error) {
	var profile Profile

	// Preload "User" untuk memastikan data User ter-load
	if err := ur.conn.WithContext(ctx).First(&profile, "id = ?", id).Error; err != nil {
		return profiles.Domain{}, err
	}

	// Update hanya jika nilai berbeda
	if profile.Name != profileDomain.Name {
		profile.Name = profileDomain.Name
	}

	if profile.Nip != profileDomain.Nip {
		profile.Nip = profileDomain.Nip
	}

	if profile.Division != profileDomain.Division {
		profile.Division = profileDomain.Division
	}

	// Simpan perubahan ke database
	if err := ur.conn.WithContext(ctx).Save(&profile).Error; err != nil {
		return profiles.Domain{}, err
	}

	// Mengembalikan profil yang telah diperbarui
	return profile.ToDomain(), nil
}

func (ur *profileRepository) UploadProfileImage(ctx context.Context, profileDomain *profiles.Domain, avatarPath string, id string) (profiles.Domain, string, error) {
	var profile Profile

	if err := ur.conn.WithContext(ctx).First(&profile, "id = ?", id).Error; err != nil {
		return profiles.Domain{}, "", err
	}

	prev_url := profile.Image_Path
	profile.Image_Path = avatarPath

	if err := ur.conn.WithContext(ctx).Save(&profile).Error; err != nil {
		return profiles.Domain{}, "", err
	}

	return profile.ToDomain(), prev_url, nil
}
