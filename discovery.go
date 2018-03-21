package discovery

import (
	peerstore "github.com/libp2p/go-libp2p-peerstore"
)

type DiscoveryCallback = func(info peerstore.PeerInfo)

//Interface for discovery mechanism
//In case you want to add a new one
//make sure to implement this interface
type DiscoveryMechanism interface {
	//Start discovery
	Start()
	//Stop discovery
	Stop()
	Discovered(cb DiscoveryCallback)
}
