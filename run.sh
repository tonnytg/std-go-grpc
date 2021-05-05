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

