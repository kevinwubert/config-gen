package main

import (
	"time"

	configgen "github.com/kevinwubert/config-gen/pkg/config-gen"
	log "github.com/sirupsen/logrus"
)

//go:generate config-gen --type Config --prefix
type Config struct {
	Filename string        `description:"filename blah blah" secret:"true"`
	Prefix   string        `description:"prefix blah blah" secret:"true"`
	Number   int           `description:"nunmber blah blah" secret:"true"`
	Boolean  bool          `description:"bool blah blah" secret:"true"`
	Duration time.Duration `description:"dur blah blah" secret:"true"`
}

var defaultConfig = Config{
	Filename: "./cmd/config-gen/main.go",
	Prefix:   "arst",
}

func main() {
	cfg := GetConfig(&defaultConfig)

	err := configgen.Generate(cfg.Filename, cfg.Prefix)
	if err != nil {
		log.WithError(err).Error("error generating config.gen.go")
	}
}
