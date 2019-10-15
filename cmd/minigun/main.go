package main

import (
	"context"
	"errors"
	"flag"
	"log"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod

	"github.com/e-zhydzetski/hlcup2017-travel/test/load/minigun"
)

type Params struct {
	TargetHostPort string
	AmmoFilePath   string
	LoadProfile    string
}

func (p Params) Validate() error {
	if p.TargetHostPort == "" {
		return errors.New("target host port not defined")
	}
	if p.AmmoFilePath == "" {
		return errors.New("ammo file path not defined")
	}
	if p.LoadProfile == "" {
		return errors.New("load profile not defined")
	}
	return nil
}

func main() {
	var params Params
	flag.StringVar(&params.TargetHostPort, "target", "", "target host port: http://127.0.0.1:8080")
	flag.StringVar(&params.AmmoFilePath, "ammo", "", "path to ammo file: /data/ammo/phase_1_get.ammo")
	flag.StringVar(&params.LoadProfile, "load", "", "load profile: line(1, 100, 30s)")
	flag.Parse()
	if err := params.Validate(); err != nil {
		log.Println(err)
		flag.PrintDefaults()
		return // TODO non-zero exit code
	}

	tl, err := minigun.Fire(context.Background(), params.TargetHostPort, params.AmmoFilePath, params.LoadProfile)
	if err != nil {
		log.Println("minigun error:", err)
		return
	}

	log.Println("Total latency:", tl.Seconds(), "seconds")
}
