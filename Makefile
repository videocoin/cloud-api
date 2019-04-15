
default: protoc

protoc: protoc-rpc \
	protoc-v1-users \
	protoc-v1-accounts \
	protoc-v1-notifications

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


