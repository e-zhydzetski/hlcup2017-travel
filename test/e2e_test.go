package test

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/e-zhydzetski/hlcup2017-travel/test/internal/helpers"
	"github.com/e-zhydzetski/hlcup2017-travel/test/internal/rawhttp"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/app"
)

func TestE2E(t *testing.T) {
	const level = "TRAIN"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	service := app.Service{
		ListenAddr:  ":8080",
		OptionsFile: "data/" + level + "/data/options.txt",
		DumpSource:  "data/" + level + "/data/data.zip",
	}
	go func() {
		_ = service.Start(ctx)
	}()

	for {
		if conn, err := net.Dial("tcp", "127.0.0.1:8080"); err == nil {
			conn.Close()
			break
		}
		time.Sleep(time.Second)
		t.Log("Wait server start listening...")
	}
	t.Log("Server is ready")

	executeTestPhase(t, level, "phase_1_get")
	executeTestPhase(t, level, "phase_2_post")
	executeTestPhase(t, level, "phase_3_get")
}

func executeTestPhase(t *testing.T, level string, phaseName string) {
	data, err := ioutil.ReadFile("data/" + level + "/ammo/" + phaseName + ".ammo")
	if err != nil {
		t.Fatal(err)
	}
	source := bytes.NewBuffer(data)
	ammo, err := helpers.ReadAmmo(source)
	if err != nil {
		t.Fatal(err)
	}

	data, err = ioutil.ReadFile("data/" + level + "/answers/" + phaseName + ".answ")
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

	for i := range ammo {
		bullet := ammo[i]
		answer := answers[i]

		t.Run(phaseName+"_"+strconv.Itoa(i)+"_"+answer.Name, func(t *testing.T) {
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
