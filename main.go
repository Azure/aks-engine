package main

import (
	"github.com/Azure/aks-engine/cmd"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(colorable.NewColorableStdout())
	if err := cmd.NewRootCmd().Execute(); err != nil {
		log.Fatalln(err)
	}
}
