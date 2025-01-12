package gmp

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"
)

type IBCMiddleware struct {
	app porttypes.IBCModule
}

func NewIBCMiddleware(app porttypes.IBCModule) IBCMiddleware {
	return IBCMiddleware{
		app: app,
	}
}

// OnChanOpenInit implements the IBCModule interface
func (im IBCMiddleware) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	// call underlying callback
	return im.app.OnChanOpenInit(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, version)
}

// OnChanOpenTry implements the IBCMiddleware interface
func (im IBCMiddleware) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	channelCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	return im.app.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, channelCap, counterparty, counterpartyVersion)
}

// OnChanOpenAck implements the IBCMiddleware interface
func (im IBCMiddleware) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID string,
	counterpartyChannelID string,
	counterpartyVersion string,
) error {
	return im.app.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, counterpartyVersion)
}

// OnChanOpenConfirm implements the IBCMiddleware interface
func (im IBCMiddleware) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return im.app.OnChanOpenConfirm(ctx, portID, channelID)
}

// OnChanCloseInit implements the IBCMiddleware interface
func (im IBCMiddleware) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return im.app.OnChanCloseInit(ctx, portID, channelID)
}

// OnChanCloseConfirm implements the IBCMiddleware interface
func (im IBCMiddleware) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	return im.app.OnChanCloseConfirm(ctx, portID, channelID)
}

// OnRecvPacket implements the IBCMiddleware interface
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return channeltypes.NewErrorAcknowledgement(fmt.Errorf("cannot unmarshal ICS-20 transfer packet data"))
	}

	var msg Message
	var err error
	err = json.Unmarshal([]byte(data.GetMemo()), &msg)
	if err != nil || len(msg.Payload) == 0 {
		// Not a packet that should be handled by the GMP middleware
		return im.app.OnRecvPacket(ctx, packet, relayer)
	}

	// TODO, figure out how to test with simulated chains over IBC
	// Since cosmos-sdk v0.50 has backward compatibility issue of not allowing to make two chains with different account prefix,
	// it is only able to run two Allora chains and communicate IBC packets.
	// So it can't send message from account with `axelar` prefix as source chain.

	//if !strings.EqualFold(data.Sender, AxelarGMPAcc) {
	//	// Not a packet that should be handled by the GMP middleware
	//	return im.app.OnRecvPacket(ctx, packet, relayer)
	//}

	logger := ctx.Logger().With("handler", "GMP")

	switch msg.Type {
	case TypeGeneralMessage:
		logger.Info("Received TypeGeneralMessage",
			"srcChain", msg.SourceChain,
			"srcAddress", msg.SourceAddress,
			"receiver", data.Receiver,
			"payload", string(msg.Payload),
			"handler", "GMP",
		)
		// let the next layer deal with this
		// the rest of the data fields should be normal
		fallthrough
	case TypeGeneralMessageWithToken:
		logger.Info("Received TypeGeneralMessageWithToken",
			"srcChain", msg.SourceChain,
			"srcAddress", msg.SourceAddress,
			"receiver", data.Receiver,
			"payload", string(msg.Payload),
			"coin", data.Denom,
			"amount", data.Amount,
			"handler", "GMP",
		)
		// we throw out the rest of the msg.Payload fields here, for better or worse
		data.Memo = string(msg.Payload)
		var dataBytes []byte
		if dataBytes, err = transfertypes.ModuleCdc.MarshalJSON(&data); err != nil {
			return channeltypes.NewErrorAcknowledgement(fmt.Errorf("cannot marshal ICS-20 post-processed transfer packet data"))
		}
		packet.Data = dataBytes
		return im.app.OnRecvPacket(ctx, packet, relayer)
	default:
		return channeltypes.NewErrorAcknowledgement(fmt.Errorf("unrecognized mesasge type: %d", msg.Type))
	}
}

// OnAcknowledgementPacket implements the IBCMiddleware interface
func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	return im.app.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

// OnTimeoutPacket implements the IBCMiddleware interface
func (im IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	return im.app.OnTimeoutPacket(ctx, packet, relayer)
}
