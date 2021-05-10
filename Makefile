default:
	go build ./cmd/tj

install: default
	go install ./cmd/tj

test:
	go test ./... --cover
