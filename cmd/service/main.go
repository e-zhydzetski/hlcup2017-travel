package main

import (
	"context"
	"github.com/e-zhydzetski/hlcup2017-travel/internal/http"
	"github.com/e-zhydzetski/hlcup2017-travel/internal/x/xhttp"
	"log"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/domain"
	"github.com/e-zhydzetski/hlcup2017-travel/internal/dump"
)

func main() {
	ctx := context.Background()

	var service domain.Service
	{
		d, err := dump.LoadFromFolder("test/data/TRAIN/data")
		if err != nil {
			log.Println("Can't load dump:", err)
			return
		}
		service = domain.NewRepositoryFromDump(d.ToDomain())
	}

	xhttp.StartServer(ctx, ":80", http.NewHandler(service))
}
