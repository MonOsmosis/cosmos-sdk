package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/MonOsmosis/cosmos-sdk/codec"
	codectypes "github.com/MonOsmosis/cosmos-sdk/codec/types"
	"github.com/MonOsmosis/cosmos-sdk/std"
	"github.com/MonOsmosis/cosmos-sdk/testutil/testdata"
	sdk "github.com/MonOsmosis/cosmos-sdk/types"
	"github.com/MonOsmosis/cosmos-sdk/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
