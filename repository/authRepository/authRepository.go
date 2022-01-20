package authrepository

import (
	_middleware "tixid/delivery/middleware"
	_entity "tixid/entity"
	_utility "tixid/utility"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) Login(login _entity.User) (string, error) {
	var user _entity.User

	if err := ar.db.Where("username = ?", login.Username).First(&user).Error; err != nil {
		return "", err
	}

	match := _utility.DecryptPassword(user.Password, login.Password)
	if !match {
		return "", nil
	}

	token, err := _middleware.CreateToken(int(login.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}
