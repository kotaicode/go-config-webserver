package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

type configData struct {
	Source string `json:"source"`
	Value  string `json:"value"`
}

func getConfigData(w http.ResponseWriter, r *http.Request) {
	configs := loadConfigs()
	configsBytes, err := json.Marshal(configs)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(configsBytes)
}

func loadConfigs() []configData {
	var configs []configData

	configs = append(configs, loadFromEnv())
	configs = append(configs, loadFromFile())

	return configs
}

func loadFromEnv() configData {
	var config configData
	envKey := getEnv("SGS_ENVVAR_NAME", "DATA_ENV")

	config.Source = fmt.Sprintf("from env var $%s", envKey)
	config.Value = os.Getenv(envKey)
	return config
}

func loadFromFile() configData {
	var config configData
	fileKey := getEnv("SGS_FILENAME", "/tmp/data/file")
	config.Source = fmt.Sprintf("from file %s", fileKey)
	data, err := os.ReadFile(fileKey)
	if err != nil {
		config.Value = ""
		return config
	}
	config.Value = string(data)
	return config
}
