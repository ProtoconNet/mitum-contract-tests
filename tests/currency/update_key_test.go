package currency

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/currency"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	"github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testUpdateKey struct {
	suite.Suite
	currency.TestUpdateKeyProcessor
	sender      []test.Account
	target      []test.Account
	amounts     []types.Amount
	senderKey   string // Private Key
	targetKey   string // Private Key
	contractKey string // Private Key
	owner       []test.Account
}

func (t *testUpdateKey) SetupTest() {
	opr := currency.NewTestUpdateKeyProcessor(util.Encoders)
	t.TestUpdateKeyProcessor = opr
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.target = make([]test.Account, 1)
	t.amounts = make([]types.Amount, 1)
	t.senderKey = t.NewPrivateKey("sender")
	t.targetKey = t.NewPrivateKey("target")
	t.contractKey = t.NewPrivateKey("contract")
}

func (t *testUpdateKey) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetAccount(t.targetKey, 100, t.GenesisCurrency, t.target, false).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.target[0].Keys(), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testUpdateKey) Test02ErrorSenderIscontract() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.owner, true).
		SetContractAccount(t.owner[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.targetKey, 100, t.GenesisCurrency, t.target, false).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.target[0].Keys(), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func TestUpdateKey(t *testing.T) {
	suite.Run(t, new(testUpdateKey))
}
