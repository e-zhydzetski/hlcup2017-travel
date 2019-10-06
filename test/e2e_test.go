package test

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/app"

	"github.com/e-zhydzetski/hlcup2017-travel/test/internal/helpers"
	"github.com/e-zhydzetski/hlcup2017-travel/test/internal/rawhttp"

	"github.com/stretchr/testify/require"
)

func TestE2E(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	service := app.Service{
		ListenAddr:  ":8080",
		OptionsFile: "data/TRAIN/data/options.txt",
		DumpFolder:  "data/TRAIN/data",
	}
	go service.Start(ctx)

	data, err := ioutil.ReadFile("data/TRAIN/ammo/phase_1_get.ammo")
	if err != nil {
		t.Fatal(err)
	}
	source := bytes.NewBuffer(data)
	ammo, err := helpers.ReadAmmo(source)
	if err != nil {
		t.Fatal(err)
	}

	data, err = ioutil.ReadFile("data/TRAIN/answers/phase_1_get.answ")
	if err != nil {
		t.Fatal(err)
	}
	source = bytes.NewBuffer(data)
	answers, err := helpers.ReadAnswers(source)
	if err != nil {
		t.Fatal(err)
	}

	if len(ammo) != len(answers) {
		t.Fatal("Ammo incompatible with answers. Ammo size:", len(ammo), ", answers size:", len(answers))
	}

	// TODO maybe check port listening (health check) before test requests

	for i := range ammo {
		bullet := ammo[i]
		answer := answers[i]

		t.Run(strconv.Itoa(i)+": "+answer.Name, func(t *testing.T) {
			code, resp, err := rawhttp.SendRequest("127.0.0.1:8080", bullet.Request)
			if err != nil {
				t.Fatal(err)
			}
			if answer.Code != code {
				t.Fatal("Unexpected code. Expected", answer.Code, ", got", code)
			}
			respMap := map[string]interface{}{}
			if len(resp) > 0 {
				err = json.Unmarshal(resp, &respMap)
				if err != nil {
					t.Fatal(err)
				}
			}

			validMap := map[string]interface{}{}
			if len(answer.Body) > 0 {
				err = json.Unmarshal(answer.Body, &validMap)
				if err != nil {
					t.Fatal(err)
				}
			}

			require.Equal(t, validMap, respMap)
		})
	}
}
