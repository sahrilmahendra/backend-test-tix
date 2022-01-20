package userController

import _entity "tixid/entity"

type CreateUserRequest struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Address  string `json:"address" form:"address"`
	Password string `json:"password" form:"password"`
}

type UserResponseFormat struct {
	ID       uint   `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Address  string `json:"address" form:"address"`
}

func FormattingUserResponse(format _entity.User) UserResponseFormat {
	return UserResponseFormat{
		ID:       format.ID,
		Username: format.Username,
		Email:    format.Email,
		Address:  format.Address,
	}
}
