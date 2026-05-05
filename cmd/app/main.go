package main

import (
	"os"

	"github.com/ghostIB/lab1-tooling/internal"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("Error reading config file")
	}

	appName := viper.GetString("app_name")
	log.Info().Str("app", appName).Msg("Configuration loaded successfully")

	res := internal.Add(10, 20)
	log.Info().Int("result", res).Msg("Math Add operation")

	val, err := internal.Divide(10, 2)
	if err != nil {
		log.Error().Err(err).Msg("Division failed")
		return
	}
	log.Info().Float64("result", val).Msg("Math Divide operation")
}
