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

type testMint struct {
	suite.Suite
	currency.TestMintProcessor
	receiver    []test.Account
	receiverKey string // Private Key
	contractKey string // Private Key
	owner       []test.Account
	currency    []types.CurrencyID
	amounts     []types.Amount
}

func (t *testMint) SetupTest() {
	opr := currency.NewTestMintProcessor(util.Encoders)
	t.TestMintProcessor = opr
	t.Setup()
	t.receiver = make([]test.Account, 1)
	t.owner = make([]test.Account, 1)
	t.currency = make([]types.CurrencyID, 1)
	t.amounts = make([]types.Amount, 1)
	t.receiverKey = t.NewPrivateKey("receiver")
	t.contractKey = t.NewPrivateKey("contract")
}

func (t *testMint) Test01ErrorCurrencyNotFound() {
	err := t.Create().
		SetCurrency("ABC", 10000, t.GenesisAddr, t.currency, false).
		SetAccount(t.receiverKey, 100, t.GenesisCurrency, t.receiver, true).
		SetAmount(100, t.currency[0], t.amounts).
		MakeItem(t.receiver[0], t.amounts[0], t.Items()).
		MakeOperation(t.Items()).Print("mint-test.json").
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testMint) Test02ErrorReceiverNotExist() {
	err := t.Create().
		SetCurrency("ABC", 10000, t.GenesisAddr, t.currency, true).
		SetAccount(t.receiverKey, 100, t.GenesisCurrency, t.receiver, false).
		SetAmount(100, t.currency[0], t.amounts).
		MakeItem(t.receiver[0], t.amounts[0], t.Items()).
		MakeOperation(t.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testMint) Test03ErrorReceiverIscontract() {
	err := t.Create().
		SetCurrency("ABC", 10000, t.GenesisAddr, t.currency, true).
		SetAccount(t.receiverKey, 1000, t.GenesisCurrency, t.owner, true).
		SetContractAccount(t.owner[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.receiver, true).
		SetAmount(100, t.currency[0], t.amounts).
		MakeItem(t.receiver[0], t.amounts[0], t.Items()).
		MakeOperation(t.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func TestMint(t *testing.T) {
	suite.Run(t, new(testMint))
}
