// TODO: add documentation
package main

import (
	"aperture/go-libs/serverconfig"
	"aperture/service-logger/database"
	"aperture/service-logger/rpc"
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

	if err := rpc.Server(); err != nil {
		panic(err)
	}
}
