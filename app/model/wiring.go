package model

type Service interface {
}

type DatabaseClient interface {
	CloseDatabaseConnections()
}

type RestClient interface {
}
