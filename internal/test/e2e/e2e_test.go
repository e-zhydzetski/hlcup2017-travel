package e2e

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/test/internal/helpers"

	"github.com/stretchr/testify/require"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/test/internal/rawhttp"
)

func TestE2E(t *testing.T) {
	data, err := ioutil.ReadFile("../../../test/data/TRAIN/ammo/phase_1_get.ammo")
	if err != nil {
		t.Fatal(err)
	}
	source := bytes.NewBuffer(data)
	ammo, err := helpers.ReadAmmo(source)
	if err != nil {
		t.Fatal(err)
	}

	data, err = ioutil.ReadFile("../../../test/data/TRAIN/answers/phase_1_get.answ")
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

		t.Run(strconv.Itoa(i)+": "+answer.Name, func(t *testing.T) {
			code, resp, err := rawhttp.SendRequest("127.0.0.1:80", bullet.Request)
			if err != nil {
				t.Fatal(err)
			}
			if answer.Code != code {
				t.Error("Unexpected code. Expected", answer.Code, ", got", code)
			}
			respMap := map[string]interface{}{}
			_ = json.Unmarshal(resp, &respMap)

			validMap := map[string]interface{}{}
			_ = json.Unmarshal(answer.Body, &validMap)

			require.Equal(t, validMap, respMap)
		})
	}
}
