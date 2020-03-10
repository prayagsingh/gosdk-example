package sdkconnector

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	providersmsp "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

//CreateChennel creates a channel .Need organization setup, channel id and channel config path.
func CreateChennel(setup *OrgSetup, channelid string, channelconfigpath string) error {
	resourceManagerClientContext := setup.sdk.Context(fabsdk.WithUser(setup.AdminName), fabsdk.WithOrg(setup.OrgName))
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return err
	}
	MSPClient, err := msp.New(setup.sdk.Context(), msp.WithOrg(setup.OrgName))
	if err != nil {
		return err
	}
	adminIdentity, err := MSPClient.GetSigningIdentity(setup.AdminName)
	if err != nil {
		return err
	}
	req := resmgmt.SaveChannelRequest{ChannelID: channelid, ChannelConfigPath: channelconfigpath, SigningIdentities: []providersmsp.SigningIdentity{adminIdentity}}
	txID, err := resMgmtClient.SaveChannel(req, resmgmt.WithOrdererEndpoint("orderer.example.com"))
	if err != nil || txID.TransactionID == "" {
		return err
	}
	return nil
}
