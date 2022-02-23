package main

import (
	"github.com/bloom42/rz-go"
	"github.com/bloom42/rz-go/log"
	"os"
)

func initLogger() {
	env := os.Getenv("GO_ENV")
	hostname, _ := os.Hostname()

	// update global logger's context fields
	log.SetLogger(log.With(rz.Fields(
		rz.String("hostname", hostname), rz.String("environment", env),
	)))

	if env == "production" {
		log.SetLogger(log.With(rz.Level(rz.InfoLevel)))
	}
}
