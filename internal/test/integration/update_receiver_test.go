package integration

import (
	"encoding/json"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/receiver-management-api/internal/application/entity"
	"github.com/vctaragao/receiver-management-api/internal/http"
	"github.com/vctaragao/receiver-management-api/internal/storage/schemas"
)

func (s *IntegrationSuite) TestUpdateReceiverIntegrationSuccess() {
	t := s.T()

	r, _ := s.seed(entity.STATUS_DRAFT)

	params := http.UpdateInputDto{
		ReceiverId:    r.ID,
		CorporateName: "Gael Bryan Aparício",
		CpfCnpj:       "011.383.228-14",
		Email:         "gael-aparicio91@uniube.br",
		PixType:       "CPF",
		PixKey:        "011.383.228-14",
	}

	reqBody, _ := json.Marshal(params)

	resp := s.Request("PATCH", "/update", reqBody)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	var result http.UpdateOutputDto
	err := s.DecodeBody(resp, &result)
	assert.NoError(t, err)

	receiver := &schemas.Receiver{}
	s.firstInDatabase(receiver, map[string]interface{}{
		"id":             result.ReceiverId,
		"corporate_name": params.CorporateName,
		"cpf_cnpj":       params.CpfCnpj,
		"email":          params.Email,
		"status":         r.Status,
	})

	assert.Equal(t, result.Email, receiver.Email)
	assert.Equal(t, result.Status, receiver.Status)
	assert.Equal(t, result.ReceiverId, receiver.ID)
	assert.Equal(t, result.CpfCnpj, receiver.CpfCnpj)
	assert.Equal(t, result.CorporateName, receiver.CorporateName)

	pix := &schemas.Pix{}
	s.firstInDatabase(pix, map[string]interface{}{
		"type":        params.PixType,
		"key":         params.PixKey,
		"receiver_id": result.ReceiverId,
	})

	assert.Equal(t, result.ReceiverId, receiver.ID)
	assert.Equal(t, result.PixType, pix.Type)
	assert.Equal(t, result.PixKey, pix.Key)
}

func (s *IntegrationSuite) TestUpdateValidatedReceiverOnlyEmailIntegrationSuccess() {
	t := s.T()

	r, _ := s.seed(entity.STATUS_VALID)

	params := http.UpdateInputDto{
		ReceiverId: r.ID,
		Email:      "gael-aparicio91@uniube.br",
	}

	reqBody, _ := json.Marshal(params)

	resp := s.Request("PATCH", "/update", reqBody)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	var result http.UpdateOutputDto
	err := s.DecodeBody(resp, &result)
	assert.NoError(t, err)

	receiver := &schemas.Receiver{}
	s.firstInDatabase(receiver, map[string]interface{}{
		"id":             r.ID,
		"corporate_name": r.CorporateName,
		"cpf_cnpj":       r.CpfCnpj,
		"email":          params.Email,
		"status":         r.Status,
	})

	assert.Equal(t, result.ReceiverId, receiver.ID)
	assert.Equal(t, result.CpfCnpj, receiver.CpfCnpj)
	assert.Equal(t, result.Email, receiver.Email)
	assert.Equal(t, result.CorporateName, receiver.CorporateName)
}

func (s *IntegrationSuite) TestUpdateReceiverIntegrationInvalidCpfBadResponse() {
	t := s.T()

	r, pix := s.seed(entity.STATUS_DRAFT)

	params := http.UpdateInputDto{
		ReceiverId:    r.ID,
		CorporateName: "Gael Bryan Aparício",
		CpfCnpj:       "011.38314",
		Email:         "gael-aparicio91@uniube.br",
		PixType:       "CPF",
		PixKey:        "011.383.228-14",
	}

	reqBody, _ := json.Marshal(params)

	resp := s.Request("PATCH", "/update", reqBody)

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	var result http.ErrorOutputDto
	err := s.DecodeBody(resp, &result)
	assert.NoError(t, err)

	assert.Equal(t, result.Message, "validating receiver: invalid cpf or cnpj")

	s.firstInDatabase(&schemas.Receiver{}, map[string]interface{}{
		"id":             r.ID,
		"corporate_name": r.CorporateName,
		"cpf_cnpj":       r.CpfCnpj,
		"email":          r.Email,
		"status":         entity.STATUS_DRAFT,
	})

	s.firstInDatabase(&schemas.Pix{}, map[string]interface{}{
		"type": pix.Type,
		"key":  pix.Key,
	})
}

func (s *IntegrationSuite) TestUpdatevALIDReceiverIntegrationBadResponse() {
	t := s.T()

	r, pix := s.seed(entity.STATUS_VALID)

	params := http.UpdateInputDto{
		ReceiverId:    r.ID,
		CorporateName: "Gael Bryan Aparício",
		CpfCnpj:       "011.38314",
		Email:         "gael-aparicio91@uniube.br",
		PixType:       "CPF",
		PixKey:        "011.383.228-14",
	}

	reqBody, _ := json.Marshal(params)

	resp := s.Request("PATCH", "/update", reqBody)

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"))

	var result http.ErrorOutputDto
	err := s.DecodeBody(resp, &result)
	assert.NoError(t, err)

	assert.Equal(t, result.Message, "updating receiver: can only update email on validated receiver")

	s.firstInDatabase(&schemas.Receiver{}, map[string]interface{}{
		"id":             r.ID,
		"corporate_name": r.CorporateName,
		"cpf_cnpj":       r.CpfCnpj,
		"email":          r.Email,
		"status":         r.Status,
	})

	s.firstInDatabase(&schemas.Pix{}, map[string]interface{}{
		"type": pix.Type,
		"key":  pix.Key,
	})
}

func (s *IntegrationSuite) seed(status string) (*schemas.Receiver, *schemas.Pix) {
	receiver := &schemas.Receiver{
		CorporateName: "Olivia Daiane Tânia Rezende",
		CpfCnpj:       "106.762.957-20",
		Email:         "bryan_barbosa@prcondominios.com.br",
		Status:        status,
	}

	s.seedDatabase(receiver)

	pix := &schemas.Pix{
		ReceiverId: receiver.ID,
		Type:       "CPF",
		Key:        "106.762.957-20",
	}

	s.seedDatabase(pix)

	return receiver, pix
}
