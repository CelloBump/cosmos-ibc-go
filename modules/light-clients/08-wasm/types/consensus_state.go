package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/ibc-go/v7/modules/core/exported"
)

var _ exported.ConsensusState = (*ConsensusState)(nil)

// NewConsensusState creates a new ConsensusState instance.
func NewConsensusState(data []byte, timestamp uint64) *ConsensusState {
	return &ConsensusState{
		Data: data,
	}
}

// ClientType returns Wasm type.
func (cs ConsensusState) ClientType() string {
	return exported.Wasm
}

// GetTimestamp returns block time in nanoseconds of the header that created consensus state.
func (cs ConsensusState) GetTimestamp() uint64 {
	return 0
}

// ValidateBasic defines a basic validation for the wasm client consensus state.
func (cs ConsensusState) ValidateBasic() error {
	if len(cs.Data) == 0 {
		return sdkerrors.Wrap(ErrInvalidData, "data cannot be empty")
	}

	return nil
}
