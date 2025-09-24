package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type PostgresConfig struct {
	UserName string `env:"POSTGRES_USER" env-required:"true"`
	Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
	Host     string `env:"POSTGRES_HOST" env-required:"true"`
	Port     string `env:"POSTGRES_PORT" env-required:"true"`
	DBName   string `env:"POSTGRES_DB" env-required:"true"`
	SSLMode  string `env:"DB_SSLMODE" env-required:"true"`
}

type AppConfig struct {
	Port     string `env:"APP_PORT" env-required:"true"`
	Address  string `env:"APP_ADDRESS" env-required:"true"`
	LogLevel string `env:"APP_LOG_LEVEL" env-required:"true"`
}

type Config struct {
	Postgres    PostgresConfig
	App         AppConfig
	MigratePath string `env:"MIGRATE_PATH" env-required:"true"`
}

func LoadConfig() *Config {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./.env", cfg)
	if err != nil {
		log.Fatalf("error reading config: %s", err.Error())
	}
	return cfg
}

func (c *Config) GetConnStr() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Postgres.Host, c.Postgres.Port, c.Postgres.UserName, c.Postgres.Password, c.Postgres.DBName, c.Postgres.SSLMode)
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		c.Postgres.UserName, c.Postgres.Password, c.Postgres.Host, c.Postgres.Port, c.Postgres.DBName, c.Postgres.SSLMode)
}
