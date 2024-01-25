package main

import (
	"chat/config"
	"chat/keys"
	"chat/logAndErrors"
)

func main() {
	logAndErrors.Init()
	//defer logAndErrors.Sync()
	//var err error
	config1 := config.LoadConfigFile("/home/pavel/workspace/secretem/config.yaml")
	config1, _ = keys.LoadKeypair(config1)
	//logAndErrors.IfErrorThenFatal(err, "")
	//logAndErrors.Info("1")
	//logAndErrors.Warn("2")
	//logAndErrors.Error("3")
	//logAndErrors.Fatal("4")
}
