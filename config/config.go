package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type (
	// Config
	Config struct {
		Redis          Redis    `yaml:"redis_setting" mapstructure:"redis_setting"`
		Logger         Logger   `yaml:"logger" mapstructure:"logger"`
		Postgres       Postgres `yaml:"postgres" mapstructure:"postgres"`
		Http           Http     `yaml:"http" mapstructure:"http"`
		RPC            RPC      `yaml:"rpc" mapstructure:"rpc"`
		Kafka          Kafka    `yaml:"kafka" mapstructure:"kafka"`
		ServiceName    string   `yaml:"service_name" mapstructure:"service_name"`
		TestOverride   string   `yaml:"test_override" mapstructure:"test_override"`
		HttpTimeout    int      `yaml:"http_timeout" mapstructure:"http_timeout"`
		SignalShutDown uint8    `yaml:"signal_shutdown" mapstructure:"signal_shutdown"`
	}

	Http struct {
		Port int `yaml:"port" mapstructure:"port"`
	}
	RPC struct {
		Port int `yaml:"port" mapstructure:"port"`
	}

	Logger struct {
		Mode              string `yaml:"mode" mapstructure:"mode"`
		Encoding          string `yaml:"encoding" mapstructure:"encoding"`
		Level             string `yaml:"level" mapstructure:"level"`
		ZapType           string `yaml:"zap_type" mapstructure:"zap_type"`
		DisableCaller     bool   `yaml:"disable_caller" mapstructure:"disable_caller"`
		DisableStacktrace bool   `yaml:"disable_stacktrace" mapstructure:"disable_stacktrace"`
		LogFile           bool   `yaml:"log_file" mapstructure:"log_file"`
		Payload           bool   `yaml:"payload" mapstructure:"payload"`
	}
	LoggerMarking struct {
		Request  map[string][]string `yaml:"marking_request" mapstructure:"marking_request"`
		Response map[string][]string `yaml:"marking_response" mapstructure:"marking_response"`
		Message  map[string][]string `yaml:"marking_message" mapstructure:"marking_message"`
	}

	// Postgres
	Postgres struct {
		Username    string `yaml:"username" mapstructure:"username"`
		Password    string `yaml:"password" mapstructure:"password"`
		Database    string `yaml:"database" mapstructure:"database"`
		Host        string `yaml:"host" mapstructure:"host"`
		SslMode     string `yaml:"sslmode" mapstructure:"sslmode"`
		Migrate     bool   `yaml:"migrate" mapstructure:"migrate"`
		MaxIdleConn int    `yaml:"max_idle_conn" mapstructure:"max_idle_conn"`
		MaxIdleTime int    `yaml:"max_idle_time" mapstructure:"max_idle_time"`
		MaxOpenConn int    `yaml:"max_open_conn" mapstructure:"max_open_conn"`
		MaxLifeTime int    `yaml:"max_life_time" mapstructure:"max_life_time"` // hour
		Port        int    `yaml:"port" mapstructure:"port"`
		IsDebug     bool   `yaml:"is_debug" mapstructure:"is_debug"`
	}

	Redis struct {
		Addrs               []string
		Password            string `yaml:"password" mapstructure:"password"`
		Database            int    `yaml:"database" mapstructure:"database"`
		PoolSize            int    `yaml:"pool_size" mapstructure:"pool_size"`
		DialTimeoutSeconds  int    `yaml:"dial_timeout_seconds" mapstructure:"dial_timeout_seconds"`
		ReadTimeoutSeconds  int    `yaml:"read_timeout_seconds" mapstructure:"read_timeout_seconds"`
		WriteTimeoutSeconds int    `yaml:"write_timeout_seconds" mapstructure:"write_timeout_seconds"`
		IdleTimeoutSeconds  int    `yaml:"idle_timeout_seconds" mapstructure:"idle_timeout_seconds"`
	}

	Kafka struct {
		KafkaReader KafkaReader `yaml:"reader" mapstructure:"reader"`
		KafkaWriter KafkaWriter `yaml:"writer" mapstructure:"writer"`
	}
	KafkaReader struct {
		BrokerAddress []string `yaml:"broker_address" mapstructure:"broker_address"`
		Topic         string   `yaml:"topic" mapstructure:"topic"`
		GroupID       string   `yaml:"group_id" mapstructure:"group_id"`
		Enable        bool     `yaml:"enable" mapstructure:"enable"`
		WorkerPool    uint8    `yaml:"worker_pool" mapstructure:"worker_pool"`
		SleepTime     uint8    `yaml:"sleep_time" mapstructure:"sleep_time"`
	}

	KafkaWriter struct {
		BrokerAddress []string `yaml:"broker_address" mapstructure:"broker_address"`
		Topic         string   `yaml:"topic" mapstructure:"topic"`
		BatchSize     int      `yaml:"batch_size" mapstructure:"batch_size"`
		Enable        bool     `yaml:"enable" mapstructure:"enable"`
	}
)

func getConfigName() string {
	mode := "dev"
	switch os.Getenv("MODE") {
	case "uat":
		mode = "uat"
	case "prod":
		mode = "prod"
	}
	return mode
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	path := "config"

	vn := viper.New()
	configName := getConfigName()
	vn.AddConfigPath(path)
	vn.SetConfigName(configName)
	vn.SetConfigType("yaml")
	vn.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	vn.AutomaticEnv()

	err := vn.ReadInConfig()
	if err != nil {
		return cfg, err

	}

	for _, key := range vn.AllKeys() {
		str := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
		if configName != "prod" {
			log.Default().Println(key, str, vn.Get(key))
		}
		vn.BindEnv(key, str)
	}

	err = vn.Unmarshal(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, err
}
