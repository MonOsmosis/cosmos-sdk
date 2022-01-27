package simulation

import (
	"math/rand"

	simtypes "github.com/MonOsmosis/cosmos-sdk/types/simulation"
	"github.com/MonOsmosis/cosmos-sdk/x/ibc/core/02-client/types"
)

// GenClientGenesis returns the default client genesis state.
func GenClientGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
