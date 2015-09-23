package api

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xiam/hyperfox/tools/capture"
)

type apiAdapter struct{}

func (a apiAdapter) Post([]byte) ([]byte, error) {
	return nil, errors.New("Some error")
}

func (a apiAdapter) ParseResponse(r *capture.Response) ([]byte, error) {
	return []byte{}, nil
}

func TestSendCapturedObjectReturnError(t *testing.T) {
	now := time.Date(2015, 9, 22, 12, 43, 51, 5, time.UTC)
	started := time.Date(2015, 9, 22, 12, 43, 51, 1, time.UTC)

	cr := capture.Response{
		Origin:        "Origin",
		Method:        "Method",
		Status:        200,
		ContentType:   "ContentType",
		ContentLength: 53421,
		Host:          "Host",
		URL:           "URL",
		Scheme:        "Scheme",
		Path:          "Path",
		Header:        capture.Header{},
		Body:          []byte("Body"),
		RequestHeader: capture.Header{},
		RequestBody:   []byte("RequestBody"),
		DateStart:     started,
		DateEnd:       time.Date(2015, 9, 22, 12, 43, 51, 2, time.UTC),
		TimeTaken:     now.UnixNano() - started.UnixNano(),
	}

	_, err := SendCapturedObject(apiAdapter{}, &cr)
	assert.NotNil(t, err)
}
