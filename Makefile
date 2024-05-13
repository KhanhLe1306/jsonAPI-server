build: 
	go build -o bin/financial_server

run: build
	./bin/financial_server

test: 
	go test -v ./...