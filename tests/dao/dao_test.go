package daotest

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/common"
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

type testDAO struct {
	suite.Suite
	*test.TestProcessor
	createDAO      dao.TestCreateDAOProcessor
	cancelProposal dao.TestCancelProposalProcessor
	execute        dao.TestExecuteProcessor
	postSnap       dao.TestPostSnapProcessor
	preSnap        dao.TestPreSnapProcessor
	propose        dao.TestProposeProcessor
	register       dao.TestRegisterProcessor
	updatePolicy   dao.TestUpdatePolicyProcessor
	vote           dao.TestVoteProcessor
	sender         []test.Account
	contract       []test.Account
	whitelist      []test.Account
	delegated      []test.Account
	blockMap       []base.BlockMap
	proposal       []daotypes.Proposal
	currency       []currencytypes.CurrencyID
	ownerKey       string // Private Key
	senderKey      string // Private Key
	contractKey    string // Private Key
	whitelistKey   string // Private Key
	delegatedKey   string // Private Key
	owner          []test.Account
}

func (t *testDAO) SetupTest() {
	tp := test.TestProcessor{Encoders: util.Encoders}
	t.TestProcessor = &tp
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.createDAO = dao.NewTestCreateDAOProcessor(&tp)
	t.cancelProposal = dao.NewTestCancelProposalProcessor(&tp)
	t.execute = dao.NewTestExecuteProcessor(&tp)
	t.postSnap = dao.NewTestPostSnapProcessor(&tp)
	t.preSnap = dao.NewTestPreSnapProcessor(&tp)
	t.propose = dao.NewTestProposeProcessor(&tp)
	t.register = dao.NewTestRegisterProcessor(&tp)
	t.updatePolicy = dao.NewTestUpdatePolicyProcessor(&tp)
	t.vote = dao.NewTestVoteProcessor(&tp)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.whitelist = make([]test.Account, 1)
	t.delegated = make([]test.Account, 1)
	t.blockMap = make([]base.BlockMap, 1)
	t.proposal = make([]daotypes.Proposal, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.whitelistKey = t.NewPrivateKey("whitelist")
	t.delegatedKey = t.NewPrivateKey("delegated")
}

func (t *testDAO) CreateDAO() {
	t.createDAO.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetCurrency("ABC", 100000, t.sender[0].Address(), t.currency, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.senderKey, 1000, t.currency[0], t.whitelist, true).
		SetWhitelist(t.whitelist, true).
		SetDAO("biz", t.currency[0], common.NewBig(100),
			currencytypes.NewAmount(common.NewBig(10), t.GenesisCurrency), 10000, 10000,
			10000, 10000, 10000, 10000, 3, 3).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(),
			t.contract[0].Address(), t.GenesisCurrency).
		RunPreProcess().
		RunProcess()
}

func (t *testDAO) CancelProposal(blockTime int64) {
	t.cancelProposal.Create(t.blockMap).
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.delegatedKey, 1000, t.GenesisCurrency, t.delegated, true).
		SetBlockMap(blockTime, t.blockMap).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(),
			t.contract[0].Address(), "proposalID", t.GenesisCurrency).
		RunPreProcess().
		RunProcess()
}

func (t *testDAO) Execute(blockTime int64) {
	t.execute.Create(t.blockMap).
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.delegatedKey, 1000, t.GenesisCurrency, t.delegated, true).
		SetBlockMap(blockTime, t.blockMap).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(),
			t.contract[0].Address(), "proposalID", t.GenesisCurrency).
		RunPreProcess().
		RunProcess()
}

func (t *testDAO) PostSnap(blockTime int64) {
	t.postSnap.Create(t.blockMap).
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetBlockMap(blockTime, t.blockMap).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(),
			t.contract[0].Address(), "proposalID", t.GenesisCurrency).
		RunPreProcess().
		RunProcess()
}

func (t *testDAO) PreSnap(blockTime int64) {
	t.preSnap.Create(t.blockMap).
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.delegatedKey, 1000, t.GenesisCurrency, t.delegated, true).
		SetBlockMap(blockTime, t.blockMap).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(),
			t.contract[0].Address(), "proposalID", t.GenesisCurrency).
		RunPreProcess().
		RunProcess()
}

func (t *testDAO) Propose(blockTime int64) {
	t.propose.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetBlockMap(1000, t.blockMap).
		SetProposal(t.sender[0].Address(), uint64(blockTime), "example.com", "hash", 4, t.proposal).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(),
			t.contract[0].Address(), "proposalID", t.proposal[0], t.GenesisCurrency).
		RunPreProcess().
		RunProcess()
}

func (t *testDAO) Register(blockTime int64) {
	t.register.Create(t.blockMap).
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.delegatedKey, 1000, t.GenesisCurrency, t.delegated, true).
		SetBlockMap(blockTime, t.blockMap).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(),
			t.contract[0].Address(), "proposalID", t.delegated[0].Address(), t.GenesisCurrency).
		RunPreProcess().
		RunProcess()
}

func (t *testDAO) UpdatePolicy(blockTime int64) {
	t.updatePolicy.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetCurrency("ABC", 100000, t.sender[0].Address(), t.currency, true).
		SetAccount(t.whitelistKey, 1000, t.currency[0], t.whitelist, true).
		SetBlockMap(blockTime, t.blockMap).
		SetWhitelist(t.whitelist, true).
		SetDAO("biz", t.currency[0], common.NewBig(10000),
			currencytypes.NewAmount(common.NewBig(10), t.GenesisCurrency), 10000, 10000,
			10000, 10000, 10000, 10000, 3, 3).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(),
			t.contract[0].Address(), t.GenesisCurrency).
		RunPreProcess().
		RunProcess()
}

func (t *testDAO) Vote(blockTime int64) {
	t.vote.Create(t.blockMap).
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetBlockMap(blockTime, t.blockMap).
		MakeOperation(t.delegated[0].Address(), t.delegated[0].Priv(),
			t.contract[0].Address(), "proposalID", 0, t.GenesisCurrency).
		RunPreProcess().
		RunProcess()
}

func (t *testDAO) Test01Error() {
	t.CreateDAO()
	t.Propose(10000)
	t.Register(20000)
	t.PreSnap(30000)
	t.Vote(40000)
	t.PostSnap(50000)
	t.Execute(70000)

	t.cancelProposal.Create(t.blockMap).
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetBlockMap(00000, t.blockMap).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(),
			t.contract[0].Address(), "proposalID", t.GenesisCurrency).
		IsValid()

	if assert.NotNil(t.Suite.T(), t.Error()) {
		t.Suite.T().Log(t.Error().Error())
	}
}

func TestDAO(t *testing.T) {
	suite.Run(t, new(testDAO))
}
