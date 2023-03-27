package integration

import (
	"math/rand"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	br_faker "github.com/dimiro1/faker"
	"github.com/google/uuid"
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

func (s *IntegrationSuite) seedInBatch(size int) []schemas.Receiver {
	brFaker, _ := br_faker.NewForLocale("pt-br")
	gofakeit.Seed(0)

	var receivers []schemas.Receiver

	for i := 0; i < size; i++ {
		cpfCnpj := brFaker.BrazilCPF()
		if rand.Intn(2) == 1 {
			cpfCnpj = brFaker.BrazilCNPJ()
		}

		status := entity.GetValidReciverStatus()[rand.Intn(2)]

		receiver := schemas.Receiver{
			CorporateName: gofakeit.Name(),
			Email:         gofakeit.Email(),
			CpfCnpj:       cpfCnpj,
			Status:        status,
		}

		t := entity.GetValidTypes()[rand.Intn(4)]

		key := ""
		switch t {
		case entity.EMAIL:
			key = gofakeit.Email()
		case entity.CPF:
			key = brFaker.BrazilCPF()
		case entity.CNPJ:
			key = brFaker.BrazilCNPJ()
		case entity.PHONE:
			key = brFaker.CellPhoneNumber()
		case entity.RANDOM_KEY:
			key = uuid.New().String()
		}

		pix := schemas.Pix{
			Type: t,
			Key:  key,
		}

		receiver.Pix = append(receiver.Pix, pix)

		s.seedDatabase(&receiver)

		receivers = append(receivers, receiver)
	}

	return receivers
}

func TestIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test suite")
	}
	suite.Run(t, new(IntegrationSuite))
}
