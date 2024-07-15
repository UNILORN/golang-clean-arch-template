# Golang - Clean Architecture Templates

## Get Started

### Install Golang & go mod download

This code is intended to create documentation on installing Go and setting up go mod. Please follow the steps below to install Go and set up go mod.

1. Download the appropriate installer for your OS from the official Go website (https://golang.org/dl/).
2. Run the downloaded installer and follow the instructions to install Go.
3. Open a terminal or command prompt and execute the following command to check the Go version:

```
$ go version
Example output: go version go1.17.2 darwin/amd64
```

4. Use go mod to install the necessary packages.

```
$ go mod download
Example output: go: finding module for package github.com/example/mypackage
     go: downloading github.com/example/mypackage v1.2.3
```

With this, the installation of Go and the setup of go mod are complete. You can now use go mod to manage packages and resolve dependencies.

### Run Server

You can run the server using the following methods:
- Execute the command `go run internal/cmd/main.go`.
- Alternatively, execute the command `docker-compose up -d`.
Both methods will start the server.

Please note that you need to run MySQL separately in order to start the server. You can use docker-compose or any other method to start MySQL.

## develop - Architecture

This repository adopts the clean architecture.

The repository that was referenced is as follows:

https://github.com/code-kakitai/code-kakitai



```
└── internal
    ├── application // UseCase layer
    │   ├── shop
    │   │   ├── fetch_shop_use_case.go  // Created for each use case
    │   │   ├── save_shop_use_case.go
    │   │   └── shop_query_service.go   // Interface for Get-related DB queries
    ├── cmd                             // Startup commands, migration commands, etc.
    │   └── main.go
    ├── config                          // Environment variable information, etc.
    │   └── config.go 
    ├── domain                          // Domain layer
    │   └── shop
    │       ├── shop.go                 // Contains entities and data processing related to entities
    │       └── shop_repository.go      // Repository interface for Infrastructure (DI is performed in the router)
    ├── infrastructure                  // Infrastructure layer
    │   └── mysql
    │       ├── db
    │       │   ├── connect_db.go       // Connection to MySQL
    │       │   ├── db_test
    │       │   │   └── container.go
    │       │   ├── dbgen               // *Auto-generated
    │       │   │   ├── db.go           // *Auto-generated
    │       │   │   ├── models.go       // *Auto-generated
    │       │   │   ├── querier.go      // *Auto-generated
    │       │   │   └── shop.sql.go     // *Auto-generated
    │       │   ├── query
    │       │   │   └── shop.sql        // Query to be used
    │       │   └── schema
    │       │       ├── migration.go   
    │       │       └── schema.sql      // Migration Schemas
    │       ├── fixtures
    │       │   └── products.yml
    │       ├── query_service
    │       │   ├── common_test.go
    │       │   ├── shop_query_service.go
    │       │   └── shop_query_service_test.go
    │       └── repository
    │           ├── common_test.go
    │           ├── shop_repository.go
    │           ├── shop_repository_test.go
    │           └── transaction_manager.go
    ├── presentation                    // Presentation layer
    │   ├── health_handler
    │   │   ├── handler.go
    │   │   └── response.go
    │   ├── settings
    │   │   ├── gin.go                  // go-gin
    │   │   └── middleware.go
    │   └── shop
    │       ├── handler.go              // Request Handler
    │       ├── request.go              // Request Types
    │       └── response.go             // Response Types
    ├── server
    │   ├── api_test
    │   │   ├── api_read_test           // API Test
    │   │   ├── api_write_test          // API Test
    │   │   └── test_utils
    │   │       └── common.go
    │   ├── route                       // Router
    │   │   └── route.go
    │   └── server.go
```

## develop - Mockgen

### Install

mockgen is a Go mock generation tool. It allows you to generate mock implementations from interfaces.

To install mockgen, follow these steps:

1. Open a terminal or command prompt.
2. Execute the following command to install mockgen:

    ```
    $ go install github.com/golang/mock/mockgen@latest
    ```

    This command installs mockgen with the latest version (v1.6.0).

3. Once the installation is complete, execute the following command to check the mockgen version:

    ```
    $ mockgen --version
    ```

    Example output: mockgen version v1.6.0

With this, the installation of mockgen is complete. You can now use mockgen to generate mock implementations from interfaces.

### Usage

With mockgen, you can generate mock implementations from Go interfaces. Follow these steps to use mockgen:

1. Open a terminal or command prompt.

2. Execute the following command to use mockgen:

    ```
    $ mockgen -source=path/to/interface.go -destination=path/to/mock.go
    ```

    This command generates mock implementations in `path/to/mock.go` based on the interface defined in `path/to/interface.go`.

    For example, if `path/to/interface.go` has the following interface:

    ```go
    package mypackage

    type MyInterface interface {
         DoSomething() error
    }
    ```

    Running the above command will generate the following mock implementation in `path/to/mock.go`:

    ```go
    package mypackage

    import (
         "github.com/golang/mock/gomock"
    )

    type MockMyInterface struct {
         ctrl     *gomock.Controller
         recorder *MockMyInterfaceMockRecorder
    }

    type MockMyInterfaceMockRecorder struct {
         mock *MockMyInterface
    }

    func NewMockMyInterface(ctrl *gomock.Controller) *MockMyInterface {
         mock := &MockMyInterface{ctrl: ctrl}
         mock.recorder = &MockMyInterfaceMockRecorder{mock}
         return mock
    }

    func (m *MockMyInterface) EXPECT() *MockMyInterfaceMockRecorder {
         return m.recorder
    }

    func (m *MockMyInterface) DoSomething() error {
         ret := m.ctrl.Call(m, "DoSomething")
         return ret.Error(0)
    }
    ```

    Now you know how to use mockgen to generate mock implementations from interfaces.

3. You can use the generated mock implementations to write test code.

These are the basic steps for using mockgen. For more options and customization, refer to the official documentation.

## develop - sqlc

### Install

sqlc is a tool for generating Go code from SQL queries. It allows you to automatically generate Go code based on SQL files.

To install sqlc, follow these steps:

1. Open a terminal or command prompt.
2. Execute the following command to install sqlc:

    ```
    $ go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
    ```

    This command installs sqlc with the latest version.

3. Once the installation is complete, execute the following command to check the sqlc version:

    ```
    $ sqlc version
    ```

    Example output: sqlc version v1.9.0

With this, the installation of sqlc is complete. You can now use sqlc to generate Go code from SQL files.

### Usage

You can use the `generate` command of sqlc to generate Go code from SQL files. Follow these steps to use the `generate` command:

1. Open a terminal or command prompt.
2. Execute the following command to use the `generate` command:

     ```
     $ sqlc generate
     ```

This command reads the `sqlc.yaml` file and generates Go code based on the specified SQL files.

> Note: The `sqlc.yaml` file should be placed in the root directory of your project. You can also configure the code generation settings in the `sqlc.yaml` file.

3. When you execute the `generate` command, Go code will be generated based on the specified SQL files.

```go
// Generated Go code
```

The generated Go code is automatically generated based on the queries and table definitions in the specified SQL files.

Now you know how to generate Go code from SQL files using the `generate` command of sqlc. For more options and usage details, refer to the official documentation.

## LICENSE

MIT LICENSE