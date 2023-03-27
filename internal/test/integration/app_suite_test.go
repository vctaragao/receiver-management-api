package integration

import (
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"github.com/vctaragao/receiver-management-api/internal/application"
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
	"github.com/vctaragao/receiver-management-api/internal/http"
	"github.com/vctaragao/receiver-management-api/internal/storage"
	"github.com/vctaragao/receiver-management-api/internal/storage/schemas"
	"github.com/vctaragao/receiver-management-api/internal/test/integration/helper"
	"gorm.io/gorm"
)

type IntegrationSuite struct {
	suite.Suite
	db *gorm.DB
	helper.Helper
}

func (s *IntegrationSuite) SetupSuite() {
	initApp(s.startRepo())
}

func (s *IntegrationSuite) startRepo() entity.Repository {
	repo := storage.NewPostgress()
	repo.Db.Unscoped().Delete(&schemas.Pix{}, "1=1")
	repo.Db.Unscoped().Delete(&schemas.Receiver{}, "1=1")

	repo.Db = repo.Db.Begin()
	repo.Db.SavePoint("init")

	s.db = repo.Db
	return repo
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
	s.db.RollbackTo("init")
}

func (s *IntegrationSuite) firstInDatabase(schema interface{}, expectedFields map[string]interface{}) {
	s.db.Where(expectedFields).First(schema)
}

func (s *IntegrationSuite) findInDatabase(schema interface{}, expectedFields map[string]interface{}) {
	s.db.Where(expectedFields).Find(&schema)
}

func (s *IntegrationSuite) seedDatabase(schemaWithData interface{}) interface{} {
	s.db.Create(schemaWithData)
	return schemaWithData
}

func TestIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test suite")
	}
	suite.Run(t, new(IntegrationSuite))
}
