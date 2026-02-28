package network

// Client handles server communication and client-side prediction.
type Client struct {
	ServerAddress string
	ServerPort    int
	Connected     bool
}

// NewClient creates a Client targeting the given server.
func NewClient(address string, port int) *Client {
	return &Client{
		ServerAddress: address,
		ServerPort:    port,
	}
}

// Connect establishes a connection to the server.
func (c *Client) Connect() error {
	// Skeleton: client connection will be implemented in Phase 4
	return nil
}

// Disconnect closes the server connection.
func (c *Client) Disconnect() error {
	// Skeleton: client disconnection will be implemented in Phase 4
	c.Connected = false
	return nil
}

// SendInput sends player input to the server.
func (c *Client) SendInput(input []byte) error {
	// Skeleton: input sending will be implemented in Phase 4
	return nil
}
