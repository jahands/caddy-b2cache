package b2

import "github.com/caddyserver/caddy"

func init() {
	caddy.RegisterPlugin("b2", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {

	return nil
}
