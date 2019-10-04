package main

import (
	"context"
	"log"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/http"
	"github.com/e-zhydzetski/hlcup2017-travel/internal/options"
	"github.com/e-zhydzetski/hlcup2017-travel/internal/x/xhttp"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/domain"
	"github.com/e-zhydzetski/hlcup2017-travel/internal/dump"
)

func main() {
	ctx := context.Background()

	opt, err := options.NewOptionsFromFile("test/data/TRAIN/data/options.txt")
	if err != nil {
		log.Println("Can't load options:", err)
		return
	}
	log.Println("Options:", *opt)

	var service domain.Service
	{
		d, err := dump.LoadFromFolder("test/data/TRAIN/data")
		if err != nil {
			log.Println("Can't load dump:", err)
			return
		}
		service = domain.NewRepositoryFromDump(d.ToDomain())
	}

	_ = xhttp.StartServer(ctx, ":80", http.NewHandler(service))
}
