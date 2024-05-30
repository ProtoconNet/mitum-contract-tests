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

type testTransfer struct {
	suite.Suite
	currency.TestTransferProcessor
	sender      []test.Account
	receiver    []test.Account
	amounts     []types.Amount
	currency    []types.CurrencyID
	senderKey   string // Private Key
	receiverKey string // Private Key
	contractKey string // Private Key
	owner       []test.Account
}

func (t *testTransfer) SetupTest() {
	opr := currency.NewTestTransferProcessor(util.Encoders)
	t.TestTransferProcessor = opr
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.receiver = make([]test.Account, 1)
	t.amounts = make([]types.Amount, 1)
	t.currency = make([]types.CurrencyID, 1)
	t.senderKey = t.NewPrivateKey("sender")
	t.receiverKey = t.NewPrivateKey("receiver")
	t.contractKey = t.NewPrivateKey("contract")
}

func (t *testTransfer) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetAccount(t.receiverKey, 100, t.GenesisCurrency, t.receiver, true).
		SetAmount(100, t.GenesisCurrency, t.amounts).
		MakeItem(t.receiver[0], t.amounts, t.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testTransfer) Test02ErrorSenderIscontract() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.owner, true).
		SetContractAccount(t.owner[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.receiverKey, 100, t.GenesisCurrency, t.receiver, true).
		SetAmount(100, t.GenesisCurrency, t.amounts).
		MakeItem(t.receiver[0], t.amounts, t.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testTransfer) Test03ErrorCurrencyNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.receiverKey, 100, t.GenesisCurrency, t.receiver, true).
		SetAmount(100, types.CurrencyID("FOO"), t.amounts).
		MakeItem(t.receiver[0], t.amounts, t.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func TestTransfer(t *testing.T) {
	suite.Run(t, new(testTransfer))
}
