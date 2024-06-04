package currencytest

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

type testCurrency struct {
	suite.Suite
	*test.TestProcessor
	CreateAccount currency.TestCreateAccountProcessor
	Mint          currency.TestMintProcessor
	Transfer      currency.TestTransferProcessor
	UpdateKey     currency.TestUpdateKeyProcessor
	senderKey     string // Private Key
	targetKey     string // Private Key
	contractKey   string // Private Key
	updaterKey    string // Private Key
	currency      []types.CurrencyID
	sender        []test.Account
	target        []test.Account
	updater       []test.Account
	amounts       []types.Amount
}

func (t *testCurrency) SetupTest() {
	tp := test.TestProcessor{Encoders: util.Encoders}
	t.TestProcessor = &tp
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.CreateAccount = currency.NewTestCreateAccountProcessor(&tp)
	t.Mint = currency.NewTestMintProcessor(&tp)
	t.Transfer = currency.NewTestTransferProcessor(&tp)
	t.UpdateKey = currency.NewTestUpdateKeyProcessor(&tp)
	t.currency = make([]types.CurrencyID, 1)
	t.updater = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.target = make([]test.Account, 1)
	t.amounts = make([]types.Amount, 1)
	t.senderKey = t.NewPrivateKey("sender")
	t.targetKey = t.NewPrivateKey("target")
	t.contractKey = t.NewPrivateKey("contract")
	t.updaterKey = t.NewPrivateKey("updater")
}

func (t *testCurrency) Test01ErrorSenderNotFound() {
	t.CreateAccount.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.targetKey, 0, t.GenesisCurrency, t.target, false).
		SetAmount(100, t.GenesisCurrency, t.amounts).
		MakeItem(t.target[0], t.amounts, t.CreateAccount.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.CreateAccount.Items()).Print("test2.json").
		RunPreProcess().RunProcess()

	assert.Nil(t.Suite.T(), t.Error())

	t.Transfer.Create().
		SetAmount(100, t.GenesisCurrency, t.amounts).
		MakeItem(t.target[0], t.amounts, t.Transfer.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Transfer.Items()).
		RunPreProcess().
		RunProcess()

	assert.Nil(t.Suite.T(), t.Error())

	t.Mint.Create().
		MakeItem(t.target[0], t.amounts[0], t.Mint.Items()).
		MakeOperation(t.Mint.Items()).
		RunPreProcess().
		RunProcess()

	assert.Nil(t.Suite.T(), t.Error())

	t.Transfer.Create().
		SetAmount(100, t.GenesisCurrency, t.amounts).
		MakeItem(t.sender[0], t.amounts, t.Transfer.Items()).
		MakeOperation(t.target[0].Address(), t.target[0].Priv(), t.Transfer.Items()).
		RunPreProcess().RunProcess()

	t.UpdateKey.Create().
		SetAccount(t.updaterKey, 0, t.GenesisCurrency, t.updater, false).
		MakeOperation(t.target[0].Address(), t.target[0].Priv(), t.updater[0].Keys(), t.GenesisCurrency).
		RunPreProcess().RunProcess()

	t.Transfer.Create().
		SetAmount(100, t.GenesisCurrency, t.amounts).
		MakeItem(t.sender[0], t.amounts, t.Transfer.Items()).
		MakeOperation(t.target[0].Address(), t.updater[0].Priv(), t.Transfer.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), t.Error()) {
		t.Suite.T().Log(t.Error().Error())
	}
}

func TestCurrency(t *testing.T) {
	suite.Run(t, new(testCurrency))
}
