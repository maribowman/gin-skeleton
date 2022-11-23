package config

type RestConfig struct {
	Debug          bool
	TimeoutSeconds int
	DemoRestClient struct {
		BaseUrl string
	}
}
