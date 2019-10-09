package main

import (
	"context"
	"log"
	"os"

	"github.com/pkg/profile"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/app"
)

func main() {
	profilePath := os.Getenv("PROFILE_PATH")
	if profilePath != "" {
		defer profile.Start(
			profile.ProfilePath(profilePath),
			profile.NoShutdownHook,
			profile.CPUProfile,
		).Stop()
	}

	OptionsFile := "test/data/TRAIN/data/options.txt"
	DumpSource := "test/data/TRAIN/data/data.zip"
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
	err := service.Start(ctx)
	log.Println("Service stopped:", err)
}
