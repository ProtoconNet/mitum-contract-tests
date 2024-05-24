package nft

import (
	"github.com/ProtoconNet/mitum-contract-tests/tests/util"
	"github.com/ProtoconNet/mitum-currency/v3/operation/test"
	currencytypes "github.com/ProtoconNet/mitum-currency/v3/types"
	"github.com/ProtoconNet/mitum-nft/v2/operation/nft"
	nfttypes "github.com/ProtoconNet/mitum-nft/v2/types"
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
	ap              nft.TestApproveProcessor
	cc              nft.TestCreateCollectionProcessor
	dl              nft.TestDelegateProcessor
	mn              nft.TestMintProcessor
	ts              nft.TestTransferProcessor
	uc              nft.TestUpdateCollectionPolicyProcessor
	sender          []test.Account
	contract        []test.Account
	approved        []test.Account
	whitelist       []test.Account
	receiver        []test.Account
	receiver2       []test.Account
	signer          []nfttypes.Signer
	signers         []nfttypes.Signers
	currency        []currencytypes.CurrencyID
	mockStateGetter *test.MockStateGetter
	senderKey       string // Private Key
	receiverKey     string // Private Key
	receiver2Key    string // Private Key
	contractKey     string // Private Key
	approvedKey     string // Private Key
	whitelistKey    string // Private Key
}

func (t *testNFT) SetupTest() {
	opr1 := nft.NewTestApproveProcessor(util.Encoders)
	t.ap = opr1
	opr2 := nft.NewTestCreateCollectionProcessor(util.Encoders)
	t.cc = opr2
	opr3 := nft.NewTestDelegateProcessor(util.Encoders)
	t.dl = opr3
	opr4 := nft.NewTestMintProcessor(util.Encoders)
	t.mn = opr4
	opr5 := nft.NewTestTransferProcessor(util.Encoders)
	t.ts = opr5
	opr6 := nft.NewTestUpdateCollectionPolicyProcessor(util.Encoders)
	t.uc = opr6
	t.mockStateGetter = test.NewMockStateGetter()
	t.ap.Setup(t.mockStateGetter)
	t.cc.Setup(t.mockStateGetter)
	t.dl.Setup(t.mockStateGetter)
	t.mn.Setup(t.mockStateGetter)
	t.ts.Setup(t.mockStateGetter)
	t.uc.Setup(t.mockStateGetter)
	t.sender = make([]test.Account, 1)
	t.receiver = make([]test.Account, 1)
	t.receiver2 = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.approved = make([]test.Account, 1)
	t.whitelist = make([]test.Account, 1)
	t.signer = make([]nfttypes.Signer, 1)
	t.signers = make([]nfttypes.Signers, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.senderKey = t.ap.NewPrivateKey("sender")
	t.contractKey = t.ap.NewPrivateKey("contract")
	t.whitelistKey = t.ap.NewPrivateKey("whitelist")
	t.approvedKey = t.ap.NewPrivateKey("approved")
	t.receiverKey = t.ap.NewPrivateKey("receiver")
	t.receiver2Key = t.ap.NewPrivateKey("receiver2")
}

func (t *testNFT) CreateCollection() {
	t.cc.Create().
		SetAccount(t.senderKey, 1000, t.cc.GenesisCurrency, t.sender, true).
		SetAccount(t.whitelistKey, 1000, t.cc.GenesisCurrency, t.whitelist, true).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.cc.GenesisCurrency, t.contract, true).
		SetDesign("abd collection", 10, "example.com").
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.whitelist, t.cc.GenesisCurrency).
		RunPreProcess().RunProcess()
}

func (t *testNFT) Mint() {
	t.mn.Create().
		SetAccount(t.senderKey, 1000, t.mn.GenesisCurrency, t.sender, true).
		SetAccount(t.receiverKey, 1000, t.mn.GenesisCurrency, t.receiver, true).
		SetAccount(t.whitelistKey, 1000, t.mn.GenesisCurrency, t.whitelist, true).
		SetSigner(t.sender[0], 10, false, t.signer).
		SetSigners(10, t.signer, t.signers).
		MakeItem(t.contract[0], t.receiver[0], "nft hash", "example.com", t.signers[0], t.mn.GenesisCurrency, t.mn.Items()).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.mn.Items()).
		RunPreProcess().RunProcess()
}

func (t *testNFT) Approve() {
	t.ap.Create().
		SetAccount(t.senderKey, 1000, t.ap.GenesisCurrency, t.sender, true).
		SetAccount(t.approvedKey, 1000, t.ap.GenesisCurrency, t.approved, true).
		SetAccount(t.whitelistKey, 1000, t.ap.GenesisCurrency, t.whitelist, true).
		SetSigner(t.sender[0], 10, false, t.signer).
		SetSigners(10, t.signer, t.signers).
		MakeItem(t.contract[0], t.approved[0], 0, t.ap.GenesisCurrency, t.ap.Items()).
		MakeOperation(t.receiver[0].Address(), t.receiver[0].Priv(), t.ap.Items()).
		RunPreProcess().RunProcess()
}

func (t *testNFT) Transfer() {
	t.ts.Create().
		SetAccount(t.senderKey, 1000, t.ts.GenesisCurrency, t.sender, true).
		SetAccount(t.receiver2Key, 1000, t.ts.GenesisCurrency, t.receiver2, true).
		SetAccount(t.whitelistKey, 1000, t.ts.GenesisCurrency, t.whitelist, true).
		SetSigner(t.sender[0], 10, false, t.signer).
		SetSigners(10, t.signer, t.signers).
		MakeItem(t.contract[0], t.receiver2[0], 0, t.ts.GenesisCurrency, t.ts.Items()).
		MakeOperation(t.receiver[0].Address(), t.receiver[0].Priv(), t.ts.Items()).
		RunPreProcess().RunProcess()
}

func (t *testNFT) UpdateCollectionPolicy() {
	t.uc.Create().
		SetAccount(t.senderKey, 1000, t.uc.GenesisCurrency, t.sender, true).
		SetAccount(t.whitelistKey, 1000, t.uc.GenesisCurrency, t.whitelist, true).
		SetDesign("abd collection", 10, "example.com").
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.whitelist, t.uc.GenesisCurrency).
		RunPreProcess().RunProcess()
}

func TestNFT(t *testing.T) {
	suite.Run(t, new(testNFT))
}
