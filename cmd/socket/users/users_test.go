package users

import (
	"dirStructureLecture/cmd/http/request"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/onsi/gomega"
	"net/url"
	"time"
)

var _ = GinkgoDescribe("User", func() {
	GinkgoIt("should be created", func() {
		e := echo.New()
		defer e.Close()

		e.GET("/user", CreateUserHandler(postgresDb))

		go func() {
			e.Start(fmt.Sprintf("%s:%s", "localhost", "3000"))
		}()

		time.Sleep(2 * time.Second)

		u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/user"}

		ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		gomega.Expect(err).Should(gomega.BeNil())

		defer ws.Close()

		user := request.User{
			Name:     "name",
			LastName: "lastName",
			Email:    "email@email.com",
		}

		b, _ := json.Marshal(user)
		err = ws.WriteMessage(websocket.TextMessage, b)
		gomega.Expect(err).Should(gomega.BeNil())

		_, _, err = ws.ReadMessage()
		gomega.Expect(err).Should(gomega.BeNil())
	})
})
