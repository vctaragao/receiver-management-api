package integration

import (
	"encoding/json"

	"github.com/stretchr/testify/assert"
	"github.com/vctaragao/receiver-management-api/internal/http"
	"github.com/vctaragao/receiver-management-api/internal/storage/schemas"
)

func (s *IntegrationSuite) TestDeleteReceiverIntegrationSuccess() {
	t := s.T()

	r := s.seedInBatch(10)[5]

	params := http.DeleteInputDto{
		ReceiversIds: []uint{r.ID},
	}

	reqBody, _ := json.Marshal(params)

	resp := s.Request("POST", "/delete", reqBody)

	assert.Equal(t, 204, resp.StatusCode)

	receivers := []schemas.Receiver{}
	s.findInDatabase(receivers, map[string]interface{}{
		"id": r.ID,
	})

	pixes := []schemas.Pix{}
	s.findInDatabase(pixes, map[string]interface{}{
		"id": r.Pix[0].ID,
	})

	assert.Equal(t, 0, len(pixes))
}
