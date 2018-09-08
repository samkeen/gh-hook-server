build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/handler cmd/handler/*

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: clean build
	sls deploy --verbose
