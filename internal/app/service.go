package app

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/x/xerror"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/domain"
	"github.com/e-zhydzetski/hlcup2017-travel/internal/dump"
	"github.com/e-zhydzetski/hlcup2017-travel/internal/http"
	"github.com/e-zhydzetski/hlcup2017-travel/internal/options"
	"github.com/e-zhydzetski/hlcup2017-travel/internal/x/xhttp"
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

	var service domain.Service
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
		repository := domain.NewRepositoryFromDump(d.ToDomain())
		repository.Opt = *opt
		service = repository
	}

	return xhttp.StartServer(ctx, s.ListenAddr, http.NewHandler(service))
}
