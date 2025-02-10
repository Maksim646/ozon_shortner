package config

type Config struct {
	DbType        string `envconfig:"DB_TYPE" default:"inmemory"`
	Addr          string `envconfig:"ADDR" default:"8000"`
	LogLevel      string `envconfig:"LOG_LEVEL"`
	MigrationsDir string `envconfig:"MIGRATIONS_DIR" default:"../../internal/database/postgresql/migrations"`
	PostgresURI   string `envconfig:"POSTGRES_URI" default:"postgres://postgres:ozon_shortner@localhost:5448/ozon_shortner_db?sslmode=disable"`
}
