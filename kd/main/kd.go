package main

import (
	"koding/kite/kd"
	"koding/kite/kd/cli"
	"koding/kite/kd/kite"
)

func main() {
	root := cli.NewCLI()
	root.AddCommand("version", kd.NewVersion())
	root.AddCommand("register", kd.NewRegister())

	k := root.AddSubCommand("kite")
	k.AddCommand("install", kite.NewInstall())
	k.AddCommand("list", kite.NewList())
	k.AddCommand("run", kite.NewRun())
	k.AddCommand("tell", kite.NewTell())
	k.AddCommand("uninstall", kite.NewUninstall())

	root.Run()
}
