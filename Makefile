include proto_gen.mk


DOCKER_IMAGE_NAME = protobuf-build-env
PROJECT_NAME = videocoin-network
REPOSITORY=gcr.io/$(PROJECT_NAME)/$(DOCKER_IMAGE_NAME)

DOCKER_BUILD_COMMAND=docker run \
    -v $(shell pwd):/go_workspace/src/github.com/videocoin/cloud-api \
    -v $(shell pwd)/proto_gen.mk:/app/proto_gen.mk \
    -w /go_workspace/src/github.com/videocoin/cloud-api \
	gcr.io/videocoin-network/protobuf-build-env

default: protoc

protoc: protoc-rpc \
	protoc-gateway-v1-users \
	protoc-gateway-v1-accounts \
	protoc-gateway-v1-streams \
	protoc-gateway-v1-profiles \
	protoc-gateway-v1-miners \
	protoc-v1-accounts \
	protoc-v1-validator \
	protoc-v1-syncer \
	protoc-v1-splitter \
	protoc-v1-notifications \
	protoc-v1-emitter \
	protoc-v1-dispatcher \
	protoc-v1-iam \
	protoc-v1-servicemanager \
	protoc-private-v1-streams \
	protoc-v1-billing \
	protoc-gateway-v1-billing \
	protoc-private-v1-billing

docker-pull-image:
	docker pull $(REPOSITORY):latest

docker-protoc: docker-protoc-rpc \
	docker-protoc-gateway-v1-users \
	docker-protoc-gateway-v1-accounts \
	docker-protoc-gateway-v1-streams \
	docker-protoc-gateway-v1-profiles \
	docker-protoc-gateway-v1-miners \
	docker-protoc-v1-accounts \
	docker-protoc-v1-validator \
	docker-protoc-v1-syncer \
	docker-protoc-v1-splitter \
	docker-protoc-v1-notifications \
	docker-protoc-v1-emitter \
	docker-protoc-v1-dispatcher \
	docker-protoc-v1-iam \
	docker-protoc-v1-servicemanager \
	docker-protoc-private-v1-streams \
	docker-protoc-manager-v1-profiles \
	docker-protoc-v1-billing \
	docker-protoc-gateway-v1-billing \
	docker-protoc-private-v1-billing

docker-protoc-rpc:
	${DOCKER_BUILD_COMMAND} --target protoc-rpc

docker-protoc-gateway-v1-%:
	${DOCKER_BUILD_COMMAND} --target protoc-gateway-v1-$*

docker-protoc-private-v1-%:
	${DOCKER_BUILD_COMMAND} --target protoc-private-v1-$*

docker-protoc-v1-%:
	${DOCKER_BUILD_COMMAND} --target protoc-v1-$*

docker-protoc-python-gateway-v1-%:
	${DOCKER_BUILD_COMMAND} --target protoc-python-gateway-v1-$*

docker-protoc-python-private-v1-%:
	${DOCKER_BUILD_COMMAND} --target protoc-python-private-v1-$*

docker-protoc-manager-v1-%:
	${DOCKER_BUILD_COMMAND} --target protoc-manager-v1-$*
