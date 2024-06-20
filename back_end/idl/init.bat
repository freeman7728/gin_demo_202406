protoc -I ./idl productService.proto --micro_out ./idl/pb --go_out=./idl/pb
protoc-go-inject-tag -input .\idl\pb\productService.pb.go