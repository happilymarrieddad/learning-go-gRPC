package repos

import (
	"github.com/go-xorm/xorm"
)

// Authrepo - the users repo interface
type Authrepo interface {
}

// NewAuthrepo - returns a new user repo
func NewAuthrepo(db *xorm.Engine) Authrepo {
	return &authRepo{db: db}
}

type authRepo struct {
	db *xorm.Engine
}