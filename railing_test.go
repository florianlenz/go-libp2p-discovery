package discovery

import (
	"testing"
	"time"

	"github.com/libp2p/go-libp2p-peerstore"
)

func TestRailingStartStop(t *testing.T) {

	r := NewRailing([]string{}, time.Second*3)

	//Default value of stopped should be false
	if r.stopped != false {
		t.Error("Expect stopped to be false")
	}

	r.Stop()

	if r.stopped != true {
		t.Error("Expect stopped to be true since we called Stop()")
	}

	r.Start(func(err error, info *peerstore.PeerInfo) {

	})

	if r.stopped != false {
		t.Error("Exepcted stopped to be false when restarting")
	}

}

func TestRailingStartWithMalformedAddress(t *testing.T) {

	r := NewRailing([]string{
		"i_am_a_malformed_address",
	}, time.Second*3)

	c := make(chan struct{})

	r.Start(func(err error, info *peerstore.PeerInfo) {

		if err.Error() != "invalid multiaddr, must begin with /" {
			t.Error(err)
		}

		c <- struct{}{}
	})

	<-c
}

func TestRailingStart(t *testing.T) {

	r := NewRailing([]string{
		"/ip4/104.131.131.82/tcp/4001/ipfs/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	}, time.Second*3)

	c := make(chan struct{})

	r.Start(func(err error, info *peerstore.PeerInfo) {

		if err != nil {
			t.Error(err)
		}

		if info.ID.String() != "<peer.ID aCpDMG>" {
			t.Errorf("Expect peer id to be <peer.ID aCpDMG> - got: %s", info.ID.String())
		}

		c <- struct{}{}
	})

	<-c

}
