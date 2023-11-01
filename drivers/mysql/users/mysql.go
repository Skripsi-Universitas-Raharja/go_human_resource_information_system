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
