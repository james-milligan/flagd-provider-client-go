package grpc

import (
	ctx "context"
	"errors"

	"github.com/james-milligan/flagd-provider-client-go/pkg/service"
	gen "github.com/james-milligan/flagd-provider-client-go/schemas/protobuf/gen/v1"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

func (c *GRPCService) ResolveBoolean(flagKey string, defaultValue bool, context map[string]interface{}, options ...service.ISercviceOption) (*gen.ResolveBooleanResponse, error) {
	client := c.GetInstance()
	if client == nil {
		return &gen.ResolveBooleanResponse{
			Value:   defaultValue,
			Reason:  "ERROR",
			Variant: "default_value",
		}, errors.New("CONNECTION_ERROR")
	}
	pbstruct, err := structpb.NewStruct(context)
	if err != nil {
		log.Error(err)
		return &gen.ResolveBooleanResponse{
			Value:   defaultValue,
			Reason:  "ERROR",
			Variant: "default_value",
		}, errors.New("INTERNAL_ERROR")
	}
	res, err := client.ResolveBoolean(ctx.TODO(), &gen.ResolveBooleanRequest{
		FlagKey: flagKey,
		Context: pbstruct,
	})
	if err != nil {
		res, ok := ParseError(err)
		if !ok {
			log.Error(err)
			return &gen.ResolveBooleanResponse{
				Value:   defaultValue,
				Reason:  "ERROR",
				Variant: "default_value",
			}, errors.New("INTERNAL_ERROR")
		}
		return &gen.ResolveBooleanResponse{
			Value:   defaultValue,
			Reason:  res.Reason,
			Variant: "default_value",
		}, errors.New(res.ErrorCode)
	}
	return res, nil
}

func (c *GRPCService) ResolveString(flagKey string, defaultValue string, context map[string]interface{}, options ...service.ISercviceOption) (*gen.ResolveStringResponse, error) {
	client := c.GetInstance()
	if client == nil {
		return &gen.ResolveStringResponse{
			Value:   defaultValue,
			Reason:  "ERROR",
			Variant: "default_value",
		}, errors.New("CONNECTION_ERROR")
	}
	pbstruct, err := structpb.NewStruct(context)
	if err != nil {
		log.Error(err)
		return &gen.ResolveStringResponse{
			Value:   defaultValue,
			Reason:  "ERROR",
			Variant: "default_value",
		}, errors.New("INTERNAL_ERROR")
	}
	res, err := client.ResolveString(ctx.TODO(), &gen.ResolveStringRequest{
		FlagKey: flagKey,
		Context: pbstruct,
	})
	if err != nil {
		res, ok := ParseError(err)
		if !ok {
			log.Error(err)
			return &gen.ResolveStringResponse{
				Value:   defaultValue,
				Reason:  "ERROR",
				Variant: "default_value",
			}, errors.New("INTERNAL_ERROR")
		}
		return &gen.ResolveStringResponse{
			Value:   defaultValue,
			Reason:  res.Reason,
			Variant: "default_value",
		}, errors.New(res.ErrorCode)
	}
	return res, nil
}

func (c *GRPCService) ResolveNumber(flagKey string, defaultValue float32, context map[string]interface{}, options ...service.ISercviceOption) (*gen.ResolveNumberResponse, error) {
	client := c.GetInstance()
	if client == nil {
		return &gen.ResolveNumberResponse{
			Value:   defaultValue,
			Reason:  "ERROR",
			Variant: "default_value",
		}, errors.New("CONNECTION_ERROR")
	}
	pbstruct, err := structpb.NewStruct(context)
	if err != nil {
		log.Error(err)
		return &gen.ResolveNumberResponse{
			Value:   defaultValue,
			Reason:  "ERROR",
			Variant: "default_value",
		}, errors.New("INTERNAL_ERROR")
	}
	res, err := client.ResolveNumber(ctx.TODO(), &gen.ResolveNumberRequest{
		FlagKey: flagKey,
		Context: pbstruct,
	})
	if err != nil {
		res, ok := ParseError(err)
		if !ok {
			log.Error(err)
			return &gen.ResolveNumberResponse{
				Value:   defaultValue,
				Reason:  "ERROR",
				Variant: "default_value",
			}, errors.New("INTERNAL_ERROR")
		}
		return &gen.ResolveNumberResponse{
			Value:   defaultValue,
			Reason:  res.Reason,
			Variant: "default_value",
		}, errors.New(res.ErrorCode)
	}
	return res, nil
}

func (c *GRPCService) ResolveObject(flagKey string, defaultValue *structpb.Struct, context map[string]interface{}, options ...service.ISercviceOption) (*gen.ResolveObjectResponse, error) {
	client := c.GetInstance()
	if client == nil {
		return &gen.ResolveObjectResponse{
			Value:   defaultValue,
			Reason:  "ERROR",
			Variant: "default_value",
		}, errors.New("CONNECTION_ERROR")
	}
	pbstruct, err := structpb.NewStruct(context)
	if err != nil {
		log.Error(err)
		return &gen.ResolveObjectResponse{
			Value:   defaultValue,
			Reason:  "ERROR",
			Variant: "default_value",
		}, errors.New("INTERNAL_ERROR")
	}
	res, err := client.ResolveObject(ctx.TODO(), &gen.ResolveObjectRequest{
		FlagKey: flagKey,
		Context: pbstruct,
	})
	if err != nil {
		res, ok := ParseError(err)
		if !ok {
			log.Error(err)
			return &gen.ResolveObjectResponse{
				Value:   defaultValue,
				Reason:  "ERROR",
				Variant: "default_value",
			}, errors.New("INTERNAL_ERROR")
		}
		return &gen.ResolveObjectResponse{
			Value:   defaultValue,
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
