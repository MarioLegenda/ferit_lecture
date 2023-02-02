package deleting

import (
	"dirStructureLecture/pkg/storage"
	"github.com/onsi/gomega"
)

var _ = GinkgoDescribe("Users handler", func() {
	GinkgoIt("should delete a user", func() {
		user := testCreateUser("name", "lastName", "email@email.com")
		handler := NewUserDeleteById(UserID{
			ID: user.ID,
		}, storage.NewRepository[*User](postgresDb))

		_, err := handler.Handle()

		gomega.Expect(err).Should(gomega.BeNil())
	})
})
