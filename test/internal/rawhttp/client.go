package rawhttp

import (
	"errors"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/e-zhydzetski/hlcup2017-travel/test/internal/reader"
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
	headers := map[string]struct{}{}
	cl := 0
	chunked := false
	var code int
	for ln := 0; ; ln++ {
		var s string
		if s, err = r.ReadStringUntil('\n'); err != nil {
			return 0, nil, err
		}
		s = strings.TrimSpace(s)
		if s == "" { // empty string before body
			break
		}
		if ln == 0 {
			s = strings.Split(s, " ")[1] // HTTP {code} {description} to {code}
			code, _ = strconv.Atoi(s)
			continue
		}
		headers[strings.Split(s, ":")[0]] = struct{}{}
		if s == "Transfer-Encoding: chunked" {
			chunked = true
			continue
		}
		if strings.HasPrefix(s, "Content-Length:") {
			s = strings.Split(s, ":")[1] // {headerName}: {headerValue} to {headerValue}
			s = strings.TrimSpace(s)
			cl, _ = strconv.Atoi(s)
		}
	}

	var body []byte
	if chunked {
		for {
			var s string
			if s, err = r.ReadStringUntil('\n'); err != nil {
				return code, nil, err
			}
			s = strings.TrimSpace(s)
			l, err := strconv.ParseInt(s, 16, 32)
			if err != nil {
				return code, nil, err
			}
			if l == 0 {
				break
			}
			chunk := make([]byte, l)
			if _, err = r.Read(chunk); err != nil {
				return code, nil, err
			}
			body = append(body, chunk...)
			if _, err = r.ReadStringUntil('\n'); err != nil {
				return code, nil, err
			}
		}
	} else {
		body = make([]byte, cl)
		if _, err = r.Read(body); err != nil {
			return code, nil, err
		}
	}

	// check after read to prevent brush in reader
	if len(body) > 0 {
		if _, exists := headers["Content-Type"]; !exists {
			return code, body, errors.New("required header not exists: Content-Type")
		}
	}
	for _, rh := range []string{"Content-Length", "Connection"} {
		if _, exists := headers[rh]; !exists {
			return code, body, errors.New("required header not exists: " + rh)
		}
	}

	return code, body, nil
}
