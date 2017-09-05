package eth

type Listener interface {
	Start()
	Stop()
}