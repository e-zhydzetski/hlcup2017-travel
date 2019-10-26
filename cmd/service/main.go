package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/e-zhydzetski/hlcup2017-travel/pkg/service"

	"github.com/pkg/profile"
)

func main() {
	profilePath := os.Getenv("PROFILE_PATH")
	if profilePath != "" {
		binaryPath := os.Args[0]
		if err := copyFile(binaryPath, profilePath+"/"+filepath.Base(binaryPath)); err != nil {
			log.Println("Can't copy profiled binary:", err)
			return
		}
		defer profile.Start(
			profile.ProfilePath(profilePath),
			profile.NoShutdownHook,
			profile.CPUProfile,
		).Stop()
	}

	log.Println("GOMAXPROCS =", runtime.GOMAXPROCS(0))

	OptionsFile := "test/data/TRAIN/data/options.txt"
	DumpSource := "test/data/TRAIN/data/data.zip"
	if os.Getenv("DOCKER") == "1" {
		OptionsFile = "/tmp/data/options.txt"
		DumpSource = "/tmp/data/data.zip"
	}

	ctx := context.Background()
	srv := service.Service{
		ListenAddr:  ":80",
		OptionsFile: OptionsFile,
		DumpSource:  DumpSource,
	}
	err := srv.Start(ctx)
	log.Println("Service stopped:", err)
}

func copyFile(src, dst string) (err error) {
	var data []byte
	if data, err = ioutil.ReadFile(src); err != nil {
		return
	}
	if err = ioutil.WriteFile(dst, data, 0644); err != nil {
		return
	}
	return
}
