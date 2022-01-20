package userRepository

import (
	_entity "tixid/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(newUser _entity.User) (_entity.User, error) {
	if err := ur.db.Save(&newUser).Error; err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (ur *UserRepository) GetAll() ([]_entity.User, error) {
	users := []_entity.User{}
	if err := ur.db.Table("users").Select("*").Where("users.deleted_at IS NULL").Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (ur *UserRepository) GetById(id int) (_entity.User, int, error) {
	user := _entity.User{}
	query := ur.db.Table("users").Select("*").Where("users.deleted_at IS NULL AND users.id = ?", id).Find(&user)
	if query.Error != nil || query.RowsAffected == 0 {
		return user, int(query.RowsAffected), query.Error
	}
	return user, int(query.RowsAffected), nil
}

func (ur *UserRepository) Update(id int, updateUser _entity.User) (_entity.User, error) {
	if err := ur.db.Where("id = ?", id).Updates(&updateUser).Error; err != nil {
		return updateUser, err
	}
	ur.db.First(&updateUser, id)
	return updateUser, nil
}

func (ur *UserRepository) Delete(id int) error {
	var user _entity.User
	if err := ur.db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
