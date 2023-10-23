package users

import (
	"backend-golang/businesses/users"
	"backend-golang/drivers/mysql/profiles"

	"context"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &userRepository{
		conn: conn,
	}
}

func (ur *userRepository) Register(ctx context.Context, userDomain *users.Domain) (users.Domain, error) {
	var profile profiles.Profile

	profile.Name = userDomain.Name
	profile.Nip = userDomain.Nip
	profile.Division = userDomain.Division

	result := ur.conn.WithContext(ctx).Create(&profile)

	if err := result.Error; err != nil {
		return users.Domain{}, err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)

	if err != nil {
		return users.Domain{}, err
	}

	record := FromDomain(userDomain)

	record.Password = string(password)
	record.Profile = profile
	record.ProfileID = profile.ID

	result = ur.conn.WithContext(ctx).Preload("Profile").Create(&record)

	if err := result.Error; err != nil {
		return users.Domain{}, err
	}

	err = result.Last(&record).Error

	if err != nil {
		return users.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (ur *userRepository) GetByEmail(ctx context.Context, userDomain *users.Domain) (users.Domain, error) {
	var user User

	err := ur.conn.WithContext(ctx).First(&user, "email = ?", userDomain.Email).Error

	if err != nil {
		return users.Domain{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDomain.Password))

	if err != nil {
		return users.Domain{}, err
	}

	return user.ToDomain(), nil
}

func (ur *userRepository) UpdateProfileUser(ctx context.Context, userDomain *users.Domain, id string) (users.Domain, error) {
	var user User

	if err := ur.conn.WithContext(ctx).Preload("Profile").First(&user, "id = ?", id).Error; err != nil {
		return users.Domain{}, err
	}

	if user.Email != user.Email {
		user.Email = userDomain.Email
	}

	var profile profiles.Profile

	if err := ur.conn.WithContext(ctx).First(&profile, "id = ?", id).Error; err != nil {
		return users.Domain{}, err
	}

	if user.Name != userDomain.Name {
		user.Name = userDomain.Name
		profile.Name = userDomain.Name
	}

	updatedProfiles := profile.ToDomain()

	updatedProfiles.Name = profile.Name
	updatedProfiles.Nip = profile.Nip
	updatedProfiles.Division = profile.Division

	if profile.Nip != userDomain.Profile.Nip {
		profile.Nip = userDomain.Profile.Nip
		updatedProfiles.Nip = profile.Nip
	}

	if profile.Division != userDomain.Profile.Division {
		profile.Division = userDomain.Profile.Division
		updatedProfiles.Division = profile.Division
	}

	if err := ur.conn.WithContext(ctx).Save(&profile).Error; err != nil {
		return users.Domain{}, err
	}

	if err := ur.conn.WithContext(ctx).Save(&user).Error; err != nil {
		return users.Domain{}, err
	}

	user.Profile = *profiles.FromDomain(&updatedProfiles)

	return user.ToDomain(), nil
}
