package main

import (
	"fmt"
	"github.com/suraboy/go-fiber-api/cmd"
)

func main() {
	fmt.Printf(`kkp-api %s, built with Go %s`, cmd.Version, cmd.GoVersion)
	cmd.Execute()
	
}
