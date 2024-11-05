package env

type EnvConfig struct {
	PostgresHost string `env:POSTGRES_HOST`
	PostgresUser string `env:POSTGRES_USER`
	PostgresPass string `env:POSTGRES_PASS`
	PostgresDB   string `env:POSTGRES_DB`
	PostgresPort string `env:POSTGRES_PORT`
}
