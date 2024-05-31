package daotest

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum-dao/operation/dao"
	daotypes "github.com/ProtoconNet/mitum-dao/types"
	"github.com/ProtoconNet/mitum2/base"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testPropose struct {
	suite.Suite
	dao.TestProposeProcessor
	sender      []test.Account
	contract    []test.Account
	creator     []test.Account
	blockMap    []base.BlockMap
	proposal    []daotypes.Proposal
	currency    []currencytypes.CurrencyID
	ownerKey    string // Private Key
	senderKey   string // Private Key
	contractKey string // Private Key
	creatorKey  string // Private Key
	owner       []test.Account
}

func (t *testPropose) SetupTest() {
	opr := dao.NewTestProposeProcessor(util.Encoders)
	t.TestProposeProcessor = opr
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.creator = make([]test.Account, 1)
	t.blockMap = make([]base.BlockMap, 1)
	t.proposal = make([]daotypes.Proposal, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.creatorKey = t.NewPrivateKey("creator")
}

func (t *testPropose) Test01ErrorSenderNotFound() {
	err := t.Create(t.blockMap).
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.creatorKey, 1000, t.GenesisCurrency, t.creator, true).
		SetBlockMap(1622547800, t.blockMap).
		SetProposal(t.sender[0].Address(), 1622547800, "example.com", "hash", 4, t.proposal).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(),
			t.contract[0].Address(), "proposalID", t.proposal[0], t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func TestPropose(t *testing.T) {
	suite.Run(t, new(testPropose))
}
