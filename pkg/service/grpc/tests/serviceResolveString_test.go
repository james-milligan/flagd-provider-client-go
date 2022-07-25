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

type TestServiceResolveStringArgs struct {
	name string
	mocks.MockResolveStringArgs

	flagKey string
	evCtx   interface{}

	value     string
	variant   string
	reason    string
	err       error
	customErr string
}

func TestServiceResolveString(t *testing.T) {
	tests := []TestServiceResolveStringArgs{
		{
			name: "happy path",
			MockResolveStringArgs: mocks.MockResolveStringArgs{
				InFK: "bool",
				InCtx: map[string]interface{}{
					"dog": "cat",
				},
				Out: schemaV1.ResolveStringResponse{
					Value:   "on",
					Variant: "on",
					Reason:  "STATIC",
				},
			},
			flagKey: "bool",
			evCtx: map[string]interface{}{
				"dog": "cat",
			},
			variant: "on",
			value:   "on",
			reason:  "STATIC",
			err:     nil,
		},
		{
			name:    "FormatAsStructpb fails",
			flagKey: "bool",
			evCtx:   "not a map[string]interface{}!",
			variant: "default_value",
			reason:  "ERROR",
			err:     errors.New(models.ParseErrorCode),
		},
		{
			name: "custom error response",
			MockResolveStringArgs: mocks.MockResolveStringArgs{
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
			variant:   "default_value",
			reason:    "ERROR",
			customErr: "CUSTOM ERROR",
			err:       errors.New("CUSTOM ERROR"),
		},
	}

	for _, test := range tests {
		if test.customErr != "" {
			st, ok := status.FromError(test.MockResolveStringArgs.OutErr)
			if !ok {
				t.Errorf("%s: malformed error status recieved, cannot attach custom properties", test.name)
			}
			stWD, err := st.WithDetails(&schemaV1.ErrorResponse{
				ErrorCode: test.customErr,
				Reason:    "ERROR",
			})
			if err != nil {
				t.Error(err)
			}
			test.MockResolveStringArgs.OutErr = stWD.Err()
		}
		srv := service.GRPCService{
			Client: &mocks.MockClient{
				RSArgs:  test.MockResolveStringArgs,
				Testing: t,
			},
		}
		res, err := srv.ResolveString(test.flagKey, test.evCtx)
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
