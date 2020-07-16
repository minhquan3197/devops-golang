#!/bin/bash

set -f
string=$PROD_DEPLOY_SERVER
array=(${string//,/})
for i in "${!array[@]}"; do 
    echo "Deploy project on server ${array[i]}"
    ssh ubuntu@${array[i]} "ls -la golang-testing-devops/"
    ssh ubuntu@${array[i]} "git -C golang-testing-devops pull origin master"
    ssh ubuntu@${array[i]} "sudo docker-compose -f golang-testing-devops/docker-compose.yml down"
    ssh ubuntu@${array[i]} "sudo docker-compose -f golang-testing-devops/docker-compose.yml up -d"
done
