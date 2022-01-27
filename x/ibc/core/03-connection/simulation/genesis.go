package simulation

import (
	"math/rand"

	simtypes "github.com/MonOsmosis/cosmos-sdk/types/simulation"
	"github.com/MonOsmosis/cosmos-sdk/x/ibc/core/03-connection/types"
)

// GenConnectionGenesis returns the default connection genesis state.
func GenConnectionGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
