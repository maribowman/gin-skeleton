package service

import (
	"github.com/maribowman/gin-skeleton/app/model"
	"github.com/maribowman/gin-skeleton/app/model/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetDemoComments(t *testing.T) {
	// given
	tests := []struct {
		testName string
		limit    int
		expected []model.DemoUserDTO
	}{
		{
			testName: "give me MJ as user",
			limit:    5,
			expected: []model.DemoUserDTO{{
				Id:     0,
				Name:   "Michael",
				Email:  "Jordan",
				Gender: "male",
				Status: "retired",
			}},
		},
		{
			testName: "second demo unit test returns empty result",
			limit:    15,
			expected: []model.DemoUserDTO{},
		},
	}
	// and
	mockDemoRestClient := new(mocks.DemoRestClient)
	mockDemoRestClient.On("GetDemoUsers", mock.Anything).
		Return([]model.DemoUserDTO{{
			Id:     0,
			Name:   "Michael",
			Email:  "Jordan",
			Gender: "male",
			Status: "retired",
		}}, nil).
		Once()
	mockDemoRestClient.On("GetDemoUsers", mock.Anything).
		Return([]model.DemoUserDTO{}, nil).
		Once()

	// and
	service := NewService(ServiceWiring{
		DatabaseClient: nil,
		RestClient:     mockDemoRestClient,
	})

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			// when
			actual, err := service.GetDemoUsers(2)

			// then
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
	// and
	assert.True(t, mockDemoRestClient.AssertNumberOfCalls(t, "GetDemoUsers", 2))
}
