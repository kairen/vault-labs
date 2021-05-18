# Secrets Store CSI
Run the following commands for running this example:

```sh
$ kubectl -n vault run vault-client --image vault:1.7.1 \
    --env=VAULT_ADDR="http://vault:8200" \
    --rm --tty -i --force -- /bin/sh

# Put and get kv into Vault:
$ vault kv put secret/db-pass password="db-secret-password"
$ vault kv get secret/db-pass

# Write out the policy named internal-app, and set auth role:
$ vault policy write nginx-app - <<EOF
path "secret/data/db-pass" {
  capabilities = ["read"]
}
EOF

$ vault write auth/kubernetes/role/nginx-app \
    bound_service_account_names=nginx-sa \
    bound_service_account_namespaces=default \
    policies=nginx-app \
    ttl=20m

$ vault read auth/kubernetes/role/nginx-app
```

## Run example
Open a new terminal tab to run example by kubectl:

```sh
$ kubectl apply -f secret-provider-class.yml
$ kubectl get secretproviderclasses vault-nginx-app -o yaml
$ kubectl apply -f nginx.yml 

$ kubectl exec -ti nginx -- cat /mnt/secrets-store/db-password
db-secret-password
```

# References

- https://github.com/kubernetes-sigs/secrets-store-csi-driver
- https://github.com/hashicorp/vault-csi-provider
- https://learn.hashicorp.com/tutorials/vault/kubernetes-secret-store-driver