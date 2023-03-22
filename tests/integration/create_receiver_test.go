package integration

import (
	"encoding/json"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
	"github.com/vctaragao/receiver-management-api/internal/http"
	"github.com/vctaragao/receiver-management-api/internal/storage/schemas"
)

func (s *IntegrationSuite) TestCreateReceiverIntegrationSuccess() {
	t := s.T()

	params := http.ReceiverInputDto{
		RazaoSocial: "Nome",
		Cpf:         "041.485.353-92",
		Cnpj:        "",
		Email:       "bryan_barbosa@prcondominios.com.br",
	}

	reqBody, _ := json.Marshal(params)

	resp := s.Request("POST", "/create", reqBody)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	var result http.ReceiverOutputDto
	err := s.DecodeBody(resp, &result)

	assert.NoError(t, err)

	receiver := &schemas.Receiver{}
	s.findInDatabase(receiver, map[string]interface{}{
		"id":           result.Id,
		"razao_social": params.RazaoSocial,
		"cpf":          params.Cpf,
		"cnpj":         params.Cnpj,
		"email":        params.Email,
		"status":       entity.STATUS_DRAFT,
	})

	assert.Equal(t, result.Id, receiver.ID)
}
