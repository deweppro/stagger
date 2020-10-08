
tests:
	docker-compose -f dockerfiles/docker-compose.yaml -p test down
	docker-compose -f dockerfiles/docker-compose.yaml -p test up app
	docker-compose -f dockerfiles/docker-compose.yaml -p test down

tests_unit:
	go test -v -race -run Unit ./...

tests_integration:
	go test -v -race -run Integration ./...