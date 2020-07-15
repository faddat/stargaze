package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Curve defines an interface that every bonding curve should implement
type Curve interface {
	// ToMint returns the number of new tokens to mint for the given amount of money
	ToMint(money sdk.Coin, tokenSupply, poolBalance sdk.Int) (sdk.Int, error)
}
