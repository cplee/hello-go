SOURCE_COMMIT := $(shell git rev-parse HEAD)

build:
	@echo $(SOURCE_COMMIT)
	@docker build -t hello-go:latest --build-arg SOURCE_COMMIT=$(SOURCE_COMMIT) .
	
test:
	-@docker rm -f hello-go-test
	@docker run -d --name hello-go-test hello-go:latest
	@while docker inspect -f "{{.State.Health.Status}}" hello-go-test | grep starting; do sleep 1; done
	@docker inspect -f "{{.State.Health.Status}}" hello-go-test | grep "^healthy$$"
	-@docker rm -f hello-go-test
	
run: 
	-@docker rm -f hello-go > /dev/null 2>&1
	@docker run -p 8080:8080 -d --name hello-go hello-go:latest
	@curl http://localhost:8080 
	@docker logs -f hello-go
	