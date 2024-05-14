package keeper

import (
	"context"

	"dereview/x/dereview/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ShowPost1(ctx context.Context, req *types.QueryShowPost1Request) (*types.QueryShowPost1Response, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: Process the query

	return &types.QueryShowPost1Response{}, nil
}
