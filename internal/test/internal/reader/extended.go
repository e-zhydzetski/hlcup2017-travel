package reader

import (
	"bytes"
	"io"
	"strings"
)

type ExtendedReader struct {
	source io.Reader
	buff   []byte
	cur    []byte
}

func NewExtendedReader(source io.Reader) *ExtendedReader {
	res := &ExtendedReader{
		source: source,
		buff:   make([]byte, 1024),
	}
	res.cur = res.buff[:0]
	return res
}

func (e *ExtendedReader) ReadStringUntil(b byte) (string, error) {
	s := strings.Builder{}
	for {
		i := bytes.IndexByte(e.cur, b)
		if i >= 0 {
			s.Write(e.cur[:i])
			e.cur = e.cur[i+1:]
			break
		}
		s.Write(e.cur)
		n, err := e.source.Read(e.buff)
		if err != nil {
			return s.String(), err
		}
		e.cur = e.buff[:n]
	}
	return s.String(), nil
}

func (e *ExtendedReader) Read(res []byte) (int, error) {
	size := len(res)
	var curSize int
	if size == 0 {
		return 0, nil
	}
	if len(e.cur) >= size {
		copy(res, e.cur[:size])
		e.cur = e.cur[size:]
		curSize = size
	} else {
		curSize = copy(res, e.cur)
		n, err := e.source.Read(res[curSize:])
		curSize += n
		if err != nil {
			return curSize, err
		}
		n, err = e.source.Read(e.buff)
		if err != nil {
			return curSize, err
		}
		e.cur = e.buff[:n]
	}
	return curSize, nil
}
