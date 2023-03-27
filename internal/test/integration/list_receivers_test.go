package integration

import (
	"math/rand"
	"net/url"

	"github.com/brianvoe/gofakeit/v6"
	br_faker "github.com/dimiro1/faker"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
	"github.com/vctaragao/receiver-management-api/internal/http"
	"github.com/vctaragao/receiver-management-api/internal/storage/schemas"
)

func (s *IntegrationSuite) TestListReceiversByCorporateNameIntegrationSuccess() {
	t := s.T()

	seed(s)

	values := url.Values{}
	values.Set("search", "Olivia Daiane Tânia Rezende")

	resp := s.Request("GET", "/list?"+values.Encode(), nil)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	var result http.ListOutputDto
	err := s.DecodeBody(resp, &result)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(result.Receivers))
	assert.Equal(t, 1, result.Total)
	assert.Equal(t, 1, result.CurrentPage)

	for _, receiver := range result.Receivers {
		assert.Equal(t, "bryan_barbosa@prcondominios.com.br", receiver.Email)
		assert.Equal(t, entity.STATUS_DRAFT, receiver.Status)
		assert.Equal(t, "106.762.957-20", receiver.CpfCnpj)
		assert.Equal(t, "Olivia Daiane Tânia Rezende", receiver.CorporateName)

		assert.Equal(t, "CPF", receiver.PixType)
		assert.Equal(t, "106.762.957-20", receiver.PixKey)
	}
}

func (s *IntegrationSuite) TestListReceiversByStatusIntegrationSuccess() {
	t := s.T()

	seed(s)

	values := url.Values{}
	values.Set("search", entity.STATUS_DRAFT)

	resp := s.Request("GET", "/list?"+values.Encode(), nil)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	var result http.ListOutputDto
	err := s.DecodeBody(resp, &result)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(result.Receivers))
	assert.Equal(t, 2, result.Total)
}

func (s *IntegrationSuite) TestListReceiversByPixTypeIntegrationSuccess() {
	t := s.T()

	seed(s)

	values := url.Values{}
	values.Set("search", entity.CNPJ)

	resp := s.Request("GET", "/list?"+values.Encode(), nil)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	var result http.ListOutputDto
	err := s.DecodeBody(resp, &result)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(result.Receivers))
}

func (s *IntegrationSuite) TestListReceiversByPixKeyIntegrationSuccess() {
	t := s.T()

	seed(s)

	values := url.Values{}
	values.Set("search", "106.762.957-20")

	resp := s.Request("GET", "/list?"+values.Encode(), nil)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	var result http.ListOutputDto
	err := s.DecodeBody(resp, &result)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(result.Receivers))
}

func (s *IntegrationSuite) TestListReceiversPaginationIntegrationSuccess() {
	t := s.T()

	seedInBatch(s, 20)

	values := url.Values{}
	values.Set("page", "2")

	resp := s.Request("GET", "/list?"+values.Encode(), nil)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	var result http.ListOutputDto
	err := s.DecodeBody(resp, &result)
	assert.NoError(t, err)

	assert.Equal(t, 10, len(result.Receivers))
	assert.Equal(t, 20, result.Total)
	assert.Equal(t, 2, result.CurrentPage)
}

func (s *IntegrationSuite) TestListReceiversPaginationNoRecordFoundIntegrationSuccess() {
	t := s.T()

	seed(s)

	values := url.Values{}
	values.Set("search", "Renan Isaac Almada")

	resp := s.Request("GET", "/list?"+values.Encode(), nil)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	var result http.ListOutputDto
	err := s.DecodeBody(resp, &result)
	assert.NoError(t, err)

	assert.Equal(t, 0, len(result.Receivers))
}

func seed(s *IntegrationSuite) (*schemas.Receiver, *schemas.Pix) {
	receiver := &schemas.Receiver{
		CorporateName: "Olivia Daiane Tânia Rezende",
		CpfCnpj:       "106.762.957-20",
		Email:         "bryan_barbosa@prcondominios.com.br",
		Status:        entity.STATUS_DRAFT,
	}

	s.seedDatabase(receiver)

	pix := &schemas.Pix{
		ReceiverId: receiver.ID,
		Type:       "CPF",
		Key:        "106.762.957-20",
	}

	s.seedDatabase(pix)

	receiver = &schemas.Receiver{
		CorporateName: "Alana Sara Silveira",
		CpfCnpj:       "366.101.352-15",
		Email:         "alana_silveira@land.com.br",
		Status:        entity.STATUS_VALID,
	}

	s.seedDatabase(receiver)

	pix = &schemas.Pix{
		ReceiverId: receiver.ID,
		Type:       "EMAIL",
		Key:        "alana_silveira@land.com.br",
	}

	s.seedDatabase(pix)

	receiver = &schemas.Receiver{
		CorporateName: "Luana Malu Alana Castro",
		CpfCnpj:       "897.243.549-03",
		Email:         "luana_castro@band.com.br",
		Status:        entity.STATUS_DRAFT,
	}
	s.seedDatabase(receiver)

	pix = &schemas.Pix{
		ReceiverId: receiver.ID,
		Type:       "EMAIL",
		Key:        "luana_castro@band.com.br",
	}

	s.seedDatabase(pix)

	receiver = &schemas.Receiver{
		CorporateName: "Malu e Davi Pizzaria Delivery ME",
		CpfCnpj:       "80.350.691/0001-92",
		Email:         "presidencia@maluedavipizzariadeliveryme.com.br",
		Status:        entity.STATUS_VALID,
	}

	s.seedDatabase(receiver)

	pix = &schemas.Pix{
		ReceiverId: receiver.ID,
		Type:       "CNPJ",
		Key:        "80.350.691/0001-92",
	}

	s.seedDatabase(pix)

	return receiver, pix
}

func seedInBatch(s *IntegrationSuite, size int) {
	brFaker, _ := br_faker.NewForLocale("pt-br")
	gofakeit.Seed(0)
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
	}
}
