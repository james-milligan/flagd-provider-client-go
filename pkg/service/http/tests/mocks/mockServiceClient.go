package mocks

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"testing"

	of "github.com/open-feature/golang-sdk/pkg/openfeature"
)

type ServiceClient struct {
	ServiceClientMockSetup

	Testing *testing.T
}

type ServiceClientMockSetup struct {
	InUrl string
	InCtx of.EvaluationContext

	OutBody interface{}
	OutSC   int
	OutErr  error
}

func (s *ServiceClient) FetchFlag(url string, ctx of.EvaluationContext, p interface{}) (io.ReadCloser, int, error) {
	outM, err := json.Marshal(s.OutBody)
	if err != nil {
		s.Testing.Error(err)
		return nil, s.OutSC, s.OutErr
	}
	out := io.NopCloser(bytes.NewReader(outM))
	if url != s.InUrl {
		s.Testing.Errorf("unexpected value for url received, expected %v got %v", s.InUrl, url)
	}
	if !reflect.DeepEqual(ctx, s.InCtx) {
		s.Testing.Errorf("unexpected value for context received, expected %v got %v", s.InCtx, ctx)
	}
	outBodyM, err := json.Marshal(s.OutBody)
	if err != nil {
		s.Testing.Error(err)
	}
	if err := json.NewDecoder(bytes.NewReader(outBodyM)).Decode(p); err != nil {
		s.Testing.Error(err)
	}
	return out, s.OutSC, s.OutErr
}
