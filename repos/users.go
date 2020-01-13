package repos

import "github.com/go-xorm/xorm"

// UsersRepo - the users repo interface
type UsersRepo interface {
}

// NewUsersRepo - returns a new user repo
func NewUsersRepo(db *xorm.Engine) UsersRepo {
	return &usersRepo{db: db}
}

type usersRepo struct {
	db *xorm.Engine
}
