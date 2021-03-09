package safecustody_request

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"time"
)

type Method string

type Header map[string]string

type Request struct {
	HttpC *http.Client

	t time.Duration

	header Header
}

func (r *Request) SetTimeOut(t time.Duration) *Request {
	r.t = t
	return r
}

func (r *Request) SetHeader(header Header) *Request {

	for k, v := range header {
		r.header[k] = v
	}
	return r
}

func (r *Request) request(method Method, url string, body []byte) (resp *http.Response, err error) {

	var req *http.Request

	if req, err = http.NewRequest(string(method), url, bytes.NewBuffer(body)); err != nil {
		return
	}

	//设置头
	for k, v := range r.header {
		req.Header.Set(k, v)
	}
	return r.HttpC.Do(req)
}

func Http() *Request {
	header := make(Header)
	header["Content-Type"] = "application/x-www-form-urlencoded"

	return &Request{
		HttpC: &http.Client{
			Timeout:   timeOutDefault,
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: InsecureSkipVerify}},
		},
		header: header,
	}
}
