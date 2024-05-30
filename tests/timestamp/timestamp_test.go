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
	ap          timestamp.TestAppendProcessor
	cs          timestamp.TestCreateServiceProcessor
	sender      []test.Account
	sender2     []test.Account
	contract    []test.Account
	creator     []test.Account
	currency    []currencytypes.CurrencyID
	ownerKey    string // Private Key
	senderKey   string // Private Key
	sender2Key  string // Private Key
	contractKey string // Private Key
	creatorKey  string // Private Key
	owner       []test.Account
}

func (t *testTimestamp) SetupTest() {
	aopr := timestamp.NewTestAppendProcessor(util.Encoders)
	t.ap = aopr
	copr := timestamp.NewTestCreateServiceProcessor(util.Encoders)
	t.cs = copr
	mockGetter := test.NewMockStateGetter()
	t.ap.Setup(mockGetter)
	t.cs.Setup(mockGetter)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.sender2 = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.creator = make([]test.Account, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.ap.NewPrivateKey("owner")
	t.senderKey = t.ap.NewPrivateKey("sender")
	t.sender2Key = t.ap.NewPrivateKey("sender2")
	t.contractKey = t.ap.NewPrivateKey("contract")
	t.creatorKey = t.ap.NewPrivateKey("creator")
}

func (t *testTimestamp) CreateService() {
	t.cs.Create().
		SetAccount(t.senderKey, 1000, t.cs.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.cs.GenesisCurrency, t.contract, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.cs.GenesisCurrency).
		RunPreProcess().RunProcess()
}

func (t *testTimestamp) Append() {
	t.ap.Create().
		SetAccount(t.senderKey, 1000, t.ap.GenesisCurrency, t.sender, true).
		SetAccount(t.creatorKey, 1000, t.ap.GenesisCurrency, t.creator, true).
		SetService(t.contract[0].Address()).
		MakeOperation(
			t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(),
			"projectId", 1000, "data", t.ap.GenesisCurrency).
		RunPreProcess().RunProcess()
}

func (t *testTimestamp) Test01CreateServiceSenderNotFound() {
	err := t.cs.Create().
		SetAccount(t.senderKey, 1000, t.cs.GenesisCurrency, t.sender, false).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.cs.GenesisCurrency, t.contract, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.cs.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testTimestamp) Test02CreateServiceServiceAlreadyExist() {
	t.CreateService()

	err := t.cs.Create().
		SetAccount(t.senderKey, 1000, t.cs.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.cs.GenesisCurrency, t.contract, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.cs.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testTimestamp) Test03AppendSenderNotFound() {
	err := t.ap.Create().
		SetAccount(t.senderKey, 1000, t.ap.GenesisCurrency, t.sender, true).
		SetAccount(t.sender2Key, 1000, t.ap.GenesisCurrency, t.sender2, false).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.ap.GenesisCurrency, t.contract, true).
		SetAccount(t.creatorKey, 1000, t.ap.GenesisCurrency, t.creator, true).
		SetService(t.contract[0].Address()).
		MakeOperation(
			t.sender2[0].Address(), t.sender2[0].Priv(), t.contract[0].Address(),
			"projectId", 1000, "data", t.ap.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testTimestamp) Test04AppendSenderNotAthorized() {
	t.CreateService()

	err := t.ap.Create().
		SetAccount(t.sender2Key, 1000, t.ap.GenesisCurrency, t.sender2, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.ap.GenesisCurrency, t.contract, true).
		SetAccount(t.creatorKey, 1000, t.ap.GenesisCurrency, t.creator, true).
		SetService(t.contract[0].Address()).
		MakeOperation(
			t.sender2[0].Address(), t.sender2[0].Priv(), t.contract[0].Address(),
			"projectId", 1000, "data", t.ap.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testTimestamp) Test04AppendSenderIsContract() {
	err := t.ap.Create().
		SetAccount(t.ownerKey, 1000, t.ap.GenesisCurrency, t.owner, true).
		SetContractAccount(t.owner[0].Address(), t.senderKey, 1000, t.ap.GenesisCurrency, t.sender, true).
		SetContractAccount(t.owner[0].Address(), t.contractKey, 1000, t.ap.GenesisCurrency, t.contract, true).
		SetAccount(t.creatorKey, 1000, t.ap.GenesisCurrency, t.creator, true).
		SetService(t.contract[0].Address()).
		MakeOperation(
			t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(),
			"projectId", 1000, "data", t.ap.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func (t *testTimestamp) Test05AppendServiceNotExist() {
	t.CreateService()

	err := t.ap.Create().
		SetAccount(t.senderKey, 1000, t.ap.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.ap.GenesisCurrency, t.contract, true).
		SetAccount(t.creatorKey, 1000, t.ap.GenesisCurrency, t.creator, true).
		MakeOperation(
			t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(),
			"projectId", 1000, "data", t.ap.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func TestTimestamp(t *testing.T) {
	suite.Run(t, new(testTimestamp))
}
