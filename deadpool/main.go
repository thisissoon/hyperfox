package deadpool

import (
	"time"

	"github.com/xiam/hyperfox/tools/capture"
)

type Header struct {
	capture.Header
}

type Response struct {
	Origin        string    `json:"origin",json"`
	Method        string    `json:"method",json"`
	Status        int       `json:"status",json"`
	ContentType   string    `json:"content_type",json"`
	ContentLength uint      `json:"content_length",json"`
	Host          string    `json:"host",json"`
	URL           string    `json:"url",json"`
	Scheme        string    `json:"scheme",json"`
	Path          string    `json:"path",path"`
	Header        Header    `json:"header",json"`
	Body          []byte    `json:"body",json"`
	RequestHeader Header    `json:"request_header",json"`
	RequestBody   []byte    `json:"request_body",json"`
	DateStart     time.Time `json:"date_start",json"`
	DateEnd       time.Time `json:"date_end",json"`
	TimeTaken     int64     `json:"time_taken",json"`
}

func FromResponse(response *capture.Response) Response {
	return Response{
		Origin:        response.Origin,
		Method:        response.Method,
		Status:        response.Status,
		ContentType:   response.ContentType,
		ContentLength: response.ContentLength,
		Host:          response.Host,
		URL:           response.URL,
		Scheme:        response.Scheme,
		Path:          response.Path,
		Header:        Header{response.Header},
		Body:          response.Body,
		RequestHeader: Header{response.RequestHeader},
		RequestBody:   response.RequestBody,
		DateStart:     response.DateStart,
		DateEnd:       response.DateEnd,
		TimeTaken:     response.TimeTaken,
	}
}

// func SendResponse(response *capture.Response) error {
// 	payload = string(&response)
// }
