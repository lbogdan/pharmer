//go:generate go-extpoints extpoints
package main

import (
	"os"

	logs "github.com/appscode/log/golog"
	"github.com/appscode/pharmer/cmds"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmds.NewRootCmd(Version).Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}