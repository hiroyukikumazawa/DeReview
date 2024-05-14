package keeper

import (
	"context"

	"dereview/x/dereview/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) UpdatePost1(ctx context.Context, msg *types.MsgUpdatePost1) (*types.MsgUpdatePost1Response, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgUpdatePost1Response{}, nil
}
