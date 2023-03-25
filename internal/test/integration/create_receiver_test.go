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
		PixType:     "CPF",
		PixKey:      "041.485.353-92",
	}

	reqBody, _ := json.Marshal(params)

	resp := s.Request("POST", "/create", reqBody)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	var result http.ReceiverOutputDto
	err := s.DecodeBody(resp, &result)

	assert.NoError(t, err)

	receiver := &schemas.Receiver{}
	s.firstInDatabase(receiver, map[string]interface{}{
		"id":             result.Id,
		"corporate_name": params.RazaoSocial,
		"cpf":            params.Cpf,
		"cnpj":           params.Cnpj,
		"email":          params.Email,
		"status":         entity.STATUS_DRAFT,
	})

	assert.Equal(t, result.Id, receiver.ID)

	pix := &schemas.Pix{}
	s.firstInDatabase(pix, map[string]interface{}{
		"type":        params.PixType,
		"key":         params.PixKey,
		"receiver_id": result.Id,
	})
}

func (s *IntegrationSuite) TestCreateReceiverIntegrationBadResponse() {
	t := s.T()

	params := http.ReceiverInputDto{
		RazaoSocial: "Nome",
		Cpf:         "041.485.353",
		Cnpj:        "",
		Email:       "bryan_barbosa@prcondominios.com.br",
		PixType:     "CPF",
		PixKey:      "041.485.353-92",
	}

	reqBody, _ := json.Marshal(params)

	resp := s.Request("POST", "/create", reqBody)

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	var result http.ErrorOutputDto
	err := s.DecodeBody(resp, &result)

	assert.NoError(t, err)

	receivers := []schemas.Receiver{}
	s.findInDatabase(receivers, map[string]interface{}{
		"corporate_name": params.RazaoSocial,
		"cpf":            params.Cpf,
		"cnpj":           params.Cnpj,
		"email":          params.Email,
		"status":         entity.STATUS_DRAFT,
	})

	assert.Equal(t, 0, len(receivers))

	pixes := []schemas.Pix{}
	s.findInDatabase(pixes, map[string]interface{}{
		"type": params.PixType,
		"key":  params.PixKey,
	})

	assert.Equal(t, 0, len(pixes))
}
