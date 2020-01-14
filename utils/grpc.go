package utils

import (
	"context"
	"errors"

	"github.com/happilymarrieddad/learning-go-gRPC/repos"
)

type key string

const (
	globalRepoKey key = "globalRepo"
)

// GetGlobalRepoFromContext - get the global repo from context
func GetGlobalRepoFromContext(
	ctx context.Context,
) (globalRepo repos.GlobalRepository, err error) {
	r, ok := ctx.Value(globalRepoKey).(repos.GlobalRepository)
	if ok {
		return r, nil
	}

	return nil, errors.New("unable to get global repo from context")
}

// SetGlobalRepoOnContext - set the global repo value in the context
func SetGlobalRepoOnContext(
	ctx context.Context,
	globalRepo repos.GlobalRepository,
) context.Context {
	return context.WithValue(
		ctx,
		globalRepoKey,
		globalRepo,
	)
}
