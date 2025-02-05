package main

import "mtls_bot_root/configuration"

var (
	config *configuration.RootConfiguration = configuration.Get()
)

func main() {
	println(config)
}
