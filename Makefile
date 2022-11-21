SERVICE		:= gin-skeleton
NAME		:= ghcr.io/maribowman/$(SERVICE)
GIT_BRANCH	:= $(shell git rev-parse --abbrev-ref HEAD)
GIT_HASH	:= $(shell git rev-parse --short HEAD)
TAG			:= $(GIT_BRANCH)-$(GIT_HASH)
IMAGE		:= $(NAME):$(TAG)
STAGE		:= prod

### DOCKER
.PHONY: build
build:
	@echo starting build...
	@docker build -q -t $(IMAGE) -t $(NAME):latest .
	@docker image prune -f --filter label=stage=builder > /dev/null

#login:
	@#docker login ghcr.io -u $(GITHUB_USER) -p $(GITHUB_PAT)

push: build
	@echo pushing images...
	@docker push $(IMAGE)
	@docker push $(NAME):latest

.PHONY: run
run: build
	@docker run -d --rm --name $(SERVICE)_$(TAG) $(NAME):latest > /dev/null

stop:
	@docker stop $$(docker ps -q) > /dev/null

.PHONY: postgres
postgres:
	@docker run -d --rm -p 5432:5432 \
          --name postgres \
          -e POSTGRES_USER=postgres \
          -e POSTGRES_PASSWORD=postgres \
         postgres:15-alpine3.16 > /dev/null

### TESTING
.PHONY: tests
tests:
	@go test -race ./...

cover:
	@go test -cover ./...
