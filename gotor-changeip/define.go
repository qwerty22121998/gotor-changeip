package gotor_changeip

import (
	"context"
	"github.com/cretz/bine/tor"
	"github.com/ipsn/go-libtor"
	"net/http"
	"os"
)

const IP_CHECK_URL = "http://checkip.amazonaws.com/"
const IP_NOT_FOUND = "NOT FOUND"

var TOR_CONFIG = tor.StartConf{
	ProcessCreator: libtor.Creator,
	DebugWriter:    os.Stderr,
	// more configs go here
}

type TorClient struct {
	Client     *http.Client
	tor        *tor.Tor
	dialCtx    context.Context
	dialCancel context.CancelFunc
	dialer     *tor.Dialer
	ip         string
}
