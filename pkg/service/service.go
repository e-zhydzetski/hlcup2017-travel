package service

import (
	"context"
	"errors"
	"github.com/e-zhydzetski/hlcup2017-travel/pkg/service/internal/app"
	"log"
	"strings"

	"github.com/e-zhydzetski/hlcup2017-travel/pkg/x/xerror"

	"github.com/e-zhydzetski/hlcup2017-travel/pkg/service/internal/infrastructure/dump"
	"github.com/e-zhydzetski/hlcup2017-travel/pkg/service/internal/infrastructure/http"
	"github.com/e-zhydzetski/hlcup2017-travel/pkg/service/internal/infrastructure/options"
	"github.com/e-zhydzetski/hlcup2017-travel/pkg/x/xhttp"
)

type Service struct {
	ListenAddr  string
	OptionsFile string
	DumpSource  string
}

func (s *Service) Start(ctx context.Context) error {
	opt, err := options.NewOptionsFromFile(s.OptionsFile)
	if err != nil {
		return xerror.Combine(err, errors.New("can't load options"))
	}
	log.Println("Options:", *opt)

	var service app.Service
	{
		var d *dump.Dump
		if strings.HasSuffix(s.DumpSource, ".zip") {
			if d, err = dump.LoadFromZip(s.DumpSource); err != nil {
				return xerror.Combine(err, errors.New("can't load dump from zip archive"))
			}
		} else {
			if d, err = dump.LoadFromFolder(s.DumpSource); err != nil {
				return xerror.Combine(err, errors.New("can't load dump from folder"))
			}
		}
		repository := app.NewRepositoryFromDump(d.ToDomain())
		repository.Opt = *opt
		service = repository
	}

	return xhttp.StartServer(ctx, s.ListenAddr, http.NewHandler(service))
}
