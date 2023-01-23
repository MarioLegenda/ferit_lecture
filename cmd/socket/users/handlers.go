package users

import (
	"dirStructureLecture/cmd/http/request"
	"dirStructureLecture/pkg/storage"
	"dirStructureLecture/pkg/users/adding"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

func CreateUserHandler(db storage.Storage) func(e echo.Context) error {
	return func(c echo.Context) error {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}
		defer ws.Close()

		for {
			_, msg, err := ws.ReadMessage()

			if err != nil {
				fmt.Println(err)

				return nil
			}

			var user request.User
			err = json.Unmarshal(msg, &user)
			if err != nil {
				fmt.Println(err)

				return nil
			}

			handler := adding.NewUserCreate(adding.User{
				Name:     user.Name,
				LastName: user.LastName,
				Email:    user.Email,
			}, storage.NewRepository[*adding.User](db))

			createdUser, err := handler.Handle()

			if err != nil {
				fmt.Println(err)

				return nil
			}

			fmt.Printf("Created a new user: %s. Sending it back...\n", msg)

			b, _ := json.Marshal(createdUser)

			err = ws.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				fmt.Println(err)

				return nil
			}
		}
	}
}
