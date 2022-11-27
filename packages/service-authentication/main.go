// TODO: add documentation
package main

import (
	"github.com/smaperture/go-libs/serverconfig"
	"github.com/smaperture/service-authentication/api"
	"github.com/smaperture/service-authentication/database"
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

	api.Server()

	// if err := rpc.Server(); err != nil {
	// 	panic(err)
	// }
}
