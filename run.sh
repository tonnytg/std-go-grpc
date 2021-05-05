# export module
export GO111MODULE=on

# Install Modules
go get google.golang.org/protobuf/cmd/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc

# update your path
export PATH="$PATH:$(go env GOPATH)/bin"

# download files to explore your test
mkdir demo
cd demo
git clone -b v1.35.0 https://github.com/grpc/grpc-go

# build proto
protoc --proto_path=proto proto/*.proto --go_out=pb

# build proto 2
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

# install Evans to test
# official repo https://github.com/ktr0731/evans
brew tap ktr0731/evans
brew install evans

# Call with Evans
evans -r repl --host localhost --port 50051
# service UserService
# call AddUser
#pb.UserService@localhost:50051> call AddUser
#id (TYPE_STRING) => 1
#name (TYPE_STRING) => Antonio
#email (TYPE_STRING) => tonnytg@gmail.com


