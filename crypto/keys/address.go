package keys

import (
	"encoding/hex"
	"fmt"

	"github.com/nivalty/nivalty-sdk/types/chain"
)

const chainAddressFormat = "/%v/%v/%v"

type Address struct {
	address []byte
}

func (addr *Address) Bytes() []byte {
	return addr.address
}

func (addr *Address) String() string {
	keyString := hex.EncodeToString(addr.address[:addressLen])
	return keyString
}

func (addr *Address) ToChainAddress(configs *chain.ChainConfigs) string {
	ca := fmt.Sprintf(chainAddressFormat, configs.ChainType, configs.CoinType, addr.Bytes())
	fmt.Println(ca)
	keyString := hex.EncodeToString([]byte(ca)[:addressLen])
	return fmt.Sprintf("%v%v", configs.Prefix, keyString)
}
