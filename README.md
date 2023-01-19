# protogo
Done's plug-in protobuf generated code for Go

# Configuration

## Dependencies

Install the code generation modules.
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
Add this to your PATH

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

Install the Protobuf compiler
```
sudo apt install protobuf-compiler
```
## Code generation
In the root of the repository, run this to generate the code.
```bash
protoc --go_out=. \
    --go-grpc_out=. \
    --experimental_allow_proto3_optional \
    proto/provider.proto
```