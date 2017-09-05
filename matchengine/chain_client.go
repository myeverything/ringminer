package matchengine

import (
	"github.com/Loopring/ringminer/types"
	"github.com/Loopring/ringminer/chainclient"
)

type ChainClient struct {
	Tokens map[types.Address]chainclient.Erc20Token

}

