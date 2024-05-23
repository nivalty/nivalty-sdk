package keys

import (
	"testing"

	"github.com/nivalty/nivalty-sdk/types/chain"
)

func TestAddress(t *testing.T) {
	addr := Address{[]byte("nivaltynivaltynivaltynivaltynivaltynivaltynivaltynivaltynivaltynivalty")}
	t.Log(addr.String())
	chainConfigs := &chain.ChainConfigs{
		CoinType:  0,
		ChainType: 0,
		Prefix:    "nivalty",
	}
	chainAddress := addr.ToChainAddress(chainConfigs)
	checkingAddress := "nivalty2f302f302f5b3131302031303520313138203937203130382031313620313231"
	if chainAddress != checkingAddress {
		t.Fatal("addresses does not match")
	}
}
