
tests:
	docker-compose -f dockerfiles/docker-compose.yaml -p test down
	docker-compose -f dockerfiles/docker-compose.yaml -p test up app
	docker-compose -f dockerfiles/docker-compose.yaml -p test down
