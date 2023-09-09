package types

// IBC events
const (
	EventTypeTimeout          = "timeout"
	EventTypeCommitmentPacket = "commitment_packet"
	EventTypePayloadPacket    = "payload_packet"
	EventTypeTlpPacket        = "tlp_packet"
	// this line is used by starport scaffolding # ibc/packet/event

	AttributeKeyAckSuccess = "success"
	AttributeKeyAck        = "acknowledgement"
	AttributeKeyAckError   = "error"
)
