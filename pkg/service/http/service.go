package http_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/james-milligan/flagd-provider-client-go/pkg/service"
	models "github.com/open-feature/flagd/pkg/model"
	of "github.com/open-feature/golang-sdk/pkg/openfeature"
	log "github.com/sirupsen/logrus"
	schemaV1 "go.buf.build/grpc/go/open-feature/flagd/schema/v1"
)

type HTTPServiceConfiguration struct {
	Port int32
	Host string
}

type HTTPService struct {
	HTTPServiceConfiguration *HTTPServiceConfiguration
	Client                   IHTTPClient
}

func (s *HTTPService) ResolveBoolean(flagKey string, context of.EvaluationContext, options ...service.IServiceOption) (*schemaV1.ResolveBooleanResponse, error) {
	url := fmt.Sprintf("http://%s:%d/flags/%s/resolve/boolean", s.HTTPServiceConfiguration.Host, s.HTTPServiceConfiguration.Port, flagKey)
	resMess := schemaV1.ResolveBooleanResponse{}
	body, stCode, err := s.Client.FetchFlag(url, context, &resMess)
	fmt.Println(resMess)
	if err != nil {
		return &schemaV1.ResolveBooleanResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, err
	}
	if stCode != 200 {
		return &schemaV1.ResolveBooleanResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, HandleNon200(body)
	}
	return &resMess, nil
}

func (s *HTTPService) ResolveString(flagKey string, context of.EvaluationContext, options ...service.IServiceOption) (*schemaV1.ResolveStringResponse, error) {
	url := fmt.Sprintf("http://%s:%d/flags/%s/resolve/string", s.HTTPServiceConfiguration.Host, s.HTTPServiceConfiguration.Port, flagKey)
	resMess := schemaV1.ResolveStringResponse{}
	body, stCode, err := s.Client.FetchFlag(url, context, &resMess)
	if err != nil {
		return &schemaV1.ResolveStringResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, err
	}
	if stCode != 200 {
		return &schemaV1.ResolveStringResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, HandleNon200(body)
	}
	return &resMess, nil
}

func (s *HTTPService) ResolveNumber(flagKey string, context of.EvaluationContext, options ...service.IServiceOption) (*schemaV1.ResolveNumberResponse, error) {
	url := fmt.Sprintf("http://%s:%d/flags/%s/resolve/number", s.HTTPServiceConfiguration.Host, s.HTTPServiceConfiguration.Port, flagKey)
	resMess := schemaV1.ResolveNumberResponse{}
	body, stCode, err := s.Client.FetchFlag(url, context, &resMess)
	if err != nil {
		return &schemaV1.ResolveNumberResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, err
	}
	if stCode != 200 {
		return &schemaV1.ResolveNumberResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, HandleNon200(body)
	}
	return &resMess, nil
}

func (s *HTTPService) ResolveObject(flagKey string, context of.EvaluationContext, options ...service.IServiceOption) (*schemaV1.ResolveObjectResponse, error) {
	url := fmt.Sprintf("http://%s:%d/flags/%s/resolve/object", s.HTTPServiceConfiguration.Host, s.HTTPServiceConfiguration.Port, flagKey)
	resMess := schemaV1.ResolveObjectResponse{}
	body, stCode, err := s.Client.FetchFlag(url, context, &resMess)
	if err != nil {
		return &schemaV1.ResolveObjectResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, err
	}
	if stCode != 200 {
		return &schemaV1.ResolveObjectResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, HandleNon200(body)
	}
	return &resMess, nil
}

func HandleNon200(body io.ReadCloser) error {
	errMess := schemaV1.ErrorResponse{}
	if err := json.NewDecoder(body).Decode(&errMess); err != nil {
		log.Error(err)
		return errors.New(models.ParseErrorCode)
	}
	if errMess.ErrorCode != "" {
		return errors.New(errMess.ErrorCode)
	}
	log.Error("unexpected error response recieved from flagd server")
	return errors.New(models.GeneralErrorCode)
}
