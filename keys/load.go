package keys

import (
	"chat/config"
	"chat/logAndErrors"
)

func LoadKeypair(config *config.Configuration, runtimeConfig *config.RuntimeConfiguration) (*config.RuntimeConfiguration, error) {
	var err error
	if config.KeyType == "ecc" {
		runtimeConfig.EccPrivateKey, err = loadECCPrivateKey(config.PrivateKeyPath)
		logAndErrors.IfErrorThenFatal(err, "")
		runtimeConfig.EccPublicKey, err = loadECCPublicKey(config.PrivateKeyPath)
		logAndErrors.IfErrorThenFatal(err, "")
	}
}
