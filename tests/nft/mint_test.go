package credential

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum-nft/v2/operation/nft"
	nfttypes "github.com/ProtoconNet/mitum-nft/v2/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testMint struct {
	suite.Suite
	nft.TestMintProcessor
	sender       []test.Account
	contract     []test.Account
	receiver     []test.Account
	whitelist    []test.Account
	signer       []nfttypes.Signer
	signers      []nfttypes.Signers
	currency     []currencytypes.CurrencyID
	ownerKey     string // Private Key
	senderKey    string // Private Key
	contractKey  string // Private Key
	receiverKey  string // Private Key
	whitelistKey string // Private Key
	owner        []test.Account
}

func (t *testMint) SetupTest() {
	opr := nft.NewTestMintProcessor(util.Encoders)
	t.TestMintProcessor = opr
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.receiver = make([]test.Account, 1)
	t.whitelist = make([]test.Account, 1)
	t.signer = make([]nfttypes.Signer, 1)
	t.signers = make([]nfttypes.Signers, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.whitelistKey = t.NewPrivateKey("whitelist")
	t.receiverKey = t.NewPrivateKey("receiver")
}

func (t *testMint) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetAccount(t.receiverKey, 1000, t.GenesisCurrency, t.receiver, true).
		SetAccount(t.whitelistKey, 1000, t.GenesisCurrency, t.whitelist, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetDesign("abd collection", 10, "example.com").
		SetService(t.sender[0].Address(), t.contract[0].Address(), t.whitelist).
		SetSigner(t.sender[0], 10, false, t.signer).
		SetSigners(10, t.signer, t.signers).
		MakeItem(t.contract[0], t.receiver[0], "nft hash", "example.com", t.signers[0], t.GenesisCurrency, t.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func TestMint(t *testing.T) {
	suite.Run(t, new(testMint))
}
