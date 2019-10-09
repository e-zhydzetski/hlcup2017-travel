package main

import (
	"bytes"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"

	"github.com/e-zhydzetski/hlcup2017-travel/test/load"
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

	var targets []vegeta.Target
	{
		data, err := ioutil.ReadFile(params.AmmoFilePath)
		if err != nil {
			log.Println(err)
			return
		}
		source := bytes.NewBuffer(data)
		targets, err = load.GenerateVegetaTargetsFromAmmo(source, params.TargetHostPort)
		if err != nil {
			log.Println(err)
			return
		}
	}

	attacker := vegeta.NewAttacker(
		vegeta.Timeout(2 * time.Second), // TODO configurable timeout
	)

	pacer, err := NewPacerFromString(params.LoadProfile)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Attack", params.TargetHostPort, "with profile", pacer, "...")

	var metrics vegeta.Metrics
	for res := range attacker.Attack(vegeta.NewStaticTargeter(targets...), pacer, pacer.DurationLimit(), "Load") {
		metrics.Add(res)
	}
	metrics.Close()

	err = vegeta.NewTextReporter(&metrics).Report(os.Stdout)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Total latency:", metrics.Latencies.Total.Seconds(), "seconds")
}

var lineRegexp = regexp.MustCompile(`line\((\d+),\s*(\d+),\s*(\w+)\)`)

func NewPacerFromString(s string) (load.LimitedPacer, error) {
	// line(1, 100, 30s) -> line from 1/s to 100s during 30s
	if g := lineRegexp.FindStringSubmatch(s); g != nil {
		var err error
		var pacer load.LinearVegetaPacer
		if pacer.From.Freq, err = strconv.Atoi(g[1]); err != nil {
			return nil, err
		}
		pacer.From.Per = time.Second
		if pacer.To.Freq, err = strconv.Atoi(g[2]); err != nil {
			return nil, err
		}
		pacer.To.Per = time.Second
		if pacer.Duration, err = time.ParseDuration(g[3]); err != nil {
			return nil, err
		}
		return pacer, nil
	}
	return nil, errors.New("invalid load profile: " + s)
}
