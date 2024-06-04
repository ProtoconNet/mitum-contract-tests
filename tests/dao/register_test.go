package daotest

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum2/base"
	"testing"

	"github.com/ProtoconNet/mitum-dao/operation/dao"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testRegister struct {
	suite.Suite
	dao.TestRegisterProcessor
	sender       []test.Account
	contract     []test.Account
	delegated    []test.Account
	blockMap     []base.BlockMap
	currency     []currencytypes.CurrencyID
	ownerKey     string // Private Key
	senderKey    string // Private Key
	contractKey  string // Private Key
	delegatedKey string // Private Key
	owner        []test.Account
}

func (t *testRegister) SetupTest() {
	tp := test.TestProcessor{Encoders: util.Encoders}
	opr := dao.NewTestRegisterProcessor(&tp)
	t.TestRegisterProcessor = opr
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.delegated = make([]test.Account, 1)
	t.blockMap = make([]base.BlockMap, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.delegatedKey = t.NewPrivateKey("delegated")
}

func (t *testRegister) Test01ErrorSenderNotFound() {
	err := t.Create(t.blockMap).
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.delegatedKey, 1000, t.GenesisCurrency, t.delegated, true).
		SetBlockMap(1622547800, t.blockMap).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(),
			t.contract[0].Address(), "proposalID", t.delegated[0].Address(), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func TestRegister(t *testing.T) {
	suite.Run(t, new(testRegister))
}
