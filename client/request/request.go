package request

import (
	"bytes"
	"math/rand"
	"net/http"
	"strconv"
)

type RequestLine struct {
	method  string
	host    string
	pattern string
}

func NewRequestLine(method, host, pattern string) RequestLine {
	return RequestLine{
		method:  method,
		host:    host,
		pattern: pattern,
	}
}

func (r *RequestLine) NewRequest() (*http.Request, error) {
	req, err := http.NewRequest(r.method, r.host, nil)
	if err != nil {
		return nil, err
	}
	req.URL.Path = generateRandomPath(r.pattern)
	return req, nil
}

func generateRandomPath(pattern string) string {
	buf := bytes.NewBuffer(nil)
	inParam := false

	for _, c := range pattern {
		switch c {
		case '/':
			inParam = false
			buf.WriteRune('/')
		case ':':
			if inParam {
				continue
			}
			inParam = true
			buf.WriteString(strconv.Itoa(rand.Intn(100)))

		default:
			if inParam {
				continue
			}
			buf.WriteRune(c)
		}
	}

	return buf.String()
}
