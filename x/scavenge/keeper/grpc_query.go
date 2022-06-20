package keeper

import (
	"github.com/irclausen/scavenge/x/scavenge/types"
)

var _ types.QueryServer = Keeper{}
