package main

import (
	"context"
	"os"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/app"
)

func main() {
	OptionsFile := "test/data/TRAIN/data/options.txt"
	DumpSource := "test/data/TRAIN/data"
	if os.Getenv("DOCKER") == "1" {
		OptionsFile = "/tmp/data/options.txt"
		DumpSource = "/tmp/data/data.zip"
	}

	ctx := context.Background()
	service := app.Service{
		ListenAddr:  ":80",
		OptionsFile: OptionsFile,
		DumpSource:  DumpSource,
	}
	_ = service.Start(ctx)
}
