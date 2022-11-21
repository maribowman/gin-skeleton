package repository

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

func NewRestClient(baseUrl string, debug bool, timeout int) *resty.Client {
	return resty.New().
		SetDebug(debug).
		SetTimeout(time.Duration(timeout) * time.Second).
		SetBaseURL(baseUrl)
}

func GetSomething(client *resty.Client, path string, headers, pathParameters, requestParameters map[string]string) ([]byte, error) {
	response, err := client.R().
		SetHeaders(headers).
		SetPathParams(pathParameters).
		SetQueryParams(requestParameters).
		Get(path)
	if err != nil {
		return nil, err
	}
	if response.StatusCode() != http.StatusOK || response.StatusCode() != http.StatusNoContent {
		return nil, errors.New(fmt.Sprintf("request failed with status %d", response.StatusCode()))
	}
	return response.Body(), nil
}
