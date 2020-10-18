#!/bin/bash

set -f
string=$PROD_DEPLOY_SERVER
array=(${string//,/})
project_path=kori_golang
for i in "${!array[@]}"; do 
    echo "Deploy project on server ${array[i]}"
    ssh ubuntu@${array[i]} "ls -la ${project_path}"
    ssh ubuntu@${array[i]} "sudo git -C ${project_path} fetch"
    ssh ubuntu@${array[i]} "sudo git -C ${project_path} pull origin master"
done
