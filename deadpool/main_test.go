package deadpool

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xiam/hyperfox/tools/capture"
)

func TestPopulateDeadPoolResponseFromCaptureResponse(t *testing.T) {
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

	r := DeadpoolApiAdapter{}.TransformResponse(&cr)
	assert.Equal(t, cr.Origin, r.Origin)
	assert.Equal(t, cr.Method, r.Method)
	assert.Equal(t, cr.Status, r.Status)
	assert.Equal(t, cr.ContentType, r.ContentType)
	assert.Equal(t, cr.ContentLength, r.ContentLength)
	assert.Equal(t, cr.Host, r.Host)
	assert.Equal(t, cr.URL, r.URL)
	assert.Equal(t, cr.Scheme, r.Scheme)
	assert.Equal(t, cr.Path, r.Path)
	// assert.Equal(t, cr.Header, r.Header)
	assert.Equal(t, cr.Body, r.Body)
	// assert.Equal(t, cr.RequestHeader, r.RequestHeader)
	assert.Equal(t, cr.RequestBody, r.RequestBody)
	assert.Equal(t, cr.DateStart, r.DateStart)
	assert.Equal(t, cr.DateEnd, r.DateEnd)
	assert.Equal(t, now.UnixNano(), r.TimeTaken.UnixNano())
}
