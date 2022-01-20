package userRepository

import _entity "tixid/entity"

type User interface {
	Create(newUser _entity.User) (_entity.User, error)
	GetAll() ([]_entity.User, error)
	GetById(id int) (_entity.User, int, error)
	Update(id int, updateUser _entity.User) (_entity.User, error)
	Delete(id int) error
}
