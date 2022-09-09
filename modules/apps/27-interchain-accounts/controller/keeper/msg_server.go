package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
)

var _ types.MsgServer = Keeper{}

<<<<<<< HEAD
// RegisterAccount defines a rpc handler for MsgRegisterAccount
func (k Keeper) RegisterAccount(goCtx context.Context, msg *types.MsgRegisterAccount) (*types.MsgRegisterAccountResponse, error) {
=======
// RegisterInterchainAccount defines a rpc handler for MsgRegisterInterchainAccount
func (k msgServer) RegisterInterchainAccount(goCtx context.Context, msg *types.MsgRegisterInterchainAccount) (*types.MsgRegisterInterchainAccountResponse, error) {
>>>>>>> f8f226d (chore: rename `RegisterAccount` rpc and msgs to `RegisterInterchainAccount` (#2253))
	ctx := sdk.UnwrapSDKContext(goCtx)

	portID, err := icatypes.NewControllerPortID(msg.Owner)
	if err != nil {
		return nil, err
	}

	channelID, err := k.registerInterchainAccount(ctx, msg.ConnectionId, portID, msg.Version)
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterInterchainAccountResponse{
		ChannelId: channelID,
	}, nil
}

// SubmitTx defines a rpc handler for MsgSubmitTx
func (k Keeper) SubmitTx(goCtx context.Context, msg *types.MsgSubmitTx) (*types.MsgSubmitTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	portID, err := icatypes.NewControllerPortID(msg.Owner)
	if err != nil {
		return nil, err
	}

	channelID, found := k.GetActiveChannelID(ctx, msg.ConnectionId, portID)
	if !found {
		return nil, sdkerrors.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel for port %s", portID)
	}

	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, channelID))
	if !found {
		return nil, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	seq, err := k.SendTx(ctx, chanCap, msg.ConnectionId, portID, msg.PacketData, msg.TimeoutTimestamp)
	if err != nil {
		return nil, err
	}

	return &types.MsgSubmitTxResponse{Sequence: seq}, nil
}
