start:
	@make clean
	@make build
	@echo "running main program..."
	@./flightify

build:
	@make test
	@go build

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