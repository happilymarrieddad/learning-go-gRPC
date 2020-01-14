package users

import (
	"context"

	pb "github.com/happilymarrieddad/learning-go-gRPC/pb"
	"github.com/happilymarrieddad/learning-go-gRPC/types"
	"github.com/happilymarrieddad/learning-go-gRPC/utils"
)

type grpcHandler struct {
}

func GetRoutes() pb.V1UsersServer {
	return &grpcHandler{}
}

func (h *grpcHandler) Create(
	ctx context.Context,
	req *pb.CreateUserRequest,
) (
	res *pb.UserReply,
	err error,
) {
	res = new(pb.UserReply)

	if err = pb.Validate(req); err != nil {
		return
	}

	globalRepo, err := utils.GetGlobalRepoFromContext(ctx)
	if err != nil {
		return
	}

	newUser, err := types.NewUser(&types.TempUser{
		FirstName:       req.GetNewUser().GetFirstName(),
		LastName:        req.GetNewUser().GetLastName(),
		Email:           req.GetNewUser().GetEmail(),
		Password:        req.GetNewUser().GetPassword(),
		ConfirmPassword: req.GetNewUser().GetConfirmPassword(),
	})

	if err = globalRepo.Users().Create(newUser); err != nil {
		return
	}

	res.User = newUser.ToProtobuf()

	return
}

func (h *grpcHandler) FindById(
	ctx context.Context,
	req *pb.FindByIdRequest,
) (
	res *pb.UserReply,
	err error,
) {
	res = new(pb.UserReply)

	if err = pb.Validate(req); err != nil {
		return
	}

	globalRepo, err := utils.GetGlobalRepoFromContext(ctx)
	if err != nil {
		return
	}

	user, err := globalRepo.Users().FindById(req.GetId())
	if err != nil {
		return
	}

	res.User = user.ToProtobuf()

	return
}

func (h *grpcHandler) FindByEmail(
	ctx context.Context,
	req *pb.FindByEmailRequest,
) (
	res *pb.UserReply,
	err error,
) {
	res = new(pb.UserReply)

	if err = pb.Validate(req); err != nil {
		return
	}

	globalRepo, err := utils.GetGlobalRepoFromContext(ctx)
	if err != nil {
		return
	}

	user, err := globalRepo.Users().FindByEmail(req.GetEmail())
	if err != nil {
		return
	}

	res.User = user.ToProtobuf()

	return
}

func (h *grpcHandler) Update(
	ctx context.Context,
	req *pb.UpdateUserRequest,
) (
	res *pb.UserReply,
	err error,
) {
	res = new(pb.UserReply)

	if err = pb.Validate(req); err != nil {
		return
	}

	globalRepo, err := utils.GetGlobalRepoFromContext(ctx)
	if err != nil {
		return
	}

	usersRepo := globalRepo.Users()

	user, err := usersRepo.FindById(req.GetId())
	if err != nil {
		return
	}

	if len(req.FirstName) > 0 {
		user.FirstName = req.GetFirstName()
	}

	if len(req.LastName) > 0 {
		user.LastName = req.GetLastName()
	}

	if len(req.GetNewPassword()) > 0 {
		if err = user.SetPassword(req.GetNewPassword()); err != nil {
			return
		}
	}

	if err = usersRepo.Update(user); err != nil {
		return
	}

	return
}
