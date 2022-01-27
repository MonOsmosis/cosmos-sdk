package types

import (
	"github.com/MonOsmosis/cosmos-sdk/codec"
	cryptocodec "github.com/MonOsmosis/cosmos-sdk/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
