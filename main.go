
package main

import (
	"github.com/kry0sc0pic/mercurius/cmd"
	"github.com/kry0sc0pic/mercurius/config"
)

func main() {
	config.LoadConfig()
	cmd.Execute()
}

