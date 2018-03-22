package discovery

import (
	"testing"
	"time"

	"github.com/libp2p/go-libp2p-peerstore"
)

func TestRailing_Stop(t *testing.T) {

	r := NewRailing([]string{}, time.Second*3)

	//Default value of stopped should be false
	if r.stopped != false {
		t.Error("Expect stopped to be false")
	}

	r.Stop()

	if r.stopped != true {
		t.Error("Expect stopped to be true since we called Stop()")
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
