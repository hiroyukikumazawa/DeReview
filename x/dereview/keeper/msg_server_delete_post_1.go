package keeper

import (
	"context"

	"dereview/x/dereview/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) DeletePost1(ctx context.Context, msg *types.MsgDeletePost1) (*types.MsgDeletePost1Response, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgDeletePost1Response{}, nil
}
