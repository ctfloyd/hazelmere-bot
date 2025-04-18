package main

import (
	"github.com/ctfloyd/hazelmere-bot/src/internal"
	"github.com/ctfloyd/hazelmere-commons/pkg/hz_config"
	"os"
	"os/signal"
)

func main() {
	cfg := hz_config.NewConfigWithAutomaticDetection()
	err := cfg.Read()
	if err != nil {
		panic(err)
	}

	app := internal.Application{}
	app.Initialize(cfg)
	defer app.Cleanup()

	app.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
