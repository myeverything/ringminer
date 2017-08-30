package p2p

type Listener interface {
	Start()
	Stop()
	Name() string
}

