
default: protoc

protoc: protoc-rpc \
	protoc-v1-users \
	protoc-v1-accounts \
	protoc-v1-notifications \
	protoc-v1-pipelines \
	protoc-v1-emitter \
	protoc-v1-profiles \
	protoc-v1-verifier \
	protoc-v1-workorder \
	protoc-v1-transcoder \
	protoc-v1-manager


protoc-rpc:
	protoc \
		-I/usr/local/include \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com \
		-I. \
		--gofast_out=plugins=grpc:. \
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


