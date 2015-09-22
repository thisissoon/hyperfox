package deadpool

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xiam/hyperfox/tools/capture"
)

func PopulateDeadPoolResponseFromCaptureResponse(t *testing.T) {
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
		DateStart:     time.Now(),
		DateEnd:       time.Now(),
		TimeTaken:     423,
	}

	r := FromResponse(&cr)
	assert.Equal(t, cr.Origin, r.Origin)
	assert.Equal(t, cr.Method, r.Method)
	assert.Equal(t, cr.Status, r.Status)
	assert.Equal(t, cr.ContentType, r.ContentType)
	assert.Equal(t, cr.ContentLength, r.ContentLength)
	assert.Equal(t, cr.Host, r.Host)
	assert.Equal(t, cr.URL, r.URL)
	assert.Equal(t, cr.Scheme, r.Scheme)
	assert.Equal(t, cr.Path, r.Path)
	assert.Equal(t, cr.Header, r.Header)
	assert.Equal(t, cr.Body, r.Body)
	assert.Equal(t, cr.RequestHeader, r.RequestHeader)
	assert.Equal(t, cr.RequestBody, r.RequestBody)
	assert.Equal(t, cr.DateStart, r.DateStart)
	assert.Equal(t, cr.DateEnd, r.DateEnd)
	assert.Equal(t, cr.TimeTaken, r.TimeTaken)
}
