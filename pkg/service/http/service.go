package http_service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/james-milligan/flagd-provider-client-go/pkg/service"
	models "github.com/open-feature/flagd/pkg/model"
	of "github.com/open-feature/golang-sdk/pkg/openfeature"
	log "github.com/sirupsen/logrus"
	schemaV1 "go.buf.build/grpc/go/open-feature/flagd/schema/v1"
)

type HTTPServiceConfiguration struct {
	Port    int32
	Address string
}

type HTTPService struct {
	HTTPServiceConfiguration *HTTPServiceConfiguration
}

func (s *HTTPService) ResolveBoolean(flagKey string, context of.EvaluationContext, options ...service.IServiceOption) (*schemaV1.ResolveBooleanResponse, error) {
	url := fmt.Sprintf("http://%s:%d/flags/%s/resolve/boolean", s.HTTPServiceConfiguration.Address, s.HTTPServiceConfiguration.Port, flagKey)
	body, err := json.Marshal(context)
	if err != nil {
		log.Error(err)
		return &schemaV1.ResolveBooleanResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New(models.ParseErrorCode)
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Error(err)
		return &schemaV1.ResolveBooleanResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New(models.GeneralErrorCode)
	}
	if res.StatusCode != 200 {
		return &schemaV1.ResolveBooleanResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, HandleNon200(res)
	}
	resMess := schemaV1.ResolveBooleanResponse{}
	if err := json.NewDecoder(res.Body).Decode(&resMess); err != nil {
		log.Error(err)
		return &schemaV1.ResolveBooleanResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New(models.ParseErrorCode)
	}
	return &resMess, nil
}

func (s *HTTPService) ResolveString(flagKey string, context of.EvaluationContext, options ...service.IServiceOption) (*schemaV1.ResolveStringResponse, error) {
	url := fmt.Sprintf("http://%s:%d/flags/%s/resolve/string", s.HTTPServiceConfiguration.Address, s.HTTPServiceConfiguration.Port, flagKey)
	body, err := json.Marshal(context)
	if err != nil {
		log.Error(err)
		return &schemaV1.ResolveStringResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New(models.ParseErrorCode)
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Error(err)
		return &schemaV1.ResolveStringResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New(models.GeneralErrorCode)
	}
	if res.StatusCode != 200 {
		return &schemaV1.ResolveStringResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, HandleNon200(res)
	}
	resMess := schemaV1.ResolveStringResponse{}
	if err := json.NewDecoder(res.Body).Decode(&resMess); err != nil {
		log.Error(err)
		return &schemaV1.ResolveStringResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New(models.ParseErrorCode)
	}
	return &resMess, nil
}

func (s *HTTPService) ResolveNumber(flagKey string, context of.EvaluationContext, options ...service.IServiceOption) (*schemaV1.ResolveNumberResponse, error) {
	url := fmt.Sprintf("http://%s:%d/flags/%s/resolve/number", s.HTTPServiceConfiguration.Address, s.HTTPServiceConfiguration.Port, flagKey)
	body, err := json.Marshal(context)
	if err != nil {
		log.Error(err)
		return &schemaV1.ResolveNumberResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New(models.ParseErrorCode)
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Error(err)
		return &schemaV1.ResolveNumberResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New(models.GeneralErrorCode)
	}
	if res.StatusCode != 200 {
		return &schemaV1.ResolveNumberResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, HandleNon200(res)
	}
	resMess := schemaV1.ResolveNumberResponse{}
	if err := json.NewDecoder(res.Body).Decode(&resMess); err != nil {
		log.Error(err)
		return &schemaV1.ResolveNumberResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New(models.ParseErrorCode)
	}
	return &resMess, nil
}

func (s *HTTPService) ResolveObject(flagKey string, context of.EvaluationContext, options ...service.IServiceOption) (*schemaV1.ResolveObjectResponse, error) {
	url := fmt.Sprintf("http://%s:%d/flags/%s/resolve/object", s.HTTPServiceConfiguration.Address, s.HTTPServiceConfiguration.Port, flagKey)
	body, err := json.Marshal(context)
	if err != nil {
		log.Error(err)
		return &schemaV1.ResolveObjectResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New(models.ParseErrorCode)
	}
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Error(err)
		return &schemaV1.ResolveObjectResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New(models.GeneralErrorCode)
	}
	if res.StatusCode != 200 {
		return &schemaV1.ResolveObjectResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, HandleNon200(res)
	}
	resMess := schemaV1.ResolveObjectResponse{}
	if err := json.NewDecoder(res.Body).Decode(&resMess); err != nil {
		log.Error(err)
		return &schemaV1.ResolveObjectResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New(models.ParseErrorCode)
	}
	return &resMess, nil
}

func HandleNon200(res *http.Response) error {
	errMess := schemaV1.ErrorResponse{}
	if err := json.NewDecoder(res.Body).Decode(&errMess); err != nil {
		log.Error(err)
		return errors.New(models.ParseErrorCode)
	}
	if errMess.ErrorCode != "" {
		return errors.New(errMess.ErrorCode)
	}
	log.Error("unexpected error response recieved from flagd server")
	return errors.New(models.GeneralErrorCode)
}
