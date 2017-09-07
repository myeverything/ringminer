package eth

import (
	"github.com/Loopring/ringminer/types"
	"github.com/Loopring/ringminer/config"
	"sync"
)

/**
区块链的listener, 得到order，
 */

// TODO(fukun): 添加相关配置
type EthClientConfig struct {
	Ip string
	Port int
}


//监听内容有：环路订单，地址的token余额变动如transfer等
// TODO(fukun):不同的channel，应当交给orderbook统一进行后续处理，可以将channel作为函数返回值、全局变量、参数等方式
type EthClientListener struct {
	config EthClientConfig
	toml config.EthClientOptions
	whisper *types.Whispers
	stop chan struct{}
	lock sync.RWMutex
}

// TODO(fukun): load default config from toml and cli
func (l *EthClientListener) loadConfig() {

}

func NewListener(whisper *types.Whispers, options config.EthClientOptions) *EthClientListener {
	var l EthClientListener
	l.toml = options
	l.loadConfig()

	l.whisper = whisper
	return &l
}

// TODO(fukun): 这里调试调不通,应当返回channel
func (l *EthClientListener) Start() {
	l.stop = make(chan struct{})

	Client.Subscribe(l.whisper.EngineOrderChan)
	go func() {
		for {
			select {
			case ord := <- l.whisper.EngineOrderChan:
				println(ord.RawOrder.SavingSharePercentage)
			}
		}
	}()
}

func (l *EthClientListener) Stop() {
	l.lock.Lock()
	defer l.lock.Unlock()

	close(l.stop)
}
