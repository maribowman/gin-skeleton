package service

import (
	"github.com/maribowman/gin-skeleton/app/model"
)

type Service struct {
	databaseClient model.DatabaseClient
	restClient     model.RestClient
}

type ServiceWiring struct {
	DatabaseClient model.DatabaseClient
	RestClient     model.RestClient
}

func NewService(wiring ServiceWiring) *Service {
	return &Service{
		databaseClient: wiring.DatabaseClient,
		restClient:     wiring.RestClient,
	}
}

func (service *Service) BusinessLogic() {
	// only smart code here
}
