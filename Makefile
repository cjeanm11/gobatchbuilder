start:
	go run cmd/launch/main.go

test:
	find tests -name "*_test.go" -exec go test -v {} \;

build:
	go build -o gobatchbuilder cmd/launch/main.go

clean:
	rm -f go