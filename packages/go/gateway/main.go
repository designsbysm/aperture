package main

import (
	"aperture/gateway/api"
	"aperture/libs/serverconfig"
)

func main() {
	if err := serverconfig.Environment(); err != nil {
		panic(err)
	}

	if err := serverconfig.Loggers(); err != nil {
		panic(err)
	}

	api.Server()
}
