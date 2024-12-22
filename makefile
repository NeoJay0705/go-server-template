GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get
GO_VET=$(GO_CMD) vet
GO_RUN=$(GO_CMD) run
GO_MOD_DEP=$(GO_CMD) mod download
ALL_PATH=./...

BINARY_NAME=app
UNPACK_PATH=$$path
MAIN_PATH=$$main_path

DOCKER_CMD=docker
DOCKER_BUILD=$(DOCKER_CMD) buildx build
DOCKER_PUSH=$(DOCKER_CMD) push
DOCKER_IMAGE_NAME=go-server-template


.PHONY: deps
deps:
	$(GO_MOD_DEP)

.PHONY: test
test:
	$(GO_TEST) -v $(ALL_PATH) -cover

.PHONY: build
build:
	$(GO_BUILD) -o $(MAIN_PATH)/$(BINARY_NAME) $(MAIN_PATH)

.PHONY: run
run:
	$(GO_RUN) $(MAIN_PATH)

.PHONY: clean
clean:
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)

.PHONY: pack
pack:
	tar -cvzf $(BINARY_NAME)-v$(VERSION).tar.gz $(MAIN_PATH)/$(BINARY_NAME) ./configs ./web

.PHONY: unpack
unpack:
	tar -zxf $(BINARY_NAME)-v$(VERSION).tar.gz -C $(UNPACK_PATH)

.PHONY: docker_build	
docker_build:
	@echo "開始打包 Docker Image - $(DOCKER_FULL_IMAGE)"
	$(DOCKER_BUILD) --platform linux/amd64 -f ./build/Dockerfile -t $(DOCKER_IMAGE_NAME) .

.PHONY: docker_push
docker_push:
	@echo "開始 push docker image - $(DOCKER_FULL_IMAGE)"
	$(DOCKER_PUSH) $(DOCKER_IMAGE_NAME)
