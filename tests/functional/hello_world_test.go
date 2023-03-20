package functional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ResponseData struct {
	Message string `json:"message"`
}

func TestHelloWorldSuccess(t *testing.T) {
	resp := Request("GET", "/", nil)

	assert.Equal(t, 200, resp.StatusCode)

	var responseData ResponseData
	err := DecodeBody(resp, &responseData)
	assert.NoError(t, err)

	assert.Equal(t, "Hello World", responseData.Message)

}
