package authrepository

import _entity "tixid/entity"

type Auth interface {
	Login(login _entity.User) (string, error)
}
