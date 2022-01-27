package codec

import (
	codectypes "github.com/MonOsmosis/cosmos-sdk/codec/types"
	"github.com/MonOsmosis/cosmos-sdk/crypto/keys/ed25519"
	"github.com/MonOsmosis/cosmos-sdk/crypto/keys/multisig"
	"github.com/MonOsmosis/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/MonOsmosis/cosmos-sdk/crypto/types"
)

// RegisterInterfaces registers the sdk.Tx interface.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface("cosmos.crypto.PubKey", (*cryptotypes.PubKey)(nil))
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ed25519.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &secp256k1.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &multisig.LegacyAminoPubKey{})
}
