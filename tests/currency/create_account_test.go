package currencytest

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
	sender      []test.Account
	target      []test.Account
	owner       []test.Account
	amounts     []types.Amount
}

func (t *testCreateAccount) SetupTest() {
	tp := test.TestProcessor{Encoders: util.Encoders}
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	opr := currency.NewTestCreateAccountProcessor(&tp)
	t.TestCreateAccountProcessor = opr
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.target = make([]test.Account, 1)
	t.amounts = make([]types.Amount, 1)
	t.senderKey = t.NewPrivateKey("sender")
	t.targetKey = t.NewPrivateKey("target")
	t.contractKey = t.NewPrivateKey("contract")
}

func (t *testCreateAccount) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetAccount(t.targetKey, 0, t.GenesisCurrency, t.target, false).
		SetAmount(100, t.GenesisCurrency, t.amounts).
		MakeItem(t.target[0], t.amounts, t.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateAccount) Test02ErrorSenderIscontract() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.owner, true).
		SetContractAccount(t.owner[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.targetKey, 100, t.GenesisCurrency, t.target, false).
		SetAmount(100, t.GenesisCurrency, t.amounts).
		MakeItem(t.target[0], t.amounts, t.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).Print("test.json").
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateAccount) Test03ErrorTargetExist() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.targetKey, 0, t.GenesisCurrency, t.target, true).
		SetAmount(100, t.GenesisCurrency, t.amounts).
		MakeItem(t.target[0], t.amounts, t.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateAccount) Test04ErrorCurrencyNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.targetKey, 0, t.GenesisCurrency, t.target, false).
		SetAmount(100, types.CurrencyID("FOO"), t.amounts).
		MakeItem(t.target[0], t.amounts, t.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateAccount) Test05ErrorIsValid() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.targetKey, 0, t.GenesisCurrency, t.target, false).
		SetAmount(-100, types.CurrencyID("FOO"), t.amounts).
		MakeItem(t.target[0], t.amounts, t.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
		IsValid()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateAccount) Test06ErrorLoadJson() {
	err := t.Create().
		LoadOperation("create-account.json").
		IsValid()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateAccount) Test06ErrorDecodeJson() {
	err := t.Create().
		Decode("create-account.json")

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func TestCreateAccount(t *testing.T) {
	suite.Run(t, new(testCreateAccount))
}
