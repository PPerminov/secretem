package main

import (
	"chat/config"
	"chat/keys"
	"chat/logAndErrors"
)

func main() {
	logAndErrors.Init()
	//var err error
	config1 := config.LoadConfigFile("/home/pavel/workspace/secretem/config.yaml")
	config1, err := keys.LoadKeypair(config1)
	logAndErrors.IfErrorThenFatal(err, "")
	logAndErrors.Fatal("shit")

}
