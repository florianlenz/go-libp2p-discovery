package discovery

import (
	peerstore "github.com/libp2p/go-libp2p-peerstore"
)

type DiscoveryCallback = func(err error, info *peerstore.PeerInfo)

//Interface for discovery mechanism
//In case you want to add a new one
//make sure to implement this interface
type DiscoveryMechanism interface {
	//Start discovery.
	//cb will be called with discovered peer
	Start(cb DiscoveryCallback)
	//Stop discovery
	Stop()
}
