To setup
```
mkdir -p third_party/googleapis
git clone https://github.com/googleapis/googleapis.git third_party/googleapis

```

# ⭐️ Proto-Test: Simple Example Repo in Go (gRPC + HTTP)

A minimal, clean, and ready-to-go example demonstrating how you can quickly set up a simple API server using:
- 🐹 **Go (Golang)**
- 📦 **Protocol Buffers (protobuf)**
- ⚡️ **gRPC**
- 🌐 **RESTful HTTP via grpc-gateway**
- 🎯 **Validation using protoc-gen-validate**

## 📁 Project Structure

```
proto-test/
│
├── api
│ └── test.proto # Proto schema definition
│
├── cmd
│ └── main.go # Entry point for the application
│
├── internal
│ └── server
│ └── server.go # API implementation logic
│
├── generated # Auto-generated proto files (committed to repo)
│ └── api
│ ├── test.pb.go
│ ├── test_grpc.pb.go
│ ├── test.pb.gw.go
│ └── test.pb.validate.go
│
├── third_party # Third-party proto dependencies (not committed)
│
├── Makefile # Commands for easy setup
├── go.mod # Go modules file
└── go.sum # Go dependency checksums
```


## 🚧 Installation & Setup

### 🛠 Prerequisites

- `Go` (v1.16 or above) - [install link](https://golang.org/dl/)
- `protoc` (Protocol compiler) - [install link](https://github.com/protocolbuffers/protobuf/releases)
- `grpcurl` (optional, recommended to test grpc) - [install link](https://github.com/fullstorydev/grpcurl)

Install Go plugins:
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/envoyproxy/protoc-gen-validate@latest

export PATH="$PATH:$(go env GOPATH)/bin"
```

### 📥 Getting the repository
```
git clone https://github.com/harshb910-mercari/proto-test.git
cd proto-test
```

### 🚀 Generate proto files
```
# First-time setup of third-party dependencies
mkdir -p third_party && git clone https://github.com/googleapis/googleapis third_party/googleapis

# Generate your proto files
make generate
```

### ▶️ Start the Server
To run your server:
```
go run cmd/main.go
```
It will start two servers:
```
Type	URL
gRPC	localhost:50051
HTTP	localhost:8080
```
### 🎯 Testing API Endpoints
👉 Testing via gRPC (grpcurl):
```
grpcurl -plaintext -d '{
    "business_id": "1234567",
    "name": {
        "value": "Haru"
    },
    "phone_number": "12345678901",
    "photo_ids": [
        "1234567",
        "2345678"
    ],
    "position_number": 12,
    "sub_paths": [
        "nulla commodo dolore ut",
        "sit ullamco labore",
        "exercitation esse voluptate"
    ]
}' localhost:50051 api.TestService/SayHello
```
Response:
```json
{
"message": "Hello, Haru!"
}
```

### 👉 Testing via HTTP REST (curl):

```
curl --location 'localhost:8080/v1/say_hello' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Haru",
    "business_id": "1234567",
    "photo_ids" : ["1234567", "2345678"],
    
    "sub_paths" : ["1", "2", "3", "4", "5"],
    
    "position_number": 2,
    "phone_number" : "12345678901"
}'
```

Response:
```json
{
"message": "Hello, Haru!"
}
```

### ✅ Request Validation Example
Proto schema validation (name must be at least 2 characters):

```
curl -X POST \
-H "Content-Type: application/json" \
-d '{"name":"M"}' \
http://localhost:8080/v1/say_hello
```
Validation Error Response:

```json
{
  "code": 3,
  "message": "Validation failed: invalid TestRequest.Name: value length must be between 3 and 150 runes, inclusive",
  "details": []
}
```

### 🔧 Useful Commands
```
Command           Task
make generate     Generate protos
make clean        Clean generated proto files
```
