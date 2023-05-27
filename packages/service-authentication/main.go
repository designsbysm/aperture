// TODO: add documentation
package main

import (
	"aperture/go-libs/serverconfig"
	"aperture/service-authentication/database"
	"aperture/service-authentication/rpc"
)

func main() {
	if err := serverconfig.Environment(); err != nil {
		panic(err)
	}

	if err := serverconfig.Loggers(); err != nil {
		panic(err)
	}

	if err := database.Init(); err != nil {
		panic(err)
	}

	// api.Server()

	if err := rpc.Server(); err != nil {
		panic(err)
	}
}
