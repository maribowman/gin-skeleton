package model

type Service interface {
}

type DatabaseClient interface {
	CloseDatabaseConnections()
}

type DemoRestClient interface {
	GetDemoComments() ([]DemoCommentDTO, error)
}
