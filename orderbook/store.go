package orderbook


type StoreVersion struct {
	BlockNumber uint64
	TxIdx uint32
}

type Store interface {
	Store()
	GetOrder(hash string)
	GetAllOrder(version StoreVersion)
	SetStoreVersion(version StoreVersion)
}
