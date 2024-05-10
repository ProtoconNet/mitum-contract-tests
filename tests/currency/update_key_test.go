package currency

import (
	"github.com/ProtoconNet/mitum-currency/v3/operation/currency"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
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
	senderKey   string // Private Key
	targetKey   string // Private Key
	contractKey string // Private Key
	owner       []test.Account
}

func (t *testUpdateKey) SetupTest() {
	opr := currency.NewTestUpdateKeyProcessor()
	t.TestUpdateKeyProcessor = opr
	t.TestProcessor.Setup()
	t.owner = make([]test.Account, 1)
	t.senderKey = t.NewPrivateKey("sender")
	t.targetKey = t.NewPrivateKey("target")
	t.contractKey = t.NewPrivateKey("contract")
}

func (t *testUpdateKey) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.Sender(), false).
		SetAccount(t.targetKey, 100, t.GenesisCurrency, t.Target(), false).
		MakeOperation().
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testUpdateKey) Test02ErrorSenderIsContract() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.owner, true).
		SetContractAccount(t.owner[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.Sender(), true).
		SetAccount(t.targetKey, 100, t.GenesisCurrency, t.Target(), false).
		MakeOperation().
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func TestUpdateKey(t *testing.T) {
	suite.Run(t, new(testUpdateKey))
}
