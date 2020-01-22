package repos

import (
	"github.com/go-xorm/xorm"
	"github.com/pascaldekloe/jwt"
)

// Authrepo - the users repo interface
type Authrepo interface {
	GetNewClaims(subject string, set map[string]interface{}) *jwt.Claims
}

// NewAuthrepo - returns a new user repo
func NewAuthrepo(db *xorm.Engine) Authrepo {
	return &authRepo{db: db}
}

type authRepo struct {
	db *xorm.Engine
}

func (a authRepo) GetNewClaims(subject string, set map[string]interface{}) *jwt.Claims {
	return &jwt.Claims{
		Registered: jwt.Registered{
			Subject: subject,
		},
		Set: set,
	}
}
