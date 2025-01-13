/*
NOTE: Usage of x/params to manage parameters is deprecated in favor of x/gov
controlled execution of MsgUpdateParams messages. These types remains solely
for migration purposes and will be removed in a future release.
*/
package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	KeyMintDenom           = []byte("MintDenom")
	KeyInflationRateChange = []byte("InflationRateChange")
	KeyInflationMax        = []byte("InflationMax")
	KeyInflationMin        = []byte("InflationMin")
	KeyGoalBonded          = []byte("GoalBonded")
	KeyBlocksPerYear       = []byte("BlocksPerYear")
	KeyMaxStakedRatio      = []byte("MaxStakedRatio")
	KeyApyAtMaxStakedRatio = []byte("ApyAtMaxStakedRatio")
	KeyMinStakedRatio      = []byte("MinStakedRatio")
	KeyApyAtMinStakedRatio = []byte("ApyAtMinStakedRatio")
	KeyDecayRate           = []byte("DecayRate")
)

// Deprecated: ParamTable for minting module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable()
}

func dummyValidate(i interface{}) error {
	// TODO: implement
	return nil
}
