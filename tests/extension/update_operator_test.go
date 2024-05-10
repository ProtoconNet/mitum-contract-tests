package extension

import (
	"github.com/ProtoconNet/mitum-currency/v3/operation/extension"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
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
	senderKey    string // Private Key
	ownerKey     string // Private Key
	operatorKey  string // Private Key
	contract1Key string // Private Key
	contract2Key string // Private Key
	owner        []test.Account
}

func (t *testUpdateOperator) SetupTest() {
	opr := extension.NewTestUpdateOperatorProcessor()
	t.TestUpdateOperatorProcessor = opr
	t.Setup()
	t.owner = make([]test.Account, 1)
	t.senderKey = t.NewPrivateKey("sender")
	t.ownerKey = t.NewPrivateKey("owner")
	t.operatorKey = t.NewPrivateKey("operator")
	t.contract1Key = t.NewPrivateKey("contract1")
	t.contract2Key = t.NewPrivateKey("contract2")
}

func (t *testUpdateOperator) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.Sender(), false).
		SetContractAccount(t.Sender()[0].Address(), t.contract1Key, 1000, t.GenesisCurrency, t.Contract(), true).
		SetAccount(t.operatorKey, 1000, t.GenesisCurrency, t.Operators(), true).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testUpdateOperator) Test02ErrorSenderIsContract() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.owner, true).
		SetContractAccount(t.owner[0].Address(), t.contract1Key, 1000, t.GenesisCurrency, t.Sender(), true).
		SetContractAccount(t.Sender()[0].Address(), t.contract2Key, 1000, t.GenesisCurrency, t.Contract(), true).
		SetAccount(t.operatorKey, 1000, t.GenesisCurrency, t.Operators(), true).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testUpdateOperator) Test03ErrorOperatorNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.Sender(), true).
		SetContractAccount(t.Sender()[0].Address(), t.contract1Key, 1000, t.GenesisCurrency, t.Contract(), true).
		SetAccount(t.operatorKey, 1000, t.GenesisCurrency, t.Operators(), false).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testUpdateOperator) Test04ErrorOperatorIsContract() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.Sender(), true).
		SetContractAccount(t.Sender()[0].Address(), t.contract1Key, 1000, t.GenesisCurrency, t.Contract(), true).
		SetContractAccount(t.Sender()[0].Address(), t.contract2Key, 1000, t.GenesisCurrency, t.Operators(), true).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func TestUpdateOperator(t *testing.T) {
	suite.Run(t, new(testUpdateOperator))
}
