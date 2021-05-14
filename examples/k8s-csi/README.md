# Kubernetes Vault CSI
Run the following commands for deploying this example:

```sh
$ kubectl apply -f secret-provider-class.yml
$ kubectl get secretproviderclasses vault-database -o yaml
$ kubectl apply -f nginx.yml 

$ kubectl exec -ti nginx -- cat /mnt/secrets-store/db-password
```