package pointtest

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"testing"

	"github.com/ProtoconNet/mitum-point/operation/point"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testToken struct {
	suite.Suite
	*test.TestProcessor
	approve       point.TestApproveProcessor
	burn          point.TestBurnProcessor
	mint          point.TestMintProcessor
	registerToken point.TestRegisterPointProcessor
	transferFrom  point.TestTransferFromProcessor
	transfer      point.TestTransferProcessor
	sender        []test.Account
	contract      []test.Account
	approved      []test.Account
	target        []test.Account
	receiver      []test.Account
	currency      []currencytypes.CurrencyID
	ownerKey      string // Private Key
	senderKey     string // Private Key
	contractKey   string // Private Key
	approvedKey   string // Private Key
	targetKey     string // Private Key
	receiverKey   string // Private Key
	owner         []test.Account
}

func (t *testToken) SetupTest() {
	tp := test.TestProcessor{Encoders: util.Encoders}
	t.TestProcessor = &tp
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.approve = point.NewTestApproveProcessor(&tp)
	t.burn = point.NewTestBurnProcessor(&tp)
	t.mint = point.NewTestMintProcessor(&tp)
	t.registerToken = point.NewTestRegisterPointProcessor(&tp)
	t.transferFrom = point.NewTestTransferFromProcessor(&tp)
	t.transfer = point.NewTestTransferProcessor(&tp)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.approved = make([]test.Account, 1)
	t.target = make([]test.Account, 1)
	t.receiver = make([]test.Account, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.approvedKey = t.NewPrivateKey("approved")
	t.targetKey = t.NewPrivateKey("target")
	t.receiverKey = t.NewPrivateKey("receiver")
}

func (t *testToken) RegisterToken() {
	t.registerToken.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(),
			"ABC", "token_name", 2500, t.GenesisCurrency).
		RunPreProcess().
		RunProcess()
}

func (t *testToken) Mint() {
	t.mint.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.targetKey, 1000, t.GenesisCurrency, t.target, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(),
			t.target[0].Address(), 1000, t.GenesisCurrency).
		RunPreProcess().
		RunProcess()
}

func (t *testToken) Approve() {
	t.approve.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.approvedKey, 1000, t.GenesisCurrency, t.approved, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(),
			t.approved[0].Address(), 1000, t.GenesisCurrency).
		RunPreProcess().
		RunProcess()
}

func (t *testToken) Burn() {
	t.burn.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(),
			t.sender[0].Address(), 1000, t.GenesisCurrency).
		RunPreProcess().
		RunProcess()
}

func (t *testToken) Transfer() {
	t.transfer.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.receiverKey, 1000, t.GenesisCurrency, t.receiver, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(),
			t.receiver[0].Address(), 1000, t.GenesisCurrency).
		RunPreProcess().
		RunProcess()
}

func (t *testToken) TransferFrom() {
	t.transferFrom.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.receiverKey, 1000, t.GenesisCurrency, t.receiver, true).
		SetAccount(t.targetKey, 1000, t.GenesisCurrency, t.target, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(),
			t.receiver[0].Address(), t.target[0].Address(), 1000, t.GenesisCurrency).
		RunPreProcess().RunProcess()
}

func (t *testToken) Test01ErrorSenderNotFound() {
	t.RegisterToken()
	t.Mint()
	t.Approve()
	t.Burn()
	t.Transfer()

	t.transferFrom.Create().
		SetAccount(t.receiverKey, 1000, t.GenesisCurrency, t.receiver, true).
		MakeOperation(t.approved[0].Address(), t.approved[0].Priv(), t.contract[0].Address(),
			t.receiver[0].Address(), t.sender[0].Address(), 1000, t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), t.Error()) {
		t.Suite.T().Log(t.Error().Error())
	}
}

func TestToken(t *testing.T) {
	suite.Run(t, new(testToken))
}
