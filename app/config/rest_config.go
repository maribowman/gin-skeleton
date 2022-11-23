package config

type RestConfig struct {
	Debug          bool
	TimeoutSeconds int
}

type DemoRestClientConfig struct {
	BaseUrl string
}
