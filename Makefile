default:
	go build ./cmd/tj

install: default
	go install ./cmd/tj

test:
	go test ./... --cover

cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out