package main

import (
	"context"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/app"
)

func main() {
	ctx := context.Background()
	service := app.Service{
		ListenAddr:  ":80",
		OptionsFile: "test/data/TRAIN/data/options.txt",
		DumpSource:  "test/data/TRAIN/data",
	}
	_ = service.Start(ctx)
}
