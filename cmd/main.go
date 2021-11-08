package main

import (
	"fmt"
	"github.com/suraboy/go-fiber-api/cmd/cmds"
)

func main() {
	fmt.Printf(`kkp-api %s, built with Go %s`, cmds.Version, cmds.GoVersion)
	cmds.Execute()
}
