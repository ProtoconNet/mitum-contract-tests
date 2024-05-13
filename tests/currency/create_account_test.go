package currency

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	"testing"

	"github.com/ProtoconNet/mitum-currency/v3/operation/currency"
	"github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testCreateAccount struct {
	suite.Suite
	currency.TestCreateAccountProcessor
	senderKey   string // Private Key
	targetKey   string // Private Key
	contractKey string // Private Key
	owner       []test.Account
}

func (t *testCreateAccount) SetupTest() {
	opr := currency.NewTestCreateAccountProcessor(util.Encoders)
	t.TestCreateAccountProcessor = opr
	t.Setup()
	t.owner = make([]test.Account, 1)
	t.senderKey = t.NewPrivateKey("sender")
	t.targetKey = t.NewPrivateKey("target")
	t.contractKey = t.NewPrivateKey("contract")
}

func (t *testCreateAccount) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.Sender(), false).
		SetAccount(t.targetKey, 0, t.GenesisCurrency, t.Target(), false).
		SetAmount(100, t.GenesisCurrency).
		MakeOperation().
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateAccount) Test02ErrorSenderIsContract() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.owner, true).
		SetContractAccount(t.owner[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.Sender(), true).
		SetAccount(t.targetKey, 100, t.GenesisCurrency, t.Target(), false).
		SetAmount(100, t.GenesisCurrency).
		MakeOperation().Print("test.json").
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateAccount) Test03ErrorTargetExist() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.Sender(), true).
		SetAccount(t.targetKey, 0, t.GenesisCurrency, t.Target(), true).
		SetAmount(100, t.GenesisCurrency).
		MakeOperation().
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateAccount) Test04ErrorCurrencyNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.Sender(), true).
		SetAccount(t.targetKey, 0, t.GenesisCurrency, t.Target(), false).
		SetAmount(100, types.CurrencyID("FOO")).
		MakeOperation().
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateAccount) Test05ErrorIsValid() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.Sender(), true).
		SetAccount(t.targetKey, 0, t.GenesisCurrency, t.Target(), false).
		SetAmount(-100, types.CurrencyID("FOO")).
		MakeOperation().
		IsValid()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateAccount) Test06ErrorLoadJson() {
	err := t.Create().
		LoadOperation("create-account.json").
		IsValid()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func TestCreateAccount(t *testing.T) {
	suite.Run(t, new(testCreateAccount))
}
