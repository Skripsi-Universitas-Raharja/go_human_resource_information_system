package drivers

import (
	userDomain "backend-golang/businesses/users"
	userDB "backend-golang/drivers/mysql/users"

	profileDomain "backend-golang/businesses/profiles"
	profileDB "backend-golang/drivers/mysql/profiles"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}

func NewProfileRepository(conn *gorm.DB) profileDomain.Repository {
	return profileDB.NewMySQLRepository(conn)
}
