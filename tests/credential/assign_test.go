package credential

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

type testAssign struct {
	suite.Suite
	credential.TestAssignProcessor
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

func (t *testAssign) SetupTest() {
	opr := credential.NewTestAssignProcessor(util.Encoders)
	t.TestAssignProcessor = opr
	t.Setup()
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

func (t *testAssign) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.holderKey, 1000, t.GenesisCurrency, t.holder, true).
		SetTemplate(
			"templateID",
			"id",
			"value",
			1000,
			2000,
			"did",
		).
		MakeItem(t.contract[0], t.holder[0], t.GenesisCurrency, t.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func TestAssign(t *testing.T) {
	suite.Run(t, new(testAssign))
}
