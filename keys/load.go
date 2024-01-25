package keys

import (
	"chat/config"
	"chat/logAndErrors"
)

func LoadKeypair(config *config.Configuration) (*config.Configuration, error) {
	var err error
	if config.KeyType == "ecc" {
		config.EccPrivateKey, err = loadECCPrivateKey(config.PrivateKeyPath)
		logAndErrors.IfErrorThenFatal(err, "")
		config.EccPublicKey, err = loadECCPublicKey(config.PrivateKeyPath)
		logAndErrors.IfErrorThenFatal(err, "")
		return config, nil
	} else if config.KeyType == "rsa" {
		config.RsaPrivateKey, err = loadRSAPrivateKey(config.PrivateKeyPath)
		logAndErrors.IfErrorThenFatal(err, "")
		config.RsaPublicKey, err = loadRSAPublicKey(config.PrivateKeyPath)
		logAndErrors.IfErrorThenFatal(err, "")
		return config, nil
	}
	return config, logAndErrors.Wrapper(nil, "key must be either RSA or ECC")
}
