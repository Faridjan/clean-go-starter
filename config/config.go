package config

import (
	"os"
	"strings"
	"sync"
	"tiny-template/pkg"
	"tiny-template/pkg/logger"

	"github.com/joho/godotenv"
)

var once sync.Once
var cfg Config

type Config struct {
	// TODO/READ:
	// - держать в одной структуре конфиги на мой взгляд удобнее: быстрый поиск переменных, удобный рефакторинг
	// - писать в камелкейсе сразу с суффиксом: легче копировать, переносить и ясность состояния нейминга переменных сред
	DB struct {
		MaxConnections int    `env:"APP_CONFIG_DB_MAX_CONNS" required:"true"`
		Host           string `env:"APP_CONFIG_DB_HOST" required:"true"`
		Port           int    `env:"APP_CONFIG_DB_PORT" required:"true"`
		Name           string `env:"APP_CONFIG_DB_NAME" required:"true"`
		User           string `env:"APP_CONFIG_DB_USER" required:"true"`
		Pass           string `env:"APP_CONFIG_DB_PASS" required:"true"`
	}
	ElasticConfig struct {
		ConnectionURL []string `env:"APP_CONFIG_ELASTIC_CONNECTION_URLS" required:"true"`
		SetSniff      bool     `env:"APP_CONFIG_ELASTIC_SET_SNIFF" required:"true"`
		Username      string   `env:"APP_CONFIG_ELASTIC_USERNAME" required:"true"`
		Password      string   `env:"APP_CONFIG_ELASTIC_PASSWORD" required:"true"`
	}
	SSOConfig struct {
		ProjectSlug     string `env:"APP_CONFIG_SSO_PROJECT_SLUG" required:"true"`
		KeycloakKeysURL string `env:"APP_CONFIG_SSO_KEYCLOAK_KEYS_URL" required:"true"`
	}
	RPC struct {
		MediaUploader string `env:"APP_CONFIG_MEDIA_UPLOADER_URL" required:"true"`
	}
	ConfigOne string   `env:"APP_CONFIG_CONFIG_ONE" required:"true"`
	ConfigTwo []string `env:"APP_CONFIG_CONFIG_TWO" required:"true"`
}

func InitConfigs() {
	once.Do(func() {
		// TODO/READ: Если APP_CONFIG_IS_LOCAL_MODE=true то environment прогрузятся с .env
		if IsLocalMode() {
			if err := godotenv.Load(); err != nil {
				logger.Fatal("godotenv.Load", err.Error())
			}
		}

		if err := pkg.LoadConfig(&cfg); err != nil {
			logger.Fatal(err.Error())
		}
	})
}

func GetConfig() Config {
	return cfg
}

func IsLocalMode() bool {
	return strings.ToLower(os.Getenv("APP_CONFIG_IS_LOCAL_MODE")) == "true"
}
