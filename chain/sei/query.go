package sei

import (
	"context"
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// rangeBlockMessages iterates through all a block's transactions and each transaction's messages yielding to f.
// Return true from f to stop iteration.
func rangeBlockMessages(ctx context.Context, interfaceRegistry codectypes.InterfaceRegistry, chainNode *ChainNode, height uint64, done func(sdk.Msg) bool) error {
	h := int64(height)

	blockRes, err := chainNode.TMClient.Block(ctx, &h)

	if err != nil {
		return fmt.Errorf("error querying block %w", err)
	}

	if err != nil {
		return fmt.Errorf("tendermint rpc get block: %w", err)
	}
	for _, txbz := range blockRes.Block.Txs {
		tx, err := decodeTX(interfaceRegistry, txbz)
		if err != nil {
			return fmt.Errorf("decode tendermint tx: %w", err)
		}
		for _, m := range tx.GetMsgs() {
			if ok := done(m); ok {
				return nil
			}
		}
	}
	return nil
}
