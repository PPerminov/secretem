```shell
openssl ecparam -name prime256v1 -out ec_param.pem

openssl genpkey -paramfile ec_param.pem -out private.pem

openssl pkey -in private.pem -pubout -out public.pem
```
