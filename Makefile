coverage:
	go test -cover ./...

docs:
	godoc -http :6060

test:
	go test -v ./...

