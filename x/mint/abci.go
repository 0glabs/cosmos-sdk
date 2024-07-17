package mint

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/mint/keeper"
	"github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/shopspring/decimal"
)

func decExp(x sdk.Dec) sdk.Dec {
	xDec := decimal.NewFromBigInt(x.BigInt(), -18)
	expDec, _ := xDec.ExpTaylor(18)
	expInt := expDec.Shift(18).BigInt()
	return sdk.NewDecFromBigIntWithPrec(expInt, 18)
}

func NextInflationRate(ctx sdk.Context, params types.Params, bondedRatio sdk.Dec, circulatingRatio sdk.Dec) sdk.Dec {
	X := bondedRatio.Quo(circulatingRatio)
	ctx.Logger().Info("NextInflationRate", "params", params)
	var apy sdk.Dec
	if X.LT(params.XMin) {
		apy = params.YMax
	} else {
		exp := params.DecayRate.Neg().Mul(params.XMax.Sub(params.XMin))
		c := decExp(exp)
		d := params.YMin.Sub(params.YMax.Mul(c)).Quo(sdk.OneDec().Sub(c))
		expBonded := params.DecayRate.Neg().Mul(X.Sub(params.XMin))
		cBonded := decExp(expBonded)
		e := params.YMax.Sub(d).Mul(cBonded)
		apy = d.Add(e)
	}

	inflation := apy.Mul(bondedRatio)

	// // The target annual inflation rate is recalculated for each previsions cycle. The
	// // inflation is also subject to a rate change (positive or negative) depending on
	// // the distance from the desired ratio (67%). The maximum rate change possible is
	// // defined to be 13% per year, however the annual inflation is capped as between
	// // 7% and 20%.

	// // (1 - bondedRatio/GoalBonded) * InflationRateChange
	// inflationRateChangePerYear := sdk.OneDec().
	// 	Sub(bondedRatio.Quo(params.GoalBonded)).
	// 	Mul(params.InflationRateChange)
	// inflationRateChange := inflationRateChangePerYear.Quo(sdk.NewDec(int64(params.BlocksPerYear)))

	// // adjust the new annual inflation for this next cycle
	// inflation := minter.Inflation.Add(inflationRateChange) // note inflationRateChange may be negative
	// if inflation.GT(params.InflationMax) {
	// 	inflation = params.InflationMax
	// }
	// if inflation.LT(params.InflationMin) {
	// 	inflation = params.InflationMin
	// }

	ctx.Logger().Info(
		"nextInflationRate",
		"bondedRatio", bondedRatio,
		"circulatingRatio", circulatingRatio,
		"apy", apy,
		"inflation", inflation,
		"params", params,
	)
	return inflation
}

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper, ic types.InflationCalculationFn) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)

	// recalculate inflation rate
	totalStakingSupply := k.StakingTokenSupply(ctx)
	bondedRatio := k.BondedRatio(ctx)
	circulatingRaio := k.CirculatingRatio(ctx)
	minter.Inflation = NextInflationRate(ctx, params, bondedRatio, circulatingRaio)
	minter.AnnualProvisions = minter.NextAnnualProvisions(params, totalStakingSupply)
	k.SetMinter(ctx, minter)

	// mint coins, update supply
	mintedCoin := minter.BlockProvision(params)
	mintedCoins := sdk.NewCoins(mintedCoin)

	err := k.MintCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	// send the minted coins to the fee collector account
	err = k.AddCollectedFees(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	if mintedCoin.Amount.IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyBondedRatio, bondedRatio.String()),
			sdk.NewAttribute(types.AttributeKeyInflation, minter.Inflation.String()),
			sdk.NewAttribute(types.AttributeKeyAnnualProvisions, minter.AnnualProvisions.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
		),
	)
}
