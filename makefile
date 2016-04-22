build:
	export GOPATH=$(shell pwd) && echo $$GOPATH && go build src/main.go
run:
	./main
