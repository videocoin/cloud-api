module github.com/videocoin/cloud-api

go 1.12

require (
	github.com/gogo/googleapis v1.2.0
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.3.1
	github.com/grpc-ecosystem/grpc-gateway v1.9.3
	github.com/mailru/dbr v3.0.0+incompatible
	github.com/videocoin/cloud-pkg v0.0.5
	google.golang.org/grpc v1.21.1
)

replace github.com/videocoin/cloud-pkg => ../cloud-pkg

replace github.com/videocoin/cloud-api => ../cloud-api
