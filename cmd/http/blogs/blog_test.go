package blogs

import (
	"dirStructureLecture/cmd/http/request"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = GinkgoDescribe("Blog", func() {
	GinkgoIt("should be created", func() {
		user := testCreateUser()

		e := echo.New()
		b, err := json.Marshal(request.Blog{
			UserID:      user.ID,
			Title:       "title",
			Content:     "content",
			Description: "description",
		})

		gomega.Expect(err).Should(gomega.BeNil())

		req := httptest.NewRequest(http.MethodPost, "/blog", strings.NewReader(string(b)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := CreateBlogHandler(postgresDb)

		err = handler(c)

		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(rec.Code).Should(gomega.Equal(http.StatusCreated))
	})
})
