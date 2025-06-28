package config

import (
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/subosito/gotenv"
)

var (
	Config EnvConfig
)

type EnvConfig struct {
	Env                   string `envconfig:"ENV"`
	KafkaBrokers          string `envconfig:"KAFKA_BROKERS"`
	KafkaTopic            string `envconfig:"KAFKA_TOPICS"`
	KafkaGroupID          string `envconfig:"KAFKA_GROUP_ID"`
	SMTPHost              string `envconfig:"SMTP_HOST"`
	SMTPUsername          string `envconfig:"SMTP_USERNAME"`
	SMTPPassword          string `envconfig:"SMTP_PASSWORD"`
	SMTPPort              int    `envconfig:"SMTP_PORT"`
	SMTPSkipInsecure      bool   `envconfig:"SMTP_SKIP_INSECURE"`
	SMTPFrom              string `envconfig:"SMTP_FROM"`
	UserServiceGrpcServer string `envconfig:"USER_SERVICE_GRPC_SERVER"`
	TaskServiceGrpcServer string `envconfig:"TASK_SERVICE_GRPC_SERVER"`
}

func LoadEnvConfig() EnvConfig {
	_ = gotenv.Load()

	var cfg EnvConfig

	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}

func (cfg *EnvConfig) GetBrokers() []string {
	if cfg.KafkaBrokers == "" {
		return nil
	}

	return strings.Split(cfg.KafkaBrokers, ",")
}

func (cfg *EnvConfig) GetTopics() []string {
	if cfg.KafkaTopic == "" {
		return nil
	}

	return strings.Split(cfg.KafkaTopic, ",")
}
