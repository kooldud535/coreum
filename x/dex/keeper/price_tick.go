package keeper

import (
	"math/big"

	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"

	cbig "github.com/CoreumFoundation/coreum/v6/pkg/math/big"
	"github.com/CoreumFoundation/coreum/v6/x/dex/types"
)

func validatePriceTick(price *big.Rat, baseURA, quoteURA sdkmath.LegacyDec, priceTickExponent int32) error {
	// Both baseURA & quoteURA are multiplied by 10^LegacyPrecision when converting to BigInt,
	// but since we divide one by other we can pass them as is.
	priceTick, exponent := ComputePriceTick(baseURA.BigInt(), quoteURA.BigInt(), priceTickExponent)
	if !isPriceTickValid(price, priceTick) {
		return sdkerrors.Wrapf(
			types.ErrInvalidInput,
			"invalid price, has to be multiple of price tick: 10^%d",
			exponent,
		)
	}

	return nil
}

func isPriceTickValid(price *big.Rat, priceTick *big.Rat) bool {
	_, remainder := cbig.RatQuoWithIntRemainder(price, priceTick)
	return cbig.IntEqZero(remainder)
}

// ComputePriceTick calculates price tick for a market by base & quote unified_ref_amount and price_tick_exponent.
func ComputePriceTick(
	baseURA,
	quoteURA *big.Int,
	priceTickExponent int32,
) (*big.Rat, int64) {
	// price_tick = 10^(price_tick_exponent + ceil(log10(unified_ref_amount(BBB)/unified_ref_amount(AAA))))
	exponent := int64(priceTickExponent) + cbig.RatLog10RoundUp(cbig.NewRatFromBigInts(quoteURA, baseURA))

	return cbig.RatTenToThePower(exponent), exponent
}
