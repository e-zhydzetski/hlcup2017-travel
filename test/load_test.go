package test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/e-zhydzetski/hlcup2017-travel/test/internal/helpers"

	vegeta "github.com/tsenart/vegeta/lib"
)

func TestLoad(t *testing.T) {
	const level = "TRAIN"

	var latency float64
	latency += load(t, level, "phase_1_get")
	//latency += load(t, level, "phase_2_post")
	//latency += load(t, level, "phase_3_get")
	t.Log(latency, "seconds")
}

func load(t *testing.T, level, phaseName string) float64 {
	data, err := ioutil.ReadFile("data/" + level + "/ammo/" + phaseName + ".ammo")
	if err != nil {
		t.Fatal(err)
	}
	source := bytes.NewBuffer(data)
	targets, err := helpers.ReadVegetaTargets(source, "http://127.0.0.1")
	if err != nil {
		t.Fatal(err)
	}

	attacker := vegeta.NewAttacker(
		vegeta.Timeout(2 * time.Second),
	)

	var metrics vegeta.Metrics
	for res := range attacker.Attack(
		vegeta.NewStaticTargeter(targets...),
		helpers.LinearVegetaPacer{
			From:     vegeta.Rate{Freq: 1, Per: time.Second},
			To:       vegeta.Rate{Freq: 1000, Per: time.Second},
			Duration: 30 * time.Second,
		},
		30*time.Second,
		"Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	err = vegeta.NewTextReporter(&metrics).Report(os.Stdout)
	if err != nil {
		t.Fatal(err)
	}

	return metrics.Latencies.Total.Seconds()
}
