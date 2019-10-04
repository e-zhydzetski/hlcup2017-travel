package test

import (
	"io"
	"io/ioutil"
	"net"
	"strconv"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/test/internal/config"
)

type AmmoEntry struct {
	Kind    string
	Request []byte
}

type AnswerEntry struct {
	Method string
	URL    string
	Code   int
	Body   []byte
}

type Case struct {
	Ammo   AmmoEntry
	Answer AnswerEntry
}

func ReadAmmo(source io.ReadCloser) ([]AmmoEntry, error) {
	tf := config.NewTokenizedReadCloser(source)
	defer tf.Close()

	var res []AmmoEntry

	for {
		s, err := tf.ReadStringUntil(' ')
		if err == io.EOF {
			break
		}
		size, _ := strconv.Atoi(s)
		kind, _ := tf.ReadStringUntil('\n')
		req := make([]byte, size)
		_, _ = tf.Read(req)

		a := AmmoEntry{
			Kind:    kind,
			Request: req,
		}
		res = append(res, a)
	}

	return res, nil
}

func SendRawRequest(addr string, request []byte) ([]byte, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	_, err = conn.Write(request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(conn) // TODO fix lock on keep-alive, analyze content-length, read and close connection
}
