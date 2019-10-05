package helpers

import (
	"io"
	"strconv"
	"strings"

	"github.com/e-zhydzetski/hlcup2017-travel/test/internal/reader"
)

type AmmoEntry struct {
	Kind    string
	Request []byte
}

type AnswerEntry struct {
	Name string
	Code int
	Body []byte
}

func ReadAmmo(source io.Reader) ([]*AmmoEntry, error) {
	r := reader.NewExtendedReader(source)

	var res []*AmmoEntry

	for {
		s, err := r.ReadStringUntil(' ')
		if err == io.EOF {
			break
		}
		size, _ := strconv.Atoi(s)
		kind, _ := r.ReadStringUntil('\n')
		req := make([]byte, size)
		_, _ = r.Read(req)

		a := &AmmoEntry{
			Kind:    kind,
			Request: req,
		}
		res = append(res, a)
	}

	return res, nil
}

func ReadAnswers(source io.Reader) ([]*AnswerEntry, error) {
	r := reader.NewExtendedReader(source)

	var res []*AnswerEntry

	for {
		s, err := r.ReadStringUntil('\n')
		if err == io.EOF {
			break
		}
		p := strings.Split(s, "\t") // {method}\t{url}\t{code}\t{?body}
		answer := &AnswerEntry{}
		answer.Name = p[0] + " " + p[1]
		answer.Code, _ = strconv.Atoi(strings.TrimSpace(p[2]))
		if len(p) > 3 {
			answer.Body = []byte(strings.TrimSpace(p[3]))
		}
		res = append(res, answer)
	}

	return res, nil
}
