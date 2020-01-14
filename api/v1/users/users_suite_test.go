package users_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUsers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Users Suite")
}
