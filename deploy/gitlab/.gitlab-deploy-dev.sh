#!/bin/bash

set -f
string=$PROD_DEPLOY_SERVER
array=(${string//,/})
for i in "${!array[@]}"; do 
    echo "Deploy project on server ${array[i]}"
    ssh ubuntu@${array[i]} "ls -la devops-golang/"
    ssh ubuntu@${array[i]} "sudo git -C devops-golang pull origin master"
    ssh ubuntu@${array[i]} "sudo docker-compose -f devops-golang/docker-compose.yml down"
    ssh ubuntu@${array[i]} "sudo docker-compose -f devops-golang/docker-compose.yml up -d"
done