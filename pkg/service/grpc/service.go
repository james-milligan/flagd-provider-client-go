package grpc

import (
	ctx "context"
	"errors"

	"github.com/james-milligan/flagd-provider-client-go/pkg/service"
	gen "github.com/james-milligan/flagd-provider-client-go/schemas/protobuf/gen/v1"
	models "github.com/open-feature/flagd/pkg/model"
	of "github.com/open-feature/golang-sdk/pkg/openfeature"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

func (c *GRPCService) ResolveBoolean(flagKey string, context of.EvaluationContext, options ...service.IServiceOption) (*gen.ResolveBooleanResponse, error) {
	client := c.GetInstance()
	if client == nil {
		return &gen.ResolveBooleanResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New("CONNECTION_ERROR") // todo: create more error codes, should these errors sit in the schema?
	}
	contextF, err := FormatAsStructpb(context)
	if err != nil {
		log.Error(err)
		return &gen.ResolveBooleanResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New(models.ParseErrorCode)
	}
	res, err := client.ResolveBoolean(ctx.TODO(), &gen.ResolveBooleanRequest{
		FlagKey: flagKey,
		Context: contextF,
	})
	if err != nil {
		res, ok := ParseError(err)
		if !ok {
			log.Error(err)
			return &gen.ResolveBooleanResponse{
				Reason:  models.ErrorReason,
				Variant: "default_value",
			}, errors.New(models.GeneralErrorCode)
		}
		return &gen.ResolveBooleanResponse{
			Reason:  res.Reason,
			Variant: "default_value",
		}, errors.New(res.ErrorCode)
	}
	return res, nil
}

func (c *GRPCService) ResolveString(flagKey string, context of.EvaluationContext, options ...service.IServiceOption) (*gen.ResolveStringResponse, error) {
	client := c.GetInstance()
	if client == nil {
		return &gen.ResolveStringResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New("CONNECTION_ERROR")
	}
	contextF, err := FormatAsStructpb(context)
	if err != nil {
		log.Error(err)
		return &gen.ResolveStringResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New(models.ParseErrorCode)
	}
	res, err := client.ResolveString(ctx.TODO(), &gen.ResolveStringRequest{
		FlagKey: flagKey,
		Context: contextF,
	})
	if err != nil {
		res, ok := ParseError(err)
		if !ok {
			log.Error(err)
			return &gen.ResolveStringResponse{
				Reason:  models.ErrorReason,
				Variant: "default_value",
			}, errors.New(models.GeneralErrorCode)
		}
		return &gen.ResolveStringResponse{
			Reason:  res.Reason,
			Variant: "default_value",
		}, errors.New(res.ErrorCode)
	}
	return res, nil
}

func (c *GRPCService) ResolveNumber(flagKey string, context of.EvaluationContext, options ...service.IServiceOption) (*gen.ResolveNumberResponse, error) {
	client := c.GetInstance()
	if client == nil {
		return &gen.ResolveNumberResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New("CONNECTION_ERROR")
	}
	contextF, err := FormatAsStructpb(context)
	if err != nil {
		log.Error(err)
		return &gen.ResolveNumberResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New(models.ParseErrorCode)
	}
	res, err := client.ResolveNumber(ctx.TODO(), &gen.ResolveNumberRequest{
		FlagKey: flagKey,
		Context: contextF,
	})
	if err != nil {
		res, ok := ParseError(err)
		if !ok {
			log.Error(err)
			return &gen.ResolveNumberResponse{
				Reason:  models.ErrorReason,
				Variant: "default_value",
			}, errors.New(models.GeneralErrorCode)
		}
		return &gen.ResolveNumberResponse{
			Reason:  res.Reason,
			Variant: "default_value",
		}, errors.New(res.ErrorCode)
	}
	return res, nil
}

func (c *GRPCService) ResolveObject(flagKey string, context of.EvaluationContext, options ...service.IServiceOption) (*gen.ResolveObjectResponse, error) {
	client := c.GetInstance()
	if client == nil {
		return &gen.ResolveObjectResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New("CONNECTION_ERROR")
	}
	contextF, err := FormatAsStructpb(context)
	if err != nil {
		log.Error(err)
		return &gen.ResolveObjectResponse{
			Reason:  models.ErrorReason,
			Variant: "default_value",
		}, errors.New(models.ParseErrorCode)
	}
	res, err := client.ResolveObject(ctx.TODO(), &gen.ResolveObjectRequest{
		FlagKey: flagKey,
		Context: contextF,
	})
	if err != nil {
		res, ok := ParseError(err)
		if !ok {
			log.Error(err)
			return &gen.ResolveObjectResponse{
				Reason:  models.ErrorReason,
				Variant: "default_value",
			}, errors.New(models.GeneralErrorCode)
		}
		return &gen.ResolveObjectResponse{
			Reason:  res.Reason,
			Variant: "default_value",
		}, errors.New(res.ErrorCode)
	}
	return res, nil
}

func ParseError(err error) (*gen.ErrorResponse, bool) {
	st := status.Convert(err)
	details := st.Details()
	if len(details) != 1 {
		log.Errorf("malformed error received by error handler, details received: %d - %v", len(details), details)
		return nil, false
	}
	res, ok := details[0].(*gen.ErrorResponse)
	return res, ok
}

func FormatAsStructpb(evCtx of.EvaluationContext) (*structpb.Struct, error) {
	evCtxM, ok := evCtx.(map[string]interface{})
	if !ok {
		return nil, errors.New("Evaluation context is not map[string]interface{}")
	}
	return structpb.NewStruct(evCtxM)
}
