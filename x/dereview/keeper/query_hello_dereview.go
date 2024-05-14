package keeper

import (
	"context"
	"fmt"

	"dereview/x/dereview/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) HelloDereview(ctx context.Context, req *types.QueryHelloDereviewRequest) (*types.QueryHelloDereviewResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Validation and Context unwrapping
	ctx = sdk.UnwrapSDKContext(ctx)
	_ = ctx

	// Custom Response
	return &types.QueryHelloDereviewResponse{Name: fmt.Sprintf("Welcom to DeReview, %s!", req.Name)}, nil
}
