package keeper

import (
	"context"

	"dereview/x/dereview/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) HelloDereview(ctx context.Context, req *types.QueryHelloDereviewRequest) (*types.QueryHelloDereviewResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// TODO: Process the query

	return &types.QueryHelloDereviewResponse{}, nil
}
