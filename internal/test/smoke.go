package test

import (
	"io"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/test/internal/config"
)

type AmmoEntry struct {
	Kind    string
	Request []byte
}

type AnswerEntry struct {
	Code int
	Body []byte
}

func (a *AnswerEntry) String() string {
	return strconv.Itoa(a.Code) + ": " + string(a.Body)
}

type Case struct {
	Ammo   AmmoEntry
	Answer AnswerEntry
}

func ReadAmmo(source io.Reader) ([]*AmmoEntry, error) {
	tf := config.NewTokenizedReader(source)

	var res []*AmmoEntry

	for {
		s, err := tf.ReadStringUntil(' ')
		if err == io.EOF {
			break
		}
		size, _ := strconv.Atoi(s)
		kind, _ := tf.ReadStringUntil('\n')
		req := make([]byte, size)
		_, _ = tf.Read(req)

		a := &AmmoEntry{
			Kind:    kind,
			Request: req,
		}
		res = append(res, a)
	}

	return res, nil
}

func SendRawRequest(addr string, ammo *AmmoEntry) (*AnswerEntry, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	_ = conn.SetDeadline(time.Now().Add(2 * time.Second))
	defer conn.Close()
	_, err = conn.Write(ammo.Request)
	if err != nil {
		return nil, err
	}

	var answer AnswerEntry

	tr := config.NewTokenizedReader(conn)
	cl := 0
	for {
		s, err := tr.ReadStringUntil('\n')
		if err != nil {
			return nil, err
		}
		if s == "\r" {
			break
		}
		if strings.HasPrefix(s, "HTTP/1.1 ") {
			s = strings.TrimPrefix(s, "HTTP/1.1 ")
			s = s[:3]
			if answer.Code, err = strconv.Atoi(s); err != nil {
				answer.Code = 0
			}
		}
		if strings.HasPrefix(s, "Content-Length:") {
			s = strings.TrimPrefix(s, "Content-Length:")
			s = strings.TrimSpace(s)
			if cl, err = strconv.Atoi(s); err != nil {
				cl = 0
			}
		}
	}
	respBody := make([]byte, cl)
	_, err = tr.Read(respBody)
	if err != nil {
		return nil, err
	}
	answer.Body = respBody
	return &answer, nil
}
