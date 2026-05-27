### RUN GOLANG
``` go
go run main.go
```

### Build
``` go
go build
```

### Run Test Integration
``` go
go test ./test/integration_test
go test -v -test.fullpath=true ./test/integration_test
```

### Install Google Wire package untuk dependency injection automatic
``` go
go install github.com/google/wire/cmd/wire@latest

// Export file bin wire nya for linux
export PATH=$PATH:$(go env GOPATH)/bin

// Cek wire berhasil add di path
wire help
```