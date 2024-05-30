package nfttest

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum-nft/v2/operation/nft"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testUpdateCollectionPolicy struct {
	suite.Suite
	nft.TestUpdateCollectionPolicyProcessor
	sender       []test.Account
	contract     []test.Account
	whitelist    []test.Account
	currency     []currencytypes.CurrencyID
	ownerKey     string // Private Key
	senderKey    string // Private Key
	contractKey  string // Private Key
	whitelistKey string
	owner        []test.Account
}

func (t *testUpdateCollectionPolicy) SetupTest() {
	opr := nft.NewTestUpdateCollectionPolicyProcessor(util.Encoders)
	t.TestUpdateCollectionPolicyProcessor = opr
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.whitelist = make([]test.Account, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.whitelistKey = t.NewPrivateKey("whitelist")
}

func (t *testUpdateCollectionPolicy) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetAccount(t.whitelistKey, 1000, t.GenesisCurrency, t.whitelist, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetDesign("abd collection", 10, "example.com").
		SetService(t.sender[0].Address(), t.contract[0].Address(), t.whitelist).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.whitelist, t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testUpdateCollectionPolicy) Test02ErrorCollectionNotExist() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.whitelistKey, 1000, t.GenesisCurrency, t.whitelist, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetDesign("abd collection", 10, "example.com").
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.whitelist, t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func TestUpdateCollectionPolicy(t *testing.T) {
	suite.Run(t, new(testUpdateCollectionPolicy))
}
