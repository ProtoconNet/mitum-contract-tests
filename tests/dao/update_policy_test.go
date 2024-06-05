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

type testUpdatePolicy struct {
	suite.Suite
	dao.TestUpdatePolicyProcessor
	sender       []test.Account
	contract     []test.Account
	whitelist    []test.Account
	blockMap     []base.BlockMap
	proposal     []daotypes.Proposal
	currency     []currencytypes.CurrencyID
	ownerKey     string // Private Key
	senderKey    string // Private Key
	contractKey  string // Private Key
	whitelistKey string // Private Key
	owner        []test.Account
}

func (t *testUpdatePolicy) SetupTest() {
	tp := test.TestProcessor{Encoders: util.Encoders}
	opr := dao.NewTestUpdatePolicyProcessor(&tp)
	t.TestUpdatePolicyProcessor = opr
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.whitelist = make([]test.Account, 1)
	t.blockMap = make([]base.BlockMap, 1)
	t.proposal = make([]daotypes.Proposal, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.whitelistKey = t.NewPrivateKey("whitelist")
}

func (t *testUpdatePolicy) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetCurrency("ABC", 100000, t.sender[0].Address(), t.currency, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.whitelistKey, 1000, t.currency[0], t.whitelist, true).
		SetBlockMap(1622547800, t.blockMap).
		SetWhitelist(t.whitelist, true).
		SetDAO("biz", t.currency[0], common.NewBig(10000),
			currencytypes.NewAmount(common.NewBig(10), t.GenesisCurrency), 10000, 10000,
			10000, 10000, 10000, 10000, 3, 3).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(),
			t.contract[0].Address(), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func TestUpdatePolicy(t *testing.T) {
	suite.Run(t, new(testUpdatePolicy))
}
