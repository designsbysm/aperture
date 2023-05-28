package main

import (
	"aperture/go-libs/serverconfig"
	"aperture/rest-api/api"
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
