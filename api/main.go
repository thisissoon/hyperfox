package api

import (
	"github.com/xiam/hyperfox/tools/capture"
)

type ApiAdapter interface {
	Post(payload []byte) ([]byte, error)
	ParseResponse(r *capture.Response) ([]byte, error)
}

func SendCapturedObject(a ApiAdapter, r *capture.Response) ([]byte, error) {
	payload, err := a.ParseResponse(r)
	if err != nil {
		return []byte{}, err
	}
	return a.Post(payload)
}
