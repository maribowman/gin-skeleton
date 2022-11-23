package mocks

import (
	"github.com/maribowman/gin-skeleton/app/model"
	"github.com/stretchr/testify/mock"
)

type DemoRestClient struct {
	mock.Mock
}

func (mock *DemoRestClient) GetDemoUsers(limit int) (comments []model.DemoUserDTO, err error) {
	returns := mock.Called(limit)

	var r0 []model.DemoUserDTO
	if returnFunc, ok := returns.Get(0).(func(int) []model.DemoUserDTO); ok {
		r0 = returnFunc(limit)
	} else {
		if returns.Get(0) != nil {
			r0 = returns.Get(0).([]model.DemoUserDTO)
		}
	}

	var r1 error
	if returnFunc, ok := returns.Get(1).(func(int) error); ok {
		r1 = returnFunc(limit)
	} else {
		r1 = returns.Error(1)
	}

	return r0, r1
}
