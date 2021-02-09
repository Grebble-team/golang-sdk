package setting

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"time"
)

type ServerConfig struct {
	GrpcPort     int           `yaml:"grpcPort" envconfig:"SERVER_GRPC_PORT"`
	ReadTimeout  time.Duration `yaml:"-" envconfig:"-"`
	WriteTimeout time.Duration `yaml:"-" envconfig:"-"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
}

var Settings = &Config{}

func init() {
	SetupSettings()
}

func SetupSettings() {
	var err error

	err = envconfig.Process("", Settings)
	if err != nil {
		log.Fatalf("setting, fail to get from env': %v", err)
	}
	if Settings.Server.GrpcPort <= 0 {
		Settings.Server.GrpcPort = 5000
	}
	Settings.Server.ReadTimeout = Settings.Server.ReadTimeout * time.Second
	Settings.Server.WriteTimeout = Settings.Server.WriteTimeout * time.Second

}
