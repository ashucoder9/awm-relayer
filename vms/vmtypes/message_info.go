// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vmtypes

import (
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ethereum/go-ethereum/common"
)

// WarpLogInfo is provided by the subscription and
// describes the transaction information emitted by the source chain,
// including the Warp Message payload bytes
type WarpLogInfo struct {
	SourceAddress    common.Address
	SourceTxID       []byte
	UnsignedMsg      *avalancheWarp.UnsignedMessage
	BlockNumber      uint64
	IsCatchUpMessage bool
}
