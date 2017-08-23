package matchengine


type FilterChange struct {

}

type BlockChainClient interface {

	SendRingHash() //发送环路凭证

	SendRing() //发送环路

	SetFilter() chan FilterChange //订阅filter，新transaction等事件

}
