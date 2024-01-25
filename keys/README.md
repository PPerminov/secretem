`openssl ecparam -name prime256v1 -out ec_param.pem`

`openssl genpkey -paramfile ec_param.pem -out ec_private_key.pem`

`openssl pkey -in ec_private_key.pem -pubout -out ec_public_key.pem`
