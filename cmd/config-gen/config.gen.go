package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func GetConfig(defaultConfig *Config) *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.Filename, "filename", LookupEnvOrString("BLAH_FILENAME", defaultConfig.Filename), "filename blah blah")
	flag.StringVar(&cfg.Prefix, "prefix", LookupEnvOrString("BLAH_PREFIX", defaultConfig.Prefix), "prefix blah blah")
	flag.DurationVar(&cfg.Number, "number", LookupEnvOrDuration("BLAH_NUMBER", defaultConfig.Number), "nunmber blah blah")

	flag.Parse()

	return cfg
}

func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func LookupEnvOrDuration(key string, defaultVal time.Duration) time.Duration {
	if val, ok := os.LookupEnv(key); ok {
		d, err := time.ParseDuration(val)
		if err != nil {
			return defaultVal
		}
		return d
	}
	return defaultVal
}

func (cfg *Config) String() string {
	output := "Config {\n"

	output += fmt.Sprintf("\t%s: %s,\n", "Filename", cfg.Filename)
	output += fmt.Sprintf("\t%s: %s,\n", "Prefix", cfg.Prefix)

	output += "}"
	return output
}
