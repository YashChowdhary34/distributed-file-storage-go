package p2p

// message holds any arbitary data that is sent over for each
// transport between the nodes in the network
type Message struct {
	payload []byte
}