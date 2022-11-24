// TODO: add documentation
package main

import (
	"github.com/smaperture/service-authentication/database"
	"github.com/smaperture/service-authentication/rpc"

	"github.com/smaperture/serverconfig"
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

	// test
}
