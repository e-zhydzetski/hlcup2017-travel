package test

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestLoadAmmo(t *testing.T) {
	data, err := ioutil.ReadFile("ammo.txt")
	if err != nil {
		t.Fatal(err)
	}
	source := bytes.NewBuffer(data)
	ammo, err := ReadAmmo(source)
	if err != nil {
		t.Fatal(err)
	}
	if len(ammo) != 2 {
		t.Error("Invalid count of loaded ammo. Expected: 2. Got:", len(ammo))
	}
	t.Log(ammo)
}

func TestSendRawRequest(t *testing.T) {
	data, err := ioutil.ReadFile("ammo.txt")
	//data, err := ioutil.ReadFile("../../test/data/TRAIN/ammo/phase_1_get.ammo")
	if err != nil {
		t.Fatal(err)
	}
	source := bytes.NewBuffer(data)
	ammo, err := ReadAmmo(source)
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range ammo {
		resp, err := SendRawRequest("127.0.0.1:80", entry)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp)
	}
}
