package args

import (
	"flag"
)

var (
	Version    *string
	NewAccount *bool
	Port       *int
)

func init() {
	Version = flag.String("v", "wl-uf-latest", "The version of the SDE to use.")
	Port = flag.Int("port", 80, "The port used to serve content")
	NewAccount = flag.Bool("newacc", false, "Doesn't start the server and querrys for a new account.")
	flag.Parse()
}
