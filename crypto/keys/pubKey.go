package keys

import "crypto/ed25519"

type PublicKey struct {
	key ed25519.PublicKey
}

func (pubKey *PublicKey) Bytes() []byte {
	return pubKey.key
}
