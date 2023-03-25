package integration

import (
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"github.com/vctaragao/receiver-management-api/internal/application"
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
	"github.com/vctaragao/receiver-management-api/internal/http"
	"github.com/vctaragao/receiver-management-api/internal/storage"
	"github.com/vctaragao/receiver-management-api/internal/test/integration/helper"
	"gorm.io/gorm"
)

type IntegrationSuite struct {
	suite.Suite
	db *gorm.DB
	helper.Helper
}

func (s *IntegrationSuite) SetupSuite() {
	os.Setenv("APP_ENV", "test")

	repo := storage.NewPostgress()
	s.db = repo.Db

	initApp(repo)
}

func initApp(repo entity.Repository) {
	rm := application.NewReceiverManagement(repo)

	e := echo.New()
	http.RegisterRouter(e, rm)

	go startServer(e)
}

func startServer(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":1323"))
}

func (s *IntegrationSuite) TearDownTest() {
	s.db.Rollback()
}

func (s *IntegrationSuite) findInDatabase(schema interface{}, expectedFields map[string]interface{}) {
	s.db.Where(expectedFields).First(schema)
}

func TestIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test suite")
	}
	suite.Run(t, new(IntegrationSuite))
}
