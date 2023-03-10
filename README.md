# Proto Go
Protobuf generated code for Go

# Usage
Import the library into your project:
```go
import provider "github.com/done-devs/proto-go/provider"
```

Then, embed `provider.UnimplementedProviderServer` into a struct.

Read [here](https://github.com/done-devs/done/blob/main/PLUGINS.md) for more details on plugins.

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
%s
```

# Examples
Find example implementations in the [`examples`](examples) directory.
