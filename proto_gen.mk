GOPATH ?= unset_please_set

protoc-rpc:
	cd /go_workspace/src/github.com/videocoin/cloud-api && protoc \
		-I . \
		-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
		-I ${GOPATH}/src/github.com/gogo/googleapis/ \
		-I ${GOPATH}/src \
		--gogofast_out=plugins=grpc:. \
		./rpc/*.proto

protoc-gateway-v1-%:
	cd /go_workspace/src/github.com/videocoin/cloud-api && protoc \
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
	cd /go_workspace/src/github.com/videocoin/cloud-api && sed -i.bak "s/empty.Empty/types.Empty/g" ./$*/v1/*.pb.gw.go && rm ./$*/v1/*.pb.gw.go.bak

protoc-v1-%:
	cd /go_workspace/src/github.com/videocoin/cloud-api && protoc \
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
	cd /go_workspace/src/github.com/videocoin/cloud-api && protoc \
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
