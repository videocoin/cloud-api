
default: protoc

protoc: protoc-rpc \
	protoc-gateway-v1-users \
	protoc-gateway-v1-pipelines \
	protoc-gateway-v1-accounts \
	protoc-v1-notifications \
	protoc-v1-emitter \
	protoc-old-v1-profiles \
	protoc-old-v1-verifier \
	protoc-old-v1-workorder \
	protoc-old-v1-transcoder \
	protoc-old-v1-manager

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
		--gofast_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--validate_out="lang=gogo:." \
		./$*/v1/*.proto

protoc-old-v1-%:
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
		./$*/v1/*.proto

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

