package credentialtest

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"testing"

	"github.com/ProtoconNet/mitum-credential/operation/credential"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testRevoke struct {
	suite.Suite
	credential.TestRevokeProcessor
	sender      []test.Account
	contract    []test.Account
	holder      []test.Account
	currency    []currencytypes.CurrencyID
	ownerKey    string // Private Key
	senderKey   string // Private Key
	contractKey string // Private Key
	holderKey   string // Private Key
	owner       []test.Account
}

func (t *testRevoke) SetupTest() {
	tp := test.TestProcessor{Encoders: util.Encoders}
	opr := credential.NewTestRevokeProcessor(&tp)
	t.TestRevokeProcessor = opr
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.holder = make([]test.Account, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.holderKey = t.NewPrivateKey("holder")
}

func (t *testRevoke) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.holderKey, 1000, t.GenesisCurrency, t.holder, true).
		SetTemplate(
			"templateID",
			"id",
		).
		MakeItem(t.contract[0], t.holder[0], t.GenesisCurrency, t.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func TestRevoke(t *testing.T) {
	suite.Run(t, new(testRevoke))
}
