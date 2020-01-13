package repos_test

import (
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/happilymarrieddad/learning-go-gRPC/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UsersRepo", func() {

	var (
		err error
		usr *User

		setupData = func() {
			usr, err = NewUser(&TempUser{
				FirstName:       "Nick",
				LastName:        "Doe2",
				Email:           "foo@bar.com",
				Password:        "1234",
				ConfirmPassword: "1234",
			})
			Ω(err).To(BeNil())
		}
	)

	BeforeEach(func() {
		clearDatabase()
		setupData()
	})

	Describe("Create", func() {
		Context("Failures", func() {
			It("should fail with a nil user", func() {
				err = gr.Users().Create(nil)
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal("validator: (nil *types.User)"))
			})
			It("should fail with a bad user", func() {
				err = gr.Users().Create(&User{
					Password: usr.Password,
					Visible:  true,
				})
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal(
					"Key: 'User.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag\n" +
						"Key: 'User.LastName' Error:Field validation for 'LastName' failed on the 'required' tag\n" +
						"Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag"),
				)
			})
			It("should fail with database error", func() {
				errMsg := "database unavailable"

				mock.ExpectExec("INSERT INTO `users` (`first_name`,`last_name`,`email`,`password`,`visible`) VALUES (?, ?, ?, ?, ?)").
					WithArgs(usr.FirstName, usr.LastName, usr.Email, usr.Password, usr.Visible).
					WillReturnError(errors.New(errMsg))

				err = gr.Users().Create(usr)
				Ω(err).NotTo(BeNil())
				Ω(err.Error()).To(Equal(errMsg))
			})
		})
		Context("Success", func() {
			It("successfully stored a user", func() {

				mock.ExpectExec("INSERT INTO `users` (`first_name`,`last_name`,`email`,`password`,`visible`) VALUES (?, ?, ?, ?, ?)").
					WithArgs(usr.FirstName, usr.LastName, usr.Email, usr.Password, usr.Visible).
					WillReturnResult(sqlmock.NewResult(1, 1))

				err = gr.Users().Create(usr)
				Ω(err).To(BeNil())
			})
		})
	})

})
