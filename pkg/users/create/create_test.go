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

		gomega.Expect(createdUser.ID).Should(gomega.BeNumerically(">=", 0))
		gomega.Expect(createdUser.UUID).ShouldNot(gomega.BeEmpty())
		gomega.Expect(createdUser.Name).Should(gomega.Equal("name"))
		gomega.Expect(createdUser.LastName).Should(gomega.Equal("lastName"))
		gomega.Expect(createdUser.Email).Should(gomega.Equal("email@email.com"))
	})
})
