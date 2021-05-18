# Go Library
Run the following commands for starting this example:

```sh
$ kubectl -n vault run vault-client --image vault:1.7.1 \
    --env=VAULT_ADDR="http://vault:8200" \
    --rm --tty -i --force -- /bin/sh

$ vault kv put secret/api/config db_username='api' \
    db_password='r00tme_my_db' \
    ttl='30s'

$ vault kv get secret/api/config
$ vault policy write internal-api - <<EOF
path "secret/data/api/*" {
    capabilities = ["read", "list"]
}
EOF

$ vault token create -display-name=api-token -policy=internal-api 
```
> Copy the `token` from console output.

(Optional) If you want to login by user/password, please follow as below:

```sh
$ vault auth enable userpass
$ vault write auth/userpass/users/api \
    password=r00tme \
    policies=internal-api

$ vault login -method=userpass \
    username=api \
    password=r00tme
```

## Run example
Open a new terminal tab to run example by Go:

```sh
$ export VAULT_ADDR="http://127.0.0.1:8200"
$ export VAULT_TOKEN=<API_TOKEN>
$ go run main.go
DB username: api
DB password: r00tme_my_db
```

(Optional) If you want to login by user/password, please follow as below:

```sh
$ export VAULT_ADDR="http://127.0.0.1:8200"
$ export VAULT_USER=api
$ export VAULT_PASSOWRD=r00tme
$ go run main.go
DB username: api
DB password: r00tme_my_db
```


# References

- https://www.vaultproject.io/api/libraries#go
- https://github.com/hashicorp/vault/tree/master/api