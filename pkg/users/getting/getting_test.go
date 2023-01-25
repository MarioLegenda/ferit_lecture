package getting

import (
	"dirStructureLecture/pkg/storage"
	"github.com/onsi/gomega"
)

var _ = GinkgoDescribe("Users handler", func() {
	GinkgoIt("should get a user", func() {
		user := testCreateUser("name", "lastName", "email@email.com")
		handler := NewGetById(UserId{
			ID: user.ID,
		}, storage.NewRepository[*User](postgresDb))

		fetchedUser, err := handler.Handle()

		gomega.Expect(err).Should(gomega.BeNil())

		gomega.Expect(fetchedUser.ID).ShouldNot(gomega.BeEmpty())
		gomega.Expect(fetchedUser.ID).ShouldNot(gomega.BeEmpty())
		gomega.Expect(fetchedUser.Name).Should(gomega.Equal("name"))
		gomega.Expect(fetchedUser.LastName).Should(gomega.Equal("lastName"))
	})
})
