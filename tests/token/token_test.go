package tokentest

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"testing"

	"github.com/ProtoconNet/mitum-token/operation/token"
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
	approve       token.TestApproveProcessor
	burn          token.TestBurnProcessor
	mint          token.TestMintProcessor
	registerToken token.TestRegisterTokenProcessor
	transferFrom  token.TestTransferFromProcessor
	transfer      token.TestTransferProcessor
	sender        []test.Account
	contract      []test.Account
	contract2     []test.Account
	approved      []test.Account
	target        []test.Account
	receiver      []test.Account
	currency      []currencytypes.CurrencyID
	ownerKey      string // Private Key
	senderKey     string // Private Key
	contractKey   string // Private Key
	contract2Key  string // Private Key
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
	t.approve = token.NewTestApproveProcessor(&tp)
	t.burn = token.NewTestBurnProcessor(&tp)
	t.mint = token.NewTestMintProcessor(&tp)
	t.registerToken = token.NewTestRegisterTokenProcessor(&tp)
	t.transferFrom = token.NewTestTransferFromProcessor(&tp)
	t.transfer = token.NewTestTransferProcessor(&tp)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.contract2 = make([]test.Account, 1)
	t.approved = make([]test.Account, 1)
	t.target = make([]test.Account, 1)
	t.receiver = make([]test.Account, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.contract2Key = t.NewPrivateKey("contract2")
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

	t.transfer.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.receiverKey, 1000, t.GenesisCurrency, t.receiver, true).
		SetContractAccount(t.sender[0].Address(), t.contract2Key, 1000, t.GenesisCurrency, t.contract2, false).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract2[0].Address(),
			t.receiver[0].Address(), 1000, t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), t.Error()) {
		t.Suite.T().Log(t.Error().Error())
	}
}

func TestToken(t *testing.T) {
	suite.Run(t, new(testToken))
}
