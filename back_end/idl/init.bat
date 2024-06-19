protoc -I ./idl producerService.proto --micro_out ./idl/pb --go_out=./idl/pb
protoc-go-inject-tag -input .\idl\pb\producerService.pb.go