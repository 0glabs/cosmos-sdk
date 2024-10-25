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
	var apy sdk.Dec
	if X.LT(params.MinStakedRatio) {
		apy = params.ApyAtMinStakedRatio
	} else if X.GT(params.MaxStakedRatio) {
		apy = params.ApyAtMaxStakedRatio
	} else {
		exp := params.DecayRate.Neg().Mul(params.MaxStakedRatio.Sub(params.MinStakedRatio))
		c := decExp(exp)
		d := params.ApyAtMaxStakedRatio.Sub(params.ApyAtMinStakedRatio.Mul(c)).Quo(sdk.OneDec().Sub(c))
		expBonded := params.DecayRate.Neg().Mul(X.Sub(params.MinStakedRatio))
		cBonded := decExp(expBonded)
		e := params.ApyAtMinStakedRatio.Sub(d).Mul(cBonded)
		apy = d.Add(e)
	}

	inflation := apy.Mul(bondedRatio)

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
