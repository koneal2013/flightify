start:
	@make clean
	@make build
	@echo "running main program..."
	@./flightify $(args)

docker-start:
	@make docker-build
	@docker run -p 8080:8080 -d --name flightify docker.io/koneal2013/flightify:latest

build:
	@make test
	@go build

docker-build:
	@docker rm flightify || true
	@docker build . -t koneal2013/flightify

clean:
	@go clean -i

testq:
	@echo "running tests..."
	go test ./server/*

test:
	@make testqv

testqv:
	@echo "running tests..."
	go test -v ./server/*.go

cover:
	@mkdir .coverage || echo "hidden coverage folder exists"
	@go test -v -cover ./*/*.go -coverprofile .coverage/coverage.out
	@go tool cover -html=.coverage/coverage.out -o .coverage/coverage.html

covero:
	@make cover
	@open .coverage/coverage.html
