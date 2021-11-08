package main

import (
	"fmt"
	"github.com/suraboy/go-fiber-api/cmd"
	"github.com/suraboy/go-fiber-api/protocol"
)

func main() {
	fmt.Printf(`kkp-api %s, built with Go %s`, cmd.Version, cmd.GoVersion)
	protocol.ServeREST()
}
