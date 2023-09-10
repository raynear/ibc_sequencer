package keeper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"ibc_sequencer/x/sequencer/types"

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

	go solveTLPAndPushTx(ctx, k, data)

	return packetAck, nil
}

func solveTLPAndPushTx(ctx sdk.Context, k Keeper, data types.TlpPacketData) {
	txs := k.GetAllTxPool(ctx)
	for i := 0; i < len(txs); i++ {
		if txs[i].Hash == data.Hash {
			key := solveTLP(data.Tlp)
			dec_tx, err := decrypt(key, txs[i].Payload)

			if err != nil {
				var newTx types.TxPool
				newTx.Payload = fmt.Sprintf("%v", dec_tx)
				newTx.Creator = txs[i].Creator
				newTx.Index = txs[i].Index
				newTx.Round = txs[i].Round
				newTx.Hash = "done"
				k.SetTxPool(ctx, newTx)
			}
		}
	}
}

func decrypt(key []byte, cipherTextHex string) ([]byte, error) {
	// Convert the hexadecimal string back to a byte slice
	cipherText, err := hex.DecodeString(cipherTextHex)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Extract the IV from the cipherText
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	// Create a new AES cipher block mode
	stream := cipher.NewCTR(block, iv)

	// Decrypt the cipherText
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	return plainText, nil
}

func solveTLP(seed string) []byte {
	numIterations := 26665975

	// Initialize a variable to hold the current hash value
	currentHash := seed

	for i := 0; i < numIterations; i++ {
		hash := sha256.New()
		hash.Write([]byte(currentHash))
		hashBytes := hash.Sum(nil)
		currentHash = hex.EncodeToString(hashBytes)
	}

	return []byte(currentHash[:32])
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
