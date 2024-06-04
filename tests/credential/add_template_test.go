package credentialtest

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
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
	sender      []test.Account
	contract    []test.Account
	creator     []test.Account
	currency    []currencytypes.CurrencyID
	ownerKey    string // Private Key
	senderKey   string // Private Key
	contractKey string // Private Key
	creatorKey  string // Private Key
	owner       []test.Account
}

func (t *testAddTemplate) SetupTest() {
	tp := test.TestProcessor{Encoders: util.Encoders}
	opr := credential.NewTestAddTemplateProcessor(&tp)
	t.TestAddTemplateProcessor = opr
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.creator = make([]test.Account, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.creatorKey = t.NewPrivateKey("creator")
}

func (t *testAddTemplate) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
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
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.creator[0].Address(), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testAddTemplate) Test02ErrorSenderIscontract() {
	err := t.Create().
		SetAccount(t.ownerKey, 1000, t.GenesisCurrency, t.owner, true).
		SetContractAccount(t.owner[0].Address(), t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.owner[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
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
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.creator[0].Address(), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testAddTemplate) Test03ErrorServiceNotExist() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
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
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.creator[0].Address(), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func TestAddTemplate(t *testing.T) {
	suite.Run(t, new(testAddTemplate))
}
