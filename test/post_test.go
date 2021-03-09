package test

import (
	"fmt"
	req "safecustody_request"
	"testing"
)

func TestPost(t *testing.T) {
	header := make(req.Header)
	res, err := req.Http().SetHeader(header).Post("", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res.Body))
}
