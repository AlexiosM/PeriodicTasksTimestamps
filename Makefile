all:
	docker build --tag docker-periodic-tasks .
	docker image tag  docker-periodic-tasks docker-periodic-tasks:latest
	docker run --publish 8080:3000 docker-periodic-tasks
clean:
	docker stop `docker ps -a -q` --time 20
	docker image rm -f docker-periodic-tasks:latest
