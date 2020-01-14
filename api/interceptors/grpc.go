package interceptors

import (
	"context"

	"github.com/go-xorm/xorm"
	"github.com/happilymarrieddad/learning-go-gRPC/repos"
	"github.com/happilymarrieddad/learning-go-gRPC/utils"
	"google.golang.org/grpc"
)

func GlobalRepoInjector(db *xorm.Engine) grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		globalRepo := repos.GlobalRepo(db)
		newContext := utils.SetGlobalRepoOnContext(ctx, globalRepo)

		// BEFORE the request

		// Make the actual request
		res, err := handler(newContext, req)

		// AFTER the request - the database has presumably already been called!!!

		return res, err
	})
}
