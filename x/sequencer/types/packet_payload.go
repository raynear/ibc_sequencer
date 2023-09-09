package types

// ValidateBasic is used for validating the packet
func (p PayloadPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p PayloadPacketData) GetBytes() ([]byte, error) {
	var modulePacket SequencerPacketData

	modulePacket.Packet = &SequencerPacketData_PayloadPacket{&p}

	return modulePacket.Marshal()
}
