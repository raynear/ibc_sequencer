package keeper

import (
	"encoding/binary"
	"errors"
	"time"

	"ibc_sequencer/x/sequencer/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
)

// TransmitTlpPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitTlpPacket(
	ctx sdk.Context,
	packetData types.TlpPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) (uint64, error) {
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return 0, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: %w", err)
	}

	return k.channelKeeper.SendPacket(ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, packetBytes)
}

// OnRecvTlpPacket processes packet reception
func (k Keeper) OnRecvTlpPacket(ctx sdk.Context, packet channeltypes.Packet, data types.TlpPacketData) (packetAck types.TlpPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: raynear
	// key := solveTLP(data.Tlp)
	// tx := k.GetTxPool(ctx, 0)
	// dec_tx := decrypt(key, tx.payload)
	// add_dec_tx(dec_tx)
	// k.SetBlock(ctx, types.Block{Txs: })

	go waitAndDoSomething(ctx, k, data)

	return packetAck, nil
}

func waitAndDoSomething(ctx sdk.Context, k Keeper, data types.TlpPacketData) {
	// Calling Sleep method
	time.Sleep(15 * time.Second)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	appendedValue := k.cdc.MustMarshal(&data)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, 0)
	store.Set(bz, appendedValue)
	// 맘대로 저장하기
}

// OnAcknowledgementTlpPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementTlpPacket(ctx sdk.Context, packet channeltypes.Packet, data types.TlpPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.TlpPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutTlpPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutTlpPacket(ctx sdk.Context, packet channeltypes.Packet, data types.TlpPacketData) error {

	// TODO: packet timeout logic

	return nil
}
