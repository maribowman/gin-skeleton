package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/maribowman/gin-skeleton/app/config"
	"github.com/maribowman/gin-skeleton/app/model"
	"net/http"
)

type DemoRestClient struct {
	restyClient *resty.Client
}

func NewDemoRestClient() DemoRestClient {
	return DemoRestClient{
		restyClient: NewRestyClient(config.Config.DemoRestClient.BaseUrl),
	}
}

func (client DemoRestClient) GetDemoUsers(limit int) (comments []model.DemoUserDTO, err error) {
	var requestParams = make(map[string]string)
	requestParams["page"] = "1"
	requestParams["per_page"] = string(rune(limit))

	response, err := GetSomething(client.restyClient, "/public/v2/comments", nil, nil, requestParams)
	if err != nil {
		return
	}
	if response.StatusCode() != http.StatusOK || response.StatusCode() != http.StatusNoContent {
		return nil, errors.New(fmt.Sprintf("request failed with status %d", response.StatusCode()))
	}
	if err = json.Unmarshal(response.Body(), &comments); err != nil {
		return nil, err
	}
	return
}
