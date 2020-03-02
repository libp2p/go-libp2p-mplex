package peerstream_multiplex

import (
	"testing"

	test "github.com/libp2p/go-libp2p-testing/suites/mux"
)

func TestDefaultTransport(t *testing.T) {
	test.SubtestAll(t, DefaultTransport)
}
