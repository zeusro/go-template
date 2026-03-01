package config

import (
	"log"

	"github.com/spf13/viper"
	"github.com/zeusro/go-template/function/web/translate/model"
)

var configPath string = ".config.yaml"

func init() {
	// 根据环境变量设定不同的配置路径，可按需开启
	// e := os.Getenv("ENV")
	// if e == "dev" {
	// 	configPath = "config.yaml"
	// }
	// if e == "prod" {
	// 	configPath = "config-prod.yaml"
	// }
	// if e == "test" {
	// 	configPath = "config-test.yaml"
	// }
}

type Config struct {
	Debug                    bool         `mapstructure:"debug"`
	Gin                      GinConfig    `mapstructure:"web"`
	Log                      LogConfig    `mapstructure:"log"`
	JWT                      JWT          `mapstructure:"jwt"`
	Cities                   []model.City `yaml:"cities"`
	MinimumDeviationDistance float64      `mapstructure:"minimum_deviation_distance"` // 最小偏差距离
	OutputFormat             string       `mapstructure:"output"`                     // 输出形式
	Database                 DatabaseConfig `mapstructure:"database"`
	Redis                    RedisConfig    `mapstructure:"redis"`
	GRPC                     GRPCConfig     `mapstructure:"grpc"`
	Observability            ObservabilityConfig `mapstructure:"observability"`
	Security                 SecurityConfig      `mapstructure:"security"`
	RateLimit                RateLimitConfig     `mapstructure:"rate_limit"`
}

type JWT struct {
	SigningKey []byte
}

type GinConfig struct {
	Port int  `mapstructure:"port"`
	CORS bool `mapstructure:"cors"`
}

type LogConfig struct {
	Path  string `mapstructure:"path"`
	Level string `mapstructure:"level"`
}

type DatabaseConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"dbname"`
	SSLMode         string `mapstructure:"sslmode"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"` // seconds
	ConnMaxIdleTime int    `mapstructure:"conn_max_idle_time"` // seconds
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type GRPCConfig struct {
	Port int `mapstructure:"port"`
}

type ObservabilityConfig struct {
	Enabled      bool   `mapstructure:"enabled"`
	Prometheus   bool   `mapstructure:"prometheus"`
	OpenTelemetry bool  `mapstructure:"open_telemetry"`
	Endpoint     string `mapstructure:"endpoint"`
}

type SecurityConfig struct {
	OAuth2 OAuth2Config `mapstructure:"oauth2"`
	Audit  AuditConfig  `mapstructure:"audit"`
}

type OAuth2Config struct {
	Enabled      bool   `mapstructure:"enabled"`
	Provider     string `mapstructure:"provider"` // oidc, oauth2-proxy, zitadel, keycloak, okta
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	IssuerURL    string `mapstructure:"issuer_url"`
	RedirectURL  string `mapstructure:"redirect_url"`
}

type AuditConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	Table   string `mapstructure:"table"`
}

type RateLimitConfig struct {
	Enabled  bool `mapstructure:"enabled"`
	Requests int  `mapstructure:"requests"` // requests per second
	Burst    int  `mapstructure:"burst"`
}

func NewFileConfig() Config {
	var config Config

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("无法读取配置文件:", err.Error())
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln("无法解析配置文件:", err.Error())
	}

	return config
}
