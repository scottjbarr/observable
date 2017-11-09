ci: clean tools prepare deps test

# Install dependencies.
deps:
	go get -t ./...

tools:
	go get github.com/golang/lint/golint

prepare:
	mkdir -p tmp

test: prepare
	go test -coverprofile=tmp/coverage.out

test-coverage: test
	go tool cover -html=tmp/coverage.out

clean:
