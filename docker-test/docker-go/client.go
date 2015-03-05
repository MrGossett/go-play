package dockergo

import d "github.com/fsouza/go-dockerclient"

// Client represents a docker client
type Client struct {
	client *d.Client
}

// NewClient returns a new docker client
func NewClient(host, tls, certPath string) *Client {
	c := &Client{}

	if host == "" {
		return c
	}

	if tls != "" {
		return NewClientTLS(host, tls, certPath)
	}

	client, err := d.NewClient(host)
	if err != nil {
		return c
	}
	c.client = client

	return c
}

// NewClientTLS returns a new docker client using TLS
func NewClientTLS(host, tls, certPath string) *Client {
	c := &Client{}

	if certPath == "" || tls == "" || host == "" {
		return c
	}

	var (
		cert = certPath + "/cert.pem"
		key  = certPath + "/key.pem"
		ca   = certPath + "/ca.pem"
	)

	client, err := d.NewTLSClient(host, cert, key, ca)
	if err != nil {
		return c
	}
	c.client = client

	return c
}

// Attach to a running container
// func (c *Client) Attach()  { return }

// Build an image from a Dockerfile
// func (c *Client) Build()   { return }

// Commit creates a new image from a container's changes
// func (c *Client) Commit()  { return }

// CP Copies files/folders from a container's filesystem to the host path
// func (c *Client) CP()      { return }

// Create a new container
func (c *Client) Create() { return }

// Diff inspects changes on a container's filesystem
// func (c *Client) Diff()    { return }

// Events get real time events from the server
func (c *Client) Events() { return }

// Exec runs a command in a running container
// func (c *Client) Exec() { return }

// Export streams the contents of a container as a tar archive
// func (c *Client) Export() { return }

// History shows the history of an image
// func (c *Client) History() { return }

// Images lists images
func (c *Client) Images() { return }

// Import creates a new filesystem image from the contents of a tarball
// func (c *Client) Import() { return }

// Info displays system-wide information
func (c *Client) Info() (map[string]string, error) {
	info, err := c.client.Info()
	if err != nil {
		return nil, err
	}

	return info.Map(), nil
}

// Inspect returns low-level information on a container or image
func (c *Client) Inspect() { return }

// Kill a running container
func (c *Client) Kill(id, signal string) error {
	opts := d.KillContainerOptions{ID: id}
	if signal != "" {
		opts.Signal = signal
	}

	return c.client.KillContainer(opts)
}

// Load an image from a tar archive
func (c *Client) Load() { return }

// Login to a Docker registry server
// func (c *Client) Login() { return }

// Logout from a Docker registry server
// func (c *Client) Logout() { return }

// Logs fetches the logs of a container
func (c *Client) LogTail(id string) (chan string, chan string) {
	opts := d.LogsOptions{Container: id, Follow: true}

	c.client.Logs(opts)
}

// Port looks up the public-facing port that is NAT-ed to PRIVATE_PORT
// func (c *Client) Port() { return }

// Pause all processes within a container
func (c *Client) Pause() { return }

// PS lists containers
func (c *Client) PS() { return }

// Pull an image or a repository from a Docker registry server
func (c *Client) Pull() { return }

// Push an image or a repository to a Docker registry server
// func (c *Client) Push() { return }

// Rename an existing container
// func (c *Client) Rename() { return }

// Restart a running container
// func (c *Client) Restart() { return }

// RM removes one or more containers
func (c *Client) RM() { return }

// RMI removes one or more images
// func (c *Client) RMI() { return }

// Run a command in a new container
// func (c *Client) Run() { return }

// Save an image to a tar archive
// func (c *Client) Save() { return }

// Search for an image on the Docker Hub
// func (c *Client) Search() { return }

// Start a stopped container
func (c *Client) Start() { return }

// Stats displays a live stream of one or more containers' resource usage statistics
func (c *Client) Stats() { return }

// Stop a running container
func (c *Client) Stop() { return }

// Tag an image into a repository
// func (c *Client) Tag() { return }

// Top looks up the running processes of a container
// func (c *Client) Top() { return }

// Unpause a paused container
func (c *Client) Unpause() { return }

// Version shows the Docker version information
func (c *Client) Version() { return }

// Wait block until a container stops, then print its exit code
func (c *Client) Wait() { return }
