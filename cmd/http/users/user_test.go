package users

import (
	"dirStructureLecture/cmd/http/request"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = GinkgoDescribe("User", func() {
	GinkgoIt("should be created", func() {
		e := echo.New()
		b, err := json.Marshal(request.User{
			Name:     "name",
			LastName: "lastName",
			Email:    "email@email.com",
		})

		gomega.Expect(err).Should(gomega.BeNil())

		req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(string(b)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		handler := CreateUserHandler(postgresDb)

		err = handler(c)

		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(rec.Code).Should(gomega.Equal(http.StatusCreated))
	})

	GinkgoIt("should get", func() {
		e := echo.New()

		user := testCreateUser("name", "lastName", "email@email.com")
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/user/%s", user.ID), nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/user/:id")
		c.SetParamNames("id")
		c.SetParamValues(user.ID)

		handler := GetUserHandler(postgresDb)

		err := handler(c)

		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(rec.Code).Should(gomega.Equal(http.StatusOK))
	})
})
