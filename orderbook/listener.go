package orderbook

/**
新order以及order更改的监听
 */

type Listener interface {
	Listen()
}