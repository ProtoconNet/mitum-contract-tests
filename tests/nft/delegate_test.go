package nft

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

type testDelegate struct {
	suite.Suite
	nft.TestDelegateProcessor
	sender       []test.Account
	contract     []test.Account
	delegatee    []test.Account
	whitelist    []test.Account
	signer       []nfttypes.Signer
	signers      []nfttypes.Signers
	currency     []currencytypes.CurrencyID
	ownerKey     string // Private Key
	senderKey    string // Private Key
	contractKey  string // Private Key
	delegateeKey string // Private Key
	whitelistKey string // Private Key
	owner        []test.Account
}

func (t *testDelegate) SetupTest() {
	opr := nft.NewTestDelegateProcessor(util.Encoders)
	t.TestDelegateProcessor = opr
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.delegatee = make([]test.Account, 1)
	t.whitelist = make([]test.Account, 1)
	t.signer = make([]nfttypes.Signer, 1)
	t.signers = make([]nfttypes.Signers, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.whitelistKey = t.NewPrivateKey("whitelist")
	t.delegateeKey = t.NewPrivateKey("delegatee")
}

func (t *testDelegate) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetAccount(t.delegateeKey, 1000, t.GenesisCurrency, t.delegatee, true).
		SetAccount(t.whitelistKey, 1000, t.GenesisCurrency, t.whitelist, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetDesign("abd collection", 10, "example.com").
		SetService(t.sender[0].Address(), t.contract[0].Address(), t.whitelist).
		SetSigner(t.sender[0], 10, false, t.signer).
		SetSigners(10, t.signer, t.signers).
		SetNFT(t.contract[0].Address(), t.sender[0].Address(), "hash", "example.com", t.signers[0]).
		MakeItem(t.contract[0], t.delegatee[0], nft.DelegateAllow, t.GenesisCurrency, t.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.Items()).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err.Error()) {
		t.Suite.T().Log(err.Error())
	}
}

func TestDelegate(t *testing.T) {
	suite.Run(t, new(testDelegate))
}
