package tests

import (
	"errors"
	"testing"

	service "github.com/james-milligan/flagd-provider-client-go/pkg/service/grpc"
	"github.com/james-milligan/flagd-provider-client-go/pkg/service/grpc/tests/mocks"
	models "github.com/open-feature/flagd/pkg/model"
	schemaV1 "go.buf.build/grpc/go/open-feature/flagd/schema/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TestServiceResolveBooleanArgs struct {
	name string
	mocks.MockResolveBooleanArgs

	flagKey string
	evCtx   interface{}

	value     bool
	variant   string
	reason    string
	err       error
	customErr string
}

func TestServiceResolveBoolean(t *testing.T) {
	tests := []TestServiceResolveBooleanArgs{
		{
			name: "happy path",
			MockResolveBooleanArgs: mocks.MockResolveBooleanArgs{
				InFK: "bool",
				InCtx: map[string]interface{}{
					"dog": "cat",
				},
				Out: schemaV1.ResolveBooleanResponse{
					Value:   true,
					Variant: "on",
					Reason:  models.StaticReason,
				},
			},
			flagKey: "bool",
			evCtx: map[string]interface{}{
				"dog": "cat",
			},
			value:   true,
			variant: "on",
			reason:  models.StaticReason,
			err:     nil,
		},
		{
			name:    "FormatAsStructpb fails",
			flagKey: "bool",
			evCtx:   "not a map[string]interface{}!",
			value:   false,
			variant: "default_value",
			reason:  models.ErrorReason,
			err:     errors.New(models.ParseErrorCode),
		},
		{
			name: "custom error response",
			MockResolveBooleanArgs: mocks.MockResolveBooleanArgs{
				InFK: "bool",
				InCtx: map[string]interface{}{
					"dog": "cat",
				},
				OutErr: status.Error(codes.NotFound, "custom message"),
			},
			flagKey: "bool",
			evCtx: map[string]interface{}{
				"dog": "cat",
			},
			value:     false,
			variant:   "default_value",
			reason:    models.ErrorReason,
			customErr: "CUSTOM ERROR",
			err:       errors.New("CUSTOM ERROR"),
		},
	}

	for _, test := range tests {
		if test.customErr != "" {
			st, ok := status.FromError(test.MockResolveBooleanArgs.OutErr)
			if !ok {
				t.Errorf("%s: malformed error status recieved, cannot attach custom properties", test.name)
			}
			stWD, err := st.WithDetails(&schemaV1.ErrorResponse{
				ErrorCode: test.customErr,
				Reason:    models.ErrorReason,
			})
			if err != nil {
				t.Error(err)
			}
			test.MockResolveBooleanArgs.OutErr = stWD.Err()
		}
		srv := service.GRPCService{
			Client: &mocks.MockClient{
				RBArgs:  test.MockResolveBooleanArgs,
				Testing: t,
			},
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
