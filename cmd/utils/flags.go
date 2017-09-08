package utils


import "gopkg.in/urfave/cli.v1"

var (
	p2pFlag = cli.StringFlag{
		Name: "network",
		Usage: "chose a p2p network, <ipfs>",
		Value: defaultP2PNetWork(),
	}

	subtopicFlag = cli.StringFlag{
		Name: "topic",
		Usage: "chose a ifps pubsub sub topic,<topic>",
		Value: defaultIpfsSubTopic(),
	}
)

const (
	DEFAULT_NET_WORK = "ipfs"
	DEFAULT_IPFS_SUB_TOPIC = "topic"
)

func defaultP2PNetWork() string {
	return DEFAULT_NET_WORK
}

func defaultIpfsSubTopic() string {
	return DEFAULT_IPFS_SUB_TOPIC
}
