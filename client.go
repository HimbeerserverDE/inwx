package inwx

import (
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/ybbus/jsonrpc/v3"
)

// An Endpoint is the server to communicate with.
type Endpoint string

const (
	Production Endpoint = "https://api.domrobot.com/jsonrpc/"
	Sandbox             = "https://api.ote.domrobot.com/jsonrpc/"
)

const requestTimeout = 30 * time.Second

// A Client is a session handle for the API.
type Client struct {
	rpcClient jsonrpc.RPCClient
	endpoint  Endpoint
}

func NewClient(endpoint Endpoint) (*Client, error) {
	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Jar:     jar,
		Timeout: requestTimeout,
	}

	clt := &Client{
		rpcClient: jsonrpc.NewClientWithOpts(string(endpoint), &jsonrpc.RPCClientOpts{
			HTTPClient: httpClient,
		}),
		endpoint: endpoint,
	}

	return clt, nil
}
