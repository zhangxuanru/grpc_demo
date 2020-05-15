#protoc --go_out=plugins=grpc:../pbfiles Prod.proto
#protoc --go_out=plugins=grpc:../pbfiles module/Prod.Module.proto


protoc --go_out=plugins=grpc:../pbfiles Order.proto
protoc --go_out=plugins=grpc:../pbfiles module/Order.Module.proto
