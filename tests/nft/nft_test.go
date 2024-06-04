package nfttest

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum-nft/operation/nft"
	nfttypes "github.com/ProtoconNet/mitum-nft/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

//Use below variables for default node configuration values
//t.NetworkID    	: network id
//t.GenesisPriv  	: genesis account private key
//t.GenesisAddr  	: genesis account address
//t.GenesisCurrency : genesis currency

type testNFT struct {
	suite.Suite
	*test.TestProcessor
	approve          nft.TestApproveProcessor
	createCollection nft.TestCreateCollectionProcessor
	delegate         nft.TestDelegateProcessor
	mint             nft.TestMintProcessor
	transfer         nft.TestTransferProcessor
	updateCollection nft.TestUpdateCollectionPolicyProcessor
	sender           []test.Account
	contract         []test.Account
	approved         []test.Account
	whitelist        []test.Account
	receiver         []test.Account
	receiver2        []test.Account
	signer           []nfttypes.Signer
	signers          []nfttypes.Signers
	currency         []currencytypes.CurrencyID
	mockStateGetter  *test.MockStateGetter
	senderKey        string // Private Key
	receiverKey      string // Private Key
	receiver2Key     string // Private Key
	contractKey      string // Private Key
	approvedKey      string // Private Key
	whitelistKey     string // Private Key
}

func (t *testNFT) SetupTest() {
	tp := test.TestProcessor{Encoders: util.Encoders}
	t.TestProcessor = &tp
	mockGetter := test.NewMockStateGetter()
	t.Setup(mockGetter)
	t.approve = nft.NewTestApproveProcessor(&tp)
	t.createCollection = nft.NewTestCreateCollectionProcessor(&tp)
	t.delegate = nft.NewTestDelegateProcessor(&tp)
	t.mint = nft.NewTestMintProcessor(&tp)
	t.transfer = nft.NewTestTransferProcessor(&tp)
	t.updateCollection = nft.NewTestUpdateCollectionPolicyProcessor(&tp)
	t.mockStateGetter = test.NewMockStateGetter()
	t.sender = make([]test.Account, 1)
	t.receiver = make([]test.Account, 1)
	t.receiver2 = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.approved = make([]test.Account, 1)
	t.whitelist = make([]test.Account, 1)
	t.signer = make([]nfttypes.Signer, 1)
	t.signers = make([]nfttypes.Signers, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
	t.whitelistKey = t.NewPrivateKey("whitelist")
	t.approvedKey = t.NewPrivateKey("approved")
	t.receiverKey = t.NewPrivateKey("receiver")
	t.receiver2Key = t.NewPrivateKey("receiver2")
}

func (t *testNFT) CreateCollection() {
	t.createCollection.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.whitelistKey, 1000, t.GenesisCurrency, t.whitelist, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		SetDesign("abd collection", 10, "example.com").
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.whitelist, t.GenesisCurrency).
		RunPreProcess().RunProcess()
}

func (t *testNFT) Mint() {
	t.mint.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.receiverKey, 1000, t.GenesisCurrency, t.receiver, true).
		SetAccount(t.whitelistKey, 1000, t.GenesisCurrency, t.whitelist, true).
		SetSigner(t.sender[0], 10, false, t.signer).
		SetSigners(t.signer, t.signers).
		MakeItem(t.contract[0], t.receiver[0], "nft hash", "example.com", t.signers[0], t.GenesisCurrency, t.mint.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.mint.Items()).
		RunPreProcess().RunProcess()
}

func (t *testNFT) Approve() {
	t.approve.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.approvedKey, 1000, t.GenesisCurrency, t.approved, true).
		SetAccount(t.whitelistKey, 1000, t.GenesisCurrency, t.whitelist, true).
		SetSigner(t.sender[0], 10, false, t.signer).
		SetSigners(t.signer, t.signers).
		MakeItem(t.contract[0], t.approved[0], 0, t.GenesisCurrency, t.approve.Items()).
		MakeOperation(t.receiver[0].Address(), t.receiver[0].Priv(), t.approve.Items()).
		RunPreProcess().RunProcess()
}

func (t *testNFT) Transfer() {
	t.transfer.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.receiver2Key, 1000, t.GenesisCurrency, t.receiver2, true).
		SetAccount(t.whitelistKey, 1000, t.GenesisCurrency, t.whitelist, true).
		SetSigner(t.sender[0], 10, false, t.signer).
		SetSigners(t.signer, t.signers).
		MakeItem(t.contract[0], t.receiver2[0], 0, t.GenesisCurrency, t.transfer.Items()).
		MakeOperation(t.receiver[0].Address(), t.receiver[0].Priv(), t.transfer.Items()).
		RunPreProcess().RunProcess()
}

func (t *testNFT) UpdateCollectionPolicy() {
	t.updateCollection.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, true).
		SetAccount(t.whitelistKey, 1000, t.GenesisCurrency, t.whitelist, true).
		SetDesign("abd collection", 10, "example.com").
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.whitelist, t.GenesisCurrency).
		RunPreProcess().RunProcess()
}

func TestNFT(t *testing.T) {
	suite.Run(t, new(testNFT))
}
