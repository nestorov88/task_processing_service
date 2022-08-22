.PHONY: help
help:		## Display help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/##//'

.PHONY: start
start:		## Start dockerised app
	docker-compose -f ./deploy/docker-compose/docker-compose.yaml up -d

.PHONY: stop
stop:		##Stop dockerised app
	docker-compose -f ./deploy/docker-compose/docker-compose.yaml down

.PHONY: test
test:		## Run tests
	@go test -coverprofile=cover.out ./...
	@go tool cover -func cover.out | grep total | awk '{print substr($$3, 1, length($3))}' | sed -e 's/^/Total Coverage: /;'

.PHONY: run
run:		## Start application locally
	@go run cmd/app/main.go

ifndef ADDR
override ADDR = http://127.0.0.1:4000/
endif
.PHONY: test-curl
test-curl:		## Send few test curls. [required ADDR=${value}]
	curl -d @mytasks.json -H "Accept: text/plain" ${ADDR} | bash
	curl -d @mytasks.json -H "Content-type: text/plain" -H "Accept: text/plain" ${ADDR}
	curl -d @mytasks.json -H "Content-type: text/application-json" ${ADDR}

