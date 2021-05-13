# Docker compose
Run the following commands in separate terminals for starting a Kubernetes cluster by minikube, and deploying Vault + Consult by docker-compose:

```sh
# first terminal
$ minikube start --cpus=4 --memory=8g
$ minikube mount ./:/home/docker/vault-workshop

# second terminal(debug)
$ ./scripts/get-compose.sh
$ minikube ssh
$ cd vault-workshop/deploy/compose && docker-compose up
```