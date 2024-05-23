package keys

import "encoding/hex"

type Address struct {
	address []byte
}

func (addr *Address) Bytes() []byte {
	return addr.address
}

func (addr *Address) String() string {
	keyString := hex.EncodeToString(addr.address)
	return keyString
}
