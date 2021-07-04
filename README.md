# Vault + Kubernetes Labs
This repo contains examples that cover various use cases and functionality for Hashicorp Vault + Kubernetes.

## What you’ll need

* 4 CPUs or more.
* 8GB of free memory.
* 20GB of free disk space.
* Internet connection.
* minikube v1.19.0+.
* vault 1.7.0+

## Quick Start
Run the following commands for starting a Kubernetes cluster by minikube, and deploying Vault + Consult + CSI into cluster:

```sh
$ minikube start --cpus=4 --memory=8g
$ kubectl get no
$ kubectl apply -f deploy/k8s/
$ kubectl apply -f deploy/k8s/vault \
    -f deploy/k8s/secrets-store-csi \
    -f deploy/k8s/vault-secret-csi \
    -f deploy/k8s/vault-injector

$ kubectl -n vault get po,svc
$ kubectl -n vault port-forward service/vault 8200:8200
```
> Access UI from [127.0.0.1:8200](http://127.0.0.1:8200/ui/)

Open a new terminal tab to init Vault by kubectl and Vault:

```sh
$ kubectl -n vault run vault-client --image vault:1.7.1 \
    --env=VAULT_ADDR="http://vault:8200" \
    --rm --tty -i --force -- /bin/sh

$ vault operator init
$ vault operator unseal
$ vault status

$ export VAULT_TOKEN=<ROOT_TOKEN>
$ vault secrets enable -version=2 -path=secret kv
$ vault auth enable kubernetes

# Configure the Kubernetes authentication method to use the service account token:
$ vault write auth/kubernetes/config \
    issuer="https://kubernetes.default.svc.cluster.local" \
    token_reviewer_jwt="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \
    kubernetes_host="https://$KUBERNETES_PORT_443_TCP_ADDR:443" \
    kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt

$ vault read auth/kubernetes/config
```

### Run examples

- [Library](examples/library)
- [Vault Agent](examples/vault-agent)
- [Secrets Store CSI](examples/secrets-store-csi)
