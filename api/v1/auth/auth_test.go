package auth

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	pb "github.com/happilymarrieddad/learning-go-gRPC/pb"
	repoMocks "github.com/happilymarrieddad/learning-go-gRPC/repos/mocks"
	types "github.com/happilymarrieddad/learning-go-gRPC/types"
	"github.com/happilymarrieddad/learning-go-gRPC/utils"
	"github.com/pascaldekloe/jwt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("grpc", func() {

	var (
		globalRepo *repoMocks.MockGlobalRepository
		authRepo   *repoMocks.MockAuthRepo
		usersRepo  *repoMocks.MockUsersRepo
		mokCtrl    *gomock.Controller
		router     pb.V1AuthServer
		ctx        context.Context

		validRequest *pb.LoginRequest
		validUser    *types.User
		validClaims  *jwt.Claims
	)

	setupRouter := func() {
		mokCtrl = gomock.NewController(GinkgoT())
		globalRepo = repoMocks.NewMockGlobalRepository(mokCtrl)
		authRepo = repoMocks.NewMockAuthRepo(mokCtrl)
		usersRepo = repoMocks.NewMockUsersRepo(mokCtrl)
		router = GetRoutes()

		ctx = utils.SetGlobalRepoOnContext(context.Background(), globalRepo)

		globalRepo.EXPECT().Auth().Return(authRepo).AnyTimes()
		globalRepo.EXPECT().Users().Return(usersRepo).AnyTimes()
	}

	setupData := func() {
		validRequest = &pb.LoginRequest{
			Email:    "foo@bar.com",
			Password: "1234",
		}
		validUser = &types.User{
			ID:        1,
			FirstName: "Nick",
			LastName:  "Doe2",
			Email:     validRequest.GetEmail(),
			Visible:   true,
		}

		validUser.SetPassword(validRequest.GetPassword())

		validClaims = &jwt.Claims{
			Registered: jwt.Registered{
				Subject: validUser.Email,
			},
			Set: map[string]interface{}{
				"user": validUser,
			},
		}
	}

	JustBeforeEach(func() {
		setupRouter()
		setupData()
	})

	JustAfterEach(func() {
		mokCtrl.Finish()
	})

	Describe("Login", func() {

		It("should fail because of a bad request", func() {
			_, err := router.Login(ctx, &pb.LoginRequest{})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal("Key: 'LoginRequest.email' Error:Field validation for 'email' failed on the 'valid-email' tag\nKey: 'LoginRequest.password' Error:Field validation for 'password' failed on the 'valid-password' tag"))
		})

		It("should fail because of a bad request", func() {
			_, err := router.Login(ctx, &pb.LoginRequest{
				Email: "foo@bar.com",
			})
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal("Key: 'LoginRequest.password' Error:Field validation for 'password' failed on the 'valid-password' tag"))
		})

		It("should fail with because it's missing the global repo", func() {
			_, err := router.Login(context.Background(), validRequest)
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal("unable to get global repo from context"))
		})

		It("should fail because of a database error", func() {
			errMsg := "database error"

			usersRepo.EXPECT().FindByEmail(validRequest.GetEmail()).
				Return(nil, errors.New(errMsg)).
				Times(1).Do(func(string) {
				defer GinkgoRecover()
			})

			_, err := router.Login(ctx, validRequest)
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})

		It("should fail because user password does not match the request", func() {
			errMsg := "crypto/bcrypt: hashedPassword is not the hash of the given password"

			validRequest.Password = "123"

			usersRepo.EXPECT().FindByEmail(validRequest.GetEmail()).
				Return(validUser, nil).
				Times(1).Do(func(string) {
				defer GinkgoRecover()
			})

			_, err := router.Login(ctx, validRequest)
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})

		It("should fail because unable to get signed token", func() {
			errMsg := "invalid key"

			usersRepo.EXPECT().FindByEmail(validRequest.GetEmail()).
				Return(validUser, nil).
				Times(1).Do(func(string) {
				defer GinkgoRecover()
			})

			authRepo.EXPECT().GetNewClaims(validUser.Email, map[string]interface{}{
				"user": validUser,
			}).Return(validClaims).Times(1).Do(func(string, map[string]interface{}) {
				defer GinkgoRecover()
			})

			authRepo.EXPECT().GetSignedToken(validClaims).
				Return("", errors.New(errMsg)).
				Times(1).Do(func(*jwt.Claims) {
				defer GinkgoRecover()
			})

			_, err := router.Login(ctx, validRequest)
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal(errMsg))
		})

		It("should successfully return a valid token", func() {
			validToken := "dfson429n9et"

			usersRepo.EXPECT().FindByEmail(validRequest.GetEmail()).
				Return(validUser, nil).
				Times(1).Do(func(string) {
				defer GinkgoRecover()
			})

			authRepo.EXPECT().GetNewClaims(validUser.Email, map[string]interface{}{
				"user": validUser,
			}).Return(validClaims).Times(1).Do(func(string, map[string]interface{}) {
				defer GinkgoRecover()
			})

			authRepo.EXPECT().GetSignedToken(validClaims).
				Return(validToken, nil).
				Times(1).Do(func(*jwt.Claims) {
				defer GinkgoRecover()
			})

			res, err := router.Login(ctx, validRequest)
			Ω(err).To(BeNil())
			Ω(res.GetToken()).To(Equal(validToken))
		})

	})

})
