package repos_test

import (
	"github.com/happilymarrieddad/learning-go-gRPC/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AuthRepo", func() {

	var (
		usr *types.User
	)

	JustBeforeEach(func() {
		usr = &types.User{
			ID:        1,
			FirstName: "Nick",
			LastName:  "Doe2",
			Email:     "foo@bar.com",
			Password:  "1234",
			Visible:   true,
		}
	})

	Describe("GetNewClaims", func() {

		It("should successfully return a claims object", func() {
			claims := gr.Auth().GetNewClaims("test", map[string]interface{}{
				"user": usr,
			})

			Ω(claims).NotTo(BeNil())
			Ω(claims.Set["user"]).NotTo(BeNil())
		})

	})

	Describe("GetSignedToken", func() {
		It("should return a signed token", func() {
			claims := gr.Auth().GetNewClaims("test", map[string]interface{}{
				"user": usr,
			})

			token, err := gr.Auth().GetSignedToken(claims)
			Ω(err).To(BeNil())
			Ω(len(token)).NotTo(Equal(0))
		})
	})

	Describe("GetDataFromToken", func() {

		It("should fail to return data because user is missing", func() {

			claims := gr.Auth().GetNewClaims("test", map[string]interface{}{})

			token, err := gr.Auth().GetSignedToken(claims)
			Ω(err).To(BeNil())
			Ω(len(token)).NotTo(Equal(0))

			_, err = gr.Auth().GetDataFromToken(token)
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal("token is valid but user data missing or corrupt"))
		})

		It("should fail to return data because user is missing ID", func() {

			claims := gr.Auth().GetNewClaims("test", map[string]interface{}{
				"user": map[string]interface{}{},
			})

			token, err := gr.Auth().GetSignedToken(claims)
			Ω(err).To(BeNil())
			Ω(len(token)).NotTo(Equal(0))

			_, err = gr.Auth().GetDataFromToken(token)
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal("token is valid but user data missing or corrupt"))
		})

		It("should fail to return data because user is missing email", func() {

			claims := gr.Auth().GetNewClaims("test", map[string]interface{}{
				"user": map[string]interface{}{
					"id": usr.ID,
				},
			})

			token, err := gr.Auth().GetSignedToken(claims)
			Ω(err).To(BeNil())
			Ω(len(token)).NotTo(Equal(0))

			_, err = gr.Auth().GetDataFromToken(token)
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal("token is valid but user data missing or corrupt"))
		})

		It("should fail to return data because user is missing visible", func() {

			claims := gr.Auth().GetNewClaims("test", map[string]interface{}{
				"user": map[string]interface{}{
					"id":    usr.ID,
					"email": usr.Email,
				},
			})

			token, err := gr.Auth().GetSignedToken(claims)
			Ω(err).To(BeNil())
			Ω(len(token)).NotTo(Equal(0))

			_, err = gr.Auth().GetDataFromToken(token)
			Ω(err).NotTo(BeNil())
			Ω(err.Error()).To(Equal("token is valid but user data missing or corrupt"))
		})

		It("should successfully return user data", func() {

			claims := gr.Auth().GetNewClaims("test", map[string]interface{}{
				"user": usr,
			})

			token, err := gr.Auth().GetSignedToken(claims)
			Ω(err).To(BeNil())
			Ω(len(token)).NotTo(Equal(0))

			user, err := gr.Auth().GetDataFromToken(token)
			Ω(err).To(BeNil())
			Ω(usr.Email).To(Equal(user.Email))
		})
	})

})
