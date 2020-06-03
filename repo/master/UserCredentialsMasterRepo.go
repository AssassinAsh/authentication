package master

import (
	"authentication/clients"
	"authentication/models"

	"github.com/jinzhu/gorm"
)

//UserCredentialsMasterRepoService - To perform write operations
type UserCredentialsMasterRepoService interface {
	SaveToUserCredentials(user *models.User) error
	UpdateToUserCredentials(user *models.User) error
}

//UserCredentialsMasterRepo - Repository
type UserCredentialsMasterRepo struct {
	database *gorm.DB
}

//NewUserCredentialsMasterRepo - returns new repo
func NewUserCredentialsMasterRepo() UserCredentialsMasterRepoService {
	var db *clients.DBInstance
	return &UserCredentialsMasterRepo{
		database: db.GetDatabaseInstance(),
	}
}

//SaveToUserCredentials - Saves new Entry
func (userRepo UserCredentialsMasterRepo) SaveToUserCredentials(data *models.User) error {
	err := userRepo.database.Debug().Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

//UpdateToUserCredentials - Updates an existing entry
func (userRepo UserCredentialsMasterRepo) UpdateToUserCredentials(data *models.User) error {
	err := userRepo.database.Debug().Where("id = ?", data.ID).Update(&data).Error
	if err != nil {
		return err
	}
	return nil
}
