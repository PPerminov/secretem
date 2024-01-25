package main

import (
	"chat/config"
	"chat/logAndErrors"
)

func main() {
	logAndErrors.Init()
	defer logAndErrors.Sync()
	logAndErrors.Dump(config.LoadConfigFile("/home/pavel/workspace/chat/config.yaml"))
	//logAndErrors.Info("1")
	//logAndErrors.Warn("2")
	//logAndErrors.Error("3")
	//logAndErrors.Fatal("4")
}
