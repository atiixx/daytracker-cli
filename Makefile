version: '3'

run:
	go run ./cmd/app/

build:
	go build -o bin/daytracker-cli ./cmd/app/

test:
	go test ./tests -v

clean:
	rm -rf bin/

