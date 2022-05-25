// Deprecated: This package has moved into go-libp2p as a sub-package: github.com/libp2p/go-libp2p/p2p/muxer/mplex.
package peerstream_multiplex

import (
	"github.com/libp2p/go-libp2p/p2p/muxer/mplex"
)

// DefaultTransport has default settings for Transport
// Deprecated: use github.com/libp2p/go-libp2p/p2p/muxer/mplex.DefaultTransport instead.
var DefaultTransport = mplex.DefaultTransport

// Transport implements mux.Multiplexer that constructs
// mplex-backed muxed connections.
// Deprecated: use github.com/libp2p/go-libp2p/p2p/muxer/mplex.Transport instead.
type Transport = mplex.Transport
