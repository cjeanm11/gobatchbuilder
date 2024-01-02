start:
	go run cmd/main.go

test:
	find tests -name "*_test.go" -exec go test -v {} \;

build:
	go build -o gobatchbuilder cmd/main.go

clean:
	rm -f go