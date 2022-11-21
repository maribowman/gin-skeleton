package config

type DatabaseConfig struct {
	Postgres postgresConfig
}

type postgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SslMode  string
	TimeZone string
}
