package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	envConfigPath = "CONFIG_PATH"
)

var (
	globalConfig Config //nolint: gochecknoglobals

	ErrConfigFileNotParsed = errors.New("config file does not parsed")
	ErrConfigPathIsEmpty   = errors.New("config path is empty")
	ErrConfigFileNotExists = errors.New("config file does not exist")
)

type Config struct {
	Common CommonConfig `yaml:"common"`
	GRPC   GRPCConfig   `yaml:"grpc"`
	AuthDB AuthDB       `yaml:"authDb"`
}

func LoadConfig() error {
	configPath, err := fetchConfigPath()
	if err != nil {
		return err
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return fmt.Errorf("%w: %s", ErrConfigFileNotParsed, err.Error())
	}

	globalConfig = cfg

	return nil
}

func SetDefault(cfg Config) {
	globalConfig = cfg
}

func Default() Config {
	return globalConfig
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func fetchConfigPath() (string, error) {
	var configPath string

	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	if configPath == "" {
		configPath = os.Getenv(envConfigPath)
	}

	if configPath == "" {
		return "", ErrConfigPathIsEmpty
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return "", fmt.Errorf("%w: %s", ErrConfigFileNotExists, configPath)
	}

	return configPath, nil
}

type CommonConfig struct {
	Version  string
	Env      string `env-default:"local" yaml:"env"`
	LogLevel string `env:"LOG_LEVEL"     env-default:"info"`
}

type AuthDB struct {
	URL                string        `env:"AUTH_DB_URL"           env-required:"true"`
	MigrationURL       string        `env:"AUTH_DB_MIGRATION_URL" env-required:"true"`
	PoolMaxConnections int32         `env-required:"true"         yaml:"poolMaxConnections"`
	HealthCheckPeriod  time.Duration `env-required:"true"         yaml:"healthCheckPeriod"`
}
