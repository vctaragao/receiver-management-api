package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	net_http "net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

type Helper struct {
}

func (Helper) Request(method, uri string, reqBody []byte) *net_http.Response {
	req := createRequest(method, uri, reqBody)
	client := &net_http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("an error accoured during request: %v", err)
		panic(err)
	}

	return resp
}

func createRequest(method, uri string, reqBody []byte) *net_http.Request {
	req := httptest.NewRequest(method, "http://localhost:1323"+uri, bytes.NewBuffer(reqBody))
	req.RequestURI = ""
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	return req
}

func (Helper) DecodeBody(resp *net_http.Response, dto interface{}) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.Unmarshal(body, &dto)
	if err != nil {
		return err
	}

	return nil
}
