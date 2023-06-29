package main

import (
	"aperture/libs/serverconfig"
	"aperture/service-auth/database"
	"aperture/service-auth/rpc"
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
