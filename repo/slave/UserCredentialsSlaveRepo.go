package slave

import (
	"authentication/clients"
	"authentication/models"

	"github.com/jinzhu/gorm"
)

//UserCredentialsSlaveRepoService - To perform read operations
type UserCredentialsSlaveRepoService interface {
	ReadUserCredentialsByUsername(username string) (*models.User, error)
	ReadUserCredentialsByPhone(phone string) (*models.User, error)
}

//UserCredentialsSlaveRepo -
type UserCredentialsSlaveRepo struct {
	database *gorm.DB
}

//NewUserCredentialsSlaveRepo -
func NewUserCredentialsSlaveRepo() UserCredentialsSlaveRepoService {
	var db *clients.DBInstance
	return &UserCredentialsSlaveRepo{
		database: db.GetDatabaseInstance(),
	}
}

//ReadUserCredentialsByUsername -
func (userRepo *UserCredentialsSlaveRepo) ReadUserCredentialsByUsername(username string) (*models.User, error) {
	var userCredentials models.User
	err := userRepo.database.Debug().Find(&userCredentials, "username = ?", username).Error

	if err != nil {
		return nil, err
	}

	return &userCredentials, nil
}

//ReadUserCredentialsByPhone -
func (userRepo *UserCredentialsSlaveRepo) ReadUserCredentialsByPhone(phone string) (*models.User, error) {
	var userCredentials models.User
	err := userRepo.database.Debug().Find(&userCredentials, "phone = ?", phone).Error

	if err != nil {
		return nil, err
	}

	return &userCredentials, nil
}
