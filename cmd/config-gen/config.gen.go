package main

import (
    "flag"
    "os"
	"strconv"
	"time"
)

func GetConfig(defaultConfig *Config) *Config {
	cfg := &Config{}
	
	flag.StringVar(&cfg.Filename, "filename", LookupEnvOrString("BLAH_FILENAME", defaultConfig.Filename), "filename blah blah")
	flag.StringVar(&cfg.Prefix, "prefix", LookupEnvOrString("BLAH_PREFIX", defaultConfig.Prefix), "prefix blah blah")
	flag.IntVar(&cfg.Number, "number", LookupEnvOrInt("BLAH_NUMBER", defaultConfig.Number), "nunmber blah blah")
	flag.BoolVar(&cfg.Boolean, "boolean", LookupEnvOrBool("BLAH_BOOLEAN", defaultConfig.Boolean), "bool blah blah")
	flag.DurationVar(&cfg.Duration, "duration", LookupEnvOrDuration("BLAH_DURATION", defaultConfig.Duration), "dur blah blah")
    
	flag.Parse()

	return cfg
}

func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func LookupEnvOrInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(val)
		if err != nil {
			return defaultVal
		}
		return i
	}
	return defaultVal
}

func LookupEnvOrBool(key string, defaultVal bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(val)
		if err != nil {
			return defaultVal
		}
		return b
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
	return "TODO"
}
