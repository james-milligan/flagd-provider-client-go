package tests

import (
	"errors"
	"testing"

	service "github.com/james-milligan/flagd-provider-client-go/pkg/service/http"
	mocks "github.com/james-milligan/flagd-provider-client-go/pkg/service/http/tests/mocks"
	models "github.com/open-feature/flagd/pkg/model"
	schemaV1 "go.buf.build/grpc/go/open-feature/flagd/schema/v1"
)

type TestServiceResolveBooleanArgs struct {
	name                     string
	ServiceClientMockSetup   mocks.ServiceClientMockSetup
	HTTPServiceConfiguration service.HTTPServiceConfiguration

	flagKey string
	evCtx   interface{}

	value   bool
	variant string
	reason  string
	err     error
}

func TestServiceResolveBoolean(t *testing.T) {
	tests := []TestServiceResolveBooleanArgs{
		{
			name: "happy path",
			ServiceClientMockSetup: mocks.ServiceClientMockSetup{
				InUrl: "http://localhost:8080/flags/bool/resolve/boolean",
				InCtx: nil,
				OutBody: schemaV1.ResolveBooleanResponse{
					Value:   true,
					Variant: "on",
					Reason:  models.StaticReason,
				},
				OutSC:  200,
				OutErr: nil,
			},
			HTTPServiceConfiguration: service.HTTPServiceConfiguration{
				Port: 8080,
				Host: "localhost",
			},
			flagKey: "bool",
			evCtx:   nil,

			value:   true,
			variant: "on",
			reason:  models.StaticReason,
			err:     nil,
		},
		{
			name: "handle non 200",
			ServiceClientMockSetup: mocks.ServiceClientMockSetup{
				InUrl: "http://localhost:8080/flags/bool/resolve/boolean",
				InCtx: nil,
				OutBody: schemaV1.ErrorResponse{
					Reason:    models.StaticReason,
					ErrorCode: models.UnknownReason,
				},
				OutSC:  500,
				OutErr: nil,
			},
			HTTPServiceConfiguration: service.HTTPServiceConfiguration{
				Port: 8080,
				Host: "localhost",
			},
			flagKey: "bool",
			evCtx:   nil,

			value:   false,
			variant: "default_value",
			reason:  models.ErrorReason,
			err:     errors.New(models.UnknownReason),
		},
		{
			name: "handle error",
			ServiceClientMockSetup: mocks.ServiceClientMockSetup{
				InUrl:  "http://localhost:8080/flags/bool/resolve/boolean",
				InCtx:  nil,
				OutSC:  0,
				OutErr: errors.New("Its all gone wrong"),
			},
			HTTPServiceConfiguration: service.HTTPServiceConfiguration{
				Port: 8080,
				Host: "localhost",
			},
			flagKey: "bool",
			evCtx:   nil,

			value:   false,
			variant: "default_value",
			reason:  models.ErrorReason,
			err:     errors.New("Its all gone wrong"),
		},
	}

	for _, test := range tests {
		srv := service.HTTPService{
			Client: &mocks.ServiceClient{
				ServiceClientMockSetup: test.ServiceClientMockSetup,
				Testing:                t,
			},
			HTTPServiceConfiguration: &test.HTTPServiceConfiguration,
		}
		res, err := srv.ResolveBoolean(test.flagKey, test.evCtx)
		if (test.err != nil && err != nil) && test.err.Error() != err.Error() {
			t.Errorf("%s: unexpected error received, expected %v, got %v", test.name, test.err, err)
		}
		if res.Reason != test.reason {
			t.Errorf("%s: unexpected reason received, expected %v, got %v", test.name, test.reason, res.Reason)
		}
		if res.Value != test.value {
			t.Errorf("%s: unexpected value received, expected %v, got %v", test.name, test.value, res.Value)
		}
		if res.Variant != test.variant {
			t.Errorf("%s: unexpected variant received, expected %v, got %v", test.name, test.variant, res.Variant)
		}
	}
}
