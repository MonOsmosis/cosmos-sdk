package ante_test

import (
	"github.com/MonOsmosis/cosmos-sdk/codec"
	"github.com/MonOsmosis/cosmos-sdk/codec/types"
	"github.com/MonOsmosis/cosmos-sdk/testutil/testdata"
	sdk "github.com/MonOsmosis/cosmos-sdk/types"
	"github.com/MonOsmosis/cosmos-sdk/x/auth/ante"
	"github.com/MonOsmosis/cosmos-sdk/x/auth/tx"
)

type setFeeGranter interface {
	SetFeeGranter(feeGranter sdk.AccAddress)
}

func (suite *AnteTestSuite) TestRejectFeeGranter() {
	suite.SetupTest(true) // setup
	txConfig := tx.NewTxConfig(codec.NewProtoCodec(types.NewInterfaceRegistry()), tx.DefaultSignModes)
	txBuilder := txConfig.NewTxBuilder()
	d := ante.NewRejectFeeGranterDecorator()
	antehandler := sdk.ChainAnteDecorators(d)

	_, err := antehandler(suite.ctx, txBuilder.GetTx(), false)
	suite.Require().NoError(err)

	setGranterTx := txBuilder.(setFeeGranter)
	_, _, addr := testdata.KeyTestPubAddr()
	setGranterTx.SetFeeGranter(addr)

	_, err = antehandler(suite.ctx, txBuilder.GetTx(), false)
	suite.Require().Error(err)
}
