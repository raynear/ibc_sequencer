package types

// ValidateBasic is used for validating the packet
func (p TlpPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p TlpPacketData) GetBytes() ([]byte, error) {
	var modulePacket SequencerPacketData

	modulePacket.Packet = &SequencerPacketData_TlpPacket{&p}

	return modulePacket.Marshal()
}
