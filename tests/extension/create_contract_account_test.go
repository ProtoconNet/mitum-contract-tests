package extensiontest

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/extension"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	"testing"

	"github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testCreateContractAccount struct {
	suite.Suite
	extension.TestCreateContractAccountProcessor
	sender      []test.Account
	target      []test.Account
	amounts     []types.Amount
	currency    []types.CurrencyID
	senderKey   string // Private Key
	targetKey   string // Private Key
	contractKey string // Private Key
	owner       []test.Account
}

func (t *testCreateContractAccount) SetupTest() {
	opr := extension.NewTestCreateContractAccountProcessor(util.Encoders)
	t.TestCreateContractAccountProcessor = opr
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.target = make([]test.Account, 1)
	t.amounts = make([]types.Amount, 1)
	t.currency = make([]types.CurrencyID, 1)
	t.senderKey = t.NewPrivateKey("sender")
	t.targetKey = t.NewPrivateKey("target")
	t.contractKey = t.NewPrivateKey("contract")
}

func (t *testCreateContractAccount) Test01ErrorSenderNotFound() {
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

func (t *testCreateContractAccount) Test02ErrorSenderIscontract() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.owner, true).
		SetContractAccount(t.owner[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.targetKey, 0, t.GenesisCurrency, t.target, false).
		SetAmount(100, t.GenesisCurrency, t.amounts).
		MakeItem(t.target[0], t.amounts, t.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testCreateContractAccount) Test03ErrorTargetExist() {
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

func (t *testCreateContractAccount) Test04ErrorCurrencyNotFound() {
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

func TestCreateContractAccount(t *testing.T) {
	suite.Run(t, new(testCreateContractAccount))
}
