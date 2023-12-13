// Package hash provides a way for managing the
// underlying hash implementations used across go-git.
package hash

import (
	"encoding/json"
	"fmt"

	"github.com/tmthrgd/go-hex"
)

var _ json.Marshaler = (*Hash)(nil)
var _ json.Unmarshaler = (*Hash)(nil)

type Hash []byte

func (hash Hash) Hex() string {
	return hex.EncodeToString(hash)
}

func (hash Hash) IsEmpty() bool {
	if hash == nil {
		return true
	}
	return len(hash) == 0
}

func (hash *Hash) UnmarshalJSON(bytes []byte) error {
	if len(bytes) < 2 {
		return fmt.Errorf("hash json must be string")
	}
	if bytes[0] != '"' || bytes[len(bytes)-1] != '"' {
		return fmt.Errorf("hash json must be string")
	}

	if len(bytes) == 2 {
		return nil
	}

	hexData, err := hex.DecodeString(string(bytes[1 : len(bytes)-1]))
	if err != nil {
		return err
	}
	*hash = hexData
	return nil
}

func (hash Hash) MarshalJSON() ([]byte, error) {
	if hash == nil {
		return []byte(`""`), nil
	}
	return []byte(`"` + hash.Hex() + `"`), nil
}

func HashesOfHexArray(hashesStr ...string) ([]Hash, error) {
	hashes := make([]Hash, len(hashesStr))
	for i, hashStr := range hashesStr {
		hash, err := hex.DecodeString(hashStr)
		if err != nil {
			return nil, err
		}
		hashes[i] = hash
	}
	return hashes, nil
}

func HexArrayOfHashes(hashes ...Hash) []string {
	hashesStr := make([]string, len(hashes))
	for i, hash := range hashes {
		hashesStr[i] = hash.Hex()
	}
	return hashesStr
}
