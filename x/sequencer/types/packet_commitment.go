package types

// ValidateBasic is used for validating the packet
func (p CommitmentPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p CommitmentPacketData) GetBytes() ([]byte, error) {
	var modulePacket SequencerPacketData

	modulePacket.Packet = &SequencerPacketData_CommitmentPacket{&p}

	return modulePacket.Marshal()
}
