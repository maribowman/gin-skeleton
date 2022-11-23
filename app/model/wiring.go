package model

type Service interface {
}

type DatabaseClient interface {
	CloseDatabaseConnections()
}

type DemoRestClient interface {
	GetDemoUsers(limit int) ([]DemoUserDTO, error)
}
