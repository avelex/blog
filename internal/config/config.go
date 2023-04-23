package config

import (
	"log"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	InsecureHTTP    bool
	Debug           bool          `env:"DEBUG" env-default:"false"`
	HttpPort        string        `env:"PORT" env-default:"8080"`
	Host            string        `env:"HOST" env-default:"localhost"`
	MongoURI        string        `env:"MONGO_URI" env-default:"mongodb://localhost:27017"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" env-default:"10s"`
	InitTimeout     time.Duration `env:"INIT_TIMEOUT" env-default:"5s"`
}

var config *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{
			InsecureHTTP: true,
		}
		if err := cleanenv.ReadEnv(config); err != nil {
			text, _ := cleanenv.GetDescription(config, nil)
			log.Fatal(text)
		}
	})

	return config
}
