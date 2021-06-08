# Sidecar Injector
Run the following commands for starting this example:

```sh
$ kubectl -n vault run vault-client --image vault:1.7.1 \
    --env=VAULT_ADDR="http://vault:8200" \
    --rm --tty -i --force -- /bin/sh

$ vault kv put secret/webapp/config username=foobaruser password=foobarbazpass
$ vault kv get secret/webapp/config 

$ vault policy write webapp-kv-ro - <<EOF
path "secret/data/webapp/*" {
    capabilities = ["read", "list"]
}
EOF

$ vault write auth/kubernetes/role/webapp \
   bound_service_account_names=webapp \
   bound_service_account_namespaces=default \
   policies=webapp-kv-ro \
   ttl=1h
```

Open a new terminal tab to run example by kubectl:

```sh
$ kubectl apply -f webapp.yml
$ kubectl exec \
    $(kubectl get pod -l app=vault-agent-demo -o jsonpath="{.items[0].metadata.name}") \
    -c webapp -- ls /vault/secrets
ls: /vault/secrets: No such file or directory

# Inject secrets into Pods
$ kubectl patch deployment webapp --patch "$(cat patch-inject-secret.yml)"
$ kubectl get po
NAME                                    READY   STATUS    RESTARTS   AGE
webapp-56c559f6c5-j87bk                 2/2     Running   0          27s

# Wait a minute for recreating Pods
$ kubectl exec \
    $(kubectl get pod -l app=vault-agent-demo -o jsonpath="{.items[0].metadata.name}") \
    -c webapp -- cat /vault/secrets/config.txt
data: map[password:foobarbazpass username:foobaruser]
metadata: map[created_time:2021-05-18T03:00:35.498411842Z deletion_time: destroyed:false version:1]

# Inject secrets as template into Pods
$ kubectl patch deployment webapp --patch "$(cat patch-inject-secrets-as-template.yml)"

# Wait a minute for recreating Pods
$ kubectl exec \
    $(kubectl get pod -l app=vault-agent-demo -o jsonpath="{.items[0].metadata.name}") \
    -c webapp -- cat /vault/secrets/config-tpml.txt
username=foobaruser, password=foobarbazpass
```

# References

- https://www.hashicorp.com/blog/injecting-vault-secrets-into-kubernetes-pods-via-a-sidecar
- https://learn.hashicorp.com/tutorials/vault/kubernetes-sidecar
- https://learn.hashicorp.com/tutorials/vault/agent-kubernetes?in=vault/kubernetes
 