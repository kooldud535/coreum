package stress

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types/query"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

	integrationtests "github.com/CoreumFoundation/coreum/v6/integration-tests"
	"github.com/CoreumFoundation/coreum/v6/pkg/client"
	"github.com/CoreumFoundation/coreum/v6/testutil/integration"
	dextypes "github.com/CoreumFoundation/coreum/v6/x/dex/types"
)

// TestLimitOrdersStressMatching tests the dex modules ability to match a lot of orders.
func TestLimitOrdersStressMatching(t *testing.T) {
	// t.Parallel() is disabled intentionally
	ctx, chain := integrationtests.NewCoreumTestingContext(t)

	preCreatedOrdersCreator := "devcore1fnrehr95flfgnzjcatv7a8hpernwufpd5zjm2v"
	baseDenom := "dexsu-" + preCreatedOrdersCreator // created VIA the genesis
	quoteDenom := chain.ChainSettings.Denom
	ordersCount := 2_000

	requireT := require.New(t)
	dexClient := dextypes.NewQueryClient(chain.ClientContext)
	bankClient := banktypes.NewQueryClient(chain.ClientContext)

	// check initial state
	creatorOrdersRes, err := dexClient.Orders(ctx, &dextypes.QueryOrdersRequest{
		Creator: preCreatedOrdersCreator,
		Pagination: &query.PageRequest{
			Limit:      query.PaginationMaxLimit,
			CountTotal: true,
		},
	})
	requireT.NoError(err)
	requireT.Len(creatorOrdersRes.Orders, ordersCount)

	sellTotal := sdkmath.ZeroInt()
	for _, order := range creatorOrdersRes.Orders {
		// check at lest one order
		requireT.Equal(baseDenom, order.BaseDenom)
		requireT.Equal(quoteDenom, order.QuoteDenom)
		requireT.Equal(dextypes.SIDE_SELL, order.Side)
		requireT.Equal(dextypes.MustNewPriceFromString("1").String(), order.Price.String())
		sellTotal = sellTotal.Add(order.Quantity)
	}

	// place order to match all orders
	dexParamsRes, err := dexClient.Params(ctx, &dextypes.QueryParamsRequest{})
	requireT.NoError(err)

	acc1 := chain.GenAccount()
	buyTotal := sellTotal

	chain.FundAccountWithOptions(ctx, t, acc1, integration.BalancesOptions{
		Amount: dexParamsRes.Params.OrderReserve.Amount.Add(buyTotal),
	})

	baseDenomBalanceRes, err := bankClient.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: acc1.String(),
		Denom:   baseDenom,
	})
	requireT.NoError(err)
	requireT.Equal(sdkmath.ZeroInt().String(), baseDenomBalanceRes.Balance.Amount.String())

	placeOrderMsg := &dextypes.MsgPlaceOrder{
		Sender:      acc1.String(),
		Type:        dextypes.ORDER_TYPE_LIMIT,
		ID:          "id1",
		BaseDenom:   baseDenom,
		QuoteDenom:  quoteDenom,
		Price:       lo.ToPtr(dextypes.MustNewPriceFromString("1")),
		Quantity:    buyTotal,
		Side:        dextypes.SIDE_BUY,
		TimeInForce: dextypes.TIME_IN_FORCE_GTC,
	}

	_, requiredGas, err := client.CalculateGas(
		ctx,
		chain.ClientContext.WithFromAddress(acc1),
		chain.TxFactory(),
		placeOrderMsg,
	)
	requireT.NoError(err)
	t.Logf("Estimated gas: %d", requiredGas)
	usedGas := uint64(float64(requiredGas) * 1.05)
	t.Logf("Used gas: %d", usedGas)

	chain.FundAccountWithOptions(ctx, t, acc1, integration.BalancesOptions{
		NondeterministicMessagesGas: usedGas,
	})

	t.Logf("Placing order to match %d orders", len(creatorOrdersRes.Orders))
	now := time.Now()
	_, err = client.BroadcastTx(
		ctx,
		chain.ClientContext.WithFromAddress(acc1),
		chain.TxFactory().WithGas(usedGas),
		placeOrderMsg,
	)
	requireT.NoError(err)
	t.Logf("Placement passed, spent %s", time.Since(now))

	// check final state
	creatorOrdersRes, err = dexClient.Orders(ctx, &dextypes.QueryOrdersRequest{
		Creator: preCreatedOrdersCreator,
		Pagination: &query.PageRequest{
			Limit:      query.PaginationMaxLimit,
			CountTotal: true,
		},
	})
	requireT.NoError(err)
	requireT.Empty(creatorOrdersRes.Orders)

	baseDenomBalanceRes, err = bankClient.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: acc1.String(),
		Denom:   baseDenom,
	})
	requireT.NoError(err)
	requireT.Equal(buyTotal.String(), baseDenomBalanceRes.Balance.Amount.String())
}
