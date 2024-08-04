generate-rpc:
	protoc \
	--go_out=genprotos \
	--go-grpc_out=genprotos \
	protos/event-protos/event.proto \
	protos/auth-protos/auth.proto \


swag:
	swag init -g api/api.go 