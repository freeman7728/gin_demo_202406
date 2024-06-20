protoc -I ./idl listService.proto --micro_out ./idl/pb --go_out=./idl/pb
protoc-go-inject-tag -input .\idl\pb\listService.pb.go