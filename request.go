package safecustody_request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Body []byte
	Resp *http.Response
}

func (r *Request) Post(url string, param interface{}) (resp *Response, err error) {
	body, err := json.Marshal(param)
	if err != nil {
		return
	}
	return r.PostByte(url, body)
}

func (r *Request) PostByte(url string, param []byte) (resp *Response, err error) {
	res, err := r.request(PostMethod, url, param)
	if err != nil {
		return
	}
	return resolveResp(res)
}

func (r *Request) Get(url string) (resp *Response, err error) {
	res, err := r.request(GetMethod, url, nil)
	if err != nil {
		return
	}
	return resolveResp(res)
}

func resolveResp(res *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(res.Body)
	return &Response{
		Body: body,
		Resp: res,
	}, err
}
