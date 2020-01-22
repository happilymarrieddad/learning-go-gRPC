package repos

import (
	"crypto/ed25519"
	"errors"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/happilymarrieddad/learning-go-gRPC/types"
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

// AuthRepo - the users repo interface
type AuthRepo interface {
	GetNewClaims(subject string, set map[string]interface{}) *jwt.Claims
	GetSignedToken(claims *jwt.Claims) (string, error)
	GetDataFromToken(token string) (*types.User, error)
}

// NewAuthRepo - returns a new user repo
func NewAuthRepo(db *xorm.Engine) AuthRepo {
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

func (a authRepo) GetDataFromToken(token string) (*types.User, error) {
	claims, err := jwt.EdDSACheck([]byte(token), pub)
	if err != nil {
		return nil, err
	}

	userDataErr := errors.New("token is valid but user data missing or corrupt")
	userData, ok := claims.Set["user"].(map[string]interface{})
	if !ok {
		return nil, userDataErr
	}

	user := new(types.User)

	id, ok := userData["id"].(float64)
	if !ok {
		return nil, userDataErr
	}
	user.ID = int64(id)

	email, ok := userData["email"].(string)
	if !ok {
		return nil, userDataErr
	}
	user.Email = email

	visible, ok := userData["id"].(bool)
	if !ok {
		return nil, userDataErr
	}
	user.Visible = visible

	return user, nil
}
