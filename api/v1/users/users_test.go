package users

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	pb "github.com/happilymarrieddad/learning-go-gRPC/pb"
	repoMocks "github.com/happilymarrieddad/learning-go-gRPC/repos/mocks"
	"github.com/happilymarrieddad/learning-go-gRPC/types"
	"github.com/happilymarrieddad/learning-go-gRPC/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("grpc", func() {

	var (
		globalRepo *repoMocks.MockGlobalRepository
		usersRepo  *repoMocks.MockUsersRepo
		mokCtrl    *gomock.Controller
		router     pb.V1UsersServer
		ctx        context.Context
	)

	setupRouter := func() {
		mokCtrl = gomock.NewController(GinkgoT())
		globalRepo = repoMocks.NewMockGlobalRepository(mokCtrl)
		usersRepo = repoMocks.NewMockUsersRepo(mokCtrl)
		router = GetRoutes()

		ctx = utils.SetGlobalRepoOnContext(context.Background(), globalRepo)

		globalRepo.EXPECT().Users().Return(usersRepo).AnyTimes()
	}

	JustBeforeEach(func() {
		setupRouter()
	})

	JustAfterEach(func() {
		mokCtrl.Finish()
	})

	Describe("Create", func() {

		It("should return error because of empty request", func() {
			errMsg := "Key: 'CreateUserRequest.newuser' Error:Field validation for 'newuser' failed on the 'valid-newuser' tag"
			_, err := router.Create(ctx, &pb.CreateUserRequest{})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})

		It("should return error because of missing email in request", func() {
			errMsg := "Key: 'CreateUserRequest.email' Error:Field validation for 'email' failed on the 'valid-email' tag\n" +
				"Key: 'CreateUserRequest.firstname' Error:Field validation for 'firstname' failed on the 'valid-firstname' tag\n" +
				"Key: 'CreateUserRequest.lastname' Error:Field validation for 'lastname' failed on the 'valid-lastname' tag\n" +
				"Key: 'CreateUserRequest.password' Error:Field validation for 'password' failed on the 'valid-password' tag\n" +
				"Key: 'CreateUserRequest.confirmpassword' Error:Field validation for 'confirmpassword' failed on the 'valid-confirmpassword' tag"
			_, err := router.Create(ctx, &pb.CreateUserRequest{
				NewUser: &pb.CreateUser{},
			})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})

		It("should fail a database test", func() {
			errMsg := "database is unavailable"

			user, err := types.NewUser(&types.TempUser{
				FirstName:       "Nick",
				LastName:        "Doe2",
				Email:           "foo@bar.com",
				Password:        "1234",
				ConfirmPassword: "1234",
			})
			Ω(err).To(BeNil())

			usersRepo.
				EXPECT().
				Create(gomock.AssignableToTypeOf(user)).
				Return(errors.New(errMsg)).
				Times(1).
				Do(func(*types.User) {
					defer GinkgoRecover()
				})

			_, err = router.Create(ctx, &pb.CreateUserRequest{
				NewUser: &pb.CreateUser{
					FirstName:       user.FirstName,
					LastName:        user.LastName,
					Email:           user.Email,
					Password:        "1234",
					ConfirmPassword: "1234",
				},
			})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})

		It("should be successful", func() {
			user, err := types.NewUser(&types.TempUser{
				FirstName:       "Nick",
				LastName:        "Doe2",
				Email:           "foo@bar.com",
				Password:        "1234",
				ConfirmPassword: "1234",
			})
			Ω(err).To(BeNil())

			usersRepo.
				EXPECT().
				Create(gomock.AssignableToTypeOf(user)).
				Return(nil).
				Times(1).
				Do(func(*types.User) {
					defer GinkgoRecover()
				})

			res, err := router.Create(ctx, &pb.CreateUserRequest{
				NewUser: &pb.CreateUser{
					FirstName:       user.FirstName,
					LastName:        user.LastName,
					Email:           user.Email,
					Password:        "1234",
					ConfirmPassword: "1234",
				},
			})
			Ω(err).To(BeNil())
			Ω(res.GetUser().GetEmail()).To(Equal(user.Email))
		})

	})

})
