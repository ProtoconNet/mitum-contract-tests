package currency

import (
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	"testing"

	"github.com/ProtoconNet/mitum-credential/operation/credential"
	"github.com/ProtoconNet/mitum-credential/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testAddTemplate struct {
	suite.Suite
	credential.TestAddTemplateProcessor
	senderKey   string // Private Key
	contractKey string // Private Key
	creatorKey  string // Private Key
	owner       []test.Account
}

func (t *testAddTemplate) SetupTest() {
	opr := credential.NewTestAddTemplateProcessor()
	t.TestAddTemplateProcessor = opr
	t.Setup()
	t.owner = make([]test.Account, 1)
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.creatorKey = t.NewPrivateKey("creator")

}

func (t *testAddTemplate) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.Sender(), false).
		SetContractAccount(t.Sender()[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.Contract(), true).
		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.Creator(), true).
		SetTemplate(
			"templateID",
			"templateName",
			types.Date("2024-01-01"),
			types.Date("2024-01-01"),
			types.Bool(true),
			types.Bool(true),
			"displayName",
			"subjectKey",
			"description",
		).
		MakeOperation().
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

//
//func (t *testAddTemplate) Test02ErrorSenderIsContract() {
//	err := t.Create().
//		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.owner, true).
//		SetContractAccount(t.owner[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.Sender(), true).
//		SetAccount(t.targetKey, 100, t.GenesisCurrency, t.Target(), false).
//		SetAmount(100, t.GenesisCurrency).RunPreProcess()
//
//	if assert.NotNil(t.Suite.T(), err) {
//		t.Suite.T().Log(err.Error())
//	}
//}
//
//func (t *testAddTemplate) Test03ErrorTargetExist() {
//	err := t.Create().
//		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.Sender(), true).
//		SetAccount(t.targetKey, 0, t.GenesisCurrency, t.Target(), true).
//		SetAmount(100, t.GenesisCurrency).RunPreProcess()
//
//	if assert.NotNil(t.Suite.T(), err) {
//		t.Suite.T().Log(err.Error())
//	}
//}
//
//func (t *testAddTemplate) Test04ErrorCurrencyNotFound() {
//	err := t.Create().
//		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.Sender(), true).
//		SetAccount(t.targetKey, 0, t.GenesisCurrency, t.Target(), false).
//		SetAmount(100, types.CurrencyID("FOO")).RunPreProcess()
//
//	if assert.NotNil(t.Suite.T(), err) {
//		t.Suite.T().Log(err.Error())
//	}
//}

func TestAddTemplate(t *testing.T) {
	suite.Run(t, new(testAddTemplate))
}
