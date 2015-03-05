package dockergo

import "os"

// DefaultClient is the default docker client, using the DOCKER_HOST, DOCKER_TLS_VERIFY, and DOCKER_CERT_PATH env vars
var DefaultClient = NewClient(os.Getenv("DOCKER_HOST"), os.Getenv("DOCKER_TLS_VERIFY"), os.Getenv("DOCKER_CERT_PATH"))

// Attach to a running container
// func Attach() { DefaultClient.Attach() }

// Build an image from a Dockerfile
// func Build() { DefaultClient.Build() }

// Commit creates a new image from a container's changes
// func Commit() { DefaultClient.Commit() }

// CP Copies files/folders from a container's filesystem to the host path
// func CP() { DefaultClient.CP() }

// Create a new container
func Create() { DefaultClient.Create() }

// Diff inspects changes on a container's filesystem
// func Diff() { DefaultClient.Diff() }

// Events get real time events from the server
func Events() { DefaultClient.Events() }

// Exec runs a command in a running container
// func Exec() { DefaultClient.Exec() }

// Export streams the contents of a container as a tar archive
// func Export() { DefaultClient.Export() }

// History shows the history of an image
// func History() { DefaultClient.History() }

// Images lists images
func Images() { DefaultClient.Images() }

// Import creates a new filesystem image from the contents of a tarball
// func Import() { DefaultClient.Import() }

// Info displays system-wide information
func Info() { DefaultClient.Info() }

// Inspect returns low-level information on a container or image
func Inspect() { DefaultClient.Inspect() }

// Kill a running container
func Kill() { DefaultClient.Kill() }

// Load an image from a tar archive
func Load() { DefaultClient.Load() }

// Login to a Docker registry server
// func Login() { DefaultClient.Login() }

// Logout from a Docker registry server
// func Logout() { DefaultClient.Logout() }

// Logs fetches the logs of a container
func Logs() { DefaultClient.Logs() }

// Port looks up the public-facing port that is NAT-ed to PRIVATE_PORT
// func Port() { DefaultClient.Port() }

// Pause all processes within a container
func Pause() { DefaultClient.Pause() }

// PS lists containers
func PS() { DefaultClient.PS() }

// Pull an image or a repository from a Docker registry server
func Pull() { DefaultClient.Pull() }

// Push an image or a repository to a Docker registry server
// func Push() { DefaultClient.Push() }

// Rename an existing container
// func Rename() { DefaultClient.Rename() }

// Restart a running container
// func Restart() { DefaultClient.Restart() }

// RM removes one or more containers
func RM() { DefaultClient.RM() }

// RMI removes one or more images
// func RMI() { DefaultClient.RMI() }

// Run a command in a new container
// func Run() { DefaultClient.Run() }

// Save an image to a tar archive
// func Save() { DefaultClient.Save() }

// Search for an image on the Docker Hub
// func Search() { DefaultClient.Search() }

// Start a stopped container
func Start() { DefaultClient.Start() }

// Stats displays a live stream of one or more containers' resource usage statistics
func Stats() { DefaultClient.Stats() }

// Stop a running container
func Stop() { DefaultClient.Stop() }

// Tag an image into a repository
// func Tag() { DefaultClient.Tag() }

// Top looks up the running processes of a container
// func Top() { DefaultClient.Top() }

// Unpause a paused container
func Unpause() { DefaultClient.Unpause() }

// Version shows the Docker version information
func Version() { DefaultClient.Version() }

// Wait block until a container stops, then print its exit code
func Wait() { DefaultClient.Wait() }
