package deadpool

import (
	"bytes"
	"encoding/json"
	"errors"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
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
	TimeTaken     time.Time `json:"time_taken",json"`
}

type DeadpoolApiAdapter struct {
}

func (a DeadpoolApiAdapter) TransformResponse(r *capture.Response) Response {
	taken := r.DateStart.UnixNano() + r.TimeTaken
	secods := int64(float64(taken) * 1e-9)

	return Response{
		Origin:        r.Origin,
		Method:        r.Method,
		Status:        r.Status,
		ContentType:   r.ContentType,
		ContentLength: r.ContentLength,
		Host:          r.Host,
		URL:           r.URL,
		Scheme:        r.Scheme,
		Path:          r.Path,
		Header:        Header{r.Header},
		Body:          r.Body,
		RequestHeader: Header{r.RequestHeader},
		RequestBody:   r.RequestBody,
		DateStart:     r.DateStart,
		DateEnd:       r.DateEnd,
		TimeTaken:     time.Unix(secods, taken-secods*1e9),
	}
}

func (a DeadpoolApiAdapter) ParseResponse(r *capture.Response) ([]byte, error) {
	res := a.TransformResponse(r)
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()
	b, err := json.Marshal(res)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}

func (a DeadpoolApiAdapter) Post(payload []byte) ([]byte, error) {
	deadpoolUrl := os.Getenv("DEADPOOL_URL")
	if deadpoolUrl == "" {
		return []byte{}, errors.New("Env variable DEADPOOL_URL cannot be empty")
	}
	response, _ := http.Post(deadpoolUrl+"/v1/report", "application/json", bytes.NewReader(payload))
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
