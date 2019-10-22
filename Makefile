
default: protoc

protoc: protoc-rpc \
	protoc-gateway-v1-users \
	protoc-gateway-v1-streams \
	protoc-gateway-v1-accounts \
	protoc-gateway-v1-streams \
	protoc-gateway-v1-profiles \
	protoc-v1-accounts \
	protoc-v1-validator \
	protoc-v1-syncer \
	protoc-v1-notifications \
	protoc-v1-emitter \
	protoc-v1-dispatcher \
	protoc-private-v1-streams

protoc-rpc:
	protoc \
		-I . \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I ${GOPATH}/src/github.com/gogo/googleapis/ \
		-I ${GOPATH}/src \
		--gogofast_out=plugins=grpc:. \
		./rpc/*.proto

protoc-v1-%:
	protoc \
		-I/usr/local/include \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com \
		-I. \
		--gogofast_out=plugins=grpc,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
. \
		./$*/v1/*.proto

protoc-private-v1-%:
	protoc \
		-I/usr/local/include \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com \
		-I. \
		--gogofast_out=plugins=grpc,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
. \
		./$*/private/v1/*.proto

protoc-gateway-v1-%:
	protoc \
		-I . \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I ${GOPATH}/src/github.com/gogo/googleapis/ \
		-I ${GOPATH}/src \
		--gogofast_out=plugins=grpc,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
. \
		--grpc-gateway_out=allow_patch_feature=false,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
. \
		--swagger_out=./openapi \
		./$*/v1/*.proto
	
	# Workaround for https://github.com/grpc-ecosystem/grpc-gateway/issues/229.
	sed -i.bak "s/empty.Empty/types.Empty/g" ./$*/v1/*.pb.gw.go && rm ./$*/v1/*.pb.gw.go.bak


DOCKER_IMAGE_NAME = protobuf-build-env
LIVE_PLANET_PROJECT_NAME = videocoin-network
REPOSITORY=gcr.io/$(LIVE_PLANET_PROJECT_NAME)/$(DOCKER_IMAGE_NAME)

DOCKER_BUILD_COMMAND=docker run \
    -v $(shell pwd):/go_workspace/src/github.com/videocoin/cloud-api -v $(shell pwd)/proto_gen.mk:/app/proto_gen.mk -w $(shell pwd) \
	gcr.io/liveplanet-cloud-internal/protobuf-build-env

docker-pull-image:
	docker pull $(REPOSITORY):latest

docker-protoc: docker-protoc-rpc \
	docker-protoc-gateway-v1-users \
	docker-protoc-gateway-v1-streams \
	docker-protoc-gateway-v1-accounts \
	docker-protoc-gateway-v1-streams \
	docker-protoc-gateway-v1-profiles \
	docker-protoc-v1-accounts \
	docker-protoc-v1-validator \
	docker-protoc-v1-syncer \
	docker-protoc-v1-notifications \
	docker-protoc-v1-emitter \
	docker-protoc-v1-dispatcher \
	docker-protoc-private-v1-streams

docker-protoc-rpc:
	${DOCKER_BUILD_COMMAND} --target protoc-rpc

docker-protoc-gateway-v1-%:
	${DOCKER_BUILD_COMMAND} --target protoc-gateway-v1-$*

docker-protoc-private-v1-%:
	${DOCKER_BUILD_COMMAND} --target protoc-private-v1-$*

docker-protoc-v1-%:
	${DOCKER_BUILD_COMMAND} --target protoc-v1-$*