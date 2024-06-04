package timestamptest

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum-timestamp/operation/timestamp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testTimestamp struct {
	suite.Suite
	*test.TestProcessor
	append        timestamp.TestAppendProcessor
	createService timestamp.TestCreateServiceProcessor
	sender        []test.Account
	sender2       []test.Account
	contract      []test.Account
	creator       []test.Account
	currency      []currencytypes.CurrencyID
	ownerKey      string // Private Key
	senderKey     string // Private Key
	sender2Key    string // Private Key
	contractKey   string // Private Key
	creatorKey    string // Private Key
	owner         []test.Account
}

func (t *testTimestamp) SetupTest() {
	tp := test.TestProcessor{Encoders: util.Encoders}
	t.TestProcessor = &tp
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.append = timestamp.NewTestAppendProcessor(&tp)
	t.createService = timestamp.NewTestCreateServiceProcessor(&tp)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.sender2 = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.creator = make([]test.Account, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.sender2Key = t.NewPrivateKey("sender2")
	t.contractKey = t.NewPrivateKey("contract")
	t.creatorKey = t.NewPrivateKey("creator")
}

func (t *testTimestamp) CreateService() {
	t.createService.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.GenesisCurrency).
		RunPreProcess().RunProcess()
}

func (t *testTimestamp) Append() {
	t.append.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
		SetService(t.contract[0].Address()).
		MakeOperation(
			t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(),
			"projectId", 1000, "data", t.GenesisCurrency).
		RunPreProcess().RunProcess()
}

func (t *testTimestamp) Test01CreateServiceSenderNotFound() {
	err := t.createService.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testTimestamp) Test02CreateServiceServiceAlreadyExist() {
	t.CreateService()

	err := t.createService.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testTimestamp) Test03AppendSenderNotFound() {
	err := t.append.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.sender2Key, 1000, t.GenesisCurrency, t.sender2, false).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
		SetService(t.contract[0].Address()).
		MakeOperation(
			t.sender2[0].Address(), t.sender2[0].Priv(), t.contract[0].Address(),
			"projectId", 1000, "data", t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testTimestamp) Test04AppendSenderNotAthorized() {
	t.CreateService()

	err := t.append.Create().
		SetAccount(t.sender2Key, 1000, t.GenesisCurrency, t.sender2, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
		SetService(t.contract[0].Address()).
		MakeOperation(
			t.sender2[0].Address(), t.sender2[0].Priv(), t.contract[0].Address(),
			"projectId", 1000, "data", t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testTimestamp) Test04AppendSenderIsContract() {
	err := t.append.Create().
		SetAccount(t.ownerKey, 1000, t.GenesisCurrency, t.owner, true).
		SetContractAccount(t.owner[0].Address(), t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.owner[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
		SetService(t.contract[0].Address()).
		MakeOperation(
			t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(),
			"projectId", 1000, "data", t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testTimestamp) Test05AppendServiceNotExist() {
	t.CreateService()

	err := t.append.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
		MakeOperation(
			t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(),
			"projectId", 1000, "data", t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func TestTimestamp(t *testing.T) {
	suite.Run(t, new(testTimestamp))
}
