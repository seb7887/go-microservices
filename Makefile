proto:
	protoc -I . ./proto/*.proto --go_out=plugins=grpc:./proto
all:
	proto