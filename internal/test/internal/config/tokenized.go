package config

import (
	"bytes"
	"io"
	"strings"
)

type TokenizedReader struct {
	source io.Reader
	buff   []byte
	cur    []byte
}

func NewTokenizedReader(source io.Reader) *TokenizedReader {
	res := &TokenizedReader{
		source: source,
		buff:   make([]byte, 1024),
	}
	res.cur = res.buff[:0]
	return res
}

func (t *TokenizedReader) ReadStringUntil(b byte) (string, error) {
	s := strings.Builder{}
	for {
		i := bytes.IndexByte(t.cur, b)
		if i >= 0 {
			s.Write(t.cur[:i])
			t.cur = t.cur[i+1:]
			break
		}
		s.Write(t.cur)
		n, err := t.source.Read(t.buff)
		if err != nil {
			return s.String(), err
		}
		t.cur = t.buff[:n]
	}
	return s.String(), nil
}

func (t *TokenizedReader) Read(res []byte) (int, error) {
	size := len(res)
	var curSize int
	if size == 0 {
		return 0, nil
	}
	if len(t.cur) >= size {
		copy(res, t.cur[:size])
		t.cur = t.cur[size:]
		curSize = size
	} else {
		curSize = copy(res, t.cur)
		n, err := t.source.Read(res[curSize:])
		curSize += n
		if err != nil {
			return curSize, err
		}
		n, err = t.source.Read(t.buff)
		if err != nil {
			return curSize, err
		}
		t.cur = t.buff[:n]
	}
	return curSize, nil
}
