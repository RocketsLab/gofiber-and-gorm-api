package repositories

import (
	"errors"
	"github.com/RocketsLab/gofiber-and-gorm-api/http/service"
	"github.com/RocketsLab/gofiber-and-gorm-api/models"
)

func UserAll() (users []models.User, err error) {
	result := service.DbConnection.Find(&users)
	err = result.Error
	return users, err
}

func UserSave(user *models.User) error {
	user.ID = GenerateId()
	user.Password = HashPassword(user.Password)

	result := service.DbConnection.Create(user)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func UserFindByID(id string) (models.User, error) {
	var user models.User
	result := service.DbConnection.Where("id = ?", id).Find(&user)
	err := result.Error
	if err != nil {
		return user, err
	}
	if result.RowsAffected == 0 {
		err = errors.New("user not found")
	}
	return user, err
}
