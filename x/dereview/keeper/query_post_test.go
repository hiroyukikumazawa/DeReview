package keeper_test

import (
	"context"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "dereview/testutil/keeper"
	"dereview/testutil/nullify"
	"dereview/x/dereview/keeper"
	"dereview/x/dereview/types"
)

func createNPost(keeper keeper.Keeper, ctx context.Context, n int) []types.Post {
	items := make([]types.Post, n)
	for i := range items {
		iu := uint64(i)
		items[i].Id = iu
		_ = keeper.Post.Set(ctx, iu, items[i])
		_ = keeper.PostSeq.Set(ctx, iu)
	}
	return items
}

func TestPostQuerySingle(t *testing.T) {
	k, ctx, _ := keepertest.DereviewKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNPost(k, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetPostRequest
		response *types.QueryGetPostResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetPostRequest{Id: msgs[0].Id},
			response: &types.QueryGetPostResponse{Post: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetPostRequest{Id: msgs[1].Id},
			response: &types.QueryGetPostResponse{Post: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetPostRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetPost(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestPostQueryPaginated(t *testing.T) {
	k, ctx, _ := keepertest.DereviewKeeper(t)
	qs := keeper.NewQueryServerImpl(k)
	msgs := createNPost(k, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPostRequest {
		return &types.QueryAllPostRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListPost(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Post), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Post),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListPost(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Post), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Post),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListPost(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Post),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListPost(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
