package keys

import (
	"crypto/ed25519"
	"errors"

	"github.com/nivalty/nivalty-sdk/types/msgs"
)

func NewPrivateKeyFromSeed(seed []byte) (*PrivateKey, error) {
	if len(seed) != seedLen {
		return nil, errors.New("len of seed incorrect")
	}
	key := ed25519.NewKeyFromSeed(seed)
	return &PrivateKey{key}, nil
}

type PrivateKey struct {
	key ed25519.PrivateKey
}

func (privKey *PrivateKey) Bytes() []byte {
	return privKey.key
}

func (privKey *PrivateKey) PublicKey() *PublicKey {
	key := make([]byte, publicKeyLen)
	copy(key, privKey.key)
	return &PublicKey{key}
}

func (privKey *PrivateKey) Sign(msg msgs.Message) *Signature {
	return &Signature{ed25519.Sign(privKey.key, msg.Bytes())}
}
