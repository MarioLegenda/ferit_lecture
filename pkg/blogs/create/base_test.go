package create

import (
	"dirStructureLecture/pkg/storage"
	"dirStructureLecture/pkg/users/create"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"log"
	"os"
	"testing"
)

func loadEnv() {
	err := godotenv.Load("../../../.env")

	if err != nil {
		log.Fatal(err)
	}
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

func testCreateUser(name string, lastName string, email string) create.User {
	handler := create.NewUserCreate(create.User{
		Name:     name,
		LastName: lastName,
		Email:    email,
	}, storage.NewRepository[*create.User](postgresDb))

	createdUser, err := handler.Handle()

	gomega.Expect(err).Should(gomega.BeNil())

	return createdUser
}
