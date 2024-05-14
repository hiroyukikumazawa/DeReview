package keeper

import (
	"context"

	"dereview/x/dereview/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListPost1(ctx context.Context, req *types.QueryListPost1Request) (*types.QueryListPost1Response, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: Process the query

	return &types.QueryListPost1Response{}, nil
}
