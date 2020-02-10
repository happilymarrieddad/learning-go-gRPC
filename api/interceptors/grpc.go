package interceptors

import (
	"context"
	"errors"
	"reflect"

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
		v := reflect.Indirect(reflect.ValueOf(req))
		vField := reflect.Indirect(v.FieldByName("JWT"))

		// If there is NO jwt field on the request, then just continue
		if !vField.IsValid() {
			return handler(newContext, req)
		}

		jwtToken := vField.String()

		user, err := globalRepo.Auth().GetDataFromToken(jwtToken)
		if err != nil {
			return nil, errors.New("unauthorized")
		}

		newContext = utils.SetUserOnContext(newContext, user)

		// Make the actual request
		res, err := handler(newContext, req)

		// AFTER the request - the database has presumably already been called!!!

		return res, err
	})
}
