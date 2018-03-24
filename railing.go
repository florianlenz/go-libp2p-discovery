package discovery

import (
	"time"

	ipfsaddr "github.com/ipfs/go-ipfs-addr"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
)

type Railing struct {
	bootstrapPeers []string
	interval       time.Duration
	stopped        bool
}

//Start railing
func (r *Railing) Start(cb DiscoveryCallback) {

	r.stopped = false

	go func() {

		for r.stopped == false {

			for _, peer := range r.bootstrapPeers {
				iAddr, err := ipfsaddr.ParseString(peer)

				if err != nil {
					cb(err, &peerstore.PeerInfo{})
					continue
				}

				pInfo, err := peerstore.InfoFromP2pAddr(iAddr.Multiaddr())

				cb(err, pInfo)
			}

			time.Sleep(r.interval)

		}

	}()

}

//Stop discovery
func (r *Railing) Stop() {
	r.stopped = true
}

//Create new railing instance
func NewRailing(bootstrapPeers []string, interval time.Duration) Railing {

	return Railing{
		bootstrapPeers: bootstrapPeers,
		interval:       interval,
		stopped:        false,
	}
}
