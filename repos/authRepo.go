package repos

import (
	"crypto/ed25519"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/pascaldekloe/jwt"
)

var (
	prv ed25519.PrivateKey
	pub ed25519.PublicKey
)

func init() {
	var err error
	pub, prv, err = ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}
}

// Authrepo - the users repo interface
type Authrepo interface {
	GetNewClaims(subject string, set map[string]interface{}) *jwt.Claims
	GetSignedToken(claims *jwt.Claims) (string, error)
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

func (a authRepo) GetSignedToken(claims *jwt.Claims) (string, error) {
	now := time.Now().Round(time.Second)

	claims.Registered.Issued = jwt.NewNumericTime(now)
	claims.Registered.Expires = jwt.NewNumericTime(now.Add(7 * time.Hour * 24))
	claims.NotBefore = jwt.NewNumericTime(now.Add(-time.Second))

	token, err := claims.EdDSASign(prv)
	if err != nil {
		return "", nil
	}

	return string(token), nil
}
