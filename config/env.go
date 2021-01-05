package config

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// GoDotEnvVariable loads a variable from the .env file
func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Error("Error loading .env file", err)
	}

	envVariable, variableExists := os.LookupEnv(key)
	if !variableExists {
		log.Error("Couldn't find variable : ", key)
	}

	return envVariable
}
