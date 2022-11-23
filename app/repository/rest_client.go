package repository

import (
	"github.com/go-resty/resty/v2"
	"github.com/maribowman/gin-skeleton/app/config"
	"time"
)

func NewRestyClient(baseUrl string) *resty.Client {
	return resty.New().
		SetDebug(config.Config.RestConfig.Debug).
		SetTimeout(time.Duration(config.Config.RestConfig.TimeoutSeconds) * time.Second).
		SetBaseURL(baseUrl)
}

func GetSomething(client *resty.Client, path string, headers, pathParameters, requestParameters map[string]string) (*resty.Response, error) {
	return client.R().
		SetHeaders(headers).
		SetPathParams(pathParameters).
		SetQueryParams(requestParameters).
		Get(path)
}
