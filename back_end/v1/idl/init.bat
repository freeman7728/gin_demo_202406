protoc -I ./idl detailService.proto --micro_out ./idl/pb --go_out=./idl/pb
protoc-go-inject-tag -input .\idl\pb\detailService.pb.go