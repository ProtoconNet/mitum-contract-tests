package extension

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

type testUpdateOperator struct {
	suite.Suite
	extension.TestUpdateOperatorProcessor
	owner        []test.Account
	sender       []test.Account
	sender2      []test.Account
	contract     []test.Account
	operators    []test.Account
	currency     []types.CurrencyID
	senderKey    string // Private Key
	sender2Key   string // Private Key
	ownerKey     string // Private Key
	operatorKey  string // Private Key
	contract1Key string // Private Key
	contract2Key string // Private Key
}

func (t *testUpdateOperator) SetupTest() {
	opr := extension.NewTestUpdateOperatorProcessor(util.Encoders)
	t.TestUpdateOperatorProcessor = opr
	t.Setup()
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.sender2 = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.operators = make([]test.Account, 1)
	t.currency = make([]types.CurrencyID, 1)
	t.senderKey = t.NewPrivateKey("sender")
	t.sender2Key = t.NewPrivateKey("sender2")
	t.ownerKey = t.NewPrivateKey("owner")
	t.operatorKey = t.NewPrivateKey("operator")
	t.contract1Key = t.NewPrivateKey("contract1")
	t.contract2Key = t.NewPrivateKey("contract2")
}

func (t *testUpdateOperator) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetContractAccount(t.sender[0].Address(), t.contract1Key, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.operatorKey, 1000, t.GenesisCurrency, t.operators, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.operators, t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testUpdateOperator) Test02ErrorSenderIscontract() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.owner, true).
		SetContractAccount(t.owner[0].Address(), t.contract1Key, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contract2Key, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.operatorKey, 1000, t.GenesisCurrency, t.operators, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.operators, t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testUpdateOperator) Test03ErrorOperatorNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contract1Key, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.operatorKey, 1000, t.GenesisCurrency, t.operators, false).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.operators, t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testUpdateOperator) Test04ErrorOperatorIscontract() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contract1Key, 1000, t.GenesisCurrency, t.contract, true).
		SetContractAccount(t.sender[0].Address(), t.contract2Key, 1000, t.GenesisCurrency, t.operators, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.operators, t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testUpdateOperator) Test05ErrorSender2() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.sender2Key, 1000, t.GenesisCurrency, t.sender2, true).
		SetContractAccount(t.sender2[0].Address(), t.contract1Key, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.operatorKey, 1000, t.GenesisCurrency, t.operators, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.operators, t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func TestUpdateOperator(t *testing.T) {
	suite.Run(t, new(testUpdateOperator))
}
