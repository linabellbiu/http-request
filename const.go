package safecustody_request

import "time"

type HeaderType string

const (
	PostMethod         Method = "POST"
	GetMethod          Method = "GET"
	timeOutDefault            = 10 * time.Second
	InsecureSkipVerify        = true
)
