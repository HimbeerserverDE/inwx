package inwx

import (
	"net"
	"net/http"
	"net/http/cookiejar"
	"sync/atomic"
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
	httpClient *http.Client
	rpcClient  jsonrpc.RPCClient
	endpoint   Endpoint

	closed  chan struct{}
	closing uint32
	err     error
}

func Login(endpoint Endpoint, user, passwd string) (*Client, error) {
	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Jar:     jar,
		Timeout: requestTimeout,
	}

	rpcClient := jsonrpc.NewClientWithOpts(string(endpoint), &jsonrpc.RPCClientOpts{
		HTTPClient: httpClient,
	})

	clt := &Client{
		httpClient: httpClient,
		rpcClient:  rpcClient,
		endpoint:   endpoint,
	}

	// TODO: login

	return clt, nil
}

// Closed returns a channel which is closed when the Client is closed.
func (c *Client) Closed() <-chan struct{} { return c.closed }

// WhyClosed returns the error that caused the Conn to be closed or nil
// if the Conn was closed using the Close method.
// WhyClosed returns nil if the Conn is not closed.
func (c *Client) WhyClosed() error {
	select {
	case <-c.Closed():
		return c.err
	default:
		return nil
	}
}

// Close closes the Client. Any blocked Calls will return net.ErrClosed.
func (c *Client) Close() error {
	return c.closeLogout(nil)
}

func (c *Client) closeLogout(err error) error {
	// TODO: logout

	return c.close(err)
}

func (c *Client) close(err error) error {
	if atomic.SwapUint32(&c.closing, 1) == 1 {
		return net.ErrClosed
	}

	c.err = err
	defer close(c.closed)

	return nil
}
