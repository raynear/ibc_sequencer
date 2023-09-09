package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TxPoolKeyPrefix is the prefix to retrieve all TxPool
	TxPoolKeyPrefix = "TxPool/value/"
)

// TxPoolKey returns the store key to retrieve a TxPool from the index fields
func TxPoolKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
