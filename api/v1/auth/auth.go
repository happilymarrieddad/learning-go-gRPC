package auth

import (
	"context"

	pb "github.com/happilymarrieddad/learning-go-gRPC/pb"
	"github.com/happilymarrieddad/learning-go-gRPC/utils"
)

type grpcHandler struct {
}

// GetRoutes init routes
func GetRoutes() pb.V1AuthServer {
	return &grpcHandler{}
}

func (h *grpcHandler) Login(
	ctx context.Context,
	req *pb.LoginRequest,
) (
	res *pb.LoginReply,
	err error,
) {
	res = new(pb.LoginReply)

	if err = pb.Validate(req); err != nil {
		return
	}

	globalRepo, err := utils.GetGlobalRepoFromContext(ctx)
	if err != nil {
		return
	}
	usersRepo := globalRepo.Users()
	authRepo := globalRepo.Auth()

	user, err := usersRepo.FindByEmail(req.GetEmail())
	if err != nil {
		return
	}

	if err = user.Authenticate(req.GetPassword()); err != nil {
		return
	}

	claims := authRepo.GetNewClaims(user.Email, map[string]interface{}{
		"user": user,
	})

	tok, err := authRepo.GetSignedToken(claims)
	if err != nil {
		return
	}

	res.Token = tok

	return
}
