path=./deployments/app/docker-compose.yml
start:
	- sudo docker-compose -f ${path} down
	- sudo docker-compose -f ${path} up
build:
	- sudo docker-compose -f ${path} down
	- sudo docker-compose -f ${path} up --build --remove-orphans
clear:
	- sudo docker stop $$(docker ps -a -q)
	- sudo docker rm $$(docker ps -a -q)
	- sudo docker rmi -f $$(docker images -a -q)
