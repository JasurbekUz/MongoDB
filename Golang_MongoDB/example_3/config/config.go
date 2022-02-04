package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment     string
	MongoURI      	string
	MongoDatabase	string
	httpPort		string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	
	// dbconfig
	c.MongoURI = cast.ToString(getOrReturnDefault("MONGO_URI", "mongodb://localhost:27017"))
	c.MongoDatabase = cast.ToString(getOrReturnDefault("MONGO_DATABASE", "users"))

	// service_port
	c.httpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
