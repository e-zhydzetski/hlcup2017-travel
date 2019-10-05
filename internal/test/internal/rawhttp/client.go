package rawhttp

import (
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/e-zhydzetski/hlcup2017-travel/internal/test/internal/reader"
)

// SendRequest returns code, body, error
func SendRequest(addr string, req []byte) (int, []byte, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return 0, nil, err
	}
	defer conn.Close()
	_ = conn.SetDeadline(time.Now().Add(2 * time.Second))
	if _, err = conn.Write(req); err != nil {
		return 0, nil, err
	}

	r := reader.NewExtendedReader(conn)
	cl := 0
	var code int
	for ln := 0; ; ln++ {
		var s string
		if s, err = r.ReadStringUntil('\n'); err != nil {
			return 0, nil, err
		}
		if s == "\r" { // empty string before body
			break
		}
		if ln == 0 {
			s = strings.Split(s, " ")[1] // HTTP {code} {description} to {code}
			code, _ = strconv.Atoi(s)
			continue
		}
		if strings.HasPrefix(s, "Content-Length:") {
			s = strings.Split(s, ":")[1] // {headerName}: {headerValue} to {headerValue}
			s = strings.TrimSpace(s)
			cl, _ = strconv.Atoi(s)
		}
	}
	body := make([]byte, cl)
	if _, err = r.Read(body); err != nil {
		return code, nil, err
	}
	return code, body, nil
}
