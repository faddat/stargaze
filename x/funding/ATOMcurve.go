package funding

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/public-awesome/stakebird/x/funding/types"
)

// ATOMCurve defines the curve against bonded ATOM
type ATOMCurve struct {
	ReserveRatio sdk.Dec
}

// ToMint returns the new tokens to be minted for the bonded amount of ATOMs
/**
  Formula:
  toMint = tokenSupply * ((1 + buyingAmount / poolBalance) ^ reserveRatio - 1)
*/
func (c *ATOMCurve) ToMint(amount sdk.Coin, tokenSupply, poolBalance sdk.Int) (sdk.Int, error) {
	temp := amount.Amount.Quo(poolBalance)
	temp2 := temp.AddRaw(1)
	temp3 := temp2.ToDec().Power(c.ReserveRatio.BigInt().Uint64())
	temp4 := temp3.TruncateInt().SubRaw(1)
	toMint := tokenSupply.Mul(temp4)

	return toMint, nil
}

// GetBondingCurve inits and returns the Bonding Curve
func GetBondingCurve() types.Curve {
	return &ATOMCurve{
		ReserveRatio: sdk.NewDecWithPrec(33, 2),
	}
}
