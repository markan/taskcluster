package whclient

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/taskcluster/webhooktunnel/wsmux"
)

// Client is used to connect to a proxy and serve endpoints
// defined in Handler
type Client struct {
	ID string
	// handler should be a mux to handle different end points
	Handler   http.Handler
	Config    wsmux.Config
	ProxyAddr string // Address of proxy server for connection
	Retry     RetryConfig
}

const (
	defaultInitialInterval = 500 * time.Millisecond
	defaultMaxInterval     = 60 * time.Second
	defaultMaxElapsedTime  = 3 * time.Minute
	defaultMultiplier      = 1.5
)

// RetryConfig contains exponential backoff parameters for retrying connections
type RetryConfig struct {
	// Retry values
	InitialInterval time.Duration // Default = 500 * time.Millisecond
	MaxInterval     time.Duration // Default = 60 * time.Second
	MaxElapsedTime  time.Duration // Default = 3 * time.Minute
	Multiplier      float64       // Default = 1.5
}

// GetSession connects to the proxy and establishes a wsmux Client session
// The session can be used as a listener to serve HTTP requests
// eg http.Serve(client)
func (c *Client) GetSession(retry bool) (*wsmux.Session, error) {
	addr := strings.TrimSuffix(c.ProxyAddr, "/") + "/register/" + c.ID

	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil && retry {
		// TODO: Reconnect logic
		conn, err = c.Reconnect()
		if err != nil {
			return nil, err
		}
	}

	client := wsmux.Client(conn, c.Config)
	return client, nil
}

// Reconnect attempts to establish a connection to the server
// using an exponential backoff algorithm
// TODO: Add randomization if required
func (c *Client) Reconnect() (*websocket.Conn, error) {
	addr := strings.TrimSuffix(c.ProxyAddr, "/") + "/register/" + c.ID

	c.initializeRetryValues()

	// planned on supporting MaxElapsedTime == 0, but no apparent use case.
	// please advise
	maxTimer := time.NewTimer(c.Retry.MaxElapsedTime)

	currentInterval := c.Retry.InitialInterval
	backoffTimer := time.NewTimer(currentInterval)

	for {
		select {
		case <-maxTimer.C:
			return nil, ErrRetryTimedOut
		case <-backoffTimer.C:
			conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
			if err == nil {
				return conn, err
			}
			// increment backoff
			nextInterval := time.Duration(float64(currentInterval) * c.Retry.Multiplier)
			if nextInterval > c.Retry.MaxInterval {
				nextInterval = c.Retry.MaxInterval
			}
			currentInterval = nextInterval
			_ = backoffTimer.Reset(currentInterval)
		}
	}
}

// initializeRetryValues sets the RetryConfig parameteres to their
// default value
func (c *Client) initializeRetryValues() {
	if c.Retry.InitialInterval == 0 {
		c.Retry.InitialInterval = defaultInitialInterval
	}
	if c.Retry.MaxInterval == 0 {
		c.Retry.MaxInterval = defaultMaxInterval
	}
	if c.Retry.MaxElapsedTime == 0 {
		c.Retry.MaxElapsedTime = defaultMaxElapsedTime
	}

	if c.Retry.Multiplier < 1.0 {
		c.Retry.Multiplier = defaultMultiplier
	}
}
