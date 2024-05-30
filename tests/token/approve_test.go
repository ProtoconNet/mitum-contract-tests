package tokentest

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/common"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"testing"

	mtoken "github.com/ProtoconNet/mitum-token/operation/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testApprove struct {
	suite.Suite
	mtoken.TestApproveProcessor
	sender      []test.Account
	contract    []test.Account
	approved    []test.Account
	currency    []currencytypes.CurrencyID
	ownerKey    string // Private Key
	senderKey   string // Private Key
	contractKey string // Private Key
	approvedKey string // Private Key
	owner       []test.Account
}

func (t *testApprove) SetupTest() {
	opr := mtoken.NewTestApproveProcessor(util.Encoders)
	t.TestApproveProcessor = opr
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.approved = make([]test.Account, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.approvedKey = t.NewPrivateKey("approved")
}

func (t *testApprove) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.approvedKey, 1000, t.GenesisCurrency, t.approved, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.approved[0].Address(), common.NewBig(1000), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

//func (t *testApprove) Test02ErrorSenderIscontract() {
//	err := t.Create().
//		SetAccount(t.ownerKey, 1000, t.GenesisCurrency, t.owner, true).
//		SetContractAccount(t.owner[0].Address(), t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
//		SetContractAccount(t.owner[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
//		SetAccount(t.approvedKey, 1000, t.GenesisCurrency, t.approved, true).
//		SetTemplate(
//			"templateID",
//			"templateName",
//			types.Date("2024-01-01"),
//			types.Date("2024-01-01"),
//			types.Bool(true),
//			types.Bool(true),
//			"displayName",
//			"subjectKey",
//			"description",
//		).
//		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.approved[0].Address(), t.GenesisCurrency).
//		RunPreProcess()
//
//	if assert.NotNil(t.Suite.T(), err.Error()) {
//		t.Suite.T().Log(err.Error())
//	}
//}
//
//func (t *testApprove) Test03ErrorServiceNotExist() {
//	err := t.Create().
//		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
//		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
//		SetAccount(t.approvedKey, 1000, t.GenesisCurrency, t.approved, true).
//		SetTemplate(
//			"templateID",
//			"templateName",
//			types.Date("2024-01-01"),
//			types.Date("2024-01-01"),
//			types.Bool(true),
//			types.Bool(true),
//			"displayName",
//			"subjectKey",
//			"description",
//		).
//		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.approved[0].Address(), t.GenesisCurrency).
//		RunPreProcess()
//
//	if assert.NotNil(t.Suite.T(), err.Error()) {
//		t.Suite.T().Log(err.Error())
//	}
//}

func TestApprove(t *testing.T) {
	suite.Run(t, new(testApprove))
}
