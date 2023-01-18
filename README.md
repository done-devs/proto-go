# proto-go
Done's plug-in protobuf generated code for Go

# Code generation
```bash
protoc --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative \
    --experimental_allow_proto3_optional \
    proto/provider.proto
```
# Configuration

## Dependencies

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
Add this to your PATH

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```


```
sudo apt install protobuf-compiler golang-goprotobuf-dev
```