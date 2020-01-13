package repos

import (
	"github.com/go-xorm/xorm"
)

var globalRepositoryInstance *globalRepository

// GlobalRepository - the global repository interface
type GlobalRepository interface {
}

// GlobalRepo - global repository func
func GlobalRepo(db *xorm.Engine) GlobalRepository {
	if globalRepositoryInstance != nil {
		return globalRepositoryInstance
	}

	globalRepositoryInstance = &globalRepository{db: db, repos: make(map[string]interface{})}
	return globalRepositoryInstance
}

type globalRepository struct {
	db    *xorm.Engine
	repos map[string]interface{}
}

func (r *globalRepository) repoFactory(key string, factory func() interface{}) interface{} {
	if iface, exists := r.repos[key]; exists {
		return iface
	}
	iface := factory()
	r.repos[key] = iface
	return iface
}
