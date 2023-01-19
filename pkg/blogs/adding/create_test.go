package adding

import (
	"dirStructureLecture/pkg/storage"
	"github.com/onsi/gomega"
)

var _ = GinkgoDescribe("Blogs handler", func() {
	GinkgoIt("should fail creating a blog", func() {
		handler := NewBlogCreate(storage.Blog{
			Title:       "",
			Description: "",
			Content:     "",
		}, storage.NewRepository[*storage.Blog](postgresDb))

		_, err := handler.Handle()

		gomega.Expect(err).ShouldNot(gomega.BeNil())
	})

	GinkgoIt("should create a user", func() {
		user := testCreateUser("name", "lastName", "email@email.com")

		handler := NewBlogCreate(storage.Blog{
			Title:       "title",
			Description: "description",
			Content:     "content",
			UserID:      user.ID,
		}, storage.NewRepository[*storage.Blog](postgresDb))

		createdBlog, err := handler.Handle()

		gomega.Expect(err).Should(gomega.BeNil())

		gomega.Expect(createdBlog.ID).ShouldNot(gomega.BeEmpty())
	})
})
