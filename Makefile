start:
	@make clean
	@make build
	@echo "running main program..."
	@./flightify

docker-start:
	@make docker-build
	@docker run -p 8080:8080 --name flightify docker.io/koneal2013/flightify:latest

build:
	@make test
	@go build

docker-build:
	@docker build . -t koneal2013/flightify

clean:
	@go clean -i

testq:
	@echo "running tests..."
	go test ./*/*

test:
	@make testqv

testqv:
	@echo "running tests..."
	go test -v ./*/*

cover:
	@mkdir .coverage || echo "hidden coverage folder exists"
	@go test -v -cover ./*/*.go -coverprofile .coverage/coverage.out
	@go tool cover -html=.coverage/coverage.out -o .coverage/coverage.html

covero:
	@make cover
	@open .coverage/coverage.html