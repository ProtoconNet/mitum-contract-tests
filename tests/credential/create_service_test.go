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

type testCreateService struct {
	suite.Suite
	credential.TestCreateServiceProcessor
	sender      []test.Account
	contract    []test.Account
	currency    []currencytypes.CurrencyID
	ownerKey    string // Private Key
	senderKey   string // Private Key
	contractKey string // Private Key
	creatorKey  string // Private Key
	owner       []test.Account
}

func (t *testCreateService) SetupTest() {
	opr := credential.NewTestCreateServiceProcessor(util.Encoders)
	t.TestCreateServiceProcessor = opr
	t.Setup()
	t.owner = make([]test.Account, 1)
	t.sender = make([]test.Account, 1)
	t.contract = make([]test.Account, 1)
	t.currency = make([]currencytypes.CurrencyID, 1)
	t.ownerKey = t.NewPrivateKey("owner")
	t.senderKey = t.NewPrivateKey("sender")
	t.contractKey = t.NewPrivateKey("contract")
}

func (t *testCreateService) Test01ErrorSenderNotFound() {
	err := t.Create().
		SetAccount(t.senderKey, 1000, t.GenesisCurrency, t.sender, false).
		SetContractAccount(t.sender[0].Address(), t.contractKey, 1000, t.GenesisCurrency, t.contract, true).
		MakeOperation(t.sender[0].Address(), t.sender[0].Priv(), t.contract[0].Address(), t.GenesisCurrency).
		RunPreProcess()

	if assert.NotNil(t.Suite.T(), err) {
		t.Suite.T().Log(err.Error())
	}
}

func TestCreateService(t *testing.T) {
	suite.Run(t, new(testCreateService))
}
