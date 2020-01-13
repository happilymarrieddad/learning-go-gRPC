package repos

import (
	"github.com/go-xorm/xorm"
	"github.com/happilymarrieddad/learning-go-gRPC/types"
)

// UsersRepo - the users repo interface
type UsersRepo interface {
	Create(*types.User) error
}

// NewUsersRepo - returns a new user repo
func NewUsersRepo(db *xorm.Engine) UsersRepo {
	return &usersRepo{db: db}
}

type usersRepo struct {
	db *xorm.Engine
}

func (u usersRepo) Create(user *types.User) (err error) {
	if err = types.Validate(user); err != nil {
		return
	}

	if _, err = u.db.Insert(user); err != nil {
		return
	}

	return
}
