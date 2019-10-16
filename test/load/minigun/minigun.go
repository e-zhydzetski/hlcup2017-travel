package minigun

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/e-zhydzetski/hlcup2017-travel/test/internal/rawhttp"

	"github.com/e-zhydzetski/hlcup2017-travel/test/internal/helpers"
	"github.com/e-zhydzetski/hlcup2017-travel/test/load"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	influxdb "github.com/influxdata/influxdb1-client/v2"
)

func Fire(ctx context.Context, target, ammoFile, profile string) (time.Duration, error) {
	data, err := ioutil.ReadFile(ammoFile)
	if err != nil {
		return 0, err
	}
	source := bytes.NewBuffer(data)
	ammo, err := helpers.ReadAmmo(source)
	if err != nil {
		return 0, err
	}

	mag := StaticMagazine{
		Ammo: ammo,
		idx:  0,
	}

	pacer, err := NewPacerFromString(profile)
	if err != nil {
		return 0, err
	}

	var wg sync.WaitGroup
	results := make(chan *Result)
	bullets := make(chan *helpers.AmmoEntry)

	for i := 0; i < 10; i++ {
		addWorker(&wg, target, bullets, results)
	}

	influxdbClient, err := influxdb.NewHTTPClient(influxdb.HTTPConfig{
		Addr: "http://192.168.99.100:8086",
	})
	if err != nil {
		log.Println("Can't create influxdb client:", err)
		return 0, err
	}
	defer influxdbClient.Close()
	bp, err := influxdb.NewBatchPoints(influxdb.BatchPointsConfig{
		Database: "metrics",
	})
	if err != nil {
		log.Println("Can't create batch points with influxdb client:", err)
		return 0, err
	}

	log.Println("Attack", target, "with profile", pacer, "...")

	go func() {
		defer close(results)
		defer wg.Wait()
		defer close(bullets)

		begin := time.Now()
		count := uint64(0)
		for {
			elapsed := time.Since(begin)
			wait, stop := pacer.Pace(elapsed, count)
			if stop {
				return
			}
			time.Sleep(wait)

			select {
			case <-ctx.Done():
				return
			case bullets <- mag.NextBullet():
				count++
				continue
			default:
				// all workers are blocked. start one more and try again
				addWorker(&wg, target, bullets, results)
			}
		}
	}()

	totalLatency := time.Duration(0)
	totalCount := 0
	for res := range results {
		totalLatency += res.Latency
		totalCount++
		p, err := influxdb.NewPoint(
			"response",
			map[string]string{
				"code":  strconv.Itoa(res.Code),
				"index": strconv.Itoa(totalCount), // to prevent duplicates
			},
			map[string]interface{}{
				"latency": res.Latency.Nanoseconds(),
			},
			res.ReqTime,
		)
		if err != nil {
			log.Println("Can't create points with influxdb client:", err)
			return 0, err
		}
		bp.AddPoint(p)
	}

	if err := influxdbClient.Write(bp); err != nil {
		log.Println("Can't write batch points with influxdb client:", err)
		return 0, err
	}

	log.Println("Total shots:", totalCount)

	return totalLatency, nil
}

func addWorker(wg *sync.WaitGroup, target string, bullets <-chan *helpers.AmmoEntry, results chan<- *Result) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for b := range bullets {
			start := time.Now()
			code, _, err := rawhttp.SendRequest(target, b.Request)
			l := time.Since(start)
			results <- &Result{
				ReqTime: start,
				Latency: l,
				Code:    code,
				Error:   err,
			}
		}
	}()
}

type Result struct {
	ReqTime time.Time
	Latency time.Duration
	Code    int
	Error   error
}

type Magazine interface {
	NextBullet() *helpers.AmmoEntry
}

type StaticMagazine struct {
	Ammo []*helpers.AmmoEntry
	idx  int32
}

func (m *StaticMagazine) NextBullet() *helpers.AmmoEntry {
	return m.Ammo[atomic.AddInt32(&m.idx, 1)%int32(len(m.Ammo))]
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
