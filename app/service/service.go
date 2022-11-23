package service

import (
	"github.com/maribowman/gin-skeleton/app/model"
)

type Service struct {
	databaseClient model.DatabaseClient
	demoRestClient model.DemoRestClient
}

type ServiceWiring struct {
	DatabaseClient model.DatabaseClient
	RestClient     model.DemoRestClient
}

func NewService(wiring ServiceWiring) *Service {
	return &Service{
		databaseClient: wiring.DatabaseClient,
		demoRestClient: wiring.RestClient,
	}
}

func (service *Service) GetDemoComments() ([]model.DemoCommentDTO, error) {
	// only smart code here
	return service.demoRestClient.GetDemoComments()
}
