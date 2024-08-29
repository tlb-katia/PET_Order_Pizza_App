package tests

import (
	"github.com/stretchr/testify/require"
	"github.com/tlb-katia/PET_Order_Pizza_App/internal/storage"
	"github.com/tlb-katia/PET_Order_Pizza_App/tests/suite"
	pizza_orderv1 "github.com/tlb-katia/protos/protos/gen/go/pizza-order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestCheckOrderStatus(t *testing.T) {
	ctx, st := suite.New(t)

	tests := []struct {
		testName     string
		orderId      string
		expectedCode codes.Code
		expectedErr  string
		mockErr      error
	}{
		{
			testName:     "Check Order Status With Valid Order ID",
			orderId:      "1",
			expectedCode: codes.OK,
			mockErr:      nil,
		},
		{
			testName:     "Check Order Status With Empty Order ID",
			orderId:      "",
			expectedCode: codes.InvalidArgument,
			expectedErr:  "Order ID is required",
		},
		{
			testName:     "Check Order Status With Invalid Order ID",
			orderId:      "invalid-order-id",
			expectedCode: codes.InvalidArgument,
			expectedErr:  "order not found",
			mockErr:      storage.ErrOrderNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			resp, err := st.OrderClient.CheckOrderStatus(ctx, &pizza_orderv1.OrderStatusRequest{
				OrderId: tt.orderId,
			})
			if tt.expectedCode == codes.OK {
				require.NoError(t, err)
				require.Equal(t, tt.orderId, resp.OrderId)
			} else {
				require.Error(t, err)
				stErr, ok := status.FromError(err)
				require.True(t, ok, "Expected a gRPC status error")
				require.Equal(t, tt.expectedCode, stErr.Code())
				require.Contains(t, stErr.Message(), tt.expectedErr)
			}
		})
	}
}
