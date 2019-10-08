package helpers

import (
	"io"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"

	"github.com/e-zhydzetski/hlcup2017-travel/test/internal/reader"
)

func ReadVegetaTargets(source io.Reader, targetHostPort string) ([]vegeta.Target, error) {
	r := reader.NewExtendedReader(source)

	var res []vegeta.Target

	for {
		_, err := r.ReadStringUntil('\n') // skip first line
		if err == io.EOF {
			break
		}

		t := vegeta.Target{
			URL:    targetHostPort, // base URL path
			Header: map[string][]string{},
		}
		cl := 0
		for ln := 0; ; ln++ {
			var s string
			if s, err = r.ReadStringUntil('\n'); err != nil {
				return nil, err
			}
			s = strings.TrimSpace(s)
			if s == "" { // empty string before body
				break
			}
			if ln == 0 {
				p := strings.Split(s, " ") // {method} {URL}
				t.Method = p[0]            // {method}
				t.URL += p[1]              // {targetHostPort}{URL}
				continue
			}
			{
				p := strings.Split(s, ":")   // {headerName}: {headerValue}
				n := strings.TrimSpace(p[0]) // {headerName}
				v := strings.TrimSpace(p[1]) // {headerValue}
				t.Header.Add(n, v)
				if n == "Content-Length" {
					cl, _ = strconv.Atoi(v)
				}
			}
		}
		if cl > 0 {
			t.Body = make([]byte, cl)
			_, err := r.Read(t.Body)
			if err != nil {
				return nil, err
			}
			_, _ = r.ReadStringUntil('\n') // skip line break after body
		}
		res = append(res, t)
	}

	log.Println("Loaded", len(res), "targets")

	return res, nil
}

type LinearVegetaPacer struct {
	From     vegeta.Rate
	To       vegeta.Rate
	Duration time.Duration
}

// Цель: зная количество уже сделанных атак и текущее время теста, определить время следующий атаки
// y = kx + b, уравнение прямой линии y - частота обстрела, x - время от начала, b - начальная частота
// k = (max - b) / tMax, max - конечная частота, tMax - время конечной частоты
// S = x * (b + y) / 2, площадь прямоугольной трапеции - количество атак
// 2*S = x * (b + k*x + b) = 2*b*x + k*x^2
// k*x^2 + 2*b*x - 2*S = 0, квадратное уравнение для определения момента времени x когда должна быть S-я атака
// решаем через дискриминант и получаем x = (-2*b + sqrt(4*b^2 +8*k*S)) / (2*k)
func (p LinearVegetaPacer) Pace(elapsed time.Duration, hits uint64) (time.Duration, bool) {
	if elapsed > p.Duration {
		return 0, true
	}

	S := float64(hits + 1)                          // номер искомой атаки
	b := float64(p.From.Freq) / float64(p.From.Per) // начальная частота атак в наносекунду
	max := float64(p.To.Freq) / float64(p.To.Per)   // конечная частота атак в наносекунду
	tMax := p.Duration.Nanoseconds()                // время конечной атаки

	k := (max - b) / float64(tMax)

	x := ((-2 * b) + math.Sqrt(4*b*b+8*k*S)) / (2 * k)

	return time.Duration(x) - elapsed, false
}
