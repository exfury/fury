package keeper_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/exfury/fury/v15/app"
	utiltx "github.com/exfury/fury/v15/testutil/tx"
	"github.com/exfury/fury/v15/utils"
	evm "github.com/exfury/fury/v15/x/evm/types"
	feemarkettypes "github.com/exfury/fury/v15/x/feemarket/types"
	"github.com/exfury/fury/v15/x/revenue/v1/types"

	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx sdk.Context

	app            *app.Fury
	queryClient    types.QueryClient
	queryClientEvm evm.QueryClient
	address        common.Address
	signer         keyring.Signer
	ethSigner      ethtypes.Signer
	consAddress    sdk.ConsAddress
	validator      stakingtypes.Validator
	denom          string
}

var s *KeeperTestSuite

var (
	contract = utiltx.GenerateAddress()
	deployer = sdk.AccAddress(utiltx.GenerateAddress().Bytes())
	withdraw = sdk.AccAddress(utiltx.GenerateAddress().Bytes())
)

func TestKeeperTestSuite(t *testing.T) {
	s = new(KeeperTestSuite)
	suite.Run(t, s)

	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keeper Suite")
}

func (suite *KeeperTestSuite) SetupTest() {
	chainID := utils.TestnetChainID + "-1"
	suite.app = app.Setup(false, feemarkettypes.DefaultGenesisState(), chainID)
	suite.SetupApp(chainID)
}
