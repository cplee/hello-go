build:
	@docker build -t hello-go:latest .
	
test:
	-@docker rm -f hello-go-test
	@docker run -d --name hello-go-test hello-go:latest
	@while docker inspect -f "{{.State.Health.Status}}" hello-go-test | grep starting; do sleep 1; done
	@docker inspect -f "{{.State.Health.Status}}" hello-go-test | grep "^healthy$$"
	-@docker rm -f hello-go-test