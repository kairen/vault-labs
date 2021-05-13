#!/bin/sh

set -eu

wget https://github.com/docker/compose/releases/download/1.29.2/docker-compose-Linux-x86_64 -O docker-compose
chmod 0755 docker-compose
minikube cp ./docker-compose /usr/bin/docker-compose
rm -rf docker-compose

minikube ssh sudo chmod 0755 /usr/bin/docker-compose
minikube ssh docker-compose version