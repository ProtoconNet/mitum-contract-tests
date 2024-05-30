package pointtest

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/common"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"testing"

	mpoint "github.com/ProtoconNet/mitum-point/operation/point"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testTransferFrom struct {
	suite.Suite
	mpoint.TestTransferFromProcessor
	sender      []test.Account
	contract    []test.Account
	target      []test.Account
	receiver    []test.Account
	currency    []currencytypes.CurrencyID
	ownerKey    string // Private Key
	senderKey   string // Private Key
	contractKey string // Private Key
	targetKey   string // Private Key
	receiverKey string // Private Key
	owner       []test.Account
}

func (t *testTransferFrom) SetupTest() {
	opr := mpoint.NewTestTransferFromProcessor(util.Encoders)
	t.TestTransferFromProcessor = opr
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.target = make([]test.Account, 1)
	t.receiver = make([]test.Account, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.targetKey = t.NewPrivateKey("target")
	t.receiverKey = t.NewPrivateKey("receiver")
}

func (t *testTransferFrom) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetAccount(t.receiverKey, 1000, t.GenesisCurrency, t.receiver, true).
		SetAccount(t.targetKey, 1000, t.GenesisCurrency, t.target, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.receiver[0].Address(), t.target[0].Address(), common.NewBig(1000), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func TestTransferFrom(t *testing.T) {
	suite.Run(t, new(testTransferFrom))
}
