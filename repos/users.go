package repos

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/happilymarrieddad/learning-go-gRPC/types"
)

// UsersRepo - the users repo interface
type UsersRepo interface {
	Create(*types.User) error
	FindById(int64) (*types.User, error)
	FindByEmail(string) (*types.User, error)
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

func (u usersRepo) FindById(id int64) (user *types.User, err error) {
	if id <= 0 {
		err = fmt.Errorf("valid positive ID is required to find a user")
		return
	}

	user = new(types.User)

	has, err := u.db.ID(id).Get(user)
	if err != nil {
		return
	}

	if !has {
		err = fmt.Errorf("unable to find user")
		return
	}

	return
}

func (u usersRepo) FindByEmail(email string) (user *types.User, err error) {
	if len(email) == 0 {
		err = fmt.Errorf("valid email is required to find a user")
		return
	}

	user = new(types.User)
	user.Email = email

	has, err := u.db.Get(user)
	if err != nil {
		return
	}

	if !has {
		err = fmt.Errorf("unable to find user")
		return
	}

	return
}
