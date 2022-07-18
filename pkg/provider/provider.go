package provider

import (
	"github.com/james-milligan/flagd-provider-client-go/pkg/service"
	of "github.com/open-feature/golang-sdk/pkg/openfeature"
)

type Provider struct {
	service service.IService
}

func (p *Provider) Metadata() of.Metadata {
	return of.Metadata{
		Name: "flagd",
	}
}

func (p *Provider) GetBooleanEvaluation(flagKey string, defaultValue bool, evalCtx of.EvaluationContext, options ...of.EvaluationOption) of.BoolResolutionDetail {
	res, err := p.service.ResolveBoolean(flagKey, evalCtx)
	if err != nil {
		return of.BoolResolutionDetail{
			Value: defaultValue,
			ResolutionDetail: of.ResolutionDetail{
				Reason:    res.Reason,
				Value:     defaultValue,
				Variant:   res.Variant,
				ErrorCode: err.Error(),
			},
		}
	}
	return of.BoolResolutionDetail{
		Value: res.Value,
		ResolutionDetail: of.ResolutionDetail{
			Reason:  res.Reason,
			Value:   res.Value,
			Variant: res.Variant,
		},
	}
}

func (p *Provider) GetStringEvaluation(flagKey string, defaultValue string, evalCtx of.EvaluationContext, options ...of.EvaluationOption) of.StringResolutionDetail {
	res, err := p.service.ResolveString(flagKey, evalCtx)
	if err != nil {
		return of.StringResolutionDetail{
			Value: defaultValue,
			ResolutionDetail: of.ResolutionDetail{
				Reason:    res.Reason,
				Value:     defaultValue,
				Variant:   res.Variant,
				ErrorCode: err.Error(),
			},
		}
	}
	return of.StringResolutionDetail{
		Value: res.Value,
		ResolutionDetail: of.ResolutionDetail{
			Reason:  res.Reason,
			Value:   res.Value,
			Variant: res.Variant,
		},
	}
}

func (p *Provider) GetNumberEvaluation(flagKey string, defaultValue float64, evalCtx of.EvaluationContext, options ...of.EvaluationOption) of.NumberResolutionDetail {
	res, err := p.service.ResolveNumber(flagKey, evalCtx)
	if err != nil {
		return of.NumberResolutionDetail{
			Value: defaultValue,
			ResolutionDetail: of.ResolutionDetail{
				Reason:    res.Reason,
				Value:     defaultValue,
				Variant:   res.Variant,
				ErrorCode: err.Error(),
			},
		}
	}
	return of.NumberResolutionDetail{
		Value: float64(res.Value), // todo - update flagd to output float64 (proto file change)
		ResolutionDetail: of.ResolutionDetail{
			Reason:  res.Reason,
			Value:   res.Value,
			Variant: res.Variant,
		},
	}
}

func (p *Provider) GetObjectEvaluation(flagKey string, defaultValue interface{}, evalCtx of.EvaluationContext, options ...of.EvaluationOption) of.ResolutionDetail {
	res, err := p.service.ResolveObject(flagKey, evalCtx)
	if err != nil {
		return of.ResolutionDetail{
			Reason:    res.Reason,
			Value:     defaultValue,
			Variant:   res.Variant,
			ErrorCode: err.Error(),
		}
	}
	return of.ResolutionDetail{
		Reason:  res.Reason,
		Value:   res.Value,
		Variant: res.Variant,
	}
}
