package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// 1. Define Config structure (represents full YAML config)
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

// 2. Define HTTPServer structure (nested YAML section)
type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

// 3. Create MustLoad() function (loads configuration before app starts)
func MustLoad() *Config {

	// 4. Inside MustLoad():

	// STEP A: Get config path from environment variable (CONFIG_PATH)
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	// STEP B: If not found in ENV, get from CLI flag (-config)
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flags

		// STEP C: If still empty → stop program
		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	// STEP D: Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {

		// STEP E: If file does not exist → stop program
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	// STEP F: Create empty Config variable
	var cfg Config

	// STEP G: Read YAML file into Config struct
	err := cleanenv.ReadConfig(configPath, &cfg)

	// STEP H: If reading fails → stop program
	if err != nil {
		log.Fatalf("Cannot read config file: %s", err.Error())
	}

	// STEP I: Return Config object
	return &cfg

}