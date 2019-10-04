package config

import (
	"bytes"
	"io"
)

type TokenizedReadCloser struct {
	source io.ReadCloser
	buff   []byte
	cur    []byte
}

func NewTokenizedReadCloser(source io.ReadCloser) *TokenizedReadCloser {
	res := &TokenizedReadCloser{
		source: source,
		buff:   make([]byte, 1024),
	}
	res.cur = res.buff[:0]
	return res
}

func (t *TokenizedReadCloser) ReadStringUntil(b byte) (string, error) {
	s := ""
	for {
		i := bytes.IndexByte(t.cur, b)
		if i >= 0 {
			s += string(t.cur[:i])
			t.cur = t.cur[i+1:]
			break
		}
		s += string(t.cur)
		n, err := t.source.Read(t.buff)
		if err != nil {
			return s, err
		}
		t.cur = t.buff[:n]
	}
	return s, nil
}

func (t *TokenizedReadCloser) Read(res []byte) (int, error) {
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

func (t *TokenizedReadCloser) Close() error {
	t.buff = nil
	t.cur = nil
	return t.source.Close()
}
