package blogs

import (
	"dirStructureLecture/cmd/http/request"
	"dirStructureLecture/cmd/http/users"
	"dirStructureLecture/pkg/storage"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func loadEnv() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal(err)
	}
}

type TestUser struct {
	Name     string `json:"Name"`
	ID       string `json:"ID"`
	LastName string `json:"LastName"`
	Email    string `json:"Email"`
}

var GomegaRegisterFailHandler = gomega.RegisterFailHandler
var GinkgoFail = ginkgo.Fail
var GinkgoRunSpecs = ginkgo.RunSpecs
var GinkgoBeforeSuite = ginkgo.BeforeSuite
var GinkgoBeforeHandler = ginkgo.BeforeEach
var GinkgoAfterHandler = ginkgo.AfterEach
var GinkgoAfterSuite = ginkgo.AfterSuite
var GinkgoDescribe = ginkgo.Describe
var GinkgoIt = ginkgo.It

var postgresDb storage.Storage

func TestApi(t *testing.T) {
	GomegaRegisterFailHandler(GinkgoFail)
	GinkgoRunSpecs(t, "API Suite")
}

var _ = GinkgoBeforeSuite(func() {
	loadEnv()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Europe/Zagreb",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)

	db, err := storage.NewStorage(dsn)

	if err != nil {
		log.Fatalln(err)
	}

	postgresDb = db
})

var _ = GinkgoAfterSuite(func() {
	sql, err := postgresDb.DB().DB()
	if err != nil {
		log.Fatalln(err)
	}

	if err := sql.Close(); err != nil {
		log.Fatalln(err)
	}
})

func testCreateUser() TestUser {
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

	handler := users.CreateUserHandler(postgresDb)

	err = handler(c)

	gomega.Expect(err).Should(gomega.BeNil())
	gomega.Expect(rec.Code).Should(gomega.Equal(http.StatusCreated))

	var user TestUser
	gomega.Expect(json.Unmarshal(rec.Body.Bytes(), &user)).Should(gomega.BeNil())

	return user
}
