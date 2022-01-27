package types

import (
	"time"

	sdk "github.com/MonOsmosis/cosmos-sdk/types"
	stakingtypes "github.com/MonOsmosis/cosmos-sdk/x/staking/types"
)

// StakingKeeper expected staking keeper
type StakingKeeper interface {
	GetHistoricalInfo(ctx sdk.Context, height int64) (stakingtypes.HistoricalInfo, bool)
	UnbondingTime(ctx sdk.Context) time.Duration
}
