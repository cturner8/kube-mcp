package main

import (
	"flag"
	"fmt"
)

var (
	host = flag.String("host", "localhost", "host to connect to/listen on")
	port = flag.Int("port", 9000, "port number to connect to/listen on")
)

func main() {
	// Parse command-line flags.
	flag.Parse()
	// Construct the server URL.
	url := fmt.Sprintf("%s:%d", *host, *port)
	// Start the MCP server.
	startServer(url)
}
