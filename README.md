# authentication
It's an authentication application written in Go.

Golang - Basic Login/Register services implemented using
- gRPC
- Database (MySQL)

### Do Insert an application.yaml file with following properties

- server:
  - host: hostname
  - port: portName

- database:
  - port: db port
  - user: db username
  - password: db user's password
  - database_name: authentication
  - address: db address

Set the go path by running

> export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))

After installing all the dependencies

> go run AuthenticationApp.go

It will start the grpc server at the address mentioned in your application.yaml file

To run the Unit tests

> go test server/GrpcServer_test.go -v
