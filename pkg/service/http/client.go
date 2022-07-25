package http_service

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	models "github.com/open-feature/flagd/pkg/model"
	of "github.com/open-feature/golang-sdk/pkg/openfeature"
	log "github.com/sirupsen/logrus"
)

type IHTTPClient interface {
	FetchFlag(url string, ctx of.EvaluationContext, p interface{}) (io.ReadCloser, int, error)
}

type HTTPClient struct {
	client *http.Client
}

func (c *HTTPClient) GetInstance() http.Client {
	if c.client == nil {
		c.client = &http.Client{}
	}
	return *c.client
}

func (c *HTTPClient) FetchFlag(url string, ctx of.EvaluationContext, p interface{}) (io.ReadCloser, int, error) {
	body, err := json.Marshal(ctx)
	if err != nil {
		log.Error(err)
		return nil, 0, errors.New(models.ParseErrorCode)
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Error(err)
		return nil, 0, errors.New(models.GeneralErrorCode)
	}
	if res.StatusCode != 200 {
		return res.Body, res.StatusCode, nil
	}
	if err := json.NewDecoder(res.Body).Decode(p); err != nil {
		log.Error(err)
		return nil, 0, errors.New(models.GeneralErrorCode)
	}
	return res.Body, 200, nil
}
