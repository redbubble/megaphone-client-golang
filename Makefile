MEGAPHONE_FLUENT_PORT = 24224

.PHONY: help create-fluent-container unit-test docker-network-create docker-network-remove all

help: ## Display this help page
	@echo "Welcome to the Golang client for Megaphone!"
	@printf "\n\033[32mEnvironment Variables\033[0m\n"
	@cat $(MAKEFILE_LIST) | egrep -o "\\$$\{[a-zA-Z0-9_]*\}" | sort | uniq | \
		sed -E 's/^[\$$\{]*|\}$$//g' | xargs -I % echo % % | \
		xargs printf "echo \"\033[93m%-30s\033[0m \$$(printenv %s)\n\"" | bash | sed "s/echo //"
	@printf "\n\033[32mMake Targets\033[0m\n"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

unit-test: ## Run the unit test suite
	@echo "--- Unit testing"
	docker run --rm \
		--name megaphone-client-golang \
		--network megaphone-network \
		-e MEGAPHONE_FLUENT_HOST="fluent-container" \
		--volume ${PWD}:/go/src/github.com/redbubble/megaphone-client-golang \
		-it golang:1.10-alpine /bin/sh -c 'cd /go/src/github.com/redbubble/megaphone-client-golang && go test -v -p 1 -cover `go list ./...`'

create-fluent-container: ## Create and start the fluent container needed for the tests
	@echo "--- Creating the fluent container"
	docker create --name fluent-container \
		--network megaphone-network \
		-p ${MEGAPHONE_FLUENT_PORT}:24224 \
		-p ${MEGAPHONE_FLUENT_PORT}:24224/udp \
		-v log:/fluentd/log \
		fluent/fluentd
	docker start fluent-container

stop-fluent-container: ## Stop the fluent container needed for the tests
	docker stop fluent-container

remove-fluent-container: ## Remove the fluent container needed for the tests
	docker rm fluent-container

docker-network-create: ## Create the docker network needed to link the unit-test container and the fluent container
	@echo "--- Creating the docker network"
	docker network create megaphone-network

docker-network-remove: ## Remove the docker network needed to link the unit-test container and the fluent container
	@echo "--- Removing the docker network"
	docker network remove megaphone-network

tests: docker-network-create create-fluent-container unit-test clean

clean: stop-fluent-container remove-fluent-container docker-network-remove ## Clean up the fluent container and the docker network

all: tests
