package p2p

// Peer is anything that represents the remote node
type Peer interface {

}

// Transport is anything that handles the communication
// between the nodes in the network. This can be of the 
// from (TCP, UDP, websockets, ...)
type transport interface {
	ListenAndAccept() error
}