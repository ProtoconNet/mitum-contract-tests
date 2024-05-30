package extensiontest

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/extension"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	"github.com/ProtoconNet/mitum-currency/v3/types"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testWithdrawOperator struct {
	suite.Suite
	extension.TestWithdrawProcessor
	sender       []test.Account
	contract     []test.Account
	operators    []test.Account
	amounts      []types.Amount
	currency     []types.CurrencyID
	senderKey    string // Private Key
	ownerKey     string // Private Key
	operatorKey  string // Private Key
	contract1Key string // Private Key
	contract2Key string // Private Key
	owner        []test.Account
}

func (t *testWithdrawOperator) SetupTest() {
	opr := extension.NewTestWithdrawProcessor(util.Encoders)
	t.TestWithdrawProcessor = opr
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.amounts = make([]types.Amount, 1)
	t.operators = make([]test.Account, 1)
	t.currency = make([]types.CurrencyID, 1)
	t.senderKey = t.NewPrivateKey("sender")
	t.ownerKey = t.NewPrivateKey("owner")
	t.operatorKey = t.NewPrivateKey("operator")
	t.contract1Key = t.NewPrivateKey("contract1")
	t.contract2Key = t.NewPrivateKey("contract2")
}

func (t *testWithdrawOperator) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetContractAccount(t.sender[0].Address(), t.contract1Key, 1000, t.GenesisCurrency, t.contract, true).
		SetAmount(100, t.GenesisCurrency, t.amounts).
		MakeItem(t.contract[0], t.amounts, t.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testWithdrawOperator) Test02ErrorSenderIscontract() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.owner, true).
		SetContractAccount(t.owner[0].Address(), t.contract1Key, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contract2Key, 1000, t.GenesisCurrency, t.contract, true).
		SetAmount(100, t.GenesisCurrency, t.amounts).
		MakeItem(t.contract[0], t.amounts, t.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func TestWithdrawOperator(t *testing.T) {
	suite.Run(t, new(testWithdrawOperator))
}
