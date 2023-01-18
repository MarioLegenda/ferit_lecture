package create

import (
	"github.com/onsi/gomega"
)

var _ = GinkgoDescribe("Users handler", func() {
	GinkgoIt("should fail creating a user", func() {
		handler := NewUserCreate(User{
			Name:     "",
			LastName: "",
			Email:    "",
		}, NewRepository[*User](postgresDb))

		_, err := handler.Handle()

		gomega.Expect(err).ShouldNot(gomega.BeNil())
	})

	GinkgoIt("should create a user", func() {
		handler := NewUserCreate(User{
			Name:     "name",
			LastName: "lastName",
			Email:    "email@email.com",
		}, NewRepository[*User](postgresDb))

		createdUser, err := handler.Handle()

		gomega.Expect(err).Should(gomega.BeNil())

		gomega.Expect(createdUser.ID).ShouldNot(gomega.BeEmpty())
	})
})
