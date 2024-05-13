package currency

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
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

type testMint struct {
	suite.Suite
	currency.TestMintProcessor
	receiverKey string // Private Key
	contractKey string // Private Key
	owner       []test.Account
}

func (t *testMint) SetupTest() {
	opr := currency.NewTestMintProcessor(util.Encoders)
	t.TestMintProcessor = opr
	t.Setup()
	t.owner = make([]test.Account, 1)
	t.receiverKey = t.NewPrivateKey("receiver")
	t.contractKey = t.NewPrivateKey("contract")
}

func (t *testMint) Test01ErrorCurrencyNotFound() {
	err := t.Create().
		SetCurrency("ABC", 10000, t.GenesisAddr, false).
		SetAccount(t.receiverKey, 100, t.GenesisCurrency, t.Receiver(), true).
		SetAmount(100, t.Currency()).
		MakeOperation().Print("mint-test.json").
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testMint) Test02ErrorReceiverNotExist() {
	err := t.Create().
		SetCurrency("ABC", 10000, t.GenesisAddr, true).
		SetAccount(t.receiverKey, 100, t.GenesisCurrency, t.Receiver(), false).
		SetAmount(100, t.Currency()).
		MakeOperation().
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testMint) Test03ErrorReceiverIsContract() {
	err := t.Create().
		SetCurrency("ABC", 10000, t.GenesisAddr, true).
		SetAccount(t.receiverKey, 1000, t.GenesisCurrency, t.owner, true).
		SetContractAccount(t.owner[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.Receiver(), true).
		SetAmount(100, t.Currency()).
		MakeOperation().
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func TestMint(t *testing.T) {
	suite.Run(t, new(testMint))
}
