package main

import (
	"context"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/app"
)

func main() {
	ctx := context.Background()
	service := app.Service{
		ListenAddr:  ":80",
		OptionsFile: "/tmp/data/options.txt",
		DumpSource:  "/tmp/data/data.zip",
	}
	_ = service.Start(ctx)
}
